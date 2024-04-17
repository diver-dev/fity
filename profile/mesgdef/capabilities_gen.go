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
	"unsafe"
)

// Capabilities is a Capabilities message.
type Capabilities struct {
	Languages             []uint8              // Array: [N]; Use language_bits_x types where x is index of array.
	Sports                []typedef.SportBits0 // Array: [N]; Use sport_bits_x types where x is index of array.
	WorkoutsSupported     typedef.WorkoutCapabilities
	ConnectivitySupported typedef.ConnectivityCapabilities

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewCapabilities creates new Capabilities struct based on given mesg.
// If mesg is nil, it will return Capabilities with all fields being set to its corresponding invalid value.
func NewCapabilities(mesg *proto.Message) *Capabilities {
	vals := [24]proto.Value{}

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

	return &Capabilities{
		Languages: vals[0].SliceUint8(),
		Sports: func() []typedef.SportBits0 {
			sliceValue := vals[1].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.SportBits0)(ptr), len(sliceValue))
		}(),
		WorkoutsSupported:     typedef.WorkoutCapabilities(vals[21].Uint32z()),
		ConnectivitySupported: typedef.ConnectivityCapabilities(vals[23].Uint32z()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Capabilities into proto.Message. If options is nil, default options will be used.
func (m *Capabilities) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumCapabilities}

	if m.Languages != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.SliceUint8(m.Languages)
		fields = append(fields, field)
	}
	if m.Sports != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint8(m.Sports)
		fields = append(fields, field)
	}
	if uint32(m.WorkoutsSupported) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint32(uint32(m.WorkoutsSupported))
		fields = append(fields, field)
	}
	if uint32(m.ConnectivitySupported) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint32(uint32(m.ConnectivitySupported))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetLanguages sets Capabilities value.
//
// Array: [N]; Use language_bits_x types where x is index of array.
func (m *Capabilities) SetLanguages(v []uint8) *Capabilities {
	m.Languages = v
	return m
}

// SetSports sets Capabilities value.
//
// Array: [N]; Use sport_bits_x types where x is index of array.
func (m *Capabilities) SetSports(v []typedef.SportBits0) *Capabilities {
	m.Sports = v
	return m
}

// SetWorkoutsSupported sets Capabilities value.
func (m *Capabilities) SetWorkoutsSupported(v typedef.WorkoutCapabilities) *Capabilities {
	m.WorkoutsSupported = v
	return m
}

// SetConnectivitySupported sets Capabilities value.
func (m *Capabilities) SetConnectivitySupported(v typedef.ConnectivityCapabilities) *Capabilities {
	m.ConnectivitySupported = v
	return m
}

// SetDeveloperFields Capabilities's DeveloperFields.
func (m *Capabilities) SetDeveloperFields(developerFields ...proto.DeveloperField) *Capabilities {
	m.DeveloperFields = developerFields
	return m
}
