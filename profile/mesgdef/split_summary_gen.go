// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// SplitSummary is a SplitSummary message.
type SplitSummary struct {
	MessageIndex    typedef.MessageIndex
	SplitType       typedef.SplitType
	NumSplits       uint16
	TotalTimerTime  uint32 // Scale: 1000; Units: s;
	TotalDistance   uint32 // Scale: 100; Units: m;
	AvgSpeed        uint32 // Scale: 1000; Units: m/s;
	MaxSpeed        uint32 // Scale: 1000; Units: m/s;
	TotalAscent     uint16 // Units: m;
	TotalDescent    uint16 // Units: m;
	AvgHeartRate    uint8  // Units: bpm;
	MaxHeartRate    uint8  // Units: bpm;
	AvgVertSpeed    int32  // Scale: 1000; Units: m/s;
	TotalCalories   uint32 // Units: kcal;
	TotalMovingTime uint32 // Scale: 1000; Units: s;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSplitSummary creates new SplitSummary struct based on given mesg.
// If mesg is nil, it will return SplitSummary with all fields being set to its corresponding invalid value.
func NewSplitSummary(mesg *proto.Message) *SplitSummary {
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

	return &SplitSummary{
		MessageIndex:    typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		SplitType:       typeconv.ToEnum[typedef.SplitType](vals[0]),
		NumSplits:       typeconv.ToUint16[uint16](vals[3]),
		TotalTimerTime:  typeconv.ToUint32[uint32](vals[4]),
		TotalDistance:   typeconv.ToUint32[uint32](vals[5]),
		AvgSpeed:        typeconv.ToUint32[uint32](vals[6]),
		MaxSpeed:        typeconv.ToUint32[uint32](vals[7]),
		TotalAscent:     typeconv.ToUint16[uint16](vals[8]),
		TotalDescent:    typeconv.ToUint16[uint16](vals[9]),
		AvgHeartRate:    typeconv.ToUint8[uint8](vals[10]),
		MaxHeartRate:    typeconv.ToUint8[uint8](vals[11]),
		AvgVertSpeed:    typeconv.ToSint32[int32](vals[12]),
		TotalCalories:   typeconv.ToUint32[uint32](vals[13]),
		TotalMovingTime: typeconv.ToUint32[uint32](vals[77]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SplitSummary into proto.Message.
func (m *SplitSummary) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumSplitSummary)
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
	if m.NumSplits != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.NumSplits
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalTimerTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.TotalTimerTime
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalDistance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.TotalDistance
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = m.AvgSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxSpeed != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.MaxSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.TotalAscent
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.TotalDescent
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.AvgHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.MaxHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = m.AvgVertSpeed
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.TotalCalories
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 77)
		field.Value = m.TotalMovingTime
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of SplitSummary's valid fields.
func (m *SplitSummary) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SplitType) != basetype.EnumInvalid {
		size++
	}
	if m.NumSplits != basetype.Uint16Invalid {
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
	if m.MaxSpeed != basetype.Uint32Invalid {
		size++
	}
	if m.TotalAscent != basetype.Uint16Invalid {
		size++
	}
	if m.TotalDescent != basetype.Uint16Invalid {
		size++
	}
	if m.AvgHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.MaxHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.AvgVertSpeed != basetype.Sint32Invalid {
		size++
	}
	if m.TotalCalories != basetype.Uint32Invalid {
		size++
	}
	if m.TotalMovingTime != basetype.Uint32Invalid {
		size++
	}
	return size
}

// SetMessageIndex sets SplitSummary value.
func (m *SplitSummary) SetMessageIndex(v typedef.MessageIndex) *SplitSummary {
	m.MessageIndex = v
	return m
}

// SetSplitType sets SplitSummary value.
func (m *SplitSummary) SetSplitType(v typedef.SplitType) *SplitSummary {
	m.SplitType = v
	return m
}

// SetNumSplits sets SplitSummary value.
func (m *SplitSummary) SetNumSplits(v uint16) *SplitSummary {
	m.NumSplits = v
	return m
}

// SetTotalTimerTime sets SplitSummary value.
//
// Scale: 1000; Units: s;
func (m *SplitSummary) SetTotalTimerTime(v uint32) *SplitSummary {
	m.TotalTimerTime = v
	return m
}

// SetTotalDistance sets SplitSummary value.
//
// Scale: 100; Units: m;
func (m *SplitSummary) SetTotalDistance(v uint32) *SplitSummary {
	m.TotalDistance = v
	return m
}

// SetAvgSpeed sets SplitSummary value.
//
// Scale: 1000; Units: m/s;
func (m *SplitSummary) SetAvgSpeed(v uint32) *SplitSummary {
	m.AvgSpeed = v
	return m
}

// SetMaxSpeed sets SplitSummary value.
//
// Scale: 1000; Units: m/s;
func (m *SplitSummary) SetMaxSpeed(v uint32) *SplitSummary {
	m.MaxSpeed = v
	return m
}

// SetTotalAscent sets SplitSummary value.
//
// Units: m;
func (m *SplitSummary) SetTotalAscent(v uint16) *SplitSummary {
	m.TotalAscent = v
	return m
}

// SetTotalDescent sets SplitSummary value.
//
// Units: m;
func (m *SplitSummary) SetTotalDescent(v uint16) *SplitSummary {
	m.TotalDescent = v
	return m
}

// SetAvgHeartRate sets SplitSummary value.
//
// Units: bpm;
func (m *SplitSummary) SetAvgHeartRate(v uint8) *SplitSummary {
	m.AvgHeartRate = v
	return m
}

// SetMaxHeartRate sets SplitSummary value.
//
// Units: bpm;
func (m *SplitSummary) SetMaxHeartRate(v uint8) *SplitSummary {
	m.MaxHeartRate = v
	return m
}

// SetAvgVertSpeed sets SplitSummary value.
//
// Scale: 1000; Units: m/s;
func (m *SplitSummary) SetAvgVertSpeed(v int32) *SplitSummary {
	m.AvgVertSpeed = v
	return m
}

// SetTotalCalories sets SplitSummary value.
//
// Units: kcal;
func (m *SplitSummary) SetTotalCalories(v uint32) *SplitSummary {
	m.TotalCalories = v
	return m
}

// SetTotalMovingTime sets SplitSummary value.
//
// Scale: 1000; Units: s;
func (m *SplitSummary) SetTotalMovingTime(v uint32) *SplitSummary {
	m.TotalMovingTime = v
	return m
}

// SetDeveloperFields SplitSummary's DeveloperFields.
func (m *SplitSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *SplitSummary {
	m.DeveloperFields = developerFields
	return m
}
