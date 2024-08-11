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

// SkinTempOvernight is a SkinTempOvernight message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SkinTempOvernight struct {
	Timestamp            time.Time
	LocalTimestamp       time.Time
	AverageDeviation     float32 // The average overnight deviation from baseline temperature in degrees C
	Average7DayDeviation float32 // The average 7 day overnight deviation from baseline temperature in degrees C
	NightlyValue         float32 // Final overnight temperature value

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSkinTempOvernight creates new SkinTempOvernight struct based on given mesg.
// If mesg is nil, it will return SkinTempOvernight with all fields being set to its corresponding invalid value.
func NewSkinTempOvernight(mesg *proto.Message) *SkinTempOvernight {
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

	return &SkinTempOvernight{
		Timestamp:            datetime.ToTime(vals[253].Uint32()),
		LocalTimestamp:       datetime.ToTime(vals[0].Uint32()),
		AverageDeviation:     vals[1].Float32(),
		Average7DayDeviation: vals[2].Float32(),
		NightlyValue:         vals[4].Float32(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SkinTempOvernight into proto.Message. If options is nil, default options will be used.
func (m *SkinTempOvernight) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumSkinTempOvernight}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(datetime.ToUint32(m.LocalTimestamp))
		fields = append(fields, field)
	}
	if math.Float32bits(m.AverageDeviation) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Float32(m.AverageDeviation)
		fields = append(fields, field)
	}
	if math.Float32bits(m.Average7DayDeviation) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Float32(m.Average7DayDeviation)
		fields = append(fields, field)
	}
	if math.Float32bits(m.NightlyValue) != basetype.Float32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Float32(m.NightlyValue)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *SkinTempOvernight) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *SkinTempOvernight) LocalTimestampUint32() uint32 { return datetime.ToUint32(m.LocalTimestamp) }

// SetTimestamp sets Timestamp value.
func (m *SkinTempOvernight) SetTimestamp(v time.Time) *SkinTempOvernight {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets LocalTimestamp value.
func (m *SkinTempOvernight) SetLocalTimestamp(v time.Time) *SkinTempOvernight {
	m.LocalTimestamp = v
	return m
}

// SetAverageDeviation sets AverageDeviation value.
//
// The average overnight deviation from baseline temperature in degrees C
func (m *SkinTempOvernight) SetAverageDeviation(v float32) *SkinTempOvernight {
	m.AverageDeviation = v
	return m
}

// SetAverage7DayDeviation sets Average7DayDeviation value.
//
// The average 7 day overnight deviation from baseline temperature in degrees C
func (m *SkinTempOvernight) SetAverage7DayDeviation(v float32) *SkinTempOvernight {
	m.Average7DayDeviation = v
	return m
}

// SetNightlyValue sets NightlyValue value.
//
// Final overnight temperature value
func (m *SkinTempOvernight) SetNightlyValue(v float32) *SkinTempOvernight {
	m.NightlyValue = v
	return m
}

// SetDeveloperFields SkinTempOvernight's DeveloperFields.
func (m *SkinTempOvernight) SetDeveloperFields(developerFields ...proto.DeveloperField) *SkinTempOvernight {
	m.DeveloperFields = developerFields
	return m
}
