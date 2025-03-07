// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Split is a Split message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Split struct {
	StartTime         time.Time
	EndTime           time.Time
	TotalElapsedTime  uint32 // Scale: 1000; Units: s
	TotalTimerTime    uint32 // Scale: 1000; Units: s
	TotalDistance     uint32 // Scale: 100; Units: m
	AvgSpeed          uint32 // Scale: 1000; Units: m/s
	StartPositionLat  int32  // Units: semicircles
	StartPositionLong int32  // Units: semicircles
	EndPositionLat    int32  // Units: semicircles
	EndPositionLong   int32  // Units: semicircles
	MaxSpeed          uint32 // Scale: 1000; Units: m/s
	AvgVertSpeed      int32  // Scale: 1000; Units: m/s
	TotalCalories     uint32 // Units: kcal
	StartElevation    uint32 // Scale: 5; Offset: 500; Units: m
	TotalMovingTime   uint32 // Scale: 1000; Units: s
	MessageIndex      typedef.MessageIndex
	TotalAscent       uint16 // Units: m
	TotalDescent      uint16 // Units: m
	SplitType         typedef.SplitType

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSplit creates new Split struct based on given mesg.
// If mesg is nil, it will return Split with all fields being set to its corresponding invalid value.
func NewSplit(mesg *proto.Message) *Split {
	m := new(Split)
	m.Reset(mesg)
	return m
}

// Reset resets all Split's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *Split) Reset(mesg *proto.Message) {
	var (
		vals            [255]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = Split{
		MessageIndex:      typedef.MessageIndex(vals[254].Uint16()),
		SplitType:         typedef.SplitType(vals[0].Uint8()),
		TotalElapsedTime:  vals[1].Uint32(),
		TotalTimerTime:    vals[2].Uint32(),
		TotalDistance:     vals[3].Uint32(),
		AvgSpeed:          vals[4].Uint32(),
		StartTime:         datetime.ToTime(vals[9].Uint32()),
		TotalAscent:       vals[13].Uint16(),
		TotalDescent:      vals[14].Uint16(),
		StartPositionLat:  vals[21].Int32(),
		StartPositionLong: vals[22].Int32(),
		EndPositionLat:    vals[23].Int32(),
		EndPositionLong:   vals[24].Int32(),
		MaxSpeed:          vals[25].Uint32(),
		AvgVertSpeed:      vals[26].Int32(),
		EndTime:           datetime.ToTime(vals[27].Uint32()),
		TotalCalories:     vals[28].Uint32(),
		StartElevation:    vals[74].Uint32(),
		TotalMovingTime:   vals[110].Uint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Split into proto.Message. If options is nil, default options will be used.
func (m *Split) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSplit}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.SplitType != typedef.SplitTypeInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.SplitType))
		fields = append(fields, field)
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint32(m.TotalElapsedTime)
		fields = append(fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.TotalTimerTime)
		fields = append(fields, field)
	}
	if m.TotalDistance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.TotalDistance)
		fields = append(fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.AvgSpeed)
		fields = append(fields, field)
	}
	if !m.StartTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint32(uint32(m.StartTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint16(m.TotalAscent)
		fields = append(fields, field)
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint16(m.TotalDescent)
		fields = append(fields, field)
	}
	if m.StartPositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Int32(m.StartPositionLat)
		fields = append(fields, field)
	}
	if m.StartPositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Int32(m.StartPositionLong)
		fields = append(fields, field)
	}
	if m.EndPositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Int32(m.EndPositionLat)
		fields = append(fields, field)
	}
	if m.EndPositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Int32(m.EndPositionLong)
		fields = append(fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint32(m.MaxSpeed)
		fields = append(fields, field)
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 26)
		field.Value = proto.Int32(m.AvgVertSpeed)
		fields = append(fields, field)
	}
	if !m.EndTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = proto.Uint32(uint32(m.EndTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = proto.Uint32(m.TotalCalories)
		fields = append(fields, field)
	}
	if m.StartElevation != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 74)
		field.Value = proto.Uint32(m.StartElevation)
		fields = append(fields, field)
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 110)
		field.Value = proto.Uint32(m.TotalMovingTime)
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

// StartTimeUint32 returns StartTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Split) StartTimeUint32() uint32 { return datetime.ToUint32(m.StartTime) }

// EndTimeUint32 returns EndTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Split) EndTimeUint32() uint32 { return datetime.ToUint32(m.EndTime) }

// TotalElapsedTimeScaled return TotalElapsedTime in its scaled value.
// If TotalElapsedTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *Split) TotalElapsedTimeScaled() float64 {
	if m.TotalElapsedTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalElapsedTime)/1000 - 0
}

// TotalTimerTimeScaled return TotalTimerTime in its scaled value.
// If TotalTimerTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *Split) TotalTimerTimeScaled() float64 {
	if m.TotalTimerTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalTimerTime)/1000 - 0
}

// TotalDistanceScaled return TotalDistance in its scaled value.
// If TotalDistance value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m
func (m *Split) TotalDistanceScaled() float64 {
	if m.TotalDistance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalDistance)/100 - 0
}

// AvgSpeedScaled return AvgSpeed in its scaled value.
// If AvgSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *Split) AvgSpeedScaled() float64 {
	if m.AvgSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AvgSpeed)/1000 - 0
}

// MaxSpeedScaled return MaxSpeed in its scaled value.
// If MaxSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *Split) MaxSpeedScaled() float64 {
	if m.MaxSpeed == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.MaxSpeed)/1000 - 0
}

// AvgVertSpeedScaled return AvgVertSpeed in its scaled value.
// If AvgVertSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *Split) AvgVertSpeedScaled() float64 {
	if m.AvgVertSpeed == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.AvgVertSpeed)/1000 - 0
}

// StartElevationScaled return StartElevation in its scaled value.
// If StartElevation value is invalid, float64 invalid value will be returned.
//
// Scale: 5; Offset: 500; Units: m
func (m *Split) StartElevationScaled() float64 {
	if m.StartElevation == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.StartElevation)/5 - 500
}

// TotalMovingTimeScaled return TotalMovingTime in its scaled value.
// If TotalMovingTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *Split) TotalMovingTimeScaled() float64 {
	if m.TotalMovingTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TotalMovingTime)/1000 - 0
}

// StartPositionLatDegrees returns StartPositionLat in degrees instead of semicircles.
// If StartPositionLat value is invalid, float64 invalid value will be returned.
func (m *Split) StartPositionLatDegrees() float64 {
	return semicircles.ToDegrees(m.StartPositionLat)
}

// StartPositionLongDegrees returns StartPositionLong in degrees instead of semicircles.
// If StartPositionLong value is invalid, float64 invalid value will be returned.
func (m *Split) StartPositionLongDegrees() float64 {
	return semicircles.ToDegrees(m.StartPositionLong)
}

// EndPositionLatDegrees returns EndPositionLat in degrees instead of semicircles.
// If EndPositionLat value is invalid, float64 invalid value will be returned.
func (m *Split) EndPositionLatDegrees() float64 {
	return semicircles.ToDegrees(m.EndPositionLat)
}

// EndPositionLongDegrees returns EndPositionLong in degrees instead of semicircles.
// If EndPositionLong value is invalid, float64 invalid value will be returned.
func (m *Split) EndPositionLongDegrees() float64 {
	return semicircles.ToDegrees(m.EndPositionLong)
}

// SetMessageIndex sets MessageIndex value.
func (m *Split) SetMessageIndex(v typedef.MessageIndex) *Split {
	m.MessageIndex = v
	return m
}

// SetSplitType sets SplitType value.
func (m *Split) SetSplitType(v typedef.SplitType) *Split {
	m.SplitType = v
	return m
}

// SetTotalElapsedTime sets TotalElapsedTime value.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalElapsedTime(v uint32) *Split {
	m.TotalElapsedTime = v
	return m
}

// SetTotalElapsedTimeScaled is similar to SetTotalElapsedTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalElapsedTimeScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalElapsedTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalElapsedTime = uint32(unscaled)
	return m
}

// SetTotalTimerTime sets TotalTimerTime value.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalTimerTime(v uint32) *Split {
	m.TotalTimerTime = v
	return m
}

// SetTotalTimerTimeScaled is similar to SetTotalTimerTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalTimerTimeScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalTimerTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalTimerTime = uint32(unscaled)
	return m
}

// SetTotalDistance sets TotalDistance value.
//
// Scale: 100; Units: m
func (m *Split) SetTotalDistance(v uint32) *Split {
	m.TotalDistance = v
	return m
}

// SetTotalDistanceScaled is similar to SetTotalDistance except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m
func (m *Split) SetTotalDistanceScaled(v float64) *Split {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalDistance = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalDistance = uint32(unscaled)
	return m
}

// SetAvgSpeed sets AvgSpeed value.
//
// Scale: 1000; Units: m/s
func (m *Split) SetAvgSpeed(v uint32) *Split {
	m.AvgSpeed = v
	return m
}

// SetAvgSpeedScaled is similar to SetAvgSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *Split) SetAvgSpeedScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.AvgSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.AvgSpeed = uint32(unscaled)
	return m
}

// SetStartTime sets StartTime value.
func (m *Split) SetStartTime(v time.Time) *Split {
	m.StartTime = v
	return m
}

// SetTotalAscent sets TotalAscent value.
//
// Units: m
func (m *Split) SetTotalAscent(v uint16) *Split {
	m.TotalAscent = v
	return m
}

// SetTotalDescent sets TotalDescent value.
//
// Units: m
func (m *Split) SetTotalDescent(v uint16) *Split {
	m.TotalDescent = v
	return m
}

// SetStartPositionLat sets StartPositionLat value.
//
// Units: semicircles
func (m *Split) SetStartPositionLat(v int32) *Split {
	m.StartPositionLat = v
	return m
}

// SetStartPositionLatDegrees is similar to SetStartPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *Split) SetStartPositionLatDegrees(degrees float64) *Split {
	m.StartPositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetStartPositionLong sets StartPositionLong value.
//
// Units: semicircles
func (m *Split) SetStartPositionLong(v int32) *Split {
	m.StartPositionLong = v
	return m
}

// SetStartPositionLongDegrees is similar to SetStartPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *Split) SetStartPositionLongDegrees(degrees float64) *Split {
	m.StartPositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetEndPositionLat sets EndPositionLat value.
//
// Units: semicircles
func (m *Split) SetEndPositionLat(v int32) *Split {
	m.EndPositionLat = v
	return m
}

// SetEndPositionLatDegrees is similar to SetEndPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *Split) SetEndPositionLatDegrees(degrees float64) *Split {
	m.EndPositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetEndPositionLong sets EndPositionLong value.
//
// Units: semicircles
func (m *Split) SetEndPositionLong(v int32) *Split {
	m.EndPositionLong = v
	return m
}

// SetEndPositionLongDegrees is similar to SetEndPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *Split) SetEndPositionLongDegrees(degrees float64) *Split {
	m.EndPositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetMaxSpeed sets MaxSpeed value.
//
// Scale: 1000; Units: m/s
func (m *Split) SetMaxSpeed(v uint32) *Split {
	m.MaxSpeed = v
	return m
}

// SetMaxSpeedScaled is similar to SetMaxSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *Split) SetMaxSpeedScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.MaxSpeed = uint32(basetype.Uint32Invalid)
		return m
	}
	m.MaxSpeed = uint32(unscaled)
	return m
}

// SetAvgVertSpeed sets AvgVertSpeed value.
//
// Scale: 1000; Units: m/s
func (m *Split) SetAvgVertSpeed(v int32) *Split {
	m.AvgVertSpeed = v
	return m
}

// SetAvgVertSpeedScaled is similar to SetAvgVertSpeed except it accepts a scaled value.
// This method automatically converts the given value to its int32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *Split) SetAvgVertSpeedScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint32Invalid) {
		m.AvgVertSpeed = int32(basetype.Sint32Invalid)
		return m
	}
	m.AvgVertSpeed = int32(unscaled)
	return m
}

// SetEndTime sets EndTime value.
func (m *Split) SetEndTime(v time.Time) *Split {
	m.EndTime = v
	return m
}

// SetTotalCalories sets TotalCalories value.
//
// Units: kcal
func (m *Split) SetTotalCalories(v uint32) *Split {
	m.TotalCalories = v
	return m
}

// SetStartElevation sets StartElevation value.
//
// Scale: 5; Offset: 500; Units: m
func (m *Split) SetStartElevation(v uint32) *Split {
	m.StartElevation = v
	return m
}

// SetStartElevationScaled is similar to SetStartElevation except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 5; Offset: 500; Units: m
func (m *Split) SetStartElevationScaled(v float64) *Split {
	unscaled := (v + 500) * 5
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.StartElevation = uint32(basetype.Uint32Invalid)
		return m
	}
	m.StartElevation = uint32(unscaled)
	return m
}

// SetTotalMovingTime sets TotalMovingTime value.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalMovingTime(v uint32) *Split {
	m.TotalMovingTime = v
	return m
}

// SetTotalMovingTimeScaled is similar to SetTotalMovingTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *Split) SetTotalMovingTimeScaled(v float64) *Split {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.TotalMovingTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.TotalMovingTime = uint32(unscaled)
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Split) SetUnknownFields(unknownFields ...proto.Field) *Split {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *Split) SetDeveloperFields(developerFields ...proto.DeveloperField) *Split {
	m.DeveloperFields = developerFields
	return m
}
