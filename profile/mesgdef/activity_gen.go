// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Activity is a Activity message.
type Activity struct {
	Timestamp      time.Time
	LocalTimestamp time.Time // timestamp epoch expressed in local time, used to convert activity timestamps to local time
	TotalTimerTime uint32    // Scale: 1000; Units: s; Exclude pauses
	NumSessions    uint16
	Type           typedef.Activity
	Event          typedef.Event
	EventType      typedef.EventType
	EventGroup     uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewActivity creates new Activity struct based on given mesg.
// If mesg is nil, it will return Activity with all fields being set to its corresponding invalid value.
func NewActivity(mesg *proto.Message) *Activity {
	vals := [254]proto.Value{}

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

	return &Activity{
		Timestamp:      datetime.ToTime(vals[253].Uint32()),
		LocalTimestamp: datetime.ToTime(vals[5].Uint32()),
		TotalTimerTime: vals[0].Uint32(),
		NumSessions:    vals[1].Uint16(),
		Type:           typedef.Activity(vals[2].Uint8()),
		Event:          typedef.Event(vals[3].Uint8()),
		EventType:      typedef.EventType(vals[4].Uint8()),
		EventGroup:     vals[6].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Activity into proto.Message. If options is nil, default options will be used.
func (m *Activity) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumActivity}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.LocalTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint32(datetime.ToUint32(m.LocalTimestamp))
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.TotalTimerTime)
		fields = append(fields, field)
	}
	if m.NumSessions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.NumSessions)
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Type))
		fields = append(fields, field)
	}
	if byte(m.Event) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Event))
		fields = append(fields, field)
	}
	if byte(m.EventType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.EventType))
		fields = append(fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.EventGroup)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Activity) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Activity) LocalTimestampUint32() uint32 { return datetime.ToUint32(m.LocalTimestamp) }

// TotalTimerTimeScaled return TotalTimerTime in its scaled value [Scale: 1000; Units: s; Exclude pauses].
//
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
func (m *Activity) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.TotalTimerTime, 1000, 0)
}

// SetTimestamp sets Activity value.
func (m *Activity) SetTimestamp(v time.Time) *Activity {
	m.Timestamp = v
	return m
}

// SetLocalTimestamp sets Activity value.
//
// timestamp epoch expressed in local time, used to convert activity timestamps to local time
func (m *Activity) SetLocalTimestamp(v time.Time) *Activity {
	m.LocalTimestamp = v
	return m
}

// SetTotalTimerTime sets Activity value.
//
// Scale: 1000; Units: s; Exclude pauses
func (m *Activity) SetTotalTimerTime(v uint32) *Activity {
	m.TotalTimerTime = v
	return m
}

// SetNumSessions sets Activity value.
func (m *Activity) SetNumSessions(v uint16) *Activity {
	m.NumSessions = v
	return m
}

// SetType sets Activity value.
func (m *Activity) SetType(v typedef.Activity) *Activity {
	m.Type = v
	return m
}

// SetEvent sets Activity value.
func (m *Activity) SetEvent(v typedef.Event) *Activity {
	m.Event = v
	return m
}

// SetEventType sets Activity value.
func (m *Activity) SetEventType(v typedef.EventType) *Activity {
	m.EventType = v
	return m
}

// SetEventGroup sets Activity value.
func (m *Activity) SetEventGroup(v uint8) *Activity {
	m.EventGroup = v
	return m
}

// SetDeveloperFields Activity's DeveloperFields.
func (m *Activity) SetDeveloperFields(developerFields ...proto.DeveloperField) *Activity {
	m.DeveloperFields = developerFields
	return m
}
