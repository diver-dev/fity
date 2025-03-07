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

// PowerZone is a PowerZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type PowerZone struct {
	Name         string
	MessageIndex typedef.MessageIndex
	HighValue    uint16 // Units: watts

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewPowerZone creates new PowerZone struct based on given mesg.
// If mesg is nil, it will return PowerZone with all fields being set to its corresponding invalid value.
func NewPowerZone(mesg *proto.Message) *PowerZone {
	m := new(PowerZone)
	m.Reset(mesg)
	return m
}

// Reset resets all PowerZone's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *PowerZone) Reset(mesg *proto.Message) {
	var (
		vals            [255]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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

	*m = PowerZone{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		HighValue:    vals[1].Uint16(),
		Name:         vals[2].String(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts PowerZone into proto.Message. If options is nil, default options will be used.
func (m *PowerZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumPowerZone}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.HighValue != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.HighValue)
		fields = append(fields, field)
	}
	if m.Name != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.String(m.Name)
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
func (m *PowerZone) SetMessageIndex(v typedef.MessageIndex) *PowerZone {
	m.MessageIndex = v
	return m
}

// SetHighValue sets HighValue value.
//
// Units: watts
func (m *PowerZone) SetHighValue(v uint16) *PowerZone {
	m.HighValue = v
	return m
}

// SetName sets Name value.
func (m *PowerZone) SetName(v string) *PowerZone {
	m.Name = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *PowerZone) SetUnknownFields(unknownFields ...proto.Field) *PowerZone {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *PowerZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *PowerZone {
	m.DeveloperFields = developerFields
	return m
}
