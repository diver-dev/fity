// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// HsaWristTemperatureData is a HsaWristTemperatureData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaWristTemperatureData struct {
	Timestamp          time.Time // Units: s
	Value              []uint16  // Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaWristTemperatureData creates new HsaWristTemperatureData struct based on given mesg.
// If mesg is nil, it will return HsaWristTemperatureData with all fields being set to its corresponding invalid value.
func NewHsaWristTemperatureData(mesg *proto.Message) *HsaWristTemperatureData {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &HsaWristTemperatureData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		ProcessingInterval: vals[0].Uint16(),
		Value:              vals[1].SliceUint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaWristTemperatureData into proto.Message. If options is nil, default options will be used.
func (m *HsaWristTemperatureData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaWristTemperatureData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}
	if m.Value != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.Value)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaWristTemperatureData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// ValueScaled return Value in its scaled value.
// If Value value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
func (m *HsaWristTemperatureData) ValueScaled() []float64 {
	if m.Value == nil {
		return nil
	}
	var vals = make([]float64, len(m.Value))
	for i := range m.Value {
		if m.Value[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.Value[i])/1000 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaWristTemperatureData) SetTimestamp(v time.Time) *HsaWristTemperatureData {
	m.Timestamp = v
	return m
}

// SetProcessingInterval sets ProcessingInterval value.
//
// Units: s; Processing interval length in seconds
func (m *HsaWristTemperatureData) SetProcessingInterval(v uint16) *HsaWristTemperatureData {
	m.ProcessingInterval = v
	return m
}

// SetValue sets Value value.
//
// Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
func (m *HsaWristTemperatureData) SetValue(v []uint16) *HsaWristTemperatureData {
	m.Value = v
	return m
}

// SetValueScaled is similar to SetValue except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: degC; Wrist temperature reading
func (m *HsaWristTemperatureData) SetValueScaled(vs []float64) *HsaWristTemperatureData {
	if vs == nil {
		m.Value = nil
		return m
	}
	m.Value = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.Value[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.Value[i] = uint16(unscaled)
	}
	return m
}

// SetDeveloperFields HsaWristTemperatureData's DeveloperFields.
func (m *HsaWristTemperatureData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaWristTemperatureData {
	m.DeveloperFields = developerFields
	return m
}
