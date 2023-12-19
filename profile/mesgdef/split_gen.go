// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// Split is a Split message.
type Split struct {
	MessageIndex      typedef.MessageIndex
	SplitType         typedef.SplitType
	TotalElapsedTime  uint32 // Scale: 1000; Units: s;
	TotalTimerTime    uint32 // Scale: 1000; Units: s;
	TotalDistance     uint32 // Scale: 100; Units: m;
	AvgSpeed          uint32 // Scale: 1000; Units: m/s;
	StartTime         time.Time
	TotalAscent       uint16 // Units: m;
	TotalDescent      uint16 // Units: m;
	StartPositionLat  int32  // Units: semicircles;
	StartPositionLong int32  // Units: semicircles;
	EndPositionLat    int32  // Units: semicircles;
	EndPositionLong   int32  // Units: semicircles;
	MaxSpeed          uint32 // Scale: 1000; Units: m/s;
	AvgVertSpeed      int32  // Scale: 1000; Units: m/s;
	EndTime           time.Time
	TotalCalories     uint32 // Units: kcal;
	StartElevation    uint32 // Scale: 5; Offset: 500; Units: m;
	TotalMovingTime   uint32 // Scale: 1000; Units: s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSplit creates new Split struct based on given mesg.
// If mesg is nil, it will return Split with all fields being set to its corresponding invalid value.
func NewSplit(mesg *proto.Message) *Split {
	vals := [255]any{}

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

	return &Split{
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		SplitType:         typeconv.ToEnum[typedef.SplitType](vals[0]),
		TotalElapsedTime:  typeconv.ToUint32[uint32](vals[1]),
		TotalTimerTime:    typeconv.ToUint32[uint32](vals[2]),
		TotalDistance:     typeconv.ToUint32[uint32](vals[3]),
		AvgSpeed:          typeconv.ToUint32[uint32](vals[4]),
		StartTime:         datetime.ToTime(vals[9]),
		TotalAscent:       typeconv.ToUint16[uint16](vals[13]),
		TotalDescent:      typeconv.ToUint16[uint16](vals[14]),
		StartPositionLat:  typeconv.ToSint32[int32](vals[21]),
		StartPositionLong: typeconv.ToSint32[int32](vals[22]),
		EndPositionLat:    typeconv.ToSint32[int32](vals[23]),
		EndPositionLong:   typeconv.ToSint32[int32](vals[24]),
		MaxSpeed:          typeconv.ToUint32[uint32](vals[25]),
		AvgVertSpeed:      typeconv.ToSint32[int32](vals[26]),
		EndTime:           datetime.ToTime(vals[27]),
		TotalCalories:     typeconv.ToUint32[uint32](vals[28]),
		StartElevation:    typeconv.ToUint32[uint32](vals[74]),
		TotalMovingTime:   typeconv.ToUint32[uint32](vals[110]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Split into proto.Message.
func (m *Split) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumSplit)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.SplitType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.SplitType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.TotalElapsedTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.TotalTimerTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalDistance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.TotalDistance
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.AvgSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = datetime.ToUint32(m.StartTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.TotalAscent
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.TotalDescent
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StartPositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = m.StartPositionLat
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StartPositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.StartPositionLong
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EndPositionLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.EndPositionLat
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.EndPositionLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.EndPositionLong
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = m.MaxSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 26)
		field.Value = m.AvgVertSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if datetime.ToUint32(m.EndTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = datetime.ToUint32(m.EndTime)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 28)
		field.Value = m.TotalCalories
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.StartElevation != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 74)
		field.Value = m.StartElevation
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 110)
		field.Value = m.TotalMovingTime
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of Split's valid fields.
func (m *Split) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SplitType) != basetype.EnumInvalid {
		size++
	}
	if m.TotalElapsedTime != basetype.Uint32Invalid {
		size++
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		size++
	}
	if m.TotalDistance != basetype.Uint32Invalid {
		size++
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		size++
	}
	if datetime.ToUint32(m.StartTime) != basetype.Uint32Invalid {
		size++
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		size++
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		size++
	}
	if m.StartPositionLat != basetype.Sint32Invalid {
		size++
	}
	if m.StartPositionLong != basetype.Sint32Invalid {
		size++
	}
	if m.EndPositionLat != basetype.Sint32Invalid {
		size++
	}
	if m.EndPositionLong != basetype.Sint32Invalid {
		size++
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		size++
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		size++
	}
	if datetime.ToUint32(m.EndTime) != basetype.Uint32Invalid {
		size++
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		size++
	}
	if m.StartElevation != basetype.Uint32Invalid {
		size++
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		size++
	}
	return size
}

// SetMessageIndex sets Split value.
func (m *Split) SetMessageIndex(v typedef.MessageIndex) *Split {
	m.MessageIndex = v
	return m
}

// SetSplitType sets Split value.
func (m *Split) SetSplitType(v typedef.SplitType) *Split {
	m.SplitType = v
	return m
}

// SetTotalElapsedTime sets Split value.
//
// Scale: 1000; Units: s;
func (m *Split) SetTotalElapsedTime(v uint32) *Split {
	m.TotalElapsedTime = v
	return m
}

// SetTotalTimerTime sets Split value.
//
// Scale: 1000; Units: s;
func (m *Split) SetTotalTimerTime(v uint32) *Split {
	m.TotalTimerTime = v
	return m
}

// SetTotalDistance sets Split value.
//
// Scale: 100; Units: m;
func (m *Split) SetTotalDistance(v uint32) *Split {
	m.TotalDistance = v
	return m
}

// SetAvgSpeed sets Split value.
//
// Scale: 1000; Units: m/s;
func (m *Split) SetAvgSpeed(v uint32) *Split {
	m.AvgSpeed = v
	return m
}

// SetStartTime sets Split value.
func (m *Split) SetStartTime(v time.Time) *Split {
	m.StartTime = v
	return m
}

// SetTotalAscent sets Split value.
//
// Units: m;
func (m *Split) SetTotalAscent(v uint16) *Split {
	m.TotalAscent = v
	return m
}

// SetTotalDescent sets Split value.
//
// Units: m;
func (m *Split) SetTotalDescent(v uint16) *Split {
	m.TotalDescent = v
	return m
}

// SetStartPositionLat sets Split value.
//
// Units: semicircles;
func (m *Split) SetStartPositionLat(v int32) *Split {
	m.StartPositionLat = v
	return m
}

// SetStartPositionLong sets Split value.
//
// Units: semicircles;
func (m *Split) SetStartPositionLong(v int32) *Split {
	m.StartPositionLong = v
	return m
}

// SetEndPositionLat sets Split value.
//
// Units: semicircles;
func (m *Split) SetEndPositionLat(v int32) *Split {
	m.EndPositionLat = v
	return m
}

// SetEndPositionLong sets Split value.
//
// Units: semicircles;
func (m *Split) SetEndPositionLong(v int32) *Split {
	m.EndPositionLong = v
	return m
}

// SetMaxSpeed sets Split value.
//
// Scale: 1000; Units: m/s;
func (m *Split) SetMaxSpeed(v uint32) *Split {
	m.MaxSpeed = v
	return m
}

// SetAvgVertSpeed sets Split value.
//
// Scale: 1000; Units: m/s;
func (m *Split) SetAvgVertSpeed(v int32) *Split {
	m.AvgVertSpeed = v
	return m
}

// SetEndTime sets Split value.
func (m *Split) SetEndTime(v time.Time) *Split {
	m.EndTime = v
	return m
}

// SetTotalCalories sets Split value.
//
// Units: kcal;
func (m *Split) SetTotalCalories(v uint32) *Split {
	m.TotalCalories = v
	return m
}

// SetStartElevation sets Split value.
//
// Scale: 5; Offset: 500; Units: m;
func (m *Split) SetStartElevation(v uint32) *Split {
	m.StartElevation = v
	return m
}

// SetTotalMovingTime sets Split value.
//
// Scale: 1000; Units: s;
func (m *Split) SetTotalMovingTime(v uint32) *Split {
	m.TotalMovingTime = v
	return m
}

// SetDeveloperFields Split's DeveloperFields.
func (m *Split) SetDeveloperFields(developerFields ...proto.DeveloperField) *Split {
	m.DeveloperFields = developerFields
	return m
}
