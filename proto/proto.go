// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"golang.org/x/exp/slices"
)

// NOTE: The term "Global Fit Profile" refers to the definition provided in the Profile.xlsx proto.

const (
	MesgDefinitionMask = 0b01000000 // Mask for determining if the message type is a message definition.

	DefaultFileHeaderSize byte   = 14     // The preferred size is 14
	DataTypeFIT           string = ".FIT" // FIT is a constant string ".FIT"
)

// Fit represents a structure for Fit Files.
type Fit struct {
	FileHeader FileHeader // File Header contains either 12 or 14 bytes
	Messages   []Message  // Messages.
	CRC        uint16     // Cyclic Redundancy Check 16-bit value to ensure the integrity of the messages.
}

// WithMessages set Messages and return the pointer to the Fit.
func (f *Fit) WithMessages(messages ...Message) *Fit {
	f.Messages = make([]Message, len(messages))
	copy(f.Messages, messages)
	return f
}

// FileHeader is a Fit's FileHeader with either 12 bytes size without CRC or a 14 bytes size with CRC, while 14 bytes size is the preferred size.
type FileHeader struct {
	Size            byte   // Header size either 12 (legacy) or 14.
	ProtocolVersion byte   // The Fit Protocol version which is being used to encode the Fit file.
	ProfileVersion  uint16 // The Fit Profile Version (associated with data defined in Global Fit Profile).
	DataSize        uint32 // The size of the messages in bytes (this field will be automatically updated by the encoder)
	DataType        string // ".FIT" (a string constant)
	CRC             uint16 // Cyclic Redundancy Check 16-bit value to ensure the integrity if the header. (this field will be automatically updated by the encoder)
}

// MessageDefinition is the definition of the upcoming data messages.
type MessageDefinition struct {
	Header                    byte                       // The message definition header with mask 0b01000000.
	LocalMesgNum              byte                       // Serves as a reference for the upcoming message during decoding and encoding.
	Reserved                  byte                       // Currently undetermined; the default value is 0.
	Architecture              byte                       // The Byte Order to be used to decode the values of both this message definition and the upcoming message. (0: Little-Endian, 1: Big-Endian)
	MesgNum                   typedef.MesgNum            // Global Message Number defined by factory (retrieved from Profile.xslx). (endianness of this 2 Byte value is defined in the Architecture byte)
	FieldDefinitions          []FieldDefinition          // List of the field definition
	DeveloperFieldDefinitions []DeveloperFieldDefinition // List of the developer field definition (only if Developer Data Flag is set in Header)
}

// Clone clones MessageDefinition
func (m MessageDefinition) Clone() MessageDefinition {
	m.FieldDefinitions = slices.Clone(m.FieldDefinitions)
	m.DeveloperFieldDefinitions = slices.Clone(m.DeveloperFieldDefinitions)
	return m
}

// FieldDefinition is the definition of the upcoming field within the message's structure.
type FieldDefinition struct {
	Num      byte              // The field definition number
	Size     byte              // The size of the upcoming value
	BaseType basetype.BaseType // The type of the upcoming value to be represented
}

// FieldDefinition is the definition of the upcoming developer field within the message's structure.
type DeveloperFieldDefinition struct { // 3 bits
	Num                byte // Map to the `field_definition_number` of a `field_description` Message.
	Size               byte // Size (in bytes) of the specified FIT message’s field
	DeveloperDataIndex byte // Maps to the `developer_data_index`` of a `developer_data_id` Message
}

// CreateMessageDefinition creates new MessageDefinition base on given Message.
func CreateMessageDefinition(mesg *Message) (mesgDef MessageDefinition) {
	mesgDef = MessageDefinition{
		Header:       MesgDefinitionMask,
		Reserved:     0, // default
		Architecture: mesg.Architecture,
		MesgNum:      mesg.Num,
	}

	fieldDefinitions := make([]FieldDefinition, 0, len(mesg.Fields))
	for _, field := range mesg.Fields {
		size := field.Size
		size *= typedef.Len(field.Value)
		if _, ok := field.Value.(string); ok {
			size += 1 // +1 utf-8 null terminator string: 0x00
		}
		fieldDefinitions = append(fieldDefinitions, FieldDefinition{
			Num:      field.Num,
			Size:     size,
			BaseType: field.Type.BaseType(),
		})
	}
	mesgDef.FieldDefinitions = fieldDefinitions

	if len(mesg.DeveloperFields) == 0 {
		return
	}

	mesgDef.Header |= DevDataMask
	developerFieldDefinitions := make([]DeveloperFieldDefinition, 0, len(mesg.DeveloperFields))
	for _, developerField := range mesg.DeveloperFields {
		developerFieldDefinitions = append(developerFieldDefinitions, DeveloperFieldDefinition{
			Num:                developerField.Num,
			Size:               developerField.Type.Size(),
			DeveloperDataIndex: developerField.DeveloperDataIndex,
		})
	}

	mesgDef.DeveloperFieldDefinitions = developerFieldDefinitions

	return
}

// Message is a FIT protocol message containing the data defined in the Message Definition
type Message struct {
	Header          byte             // Message Header serves to distinguish whether the message is a Normal Data or a Compressed Timestamp Data. Unlike MessageDefinition, Message's Header should not contain Developer Data Flag.
	Name            string           // The Name of the Global Message Number.
	Num             typedef.MesgNum  // Global Message Number defined in Global Fit Profile, except number within range 0xFF00 - 0xFFFE are manufacturer specific number.
	LocalNum        byte             // Only used as a reference for its definition message during decoding and encoding.
	Architecture    byte             // Architecture type / Endianess. Must be the same
	Fields          []Field          // List of Field
	DeveloperFields []DeveloperField // List of DeveloperField
}

// WithFields copies the provided fields into the message's fields.
func (m Message) WithFields(fields ...Field) Message {
	m.Fields = make([]Field, len(fields))
	copy(m.Fields, fields)
	return m
}

// WithFieldValues assigns the values of the targeted fields with the given map,
// where map[byte]any represents the field numbers and their respective values.
func (m Message) WithFieldValues(fieldNumValues map[byte]any) Message {
	for i := range m.Fields {
		value, ok := fieldNumValues[m.Fields[i].Num]
		if !ok {
			continue
		}
		if value != nil { // message created from factory should be non-nil, extra check for user-defined message.
			m.Fields[i].Value = value
		}
	}
	return m
}

// WithFields copies the provided fields into the message's fields.
func (m Message) WithDeveloperFields(developerFields ...DeveloperField) Message {
	m.DeveloperFields = make([]DeveloperField, len(developerFields))
	copy(m.DeveloperFields, developerFields)
	return m
}

// FieldByNum returns a pointer to the Field in a Message, if not found return false.
func (m *Message) FieldByNum(num byte) (*Field, bool) {
	for i := range m.Fields {
		if m.Fields[i].Num == num {
			return &m.Fields[i], true
		}
	}
	return nil, false
}

// RemoveFieldByNum removes Field in a Message by num.
func (m *Message) RemoveFieldByNum(num byte) {
	var idx *int
	for i := range m.Fields {
		if m.Fields[i].Num != num {
			idx = &i
			break
		}
	}

	if idx != nil {
		m.Fields = append(m.Fields[:*idx], m.Fields[*idx+1:]...)
	}
}

// Clone clones Message.
func (m Message) Clone() Message {
	fields := make([]Field, 0)
	for i := range m.Fields {
		if m.Fields[i].FieldBase == nil { // sanitize fields.
			continue
		}
		fields = append(fields, m.Fields[i].Clone())
	}
	m.Fields = fields

	developerFields := make([]DeveloperField, 0)
	for i := range m.DeveloperFields {
		if m.DeveloperFields[i] == (DeveloperField{}) {
			continue
		}
		developerFields = append(developerFields, m.DeveloperFields[i].Clone())
	}
	m.DeveloperFields = developerFields

	return m
}

// FieldBase acts as a fundamental representation of a field as defined in the Global Fit Profile.
// The value of this representation should not be altered,except in the case of an unknown field.
type FieldBase struct {
	Name       string              // Defined in the Global FIT profile for the specified FIT message, otherwise its a manufaturer specific name (defined by manufacturer).
	Num        byte                // Defined in the Global FIT profile for the specified FIT message, otherwise its a manufaturer specific number (defined by manufacturer). (255 == invalid)
	Type       profile.ProfileType // Type is defined type that serves as an abstraction layer above base types (primitive-types), e.g. DateTime is a time representation in uint32.
	Size       byte                // Size of decoded Value in bytes.
	Scale      float64             // A scale or offset specified in the FIT profile for binary fields (sint/uint etc.) only. the binary quantity is divided by the scale factor and then the offset is subtracted. (default: 1)
	Offset     float64             // A scale or offset specified in the FIT profile for binary fields (sint/uint etc.) only. the binary quantity is divided by the scale factor and then the offset is subtracted. (default: 0)
	Units      string              // Units of the value, such as m (meter), m/s (meter per second), s (second), etc.
	Accumulate bool                // Flag to indicate if the value of the field is accumulable.
	Components []Component         // List of components
	SubFields  []SubField          // List of sub-fields
}

// Field represents the full representation of a field, as specified in the Global Fit Profile.
type Field struct {
	*FieldBase           // PERF: Embedding the struct as a pointer to avoid runtime duffcopy when creating a field since FieldBase should not be altered.
	Value           any  // The decoded value, should always be any of these: any, []any, []byte. The underlying value of any will always a primitive-types: int8, uint16, etc.
	IsExpandedField bool // A flag to detect whether this field is generated through component expansion.
}

// WithValue returns a Field containing v value.
func (f Field) WithValue(v any) Field {
	f.Value = v
	return f
}

// SubFieldSubstitution returns any sub-field that can substitute the properties interpretation of the parent Field (Dynamic Field).
func (f *Field) SubFieldSubtitution(mesgRef *Message) (*SubField, bool) {
	for i := range f.SubFields {
		subField := &f.SubFields[i]
		for j := range subField.Maps {
			smap := &subField.Maps[j]
			fieldRef, ok := mesgRef.FieldByNum(smap.RefFieldNum)
			if !ok {
				return nil, false
			}
			if fieldRef.isValueEqualTo(smap.RefFieldValue) {
				return subField, true
			}
		}
	}
	return nil, false
}

// isValueEqualTo compare if Value == SubField's Map RefFieldValue.
// fit documentation on dynamic fields says: reference fields must be of integer type, floating point reference values are not supported.
func (f *Field) isValueEqualTo(refFieldValue int64) bool {
	v, ok := typeconv.IntegerToInt64(f.Value)
	if !ok {
		return ok
	}
	return v == refFieldValue
}

// Clone clones Field
func (f Field) Clone() Field {
	if f.FieldBase == nil {
		return f
	}

	fieldBase := *f.FieldBase // also include FieldBase, clone is meant to be a deep copy
	fieldBase.Components = slices.Clone(fieldBase.Components)
	fieldBase.SubFields = slices.Clone(fieldBase.SubFields)
	for i := range fieldBase.SubFields {
		fieldBase.SubFields[i] = fieldBase.SubFields[i].Clone()
	}
	f.FieldBase = &fieldBase

	return f
}

// Developer Field is a way to add custom data fields to existing messages. Developer Data Fields can be added to
// any message at runtime by providing a self-describing field definition. Developer Data Fields are also used by
// the Connect IQ FIT Contributor library, allowing Connect IQ apps and data fields to include custom data in
// FIT Activity files during the recording of activities. [Added since protocol version 2.0]
//
// NOTE: If Developer Field contains a valid NativeMesgNum and NativeFieldNum,
// the value should be treated as native value (scale, offset, etc shall apply).
type DeveloperField struct {
	Num                byte
	DeveloperDataIndex byte
	Size               byte
	NativeMesgNum      typedef.MesgNum
	NativeFieldNum     byte
	Name               string
	Type               basetype.BaseType
	Units              string
	Value              any
}

// Clone clones DeveloperField
func (f DeveloperField) Clone() DeveloperField {
	return f
}

// Component is a way of compressing one or more fields into a bit field expressed in a single containing field.
// The component can be expanded as a main Field in a Message or to update the value of the destination main Field.
type Component struct {
	FieldNum   byte
	Scale      float64
	Offset     float64
	Accumulate bool
	Bits       byte // bit value max 32
}

// SubField is a dynamic interpretation of the main Field in a Message when the SubFieldMap mapping match. See SubFieldMap's docs.
type SubField struct {
	Name       string
	Type       profile.ProfileType
	Scale      float64
	Offset     float64
	Units      string
	Maps       []SubFieldMap
	Components []Component
}

// Clone clones SubField
func (s SubField) Clone() SubField {
	s.Components = slices.Clone(s.Components)
	s.Maps = slices.Clone(s.Maps)
	return s
}

// SubFieldMap is the mapping between SubField and and the corresponding main Field in a Message.
// When any Field in a Message has Field.Num == RefFieldNum and Field.Value == RefFieldValue, then the SubField containing
// this mapping can be interpreted as the main Field's properties (name, scale, type etc.)
type SubFieldMap struct {
	RefFieldNum   byte
	RefFieldValue int64
}