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
	"math"
)

// Hrv is a Hrv message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Hrv struct {
	Time []uint16 // Array: [N]; Scale: 1000; Units: s; Time between beats

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrv creates new Hrv struct based on given mesg.
// If mesg is nil, it will return Hrv with all fields being set to its corresponding invalid value.
func NewHrv(mesg *proto.Message) *Hrv {
	vals := [1]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 0 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Hrv{
		Time: vals[0].SliceUint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Hrv into proto.Message. If options is nil, default options will be used.
func (m *Hrv) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHrv}

	if m.Time != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.SliceUint16(m.Time)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimeScaled return Time in its scaled value.
// If Time value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s; Time between beats
func (m *Hrv) TimeScaled() []float64 {
	if m.Time == nil {
		return nil
	}
	var vals = make([]float64, len(m.Time))
	for i := range m.Time {
		if m.Time[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Time[i])/1000 - 0
	}
	return vals
}

// SetTime sets Time value.
//
// Array: [N]; Scale: 1000; Units: s; Time between beats
func (m *Hrv) SetTime(v []uint16) *Hrv {
	m.Time = v
	return m
}

// SetTimeScaled is similar to SetTime except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s; Time between beats
func (m *Hrv) SetTimeScaled(vs []float64) *Hrv {
	if vs == nil {
		m.Time = nil
		return m
	}
	m.Time = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.Time[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.Time[i] = uint16(unscaled)
	}
	return m
}

// SetDeveloperFields Hrv's DeveloperFields.
func (m *Hrv) SetDeveloperFields(developerFields ...proto.DeveloperField) *Hrv {
	m.DeveloperFields = developerFields
	return m
}
