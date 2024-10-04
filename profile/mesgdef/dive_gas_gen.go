// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// DiveGas is a DiveGas message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type DiveGas struct {
	MessageIndex  typedef.MessageIndex
	HeliumContent uint8 // Units: percent
	OxygenContent uint8 // Units: percent
	Status        typedef.DiveGasStatus
	Mode          typedef.DiveGasMode

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewDiveGas creates new DiveGas struct based on given mesg.
// If mesg is nil, it will return DiveGas with all fields being set to its corresponding invalid value.
func NewDiveGas(mesg *proto.Message) *DiveGas {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &DiveGas{
		MessageIndex:  typedef.MessageIndex(vals[254].Uint16()),
		HeliumContent: vals[0].Uint8(),
		OxygenContent: vals[1].Uint8(),
		Status:        typedef.DiveGasStatus(vals[2].Uint8()),
		Mode:          typedef.DiveGasMode(vals[3].Uint8()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts DiveGas into proto.Message. If options is nil, default options will be used.
func (m *DiveGas) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumDiveGas}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.HeliumContent != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.HeliumContent)
		fields = append(fields, field)
	}
	if m.OxygenContent != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.OxygenContent)
		fields = append(fields, field)
	}
	if m.Status != typedef.DiveGasStatusInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Status))
		fields = append(fields, field)
	}
	if m.Mode != typedef.DiveGasModeInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Mode))
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
func (m *DiveGas) SetMessageIndex(v typedef.MessageIndex) *DiveGas {
	m.MessageIndex = v
	return m
}

// SetHeliumContent sets HeliumContent value.
//
// Units: percent
func (m *DiveGas) SetHeliumContent(v uint8) *DiveGas {
	m.HeliumContent = v
	return m
}

// SetOxygenContent sets OxygenContent value.
//
// Units: percent
func (m *DiveGas) SetOxygenContent(v uint8) *DiveGas {
	m.OxygenContent = v
	return m
}

// SetStatus sets Status value.
func (m *DiveGas) SetStatus(v typedef.DiveGasStatus) *DiveGas {
	m.Status = v
	return m
}

// SetMode sets Mode value.
func (m *DiveGas) SetMode(v typedef.DiveGasMode) *DiveGas {
	m.Mode = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *DiveGas) SetUnknownFields(unknownFields ...proto.Field) *DiveGas {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *DiveGas) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveGas {
	m.DeveloperFields = developerFields
	return m
}
