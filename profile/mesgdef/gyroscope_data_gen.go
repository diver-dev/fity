// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// GyroscopeData is a GyroscopeData message.
type GyroscopeData struct {
	Timestamp        typedef.DateTime // Units: s; Whole second part of the timestamp
	TimestampMs      uint16           // Units: ms; Millisecond part of the timestamp.
	SampleTimeOffset []uint16         // Array: [N]; Units: ms; Each time in the array describes the time at which the gyro sample with the corrosponding index was taken. Limited to 30 samples in each message. The samples may span across seconds. Array size must match the number of samples in gyro_x and gyro_y and gyro_z
	GyroX            []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	GyroY            []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	GyroZ            []uint16         // Array: [N]; Units: counts; These are the raw ADC reading. Maximum number of samples is 30 in each message. The samples may span across seconds. A conversion will need to be done on this data once read.
	CalibratedGyroX  []float32        // Array: [N]; Units: deg/s; Calibrated gyro reading
	CalibratedGyroY  []float32        // Array: [N]; Units: deg/s; Calibrated gyro reading
	CalibratedGyroZ  []float32        // Array: [N]; Units: deg/s; Calibrated gyro reading

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewGyroscopeData creates new GyroscopeData struct based on given mesg. If mesg is nil or mesg.Num is not equal to GyroscopeData mesg number, it will return nil.
func NewGyroscopeData(mesg proto.Message) *GyroscopeData {
	if mesg.Num != typedef.MesgNumGyroscopeData {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* TimestampMs */
		1:   nil,                    /* SampleTimeOffset */
		2:   nil,                    /* GyroX */
		3:   nil,                    /* GyroY */
		4:   nil,                    /* GyroZ */
		5:   nil,                    /* CalibratedGyroX */
		6:   nil,                    /* CalibratedGyroY */
		7:   nil,                    /* CalibratedGyroZ */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &GyroscopeData{
		Timestamp:        typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs:      typeconv.ToUint16[uint16](vals[0]),
		SampleTimeOffset: typeconv.ToSliceUint16[uint16](vals[1]),
		GyroX:            typeconv.ToSliceUint16[uint16](vals[2]),
		GyroY:            typeconv.ToSliceUint16[uint16](vals[3]),
		GyroZ:            typeconv.ToSliceUint16[uint16](vals[4]),
		CalibratedGyroX:  typeconv.ToSliceFloat32[float32](vals[5]),
		CalibratedGyroY:  typeconv.ToSliceFloat32[float32](vals[6]),
		CalibratedGyroZ:  typeconv.ToSliceFloat32[float32](vals[7]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to GyroscopeData mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumGyroscopeData)
func (m GyroscopeData) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumGyroscopeData {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.SampleTimeOffset,
		2:   m.GyroX,
		3:   m.GyroY,
		4:   m.GyroZ,
		5:   m.CalibratedGyroX,
		6:   m.CalibratedGyroY,
		7:   m.CalibratedGyroZ,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
