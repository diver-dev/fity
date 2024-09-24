// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// CadenceZone is a CadenceZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type CadenceZone struct {
	Name         string
	MessageIndex typedef.MessageIndex
	HighValue    uint8 // Units: rpm

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewCadenceZone creates new CadenceZone struct based on given mesg.
// If mesg is nil, it will return CadenceZone with all fields being set to its corresponding invalid value.
func NewCadenceZone(mesg *proto.Message) *CadenceZone {
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
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &CadenceZone{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		HighValue:    vals[0].Uint8(),
		Name:         vals[1].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts CadenceZone into proto.Message. If options is nil, default options will be used.
func (m *CadenceZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumCadenceZone}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.HighValue != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.HighValue)
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.String(m.Name)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets MessageIndex value.
func (m *CadenceZone) SetMessageIndex(v typedef.MessageIndex) *CadenceZone {
	m.MessageIndex = v
	return m
}

// SetHighValue sets HighValue value.
//
// Units: rpm
func (m *CadenceZone) SetHighValue(v uint8) *CadenceZone {
	m.HighValue = v
	return m
}

// SetName sets Name value.
func (m *CadenceZone) SetName(v string) *CadenceZone {
	m.Name = v
	return m
}

// SetDeveloperFields CadenceZone's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *CadenceZone) SetUnknownFields(unknownFields ...proto.Field) *CadenceZone {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields CadenceZone's DeveloperFields.
func (m *CadenceZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *CadenceZone {
	m.DeveloperFields = developerFields
	return m
}
