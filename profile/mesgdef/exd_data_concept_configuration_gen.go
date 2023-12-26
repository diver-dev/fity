// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// ExdDataConceptConfiguration is a ExdDataConceptConfiguration message.
type ExdDataConceptConfiguration struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptIndex uint8
	DataPage     uint8
	ConceptKey   uint8
	Scaling      uint8
	DataUnits    typedef.ExdDataUnits
	Qualifier    typedef.ExdQualifiers
	Descriptor   typedef.ExdDescriptors
	IsSigned     bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField

	IsExpandedFields [4]bool // Used for tracking expanded fields, field.Num as index.
}

// NewExdDataConceptConfiguration creates new ExdDataConceptConfiguration struct based on given mesg.
// If mesg is nil, it will return ExdDataConceptConfiguration with all fields being set to its corresponding invalid value.
func NewExdDataConceptConfiguration(mesg *proto.Message) *ExdDataConceptConfiguration {
	vals := [12]any{}
	isExpandedFields := [4]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &ExdDataConceptConfiguration{
		ScreenIndex:  typeconv.ToUint8[uint8](vals[0]),
		ConceptField: typeconv.ToByte[byte](vals[1]),
		FieldId:      typeconv.ToUint8[uint8](vals[2]),
		ConceptIndex: typeconv.ToUint8[uint8](vals[3]),
		DataPage:     typeconv.ToUint8[uint8](vals[4]),
		ConceptKey:   typeconv.ToUint8[uint8](vals[5]),
		Scaling:      typeconv.ToUint8[uint8](vals[6]),
		DataUnits:    typeconv.ToEnum[typedef.ExdDataUnits](vals[8]),
		Qualifier:    typeconv.ToEnum[typedef.ExdQualifiers](vals[9]),
		Descriptor:   typeconv.ToEnum[typedef.ExdDescriptors](vals[10]),
		IsSigned:     typeconv.ToBool[bool](vals[11]),

		DeveloperFields: developerFields,

		IsExpandedFields: isExpandedFields,
	}
}

// ToMesg converts ExdDataConceptConfiguration into proto.Message.
func (m *ExdDataConceptConfiguration) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumExdDataConceptConfiguration)

	if m.ScreenIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ScreenIndex
		fields = append(fields, field)
	}
	if m.ConceptField != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.ConceptField
		fields = append(fields, field)
	}
	if m.FieldId != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.FieldId
		field.IsExpandedField = m.IsExpandedFields[2]
		fields = append(fields, field)
	}
	if m.ConceptIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ConceptIndex
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if m.DataPage != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.DataPage
		fields = append(fields, field)
	}
	if m.ConceptKey != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.ConceptKey
		fields = append(fields, field)
	}
	if m.Scaling != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.Scaling
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.DataUnits) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = typeconv.ToEnum[byte](m.DataUnits)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.Qualifier) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToEnum[byte](m.Qualifier)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.Descriptor) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToEnum[byte](m.Descriptor)
		fields = append(fields, field)
	}
	if m.IsSigned != false {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.IsSigned
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetScreenIndex sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetScreenIndex(v uint8) *ExdDataConceptConfiguration {
	m.ScreenIndex = v
	return m
}

// SetConceptField sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptField(v byte) *ExdDataConceptConfiguration {
	m.ConceptField = v
	return m
}

// SetFieldId sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetFieldId(v uint8) *ExdDataConceptConfiguration {
	m.FieldId = v
	return m
}

// SetConceptIndex sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptIndex(v uint8) *ExdDataConceptConfiguration {
	m.ConceptIndex = v
	return m
}

// SetDataPage sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDataPage(v uint8) *ExdDataConceptConfiguration {
	m.DataPage = v
	return m
}

// SetConceptKey sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptKey(v uint8) *ExdDataConceptConfiguration {
	m.ConceptKey = v
	return m
}

// SetScaling sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetScaling(v uint8) *ExdDataConceptConfiguration {
	m.Scaling = v
	return m
}

// SetDataUnits sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDataUnits(v typedef.ExdDataUnits) *ExdDataConceptConfiguration {
	m.DataUnits = v
	return m
}

// SetQualifier sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetQualifier(v typedef.ExdQualifiers) *ExdDataConceptConfiguration {
	m.Qualifier = v
	return m
}

// SetDescriptor sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDescriptor(v typedef.ExdDescriptors) *ExdDataConceptConfiguration {
	m.Descriptor = v
	return m
}

// SetIsSigned sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetIsSigned(v bool) *ExdDataConceptConfiguration {
	m.IsSigned = v
	return m
}

// SetDeveloperFields ExdDataConceptConfiguration's DeveloperFields.
func (m *ExdDataConceptConfiguration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExdDataConceptConfiguration {
	m.DeveloperFields = developerFields
	return m
}
