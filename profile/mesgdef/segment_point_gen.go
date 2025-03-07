// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
)

// SegmentPoint is a SegmentPoint message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type SegmentPoint struct {
	LeaderTime       []uint32 // Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
	PositionLat      int32    // Units: semicircles
	PositionLong     int32    // Units: semicircles
	Distance         uint32   // Scale: 100; Units: m; Accumulated distance along the segment at the described point
	EnhancedAltitude uint32   // Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
	MessageIndex     typedef.MessageIndex
	Altitude         uint16 // Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point

	state [1]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewSegmentPoint creates new SegmentPoint struct based on given mesg.
// If mesg is nil, it will return SegmentPoint with all fields being set to its corresponding invalid value.
func NewSegmentPoint(mesg *proto.Message) *SegmentPoint {
	m := new(SegmentPoint)
	m.Reset(mesg)
	return m
}

// Reset resets all SegmentPoint's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *SegmentPoint) Reset(mesg *proto.Message) {
	var (
		vals            [255]proto.Value
		state           [1]uint8
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
			if mesg.Fields[i].Num < 7 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	*m = SegmentPoint{
		MessageIndex:     typedef.MessageIndex(vals[254].Uint16()),
		PositionLat:      vals[1].Int32(),
		PositionLong:     vals[2].Int32(),
		Distance:         vals[3].Uint32(),
		Altitude:         vals[4].Uint16(),
		LeaderTime:       vals[5].SliceUint32(),
		EnhancedAltitude: vals[6].Uint32(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts SegmentPoint into proto.Message. If options is nil, default options will be used.
func (m *SegmentPoint) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumSegmentPoint}

	if m.MessageIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = proto.Uint16(uint16(m.MessageIndex))
		fields = append(fields, field)
	}
	if m.PositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int32(m.PositionLat)
		fields = append(fields, field)
	}
	if m.PositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Int32(m.PositionLong)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Altitude != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Altitude)
		fields = append(fields, field)
	}
	if m.LeaderTime != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceUint32(m.LeaderTime)
		fields = append(fields, field)
	}
	if m.EnhancedAltitude != basetype.Uint32Invalid {
		if expanded := m.IsExpandedField(6); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 6)
			field.Value = proto.Uint32(m.EnhancedAltitude)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
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

// DistanceScaled return Distance in its scaled value.
// If Distance value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m; Accumulated distance along the segment at the described point
func (m *SegmentPoint) DistanceScaled() float64 {
	if m.Distance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Distance)/100 - 0
}

// AltitudeScaled return Altitude in its scaled value.
// If Altitude value is invalid, float64 invalid value will be returned.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) AltitudeScaled() float64 {
	if m.Altitude == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Altitude)/5 - 500
}

// LeaderTimeScaled return LeaderTime in its scaled value.
// If LeaderTime value is invalid, nil will be returned.
//
// Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
func (m *SegmentPoint) LeaderTimeScaled() []float64 {
	if m.LeaderTime == nil {
		return nil
	}
	var vals = make([]float64, len(m.LeaderTime))
	for i := range m.LeaderTime {
		if m.LeaderTime[i] == basetype.Uint32Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.LeaderTime[i])/1000 - 0
	}
	return vals
}

// EnhancedAltitudeScaled return EnhancedAltitude in its scaled value.
// If EnhancedAltitude value is invalid, float64 invalid value will be returned.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) EnhancedAltitudeScaled() float64 {
	if m.EnhancedAltitude == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.EnhancedAltitude)/5 - 500
}

// PositionLatDegrees returns PositionLat in degrees instead of semicircles.
// If PositionLat value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) PositionLatDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLat)
}

// PositionLongDegrees returns PositionLong in degrees instead of semicircles.
// If PositionLong value is invalid, float64 invalid value will be returned.
func (m *SegmentPoint) PositionLongDegrees() float64 {
	return semicircles.ToDegrees(m.PositionLong)
}

// SetMessageIndex sets MessageIndex value.
func (m *SegmentPoint) SetMessageIndex(v typedef.MessageIndex) *SegmentPoint {
	m.MessageIndex = v
	return m
}

// SetPositionLat sets PositionLat value.
//
// Units: semicircles
func (m *SegmentPoint) SetPositionLat(v int32) *SegmentPoint {
	m.PositionLat = v
	return m
}

// SetPositionLatDegrees is similar to SetPositionLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *SegmentPoint) SetPositionLatDegrees(degrees float64) *SegmentPoint {
	m.PositionLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetPositionLong sets PositionLong value.
//
// Units: semicircles
func (m *SegmentPoint) SetPositionLong(v int32) *SegmentPoint {
	m.PositionLong = v
	return m
}

// SetPositionLongDegrees is similar to SetPositionLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *SegmentPoint) SetPositionLongDegrees(degrees float64) *SegmentPoint {
	m.PositionLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetDistance sets Distance value.
//
// Scale: 100; Units: m; Accumulated distance along the segment at the described point
func (m *SegmentPoint) SetDistance(v uint32) *SegmentPoint {
	m.Distance = v
	return m
}

// SetDistanceScaled is similar to SetDistance except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m; Accumulated distance along the segment at the described point
func (m *SegmentPoint) SetDistanceScaled(v float64) *SegmentPoint {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Distance = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Distance = uint32(unscaled)
	return m
}

// SetAltitude sets Altitude value.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetAltitude(v uint16) *SegmentPoint {
	m.Altitude = v
	return m
}

// SetAltitudeScaled is similar to SetAltitude except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetAltitudeScaled(v float64) *SegmentPoint {
	unscaled := (v + 500) * 5
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Altitude = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Altitude = uint16(unscaled)
	return m
}

// SetLeaderTime sets LeaderTime value.
//
// Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
func (m *SegmentPoint) SetLeaderTime(v []uint32) *SegmentPoint {
	m.LeaderTime = v
	return m
}

// SetLeaderTimeScaled is similar to SetLeaderTime except it accepts a scaled value.
// This method automatically converts the given value to its []uint32 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 1000; Units: s; Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
func (m *SegmentPoint) SetLeaderTimeScaled(vs []float64) *SegmentPoint {
	if vs == nil {
		m.LeaderTime = nil
		return m
	}
	m.LeaderTime = make([]uint32, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 1000
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
			m.LeaderTime[i] = uint32(basetype.Uint32Invalid)
			continue
		}
		m.LeaderTime[i] = uint32(unscaled)
	}
	return m
}

// SetEnhancedAltitude sets EnhancedAltitude value.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetEnhancedAltitude(v uint32) *SegmentPoint {
	m.EnhancedAltitude = v
	return m
}

// SetEnhancedAltitudeScaled is similar to SetEnhancedAltitude except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 5; Offset: 500; Units: m; Accumulated altitude along the segment at the described point
func (m *SegmentPoint) SetEnhancedAltitudeScaled(v float64) *SegmentPoint {
	unscaled := (v + 500) * 5
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.EnhancedAltitude = uint32(basetype.Uint32Invalid)
		return m
	}
	m.EnhancedAltitude = uint32(unscaled)
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *SegmentPoint) SetUnknownFields(unknownFields ...proto.Field) *SegmentPoint {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *SegmentPoint) SetDeveloperFields(developerFields ...proto.DeveloperField) *SegmentPoint {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 6.
func (m *SegmentPoint) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 6:
	default:
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	m.state[pos] &^= bit
	if flag {
		m.state[pos] |= bit
	}
	return true
}

// IsExpandedField checks whether given fieldNum is a field generated through
// a component expansion. Eligible for field number: 6.
func (m *SegmentPoint) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 7 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
