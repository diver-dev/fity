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

// TimestampCorrelation is a TimestampCorrelation message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TimestampCorrelation struct {
	Timestamp                 time.Time // Units: s; Whole second part of UTC timestamp at the time the system timestamp was recorded.
	SystemTimestamp           time.Time // Units: s; Whole second part of the system timestamp
	LocalTimestamp            time.Time // Units: s; timestamp epoch expressed in local time used to convert timestamps to local time
	FractionalTimestamp       uint16    // Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
	FractionalSystemTimestamp uint16    // Scale: 32768; Units: s; Fractional part of the system timestamp
	TimestampMs               uint16    // Units: ms; Millisecond part of the UTC timestamp at the time the system timestamp was recorded.
	SystemTimestampMs         uint16    // Units: ms; Millisecond part of the system timestamp

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewTimestampCorrelation creates new TimestampCorrelation struct based on given mesg.
// If mesg is nil, it will return TimestampCorrelation with all fields being set to its corresponding invalid value.
func NewTimestampCorrelation(mesg *proto.Message) *TimestampCorrelation {
	m := new(TimestampCorrelation)
	m.Reset(mesg)
	return m
}

// Reset resets all TimestampCorrelation's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *TimestampCorrelation) Reset(mesg *proto.Message) {
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

	*m = TimestampCorrelation{
		Timestamp:                 datetime.ToTime(vals[253].Uint32()),
		FractionalTimestamp:       vals[0].Uint16(),
		SystemTimestamp:           datetime.ToTime(vals[1].Uint32()),
		FractionalSystemTimestamp: vals[2].Uint16(),
		LocalTimestamp:            datetime.ToTime(vals[3].Uint32()),
		TimestampMs:               vals[4].Uint16(),
		SystemTimestampMs:         vals[5].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts TimestampCorrelation into proto.Message. If options is nil, default options will be used.
func (m *TimestampCorrelation) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumTimestampCorrelation}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.FractionalTimestamp)
		fields = append(fields, field)
	}
	if !m.SystemTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(uint32(m.SystemTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.FractionalSystemTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.FractionalSystemTimestamp)
		fields = append(fields, field)
	}
	if !m.LocalTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(uint32(m.LocalTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.TimestampMs)
		fields = append(fields, field)
	}
	if m.SystemTimestampMs != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.SystemTimestampMs)
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
func (m *TimestampCorrelation) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SystemTimestampUint32 returns SystemTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TimestampCorrelation) SystemTimestampUint32() uint32 {
	return datetime.ToUint32(m.SystemTimestamp)
}

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TimestampCorrelation) LocalTimestampUint32() uint32 {
	return datetime.ToUint32(m.LocalTimestamp)
}

// FractionalTimestampScaled return FractionalTimestamp in its scaled value.
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
//
// Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.FractionalTimestamp)/32768 - 0
}

// FractionalSystemTimestampScaled return FractionalSystemTimestamp in its scaled value.
// If FractionalSystemTimestamp value is invalid, float64 invalid value will be returned.
//
// Scale: 32768; Units: s; Fractional part of the system timestamp
func (m *TimestampCorrelation) FractionalSystemTimestampScaled() float64 {
	if m.FractionalSystemTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.FractionalSystemTimestamp)/32768 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s; Whole second part of UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetTimestamp(v time.Time) *TimestampCorrelation {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets FractionalTimestamp value.
//
// Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetFractionalTimestamp(v uint16) *TimestampCorrelation {
	m.FractionalTimestamp = v
	return m
}

// SetFractionalTimestampScaled is similar to SetFractionalTimestamp except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 32768; Units: s; Fractional part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetFractionalTimestampScaled(v float64) *TimestampCorrelation {
	unscaled := (v + 0) * 32768
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.FractionalTimestamp = uint16(basetype.Uint16Invalid)
		return m
	}
	m.FractionalTimestamp = uint16(unscaled)
	return m
}

// SetSystemTimestamp sets SystemTimestamp value.
//
// Units: s; Whole second part of the system timestamp
func (m *TimestampCorrelation) SetSystemTimestamp(v time.Time) *TimestampCorrelation {
	m.SystemTimestamp = v
	return m
}

// SetFractionalSystemTimestamp sets FractionalSystemTimestamp value.
//
// Scale: 32768; Units: s; Fractional part of the system timestamp
func (m *TimestampCorrelation) SetFractionalSystemTimestamp(v uint16) *TimestampCorrelation {
	m.FractionalSystemTimestamp = v
	return m
}

// SetFractionalSystemTimestampScaled is similar to SetFractionalSystemTimestamp except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 32768; Units: s; Fractional part of the system timestamp
func (m *TimestampCorrelation) SetFractionalSystemTimestampScaled(v float64) *TimestampCorrelation {
	unscaled := (v + 0) * 32768
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.FractionalSystemTimestamp = uint16(basetype.Uint16Invalid)
		return m
	}
	m.FractionalSystemTimestamp = uint16(unscaled)
	return m
}

// SetLocalTimestamp sets LocalTimestamp value.
//
// Units: s; timestamp epoch expressed in local time used to convert timestamps to local time
func (m *TimestampCorrelation) SetLocalTimestamp(v time.Time) *TimestampCorrelation {
	m.LocalTimestamp = v
	return m
}

// SetTimestampMs sets TimestampMs value.
//
// Units: ms; Millisecond part of the UTC timestamp at the time the system timestamp was recorded.
func (m *TimestampCorrelation) SetTimestampMs(v uint16) *TimestampCorrelation {
	m.TimestampMs = v
	return m
}

// SetSystemTimestampMs sets SystemTimestampMs value.
//
// Units: ms; Millisecond part of the system timestamp
func (m *TimestampCorrelation) SetSystemTimestampMs(v uint16) *TimestampCorrelation {
	m.SystemTimestampMs = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *TimestampCorrelation) SetUnknownFields(unknownFields ...proto.Field) *TimestampCorrelation {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *TimestampCorrelation) SetDeveloperFields(developerFields ...proto.DeveloperField) *TimestampCorrelation {
	m.DeveloperFields = developerFields
	return m
}
