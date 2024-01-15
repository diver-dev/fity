// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// GyroscopeData is a GyroscopeData message.
type GyroscopeData struct {
	Timestamp        time.Time // Units: s; Whole second part of the timestamp
	SampleTimeOffset []uint16  // Array: [N]; Units: ms; Each time in the array describes the time at which the gyro sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in gyro_x and gyro_y and gyro_z
	GyroX            []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	GyroY            []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	GyroZ            []uint16  // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedGyroX  []float32 // Array: [N]; Units: deg/s; Calibrated gyro reading
	CalibratedGyroY  []float32 // Array: [N]; Units: deg/s; Calibrated gyro reading
	CalibratedGyroZ  []float32 // Array: [N]; Units: deg/s; Calibrated gyro reading
	TimestampMs      uint16    // Units: ms; Millisecond part of the timestamp.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewGyroscopeData creates new GyroscopeData struct based on given mesg.
// If mesg is nil, it will return GyroscopeData with all fields being set to its corresponding invalid value.
func NewGyroscopeData(mesg *proto.Message) *GyroscopeData {
	vals := [254]any{}

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

	return &GyroscopeData{
		Timestamp:        datetime.ToTime(vals[253]),
		SampleTimeOffset: typeconv.ToSliceUint16[uint16](vals[1]),
		GyroX:            typeconv.ToSliceUint16[uint16](vals[2]),
		GyroY:            typeconv.ToSliceUint16[uint16](vals[3]),
		GyroZ:            typeconv.ToSliceUint16[uint16](vals[4]),
		CalibratedGyroX:  typeconv.ToSliceFloat32[float32](vals[5]),
		CalibratedGyroY:  typeconv.ToSliceFloat32[float32](vals[6]),
		CalibratedGyroZ:  typeconv.ToSliceFloat32[float32](vals[7]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts GyroscopeData into proto.Message.
func (m *GyroscopeData) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumGyroscopeData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.SampleTimeOffset != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.SampleTimeOffset
		fields = append(fields, field)
	}
	if m.GyroX != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.GyroX
		fields = append(fields, field)
	}
	if m.GyroY != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.GyroY
		fields = append(fields, field)
	}
	if m.GyroZ != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.GyroZ
		fields = append(fields, field)
	}
	if m.CalibratedGyroX != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.CalibratedGyroX
		fields = append(fields, field)
	}
	if m.CalibratedGyroY != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.CalibratedGyroY
		fields = append(fields, field)
	}
	if m.CalibratedGyroZ != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.CalibratedGyroZ
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.TimestampMs
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets GyroscopeData value.
//
// Units: s; Whole second part of the timestamp
func (m *GyroscopeData) SetTimestamp(v time.Time) *GyroscopeData {
	m.Timestamp = v
	return m
}

// SetSampleTimeOffset sets GyroscopeData value.
//
// Array: [N]; Units: ms; Each time in the array describes the time at which the gyro sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in gyro_x and gyro_y and gyro_z
func (m *GyroscopeData) SetSampleTimeOffset(v []uint16) *GyroscopeData {
	m.SampleTimeOffset = v
	return m
}

// SetGyroX sets GyroscopeData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *GyroscopeData) SetGyroX(v []uint16) *GyroscopeData {
	m.GyroX = v
	return m
}

// SetGyroY sets GyroscopeData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *GyroscopeData) SetGyroY(v []uint16) *GyroscopeData {
	m.GyroY = v
	return m
}

// SetGyroZ sets GyroscopeData value.
//
// Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
func (m *GyroscopeData) SetGyroZ(v []uint16) *GyroscopeData {
	m.GyroZ = v
	return m
}

// SetCalibratedGyroX sets GyroscopeData value.
//
// Array: [N]; Units: deg/s; Calibrated gyro reading
func (m *GyroscopeData) SetCalibratedGyroX(v []float32) *GyroscopeData {
	m.CalibratedGyroX = v
	return m
}

// SetCalibratedGyroY sets GyroscopeData value.
//
// Array: [N]; Units: deg/s; Calibrated gyro reading
func (m *GyroscopeData) SetCalibratedGyroY(v []float32) *GyroscopeData {
	m.CalibratedGyroY = v
	return m
}

// SetCalibratedGyroZ sets GyroscopeData value.
//
// Array: [N]; Units: deg/s; Calibrated gyro reading
func (m *GyroscopeData) SetCalibratedGyroZ(v []float32) *GyroscopeData {
	m.CalibratedGyroZ = v
	return m
}

// SetTimestampMs sets GyroscopeData value.
//
// Units: ms; Millisecond part of the timestamp.
func (m *GyroscopeData) SetTimestampMs(v uint16) *GyroscopeData {
	m.TimestampMs = v
	return m
}

// SetDeveloperFields GyroscopeData's DeveloperFields.
func (m *GyroscopeData) SetDeveloperFields(developerFields ...proto.DeveloperField) *GyroscopeData {
	m.DeveloperFields = developerFields
	return m
}
