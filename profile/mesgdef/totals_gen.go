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
	"time"
)

// Totals is a Totals message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Totals struct {
	Timestamp    time.Time // Units: s
	TimerTime    uint32    // Units: s; Excludes pauses
	Distance     uint32    // Units: m
	Calories     uint32    // Units: kcal
	ElapsedTime  uint32    // Units: s; Includes pauses
	ActiveTime   uint32    // Units: s
	MessageIndex typedef.MessageIndex
	Sessions     uint16
	Sport        typedef.Sport
	SportIndex   uint8

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewTotals creates new Totals struct based on given mesg.
// If mesg is nil, it will return Totals with all fields being set to its corresponding invalid value.
func NewTotals(mesg *proto.Message) *Totals {
	vals := [255]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 254 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Totals{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Timestamp:    datetime.ToTime(vals[253].Uint32()),
		TimerTime:    vals[0].Uint32(),
		Distance:     vals[1].Uint32(),
		Calories:     vals[2].Uint32(),
		Sport:        typedef.Sport(vals[3].Uint8()),
		ElapsedTime:  vals[4].Uint32(),
		Sessions:     vals[5].Uint16(),
		ActiveTime:   vals[6].Uint32(),
		SportIndex:   vals[9].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Totals into proto.Message. If options is nil, default options will be used.
func (m *Totals) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumTotals}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.TimerTime)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.Calories)
		fields = append(fields, field)
	}
	if m.Sport != typedef.SportInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if m.ElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.ElapsedTime)
		fields = append(fields, field)
	}
	if m.Sessions != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.Sessions)
		fields = append(fields, field)
	}
	if m.ActiveTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.ActiveTime)
		fields = append(fields, field)
	}
	if m.SportIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(m.SportIndex)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Totals) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetMessageIndex sets MessageIndex value.
func (m *Totals) SetMessageIndex(v typedef.MessageIndex) *Totals {
	m.MessageIndex = v
	return m
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *Totals) SetTimestamp(v time.Time) *Totals {
	m.Timestamp = v
	return m
}

// SetTimerTime sets TimerTime value.
//
// Units: s; Excludes pauses
func (m *Totals) SetTimerTime(v uint32) *Totals {
	m.TimerTime = v
	return m
}

// SetDistance sets Distance value.
//
// Units: m
func (m *Totals) SetDistance(v uint32) *Totals {
	m.Distance = v
	return m
}

// SetCalories sets Calories value.
//
// Units: kcal
func (m *Totals) SetCalories(v uint32) *Totals {
	m.Calories = v
	return m
}

// SetSport sets Sport value.
func (m *Totals) SetSport(v typedef.Sport) *Totals {
	m.Sport = v
	return m
}

// SetElapsedTime sets ElapsedTime value.
//
// Units: s; Includes pauses
func (m *Totals) SetElapsedTime(v uint32) *Totals {
	m.ElapsedTime = v
	return m
}

// SetSessions sets Sessions value.
func (m *Totals) SetSessions(v uint16) *Totals {
	m.Sessions = v
	return m
}

// SetActiveTime sets ActiveTime value.
//
// Units: s
func (m *Totals) SetActiveTime(v uint32) *Totals {
	m.ActiveTime = v
	return m
}

// SetSportIndex sets SportIndex value.
func (m *Totals) SetSportIndex(v uint8) *Totals {
	m.SportIndex = v
	return m
}

// SetDeveloperFields Totals's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Totals) SetUnknownFields(unknownFields ...proto.Field) *Totals {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Totals's DeveloperFields.
func (m *Totals) SetDeveloperFields(developerFields ...proto.DeveloperField) *Totals {
	m.DeveloperFields = developerFields
	return m
}
