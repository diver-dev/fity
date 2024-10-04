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
	"time"
)

// AccelerometerData is a AccelerometerData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type AccelerometerData struct {
	Timestamp                  time.Time // Units: s; Whole second part of the timestamp
	SampleTimeOffset           []uint16  // Array: [N]; Units: ms; Each time in the array describes the time at which the accelerometer sample with the corresponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z
	AccelX                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelY                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	AccelZ                     []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedAccelX           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelY           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CalibratedAccelZ           []float32 // Array: [N]; Units: g; Calibrated accel reading
	CompressedCalibratedAccelX []int16   // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelY []int16   // Array: [N]; Units: mG; Calibrated accel reading
	CompressedCalibratedAccelZ []int16   // Array: [N]; Units: mG; Calibrated accel reading
	TimestampMs                uint16    // Units: ms; Millisecond part of the timestamp.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewAccelerometerData creates new AccelerometerData struct based on given mesg.
// If mesg is nil, it will return AccelerometerData with all fields being set to its corresponding invalid value.
func NewAccelerometerData(mesg *proto.Message) *AccelerometerData {
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

	return &AccelerometerData{
		Timestamp:                  datetime.ToTime(vals[253].Uint32()),
		TimestampMs:                vals[0].Uint16(),
		SampleTimeOffset:           vals[1].SliceUint16(),
		AccelX:                     vals[2].SliceUint16(),
		AccelY:                     vals[3].SliceUint16(),
		AccelZ:                     vals[4].SliceUint16(),
		CalibratedAccelX:           vals[5].SliceFloat32(),
		CalibratedAccelY:           vals[6].SliceFloat32(),
		CalibratedAccelZ:           vals[7].SliceFloat32(),
		CompressedCalibratedAccelX: vals[8].SliceInt16(),
		CompressedCalibratedAccelY: vals[9].SliceInt16(),
		CompressedCalibratedAccelZ: vals[10].SliceInt16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts AccelerometerData into proto.Message. If options is nil, default options will be used.
func (m *AccelerometerData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumAccelerometerData}

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
	if m.SampleTimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint16(m.SampleTimeOffset)
		fields = append(fields, field)
	}
	if m.AccelX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint16(m.AccelX)
		fields = append(fields, field)
	}
	if m.AccelY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint16(m.AccelY)
		fields = append(fields, field)
	}
	if m.AccelZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint16(m.AccelZ)
		fields = append(fields, field)
	}
	if m.CalibratedAccelX != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceFloat32(m.CalibratedAccelX)
		fields = append(fields, field)
	}
	if m.CalibratedAccelY != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceFloat32(m.CalibratedAccelY)
		fields = append(fields, field)
	}
	if m.CalibratedAccelZ != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceFloat32(m.CalibratedAccelZ)
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelX != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceInt16(m.CompressedCalibratedAccelX)
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelY != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceInt16(m.CompressedCalibratedAccelY)
		fields = append(fields, field)
	}
	if m.CompressedCalibratedAccelZ != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.SliceInt16(m.CompressedCalibratedAccelZ)
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
func (m *AccelerometerData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of the timestamp
func (m *AccelerometerData) SetTimestamp(v time.Time) *AccelerometerData {
	m.Timestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *AccelerometerData) SetTimestampMs(v uint16) *AccelerometerData {
	m.TimestampMs = v
	return m
}

// SetSampleTimeOffset sets SampleTimeOffset value.
//
// Array: [N]; Units: ms; Each time in the array describes the time at which the accelerometer sample with the corresponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in accel_x and accel_y and accel_z
func (m *AccelerometerData) SetSampleTimeOffset(v []uint16) *AccelerometerData {
	m.SampleTimeOffset = v
	return m
}

// SetAccelX sets AccelX value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelX(v []uint16) *AccelerometerData {
	m.AccelX = v
	return m
}

// SetAccelY sets AccelY value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelY(v []uint16) *AccelerometerData {
	m.AccelY = v
	return m
}

// SetAccelZ sets AccelZ value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *AccelerometerData) SetAccelZ(v []uint16) *AccelerometerData {
	m.AccelZ = v
	return m
}

// SetCalibratedAccelX sets CalibratedAccelX value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelX(v []float32) *AccelerometerData {
	m.CalibratedAccelX = v
	return m
}

// SetCalibratedAccelY sets CalibratedAccelY value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelY(v []float32) *AccelerometerData {
	m.CalibratedAccelY = v
	return m
}

// SetCalibratedAccelZ sets CalibratedAccelZ value.
//
// Array: [N]; Units: g; Calibrated accel reading
func (m *AccelerometerData) SetCalibratedAccelZ(v []float32) *AccelerometerData {
	m.CalibratedAccelZ = v
	return m
}

// SetCompressedCalibratedAccelX sets CompressedCalibratedAccelX value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelX(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelX = v
	return m
}

// SetCompressedCalibratedAccelY sets CompressedCalibratedAccelY value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelY(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelY = v
	return m
}

// SetCompressedCalibratedAccelZ sets CompressedCalibratedAccelZ value.
//
// Array: [N]; Units: mG; Calibrated accel reading
func (m *AccelerometerData) SetCompressedCalibratedAccelZ(v []int16) *AccelerometerData {
	m.CompressedCalibratedAccelZ = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *AccelerometerData) SetUnknownFields(unknownFields ...proto.Field) *AccelerometerData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *AccelerometerData) SetDeveloperFields(developerFields ...proto.DeveloperField) *AccelerometerData {
	m.DeveloperFields = developerFields
	return m
}
