// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaHeartRateData is a HsaHeartRateData message.
type HsaHeartRateData struct {
	Timestamp          time.Time // Units: s
	HeartRate          []uint8   // Array: [N]; Units: bpm; Beats / min
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds
	Status             uint8     // Status of measurements in buffer - 0 indicates SEARCHING 1 indicates LOCKED

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaHeartRateData creates new HsaHeartRateData struct based on given mesg.
// If mesg is nil, it will return HsaHeartRateData with all fields being set to its corresponding invalid value.
func NewHsaHeartRateData(mesg *proto.Message) *HsaHeartRateData {
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

	return &HsaHeartRateData{
		Timestamp:          datetime.ToTime(vals[253]),
		HeartRate:          typeconv.ToSliceUint8[uint8](vals[2]),
		ProcessingInterval: typeconv.ToUint16[uint16](vals[0]),
		Status:             typeconv.ToUint8[uint8](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaHeartRateData into proto.Message. If options is nil, default options will be used.
func (m *HsaHeartRateData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumHsaHeartRateData)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.HeartRate != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.HeartRate
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.ProcessingInterval
		fields = append(fields, field)
	}
	if m.Status != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Status
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets HsaHeartRateData value.
//
// Units: s
func (m *HsaHeartRateData) SetTimestamp(v time.Time) *HsaHeartRateData {
	m.Timestamp = v
	return m
}

// SetHeartRate sets HsaHeartRateData value.
//
// Array: [N]; Units: bpm; Beats / min
func (m *HsaHeartRateData) SetHeartRate(v []uint8) *HsaHeartRateData {
	m.HeartRate = v
	return m
}

// SetProcessingInterval sets HsaHeartRateData value.
//
// Units: s; Processing interval length in seconds
func (m *HsaHeartRateData) SetProcessingInterval(v uint16) *HsaHeartRateData {
	m.ProcessingInterval = v
	return m
}

// SetStatus sets HsaHeartRateData value.
//
// Status of measurements in buffer - 0 indicates SEARCHING 1 indicates LOCKED
func (m *HsaHeartRateData) SetStatus(v uint8) *HsaHeartRateData {
	m.Status = v
	return m
}

// SetDeveloperFields HsaHeartRateData's DeveloperFields.
func (m *HsaHeartRateData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaHeartRateData {
	m.DeveloperFields = developerFields
	return m
}