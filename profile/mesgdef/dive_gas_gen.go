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

// DiveGas is a DiveGas message.
type DiveGas struct {
	MessageIndex  typedef.MessageIndex
	HeliumContent uint8 // Units: percent
	OxygenContent uint8 // Units: percent
	Status        typedef.DiveGasStatus
	Mode          typedef.DiveGasMode

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveGas creates new DiveGas struct based on given mesg.
// If mesg is nil, it will return DiveGas with all fields being set to its corresponding invalid value.
func NewDiveGas(mesg *proto.Message) *DiveGas {
	vals := [255]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &DiveGas{
		MessageIndex:  typedef.MessageIndex(vals[254].Uint16()),
		HeliumContent: vals[0].Uint8(),
		OxygenContent: vals[1].Uint8(),
		Status:        typedef.DiveGasStatus(vals[2].Uint8()),
		Mode:          typedef.DiveGasMode(vals[3].Uint8()),

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

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumDiveGas}

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
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
	if byte(m.Status) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Status))
		fields = append(fields, field)
	}
	if byte(m.Mode) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Mode))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetMessageIndex sets DiveGas value.
func (m *DiveGas) SetMessageIndex(v typedef.MessageIndex) *DiveGas {
	m.MessageIndex = v
	return m
}

// SetHeliumContent sets DiveGas value.
//
// Units: percent
func (m *DiveGas) SetHeliumContent(v uint8) *DiveGas {
	m.HeliumContent = v
	return m
}

// SetOxygenContent sets DiveGas value.
//
// Units: percent
func (m *DiveGas) SetOxygenContent(v uint8) *DiveGas {
	m.OxygenContent = v
	return m
}

// SetStatus sets DiveGas value.
func (m *DiveGas) SetStatus(v typedef.DiveGasStatus) *DiveGas {
	m.Status = v
	return m
}

// SetMode sets DiveGas value.
func (m *DiveGas) SetMode(v typedef.DiveGasMode) *DiveGas {
	m.Mode = v
	return m
}

// SetDeveloperFields DiveGas's DeveloperFields.
func (m *DiveGas) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveGas {
	m.DeveloperFields = developerFields
	return m
}
