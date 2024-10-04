// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// HsaAccelerometerData is a HsaAccelerometerData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaAccelerometerData struct {
	Timestamp        time.Time // Units: s
	AccelX           []int16   // Array: [N]; Scale: 1.024; Units: mG; X-Axis Measurement
	AccelY           []int16   // Array: [N]; Scale: 1.024; Units: mG; Y-Axis Measurement
	AccelZ           []int16   // Array: [N]; Scale: 1.024; Units: mG; Z-Axis Measurement
	Timestamp32K     uint32    // 32 kHz timestamp
	TimestampMs      uint16    // Units: ms; Millisecond resolution of the timestamp
	SamplingInterval uint16    // Units: ms; Sampling Interval in Milliseconds

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHsaAccelerometerData creates new HsaAccelerometerData struct based on given mesg.
// If mesg is nil, it will return HsaAccelerometerData with all fields being set to its corresponding invalid value.
func NewHsaAccelerometerData(mesg *proto.Message) *HsaAccelerometerData {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		clear(arr[:len(unknownFields)])
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &HsaAccelerometerData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		SamplingInterval: vals[1].Uint16(),
		AccelX:           vals[2].SliceInt16(),
		AccelY:           vals[3].SliceInt16(),
		AccelZ:           vals[4].SliceInt16(),
		Timestamp32K:     vals[5].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaAccelerometerData into proto.Message. If options is nil, default options will be used.
func (m *HsaAccelerometerData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHsaAccelerometerData}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SamplingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.SamplingInterval)
		fields = append(fields, field)
	}
	if m.AccelX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceInt16(m.AccelX)
		fields = append(fields, field)
	}
	if m.AccelY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceInt16(m.AccelY)
		fields = append(fields, field)
	}
	if m.AccelZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceInt16(m.AccelZ)
		fields = append(fields, field)
	}
	if m.Timestamp32K != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(m.Timestamp32K)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	clear(fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaAccelerometerData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// AccelXScaled return AccelX in its scaled value.
// If AccelX value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1.024; Units: mG; X-Axis Measurement
func (m *HsaAccelerometerData) AccelXScaled() []float64 {
	if m.AccelX == nil {
		return nil
	}
	var vals = make([]float64, len(m.AccelX))
	for i := range m.AccelX {
		if m.AccelX[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.AccelX[i])/1.024 - 0
	}
	return vals
}

// AccelYScaled return AccelY in its scaled value.
// If AccelY value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1.024; Units: mG; Y-Axis Measurement
func (m *HsaAccelerometerData) AccelYScaled() []float64 {
	if m.AccelY == nil {
		return nil
	}
	var vals = make([]float64, len(m.AccelY))
	for i := range m.AccelY {
		if m.AccelY[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.AccelY[i])/1.024 - 0
	}
	return vals
}

// AccelZScaled return AccelZ in its scaled value.
// If AccelZ value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1.024; Units: mG; Z-Axis Measurement
func (m *HsaAccelerometerData) AccelZScaled() []float64 {
	if m.AccelZ == nil {
		return nil
	}
	var vals = make([]float64, len(m.AccelZ))
	for i := range m.AccelZ {
		if m.AccelZ[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.AccelZ[i])/1.024 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaAccelerometerData) SetTimestamp(v time.Time) *HsaAccelerometerData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond resolution of the timestamp
func (m *HsaAccelerometerData) SetTimestampMs(v uint16) *HsaAccelerometerData {
	m.TimestampMs = v
	return m
}

// SetSamplingInterval sets SamplingInterval value.
//
// Units: ms; Sampling Interval in Milliseconds
func (m *HsaAccelerometerData) SetSamplingInterval(v uint16) *HsaAccelerometerData {
	m.SamplingInterval = v
	return m
}

// SetAccelX sets AccelX value.
//
// Array: [N]; Scale: 1.024; Units: mG; X-Axis Measurement
func (m *HsaAccelerometerData) SetAccelX(v []int16) *HsaAccelerometerData {
	m.AccelX = v
	return m
}

// SetAccelXScaled is similar to SetAccelX except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1.024; Units: mG; X-Axis Measurement
func (m *HsaAccelerometerData) SetAccelXScaled(vs []float64) *HsaAccelerometerData {
	if vs == nil {
		m.AccelX = nil
		return m
	}
	m.AccelX = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1.024
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.AccelX[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.AccelX[i] = int16(unscaled)
	}
	return m
}

// SetAccelY sets AccelY value.
//
// Array: [N]; Scale: 1.024; Units: mG; Y-Axis Measurement
func (m *HsaAccelerometerData) SetAccelY(v []int16) *HsaAccelerometerData {
	m.AccelY = v
	return m
}

// SetAccelYScaled is similar to SetAccelY except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1.024; Units: mG; Y-Axis Measurement
func (m *HsaAccelerometerData) SetAccelYScaled(vs []float64) *HsaAccelerometerData {
	if vs == nil {
		m.AccelY = nil
		return m
	}
	m.AccelY = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1.024
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.AccelY[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.AccelY[i] = int16(unscaled)
	}
	return m
}

// SetAccelZ sets AccelZ value.
//
// Array: [N]; Scale: 1.024; Units: mG; Z-Axis Measurement
func (m *HsaAccelerometerData) SetAccelZ(v []int16) *HsaAccelerometerData {
	m.AccelZ = v
	return m
}

// SetAccelZScaled is similar to SetAccelZ except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1.024; Units: mG; Z-Axis Measurement
func (m *HsaAccelerometerData) SetAccelZScaled(vs []float64) *HsaAccelerometerData {
	if vs == nil {
		m.AccelZ = nil
		return m
	}
	m.AccelZ = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1.024
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.AccelZ[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.AccelZ[i] = int16(unscaled)
	}
	return m
}

// SetTimestamp32K sets Timestamp32K value.
//
// 32 kHz timestamp
func (m *HsaAccelerometerData) SetTimestamp32K(v uint32) *HsaAccelerometerData {
	m.Timestamp32K = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *HsaAccelerometerData) SetUnknownFields(unknownFields ...proto.Field) *HsaAccelerometerData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *HsaAccelerometerData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaAccelerometerData {
	m.DeveloperFields = developerFields
	return m
}
