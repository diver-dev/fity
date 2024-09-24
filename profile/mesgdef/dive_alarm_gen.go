// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"unsafe"
)

// DiveAlarm is a DiveAlarm message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type DiveAlarm struct {
	DiveTypes        []typedef.SubSport    // Array: [N]; Dive types the alarm will trigger on
	Depth            uint32                // Scale: 1000; Units: m; Depth setting (m) for depth type alarms
	Time             int32                 // Units: s; Time setting (s) for time type alarms
	Id               uint32                // Alarm ID
	Speed            int32                 // Scale: 1000; Units: mps; Ascent/descent rate (mps) setting for speed type alarms
	MessageIndex     typedef.MessageIndex  // Index of the alarm
	Enabled          typedef.Bool          // Enablement flag
	AlarmType        typedef.DiveAlarmType // Alarm type setting
	Sound            typedef.Tone          // Tone and Vibe setting for the alarm
	PopupEnabled     typedef.Bool          // Show a visible pop-up for this alarm
	TriggerOnDescent typedef.Bool          // Trigger the alarm on descent
	TriggerOnAscent  typedef.Bool          // Trigger the alarm on ascent
	Repeating        typedef.Bool          // Repeat alarm each time threshold is crossed?

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewDiveAlarm creates new DiveAlarm struct based on given mesg.
// If mesg is nil, it will return DiveAlarm with all fields being set to its corresponding invalid value.
func NewDiveAlarm(mesg *proto.Message) *DiveAlarm {
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

	return &DiveAlarm{
		MessageIndex: typedef.MessageIndex(vals[254].Uint16()),
		Depth:        vals[0].Uint32(),
		Time:         vals[1].Int32(),
		Enabled:      vals[2].Bool(),
		AlarmType:    typedef.DiveAlarmType(vals[3].Uint8()),
		Sound:        typedef.Tone(vals[4].Uint8()),
		DiveTypes: func() []typedef.SubSport {
			sliceValue := vals[5].SliceUint8()
			ptr := unsafe.SliceData(sliceValue)
			return unsafe.Slice((*typedef.SubSport)(ptr), len(sliceValue))
		}(),
		Id:               vals[6].Uint32(),
		PopupEnabled:     vals[7].Bool(),
		TriggerOnDescent: vals[8].Bool(),
		TriggerOnAscent:  vals[9].Bool(),
		Repeating:        vals[10].Bool(),
		Speed:            vals[11].Int32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts DiveAlarm into proto.Message. If options is nil, default options will be used.
func (m *DiveAlarm) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumDiveAlarm}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.Depth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(m.Depth)
		fields = append(fields, field)
	}
	if m.Time != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int32(m.Time)
		fields = append(fields, field)
	}
	if m.Enabled < 2 {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Bool(m.Enabled)
		fields = append(fields, field)
	}
	if m.AlarmType != typedef.DiveAlarmTypeInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(byte(m.AlarmType))
		fields = append(fields, field)
	}
	if m.Sound != typedef.ToneInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.Sound))
		fields = append(fields, field)
	}
	if m.DiveTypes != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint8(m.DiveTypes)
		fields = append(fields, field)
	}
	if m.Id != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint32(m.Id)
		fields = append(fields, field)
	}
	if m.PopupEnabled < 2 {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Bool(m.PopupEnabled)
		fields = append(fields, field)
	}
	if m.TriggerOnDescent < 2 {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Bool(m.TriggerOnDescent)
		fields = append(fields, field)
	}
	if m.TriggerOnAscent < 2 {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Bool(m.TriggerOnAscent)
		fields = append(fields, field)
	}
	if m.Repeating < 2 {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Bool(m.Repeating)
		fields = append(fields, field)
	}
	if m.Speed != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Int32(m.Speed)
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

// DepthScaled return Depth in its scaled value.
// If Depth value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m; Depth setting (m) for depth type alarms
func (m *DiveAlarm) DepthScaled() float64 {
	if m.Depth == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Depth)/1000 - 0
}

// SpeedScaled return Speed in its scaled value.
// If Speed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: mps; Ascent/descent rate (mps) setting for speed type alarms
func (m *DiveAlarm) SpeedScaled() float64 {
	if m.Speed == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Speed)/1000 - 0
}

// SetMessageIndex sets MessageIndex value.
//
// Index of the alarm
func (m *DiveAlarm) SetMessageIndex(v typedef.MessageIndex) *DiveAlarm {
	m.MessageIndex = v
	return m
}

// SetDepth sets Depth value.
//
// Scale: 1000; Units: m; Depth setting (m) for depth type alarms
func (m *DiveAlarm) SetDepth(v uint32) *DiveAlarm {
	m.Depth = v
	return m
}

// SetDepthScaled is similar to SetDepth except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m; Depth setting (m) for depth type alarms
func (m *DiveAlarm) SetDepthScaled(v float64) *DiveAlarm {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Depth = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Depth = uint32(unscaled)
	return m
}

// SetTime sets Time value.
//
// Units: s; Time setting (s) for time type alarms
func (m *DiveAlarm) SetTime(v int32) *DiveAlarm {
	m.Time = v
	return m
}

// SetEnabled sets Enabled value.
//
// Enablement flag
func (m *DiveAlarm) SetEnabled(v typedef.Bool) *DiveAlarm {
	m.Enabled = v
	return m
}

// SetAlarmType sets AlarmType value.
//
// Alarm type setting
func (m *DiveAlarm) SetAlarmType(v typedef.DiveAlarmType) *DiveAlarm {
	m.AlarmType = v
	return m
}

// SetSound sets Sound value.
//
// Tone and Vibe setting for the alarm
func (m *DiveAlarm) SetSound(v typedef.Tone) *DiveAlarm {
	m.Sound = v
	return m
}

// SetDiveTypes sets DiveTypes value.
//
// Array: [N]; Dive types the alarm will trigger on
func (m *DiveAlarm) SetDiveTypes(v []typedef.SubSport) *DiveAlarm {
	m.DiveTypes = v
	return m
}

// SetId sets Id value.
//
// Alarm ID
func (m *DiveAlarm) SetId(v uint32) *DiveAlarm {
	m.Id = v
	return m
}

// SetPopupEnabled sets PopupEnabled value.
//
// Show a visible pop-up for this alarm
func (m *DiveAlarm) SetPopupEnabled(v typedef.Bool) *DiveAlarm {
	m.PopupEnabled = v
	return m
}

// SetTriggerOnDescent sets TriggerOnDescent value.
//
// Trigger the alarm on descent
func (m *DiveAlarm) SetTriggerOnDescent(v typedef.Bool) *DiveAlarm {
	m.TriggerOnDescent = v
	return m
}

// SetTriggerOnAscent sets TriggerOnAscent value.
//
// Trigger the alarm on ascent
func (m *DiveAlarm) SetTriggerOnAscent(v typedef.Bool) *DiveAlarm {
	m.TriggerOnAscent = v
	return m
}

// SetRepeating sets Repeating value.
//
// Repeat alarm each time threshold is crossed?
func (m *DiveAlarm) SetRepeating(v typedef.Bool) *DiveAlarm {
	m.Repeating = v
	return m
}

// SetSpeed sets Speed value.
//
// Scale: 1000; Units: mps; Ascent/descent rate (mps) setting for speed type alarms
func (m *DiveAlarm) SetSpeed(v int32) *DiveAlarm {
	m.Speed = v
	return m
}

// SetSpeedScaled is similar to SetSpeed except it accepts a scaled value.
// This method automatically converts the given value to its int32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: mps; Ascent/descent rate (mps) setting for speed type alarms
func (m *DiveAlarm) SetSpeedScaled(v float64) *DiveAlarm {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint32Invalid) {
		m.Speed = int32(basetype.Sint32Invalid)
		return m
	}
	m.Speed = int32(unscaled)
	return m
}

// SetDeveloperFields DiveAlarm's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *DiveAlarm) SetUnknownFields(unknownFields ...proto.Field) *DiveAlarm {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields DiveAlarm's DeveloperFields.
func (m *DiveAlarm) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveAlarm {
	m.DeveloperFields = developerFields
	return m
}
