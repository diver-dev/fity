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

// HsaGyroscopeData is a HsaGyroscopeData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaGyroscopeData struct {
	Timestamp        time.Time // Units: s
	GyroX            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
	GyroY            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
	GyroZ            []int16   // Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
	Timestamp32K     uint32    // Units: 1/32768 s; 32 kHz timestamp
	TimestampMs      uint16    // Units: ms; Millisecond resolution of the timestamp
	SamplingInterval uint16    // Units: 1/32768 s; Sampling Interval in 32 kHz timescale

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHsaGyroscopeData creates new HsaGyroscopeData struct based on given mesg.
// If mesg is nil, it will return HsaGyroscopeData with all fields being set to its corresponding invalid value.
func NewHsaGyroscopeData(mesg *proto.Message) *HsaGyroscopeData {
	m := new(HsaGyroscopeData)
	m.Reset(mesg)
	return m
}

// Reset resets all HsaGyroscopeData's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *HsaGyroscopeData) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = HsaGyroscopeData{
		Timestamp:        datetime.ToTime(vals[253].Uint32()),
		TimestampMs:      vals[0].Uint16(),
		SamplingInterval: vals[1].Uint16(),
		GyroX:            vals[2].SliceInt16(),
		GyroY:            vals[3].SliceInt16(),
		GyroZ:            vals[4].SliceInt16(),
		Timestamp32K:     vals[5].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaGyroscopeData into proto.Message. If options is nil, default options will be used.
func (m *HsaGyroscopeData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHsaGyroscopeData}

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
	if m.GyroX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceInt16(m.GyroX)
		fields = append(fields, field)
	}
	if m.GyroY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceInt16(m.GyroY)
		fields = append(fields, field)
	}
	if m.GyroZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceInt16(m.GyroZ)
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
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaGyroscopeData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// GyroXScaled return GyroX in its scaled value.
// If GyroX value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) GyroXScaled() []float64 {
	if m.GyroX == nil {
		return nil
	}
	var vals = make([]float64, len(m.GyroX))
	for i := range m.GyroX {
		if m.GyroX[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.GyroX[i])/28.57143 - 0
	}
	return vals
}

// GyroYScaled return GyroY in its scaled value.
// If GyroY value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) GyroYScaled() []float64 {
	if m.GyroY == nil {
		return nil
	}
	var vals = make([]float64, len(m.GyroY))
	for i := range m.GyroY {
		if m.GyroY[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.GyroY[i])/28.57143 - 0
	}
	return vals
}

// GyroZScaled return GyroZ in its scaled value.
// If GyroZ value is invalid, nil will be returned.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) GyroZScaled() []float64 {
	if m.GyroZ == nil {
		return nil
	}
	var vals = make([]float64, len(m.GyroZ))
	for i := range m.GyroZ {
		if m.GyroZ[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.GyroZ[i])/28.57143 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaGyroscopeData) SetTimestamp(v time.Time) *HsaGyroscopeData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond resolution of the timestamp
func (m *HsaGyroscopeData) SetTimestampMs(v uint16) *HsaGyroscopeData {
	m.TimestampMs = v
	return m
}

// SetSamplingInterval sets SamplingInterval value.
//
// Units: 1/32768 s; Sampling Interval in 32 kHz timescale
func (m *HsaGyroscopeData) SetSamplingInterval(v uint16) *HsaGyroscopeData {
	m.SamplingInterval = v
	return m
}

// SetGyroX sets GyroX value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) SetGyroX(v []int16) *HsaGyroscopeData {
	m.GyroX = v
	return m
}

// SetGyroXScaled is similar to SetGyroX except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; X-Axis Measurement
func (m *HsaGyroscopeData) SetGyroXScaled(vs []float64) *HsaGyroscopeData {
	if vs == nil {
		m.GyroX = nil
		return m
	}
	m.GyroX = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 28.57143
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.GyroX[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.GyroX[i] = int16(unscaled)
	}
	return m
}

// SetGyroY sets GyroY value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) SetGyroY(v []int16) *HsaGyroscopeData {
	m.GyroY = v
	return m
}

// SetGyroYScaled is similar to SetGyroY except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Y-Axis Measurement
func (m *HsaGyroscopeData) SetGyroYScaled(vs []float64) *HsaGyroscopeData {
	if vs == nil {
		m.GyroY = nil
		return m
	}
	m.GyroY = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 28.57143
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.GyroY[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.GyroY[i] = int16(unscaled)
	}
	return m
}

// SetGyroZ sets GyroZ value.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) SetGyroZ(v []int16) *HsaGyroscopeData {
	m.GyroZ = v
	return m
}

// SetGyroZScaled is similar to SetGyroZ except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 28.57143; Units: deg/s; Z-Axis Measurement
func (m *HsaGyroscopeData) SetGyroZScaled(vs []float64) *HsaGyroscopeData {
	if vs == nil {
		m.GyroZ = nil
		return m
	}
	m.GyroZ = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 28.57143
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.GyroZ[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.GyroZ[i] = int16(unscaled)
	}
	return m
}

// SetTimestamp32K sets Timestamp32K value.
//
// Units: 1/32768 s; 32 kHz timestamp
func (m *HsaGyroscopeData) SetTimestamp32K(v uint32) *HsaGyroscopeData {
	m.Timestamp32K = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *HsaGyroscopeData) SetUnknownFields(unknownFields ...proto.Field) *HsaGyroscopeData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *HsaGyroscopeData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaGyroscopeData {
	m.DeveloperFields = developerFields
	return m
}
