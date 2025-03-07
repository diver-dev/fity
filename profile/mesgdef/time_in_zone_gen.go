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

// TimeInZone is a TimeInZone message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type TimeInZone struct {
	Timestamp                time.Time // Units: s
	TimeInHrZone             []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInSpeedZone          []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInCadenceZone        []uint32  // Array: [N]; Scale: 1000; Units: s
	TimeInPowerZone          []uint32  // Array: [N]; Scale: 1000; Units: s
	HrZoneHighBoundary       []uint8   // Array: [N]; Units: bpm
	SpeedZoneHighBoundary    []uint16  // Array: [N]; Scale: 1000; Units: m/s
	CadenceZoneHighBoundary  []uint8   // Array: [N]; Units: rpm
	PowerZoneHighBoundary    []uint16  // Array: [N]; Units: watts
	ReferenceMesg            typedef.MesgNum
	ReferenceIndex           typedef.MessageIndex
	FunctionalThresholdPower uint16
	HrCalcType               typedef.HrZoneCalc
	MaxHeartRate             uint8
	RestingHeartRate         uint8
	ThresholdHeartRate       uint8
	PwrCalcType              typedef.PwrZoneCalc

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewTimeInZone creates new TimeInZone struct based on given mesg.
// If mesg is nil, it will return TimeInZone with all fields being set to its corresponding invalid value.
func NewTimeInZone(mesg *proto.Message) *TimeInZone {
	m := new(TimeInZone)
	m.Reset(mesg)
	return m
}

// Reset resets all TimeInZone's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *TimeInZone) Reset(mesg *proto.Message) {
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

	*m = TimeInZone{
		Timestamp:                datetime.ToTime(vals[253].Uint32()),
		ReferenceMesg:            typedef.MesgNum(vals[0].Uint16()),
		ReferenceIndex:           typedef.MessageIndex(vals[1].Uint16()),
		TimeInHrZone:             vals[2].SliceUint32(),
		TimeInSpeedZone:          vals[3].SliceUint32(),
		TimeInCadenceZone:        vals[4].SliceUint32(),
		TimeInPowerZone:          vals[5].SliceUint32(),
		HrZoneHighBoundary:       vals[6].SliceUint8(),
		SpeedZoneHighBoundary:    vals[7].SliceUint16(),
		CadenceZoneHighBoundary:  vals[8].SliceUint8(),
		PowerZoneHighBoundary:    vals[9].SliceUint16(),
		HrCalcType:               typedef.HrZoneCalc(vals[10].Uint8()),
		MaxHeartRate:             vals[11].Uint8(),
		RestingHeartRate:         vals[12].Uint8(),
		ThresholdHeartRate:       vals[13].Uint8(),
		PwrCalcType:              typedef.PwrZoneCalc(vals[14].Uint8()),
		FunctionalThresholdPower: vals[15].Uint16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts TimeInZone into proto.Message. If options is nil, default options will be used.
func (m *TimeInZone) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumTimeInZone}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ReferenceMesg != typedef.MesgNumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.ReferenceMesg))
		fields = append(fields, field)
	}
	if m.ReferenceIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.ReferenceIndex))
		fields = append(fields, field)
	}
	if m.TimeInHrZone != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.SliceUint32(m.TimeInHrZone)
		fields = append(fields, field)
	}
	if m.TimeInSpeedZone != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceUint32(m.TimeInSpeedZone)
		fields = append(fields, field)
	}
	if m.TimeInCadenceZone != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.SliceUint32(m.TimeInCadenceZone)
		fields = append(fields, field)
	}
	if m.TimeInPowerZone != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint32(m.TimeInPowerZone)
		fields = append(fields, field)
	}
	if m.HrZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceUint8(m.HrZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.SpeedZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.SliceUint16(m.SpeedZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.CadenceZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceUint8(m.CadenceZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.PowerZoneHighBoundary != nil {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceUint16(m.PowerZoneHighBoundary)
		fields = append(fields, field)
	}
	if m.HrCalcType != typedef.HrZoneCalcInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(byte(m.HrCalcType))
		fields = append(fields, field)
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(m.MaxHeartRate)
		fields = append(fields, field)
	}
	if m.RestingHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(m.RestingHeartRate)
		fields = append(fields, field)
	}
	if m.ThresholdHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint8(m.ThresholdHeartRate)
		fields = append(fields, field)
	}
	if m.PwrCalcType != typedef.PwrZoneCalcInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint8(byte(m.PwrCalcType))
		fields = append(fields, field)
	}
	if m.FunctionalThresholdPower != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint16(m.FunctionalThresholdPower)
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
func (m *TimeInZone) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// TimeInHrZoneScaled return TimeInHrZone in its scaled value.
// If TimeInHrZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInHrZoneScaled() []float64 {
	if m.TimeInHrZone == nil {
		return nil
	}
	var vals = make([]float64, len(m.TimeInHrZone))
	for i := range m.TimeInHrZone {
		if m.TimeInHrZone[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.TimeInHrZone[i])/1000 - 0
	}
	return vals
}

// TimeInSpeedZoneScaled return TimeInSpeedZone in its scaled value.
// If TimeInSpeedZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInSpeedZoneScaled() []float64 {
	if m.TimeInSpeedZone == nil {
		return nil
	}
	var vals = make([]float64, len(m.TimeInSpeedZone))
	for i := range m.TimeInSpeedZone {
		if m.TimeInSpeedZone[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.TimeInSpeedZone[i])/1000 - 0
	}
	return vals
}

// TimeInCadenceZoneScaled return TimeInCadenceZone in its scaled value.
// If TimeInCadenceZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInCadenceZoneScaled() []float64 {
	if m.TimeInCadenceZone == nil {
		return nil
	}
	var vals = make([]float64, len(m.TimeInCadenceZone))
	for i := range m.TimeInCadenceZone {
		if m.TimeInCadenceZone[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.TimeInCadenceZone[i])/1000 - 0
	}
	return vals
}

// TimeInPowerZoneScaled return TimeInPowerZone in its scaled value.
// If TimeInPowerZone value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) TimeInPowerZoneScaled() []float64 {
	if m.TimeInPowerZone == nil {
		return nil
	}
	var vals = make([]float64, len(m.TimeInPowerZone))
	for i := range m.TimeInPowerZone {
		if m.TimeInPowerZone[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.TimeInPowerZone[i])/1000 - 0
	}
	return vals
}

// SpeedZoneHighBoundaryScaled return SpeedZoneHighBoundary in its scaled value.
// If SpeedZoneHighBoundary value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SpeedZoneHighBoundaryScaled() []float64 {
	if m.SpeedZoneHighBoundary == nil {
		return nil
	}
	var vals = make([]float64, len(m.SpeedZoneHighBoundary))
	for i := range m.SpeedZoneHighBoundary {
		if m.SpeedZoneHighBoundary[i] == basetype.Uint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.SpeedZoneHighBoundary[i])/1000 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *TimeInZone) SetTimestamp(v time.Time) *TimeInZone {
	m.Timestamp = v
	return m
}

// SetReferenceMesg sets ReferenceMesg value.
func (m *TimeInZone) SetReferenceMesg(v typedef.MesgNum) *TimeInZone {
	m.ReferenceMesg = v
	return m
}

// SetReferenceIndex sets ReferenceIndex value.
func (m *TimeInZone) SetReferenceIndex(v typedef.MessageIndex) *TimeInZone {
	m.ReferenceIndex = v
	return m
}

// SetTimeInHrZone sets TimeInHrZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInHrZone(v []uint32) *TimeInZone {
	m.TimeInHrZone = v
	return m
}

// SetTimeInHrZoneScaled is similar to SetTimeInHrZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInHrZoneScaled(vs []float64) *TimeInZone {
	if vs == nil {
		m.TimeInHrZone = nil
		return m
	}
	m.TimeInHrZone = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.TimeInHrZone[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.TimeInHrZone[i] = uint32(unscaled)
	}
	return m
}

// SetTimeInSpeedZone sets TimeInSpeedZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInSpeedZone(v []uint32) *TimeInZone {
	m.TimeInSpeedZone = v
	return m
}

// SetTimeInSpeedZoneScaled is similar to SetTimeInSpeedZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInSpeedZoneScaled(vs []float64) *TimeInZone {
	if vs == nil {
		m.TimeInSpeedZone = nil
		return m
	}
	m.TimeInSpeedZone = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.TimeInSpeedZone[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.TimeInSpeedZone[i] = uint32(unscaled)
	}
	return m
}

// SetTimeInCadenceZone sets TimeInCadenceZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInCadenceZone(v []uint32) *TimeInZone {
	m.TimeInCadenceZone = v
	return m
}

// SetTimeInCadenceZoneScaled is similar to SetTimeInCadenceZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInCadenceZoneScaled(vs []float64) *TimeInZone {
	if vs == nil {
		m.TimeInCadenceZone = nil
		return m
	}
	m.TimeInCadenceZone = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.TimeInCadenceZone[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.TimeInCadenceZone[i] = uint32(unscaled)
	}
	return m
}

// SetTimeInPowerZone sets TimeInPowerZone value.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInPowerZone(v []uint32) *TimeInZone {
	m.TimeInPowerZone = v
	return m
}

// SetTimeInPowerZoneScaled is similar to SetTimeInPowerZone except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s
func (m *TimeInZone) SetTimeInPowerZoneScaled(vs []float64) *TimeInZone {
	if vs == nil {
		m.TimeInPowerZone = nil
		return m
	}
	m.TimeInPowerZone = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.TimeInPowerZone[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.TimeInPowerZone[i] = uint32(unscaled)
	}
	return m
}

// SetHrZoneHighBoundary sets HrZoneHighBoundary value.
//
// Array: [N]; Units: bpm
func (m *TimeInZone) SetHrZoneHighBoundary(v []uint8) *TimeInZone {
	m.HrZoneHighBoundary = v
	return m
}

// SetSpeedZoneHighBoundary sets SpeedZoneHighBoundary value.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SetSpeedZoneHighBoundary(v []uint16) *TimeInZone {
	m.SpeedZoneHighBoundary = v
	return m
}

// SetSpeedZoneHighBoundaryScaled is similar to SetSpeedZoneHighBoundary except it accepts a scaled value.
// This method automatically converts the given value to its []uint16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: m/s
func (m *TimeInZone) SetSpeedZoneHighBoundaryScaled(vs []float64) *TimeInZone {
	if vs == nil {
		m.SpeedZoneHighBoundary = nil
		return m
	}
	m.SpeedZoneHighBoundary = make([]uint16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
			m.SpeedZoneHighBoundary[i] = uint16(basetype.Uint16Invalid)
			continue
		}
		m.SpeedZoneHighBoundary[i] = uint16(unscaled)
	}
	return m
}

// SetCadenceZoneHighBoundary sets CadenceZoneHighBoundary value.
//
// Array: [N]; Units: rpm
func (m *TimeInZone) SetCadenceZoneHighBoundary(v []uint8) *TimeInZone {
	m.CadenceZoneHighBoundary = v
	return m
}

// SetPowerZoneHighBoundary sets PowerZoneHighBoundary value.
//
// Array: [N]; Units: watts
func (m *TimeInZone) SetPowerZoneHighBoundary(v []uint16) *TimeInZone {
	m.PowerZoneHighBoundary = v
	return m
}

// SetHrCalcType sets HrCalcType value.
func (m *TimeInZone) SetHrCalcType(v typedef.HrZoneCalc) *TimeInZone {
	m.HrCalcType = v
	return m
}

// SetMaxHeartRate sets MaxHeartRate value.
func (m *TimeInZone) SetMaxHeartRate(v uint8) *TimeInZone {
	m.MaxHeartRate = v
	return m
}

// SetRestingHeartRate sets RestingHeartRate value.
func (m *TimeInZone) SetRestingHeartRate(v uint8) *TimeInZone {
	m.RestingHeartRate = v
	return m
}

// SetThresholdHeartRate sets ThresholdHeartRate value.
func (m *TimeInZone) SetThresholdHeartRate(v uint8) *TimeInZone {
	m.ThresholdHeartRate = v
	return m
}

// SetPwrCalcType sets PwrCalcType value.
func (m *TimeInZone) SetPwrCalcType(v typedef.PwrZoneCalc) *TimeInZone {
	m.PwrCalcType = v
	return m
}

// SetFunctionalThresholdPower sets FunctionalThresholdPower value.
func (m *TimeInZone) SetFunctionalThresholdPower(v uint16) *TimeInZone {
	m.FunctionalThresholdPower = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *TimeInZone) SetUnknownFields(unknownFields ...proto.Field) *TimeInZone {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *TimeInZone) SetDeveloperFields(developerFields ...proto.DeveloperField) *TimeInZone {
	m.DeveloperFields = developerFields
	return m
}
