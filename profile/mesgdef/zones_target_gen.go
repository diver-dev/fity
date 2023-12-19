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

// ZonesTarget is a ZonesTarget message.
type ZonesTarget struct {
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	FunctionalThresholdPower uint16
	HrCalcType               typedef.HrZoneCalc
	PwrCalcType              typedef.PwrZoneCalc

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewZonesTarget creates new ZonesTarget struct based on given mesg.
// If mesg is nil, it will return ZonesTarget with all fields being set to its corresponding invalid value.
func NewZonesTarget(mesg *proto.Message) *ZonesTarget {
	vals := [8]any{}

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

	return &ZonesTarget{
		MaxHeartRate:             typeconv.ToUint8[uint8](vals[1]),
		ThresholdHeartRate:       typeconv.ToUint8[uint8](vals[2]),
		FunctionalThresholdPower: typeconv.ToUint16[uint16](vals[3]),
		HrCalcType:               typeconv.ToEnum[typedef.HrZoneCalc](vals[5]),
		PwrCalcType:              typeconv.ToEnum[typedef.PwrZoneCalc](vals[7]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ZonesTarget into proto.Message.
func (m *ZonesTarget) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumZonesTarget)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if m.MaxHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.MaxHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ThresholdHeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.ThresholdHeartRate
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.FunctionalThresholdPower != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.FunctionalThresholdPower
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.HrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = typeconv.ToEnum[byte](m.HrCalcType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.PwrCalcType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToEnum[byte](m.PwrCalcType)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of ZonesTarget's valid fields.
func (m *ZonesTarget) size() byte {
	var size byte
	if m.MaxHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.ThresholdHeartRate != basetype.Uint8Invalid {
		size++
	}
	if m.FunctionalThresholdPower != basetype.Uint16Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.HrCalcType) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.PwrCalcType) != basetype.EnumInvalid {
		size++
	}
	return size
}

// SetMaxHeartRate sets ZonesTarget value.
func (m *ZonesTarget) SetMaxHeartRate(v uint8) *ZonesTarget {
	m.MaxHeartRate = v
	return m
}

// SetThresholdHeartRate sets ZonesTarget value.
func (m *ZonesTarget) SetThresholdHeartRate(v uint8) *ZonesTarget {
	m.ThresholdHeartRate = v
	return m
}

// SetFunctionalThresholdPower sets ZonesTarget value.
func (m *ZonesTarget) SetFunctionalThresholdPower(v uint16) *ZonesTarget {
	m.FunctionalThresholdPower = v
	return m
}

// SetHrCalcType sets ZonesTarget value.
func (m *ZonesTarget) SetHrCalcType(v typedef.HrZoneCalc) *ZonesTarget {
	m.HrCalcType = v
	return m
}

// SetPwrCalcType sets ZonesTarget value.
func (m *ZonesTarget) SetPwrCalcType(v typedef.PwrZoneCalc) *ZonesTarget {
	m.PwrCalcType = v
	return m
}

// SetDeveloperFields ZonesTarget's DeveloperFields.
func (m *ZonesTarget) SetDeveloperFields(developerFields ...proto.DeveloperField) *ZonesTarget {
	m.DeveloperFields = developerFields
	return m
}
