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

// BikeProfile is a BikeProfile message.
type BikeProfile struct {
	MessageIndex             typedef.MessageIndex
	Name                     string
	Sport                    typedef.Sport
	SubSport                 typedef.SubSport
	Odometer                 uint32 // Scale: 100; Units: m;
	BikeSpdAntId             uint16
	BikeCadAntId             uint16
	BikeSpdcadAntId          uint16
	BikePowerAntId           uint16
	CustomWheelsize          uint16 // Scale: 1000; Units: m;
	AutoWheelsize            uint16 // Scale: 1000; Units: m;
	BikeWeight               uint16 // Scale: 10; Units: kg;
	PowerCalFactor           uint16 // Scale: 10; Units: %;
	AutoWheelCal             bool
	AutoPowerZero            bool
	Id                       uint8
	SpdEnabled               bool
	CadEnabled               bool
	SpdcadEnabled            bool
	PowerEnabled             bool
	CrankLength              uint8 // Scale: 2; Offset: -110; Units: mm;
	Enabled                  bool
	BikeSpdAntIdTransType    uint8
	BikeCadAntIdTransType    uint8
	BikeSpdcadAntIdTransType uint8
	BikePowerAntIdTransType  uint8
	OdometerRollover         uint8   // Rollover counter that can be used to extend the odometer
	FrontGearNum             uint8   // Number of front gears
	FrontGear                []uint8 // Array: [N]; Number of teeth on each gear 0 is innermost
	RearGearNum              uint8   // Number of rear gears
	RearGear                 []uint8 // Array: [N]; Number of teeth on each gear 0 is innermost
	ShimanoDi2Enabled        bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewBikeProfile creates new BikeProfile struct based on given mesg.
// If mesg is nil, it will return BikeProfile with all fields being set to its corresponding invalid value.
func NewBikeProfile(mesg *proto.Message) *BikeProfile {
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

	return &BikeProfile{
		MessageIndex:             typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Name:                     typeconv.ToString[string](vals[0]),
		Sport:                    typeconv.ToEnum[typedef.Sport](vals[1]),
		SubSport:                 typeconv.ToEnum[typedef.SubSport](vals[2]),
		Odometer:                 typeconv.ToUint32[uint32](vals[3]),
		BikeSpdAntId:             typeconv.ToUint16z[uint16](vals[4]),
		BikeCadAntId:             typeconv.ToUint16z[uint16](vals[5]),
		BikeSpdcadAntId:          typeconv.ToUint16z[uint16](vals[6]),
		BikePowerAntId:           typeconv.ToUint16z[uint16](vals[7]),
		CustomWheelsize:          typeconv.ToUint16[uint16](vals[8]),
		AutoWheelsize:            typeconv.ToUint16[uint16](vals[9]),
		BikeWeight:               typeconv.ToUint16[uint16](vals[10]),
		PowerCalFactor:           typeconv.ToUint16[uint16](vals[11]),
		AutoWheelCal:             typeconv.ToBool[bool](vals[12]),
		AutoPowerZero:            typeconv.ToBool[bool](vals[13]),
		Id:                       typeconv.ToUint8[uint8](vals[14]),
		SpdEnabled:               typeconv.ToBool[bool](vals[15]),
		CadEnabled:               typeconv.ToBool[bool](vals[16]),
		SpdcadEnabled:            typeconv.ToBool[bool](vals[17]),
		PowerEnabled:             typeconv.ToBool[bool](vals[18]),
		CrankLength:              typeconv.ToUint8[uint8](vals[19]),
		Enabled:                  typeconv.ToBool[bool](vals[20]),
		BikeSpdAntIdTransType:    typeconv.ToUint8z[uint8](vals[21]),
		BikeCadAntIdTransType:    typeconv.ToUint8z[uint8](vals[22]),
		BikeSpdcadAntIdTransType: typeconv.ToUint8z[uint8](vals[23]),
		BikePowerAntIdTransType:  typeconv.ToUint8z[uint8](vals[24]),
		OdometerRollover:         typeconv.ToUint8[uint8](vals[37]),
		FrontGearNum:             typeconv.ToUint8z[uint8](vals[38]),
		FrontGear:                typeconv.ToSliceUint8z[uint8](vals[39]),
		RearGearNum:              typeconv.ToUint8z[uint8](vals[40]),
		RearGear:                 typeconv.ToSliceUint8z[uint8](vals[41]),
		ShimanoDi2Enabled:        typeconv.ToBool[bool](vals[44]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts BikeProfile into proto.Message.
func (m *BikeProfile) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumBikeProfile)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Name != basetype.StringInvalid && m.Name != "" {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Name
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToEnum[byte](m.Sport)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = typeconv.ToEnum[byte](m.SubSport)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Odometer != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Odometer
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16z[uint16](m.BikeSpdAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = typeconv.ToUint16z[uint16](m.BikeSpdAntId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16z[uint16](m.BikeCadAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = typeconv.ToUint16z[uint16](m.BikeCadAntId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16z[uint16](m.BikeSpdcadAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = typeconv.ToUint16z[uint16](m.BikeSpdcadAntId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16z[uint16](m.BikePowerAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = typeconv.ToUint16z[uint16](m.BikePowerAntId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.CustomWheelsize != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.CustomWheelsize
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AutoWheelsize != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = m.AutoWheelsize
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.BikeWeight != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = m.BikeWeight
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.PowerCalFactor != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = m.PowerCalFactor
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AutoWheelCal != false {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = m.AutoWheelCal
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.AutoPowerZero != false {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = m.AutoPowerZero
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Id != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = m.Id
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.SpdEnabled != false {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = m.SpdEnabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.CadEnabled != false {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = m.CadEnabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.SpdcadEnabled != false {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = m.SpdcadEnabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.PowerEnabled != false {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = m.PowerEnabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.CrankLength != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = m.CrankLength
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = m.Enabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.BikeSpdAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = typeconv.ToUint8z[uint8](m.BikeSpdAntIdTransType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.BikeCadAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = typeconv.ToUint8z[uint8](m.BikeCadAntIdTransType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.BikeSpdcadAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = typeconv.ToUint8z[uint8](m.BikeSpdcadAntIdTransType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.BikePowerAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = typeconv.ToUint8z[uint8](m.BikePowerAntIdTransType)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.OdometerRollover != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 37)
		field.Value = m.OdometerRollover
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.FrontGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 38)
		field.Value = typeconv.ToUint8z[uint8](m.FrontGearNum)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToSliceUint8z[uint8](m.FrontGear) != nil {
		field := fac.CreateField(mesg.Num, 39)
		field.Value = typeconv.ToSliceUint8z[uint8](m.FrontGear)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.RearGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 40)
		field.Value = typeconv.ToUint8z[uint8](m.RearGearNum)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToSliceUint8z[uint8](m.RearGear) != nil {
		field := fac.CreateField(mesg.Num, 41)
		field.Value = typeconv.ToSliceUint8z[uint8](m.RearGear)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.ShimanoDi2Enabled != false {
		field := fac.CreateField(mesg.Num, 44)
		field.Value = m.ShimanoDi2Enabled
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of BikeProfile's valid fields.
func (m *BikeProfile) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if m.Name != basetype.StringInvalid && m.Name != "" {
		size++
	}
	if typeconv.ToEnum[byte](m.Sport) != basetype.EnumInvalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SubSport) != basetype.EnumInvalid {
		size++
	}
	if m.Odometer != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToUint16z[uint16](m.BikeSpdAntId) != basetype.Uint16zInvalid {
		size++
	}
	if typeconv.ToUint16z[uint16](m.BikeCadAntId) != basetype.Uint16zInvalid {
		size++
	}
	if typeconv.ToUint16z[uint16](m.BikeSpdcadAntId) != basetype.Uint16zInvalid {
		size++
	}
	if typeconv.ToUint16z[uint16](m.BikePowerAntId) != basetype.Uint16zInvalid {
		size++
	}
	if m.CustomWheelsize != basetype.Uint16Invalid {
		size++
	}
	if m.AutoWheelsize != basetype.Uint16Invalid {
		size++
	}
	if m.BikeWeight != basetype.Uint16Invalid {
		size++
	}
	if m.PowerCalFactor != basetype.Uint16Invalid {
		size++
	}
	if m.AutoWheelCal != false {
		size++
	}
	if m.AutoPowerZero != false {
		size++
	}
	if m.Id != basetype.Uint8Invalid {
		size++
	}
	if m.SpdEnabled != false {
		size++
	}
	if m.CadEnabled != false {
		size++
	}
	if m.SpdcadEnabled != false {
		size++
	}
	if m.PowerEnabled != false {
		size++
	}
	if m.CrankLength != basetype.Uint8Invalid {
		size++
	}
	if m.Enabled != false {
		size++
	}
	if typeconv.ToUint8z[uint8](m.BikeSpdAntIdTransType) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.BikeCadAntIdTransType) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.BikeSpdcadAntIdTransType) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.BikePowerAntIdTransType) != basetype.Uint8zInvalid {
		size++
	}
	if m.OdometerRollover != basetype.Uint8Invalid {
		size++
	}
	if typeconv.ToUint8z[uint8](m.FrontGearNum) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToSliceUint8z[uint8](m.FrontGear) != nil {
		size++
	}
	if typeconv.ToUint8z[uint8](m.RearGearNum) != basetype.Uint8zInvalid {
		size++
	}
	if typeconv.ToSliceUint8z[uint8](m.RearGear) != nil {
		size++
	}
	if m.ShimanoDi2Enabled != false {
		size++
	}
	return size
}

// SetMessageIndex sets BikeProfile value.
func (m *BikeProfile) SetMessageIndex(v typedef.MessageIndex) *BikeProfile {
	m.MessageIndex = v
	return m
}

// SetName sets BikeProfile value.
func (m *BikeProfile) SetName(v string) *BikeProfile {
	m.Name = v
	return m
}

// SetSport sets BikeProfile value.
func (m *BikeProfile) SetSport(v typedef.Sport) *BikeProfile {
	m.Sport = v
	return m
}

// SetSubSport sets BikeProfile value.
func (m *BikeProfile) SetSubSport(v typedef.SubSport) *BikeProfile {
	m.SubSport = v
	return m
}

// SetOdometer sets BikeProfile value.
//
// Scale: 100; Units: m;
func (m *BikeProfile) SetOdometer(v uint32) *BikeProfile {
	m.Odometer = v
	return m
}

// SetBikeSpdAntId sets BikeProfile value.
func (m *BikeProfile) SetBikeSpdAntId(v uint16) *BikeProfile {
	m.BikeSpdAntId = v
	return m
}

// SetBikeCadAntId sets BikeProfile value.
func (m *BikeProfile) SetBikeCadAntId(v uint16) *BikeProfile {
	m.BikeCadAntId = v
	return m
}

// SetBikeSpdcadAntId sets BikeProfile value.
func (m *BikeProfile) SetBikeSpdcadAntId(v uint16) *BikeProfile {
	m.BikeSpdcadAntId = v
	return m
}

// SetBikePowerAntId sets BikeProfile value.
func (m *BikeProfile) SetBikePowerAntId(v uint16) *BikeProfile {
	m.BikePowerAntId = v
	return m
}

// SetCustomWheelsize sets BikeProfile value.
//
// Scale: 1000; Units: m;
func (m *BikeProfile) SetCustomWheelsize(v uint16) *BikeProfile {
	m.CustomWheelsize = v
	return m
}

// SetAutoWheelsize sets BikeProfile value.
//
// Scale: 1000; Units: m;
func (m *BikeProfile) SetAutoWheelsize(v uint16) *BikeProfile {
	m.AutoWheelsize = v
	return m
}

// SetBikeWeight sets BikeProfile value.
//
// Scale: 10; Units: kg;
func (m *BikeProfile) SetBikeWeight(v uint16) *BikeProfile {
	m.BikeWeight = v
	return m
}

// SetPowerCalFactor sets BikeProfile value.
//
// Scale: 10; Units: %;
func (m *BikeProfile) SetPowerCalFactor(v uint16) *BikeProfile {
	m.PowerCalFactor = v
	return m
}

// SetAutoWheelCal sets BikeProfile value.
func (m *BikeProfile) SetAutoWheelCal(v bool) *BikeProfile {
	m.AutoWheelCal = v
	return m
}

// SetAutoPowerZero sets BikeProfile value.
func (m *BikeProfile) SetAutoPowerZero(v bool) *BikeProfile {
	m.AutoPowerZero = v
	return m
}

// SetId sets BikeProfile value.
func (m *BikeProfile) SetId(v uint8) *BikeProfile {
	m.Id = v
	return m
}

// SetSpdEnabled sets BikeProfile value.
func (m *BikeProfile) SetSpdEnabled(v bool) *BikeProfile {
	m.SpdEnabled = v
	return m
}

// SetCadEnabled sets BikeProfile value.
func (m *BikeProfile) SetCadEnabled(v bool) *BikeProfile {
	m.CadEnabled = v
	return m
}

// SetSpdcadEnabled sets BikeProfile value.
func (m *BikeProfile) SetSpdcadEnabled(v bool) *BikeProfile {
	m.SpdcadEnabled = v
	return m
}

// SetPowerEnabled sets BikeProfile value.
func (m *BikeProfile) SetPowerEnabled(v bool) *BikeProfile {
	m.PowerEnabled = v
	return m
}

// SetCrankLength sets BikeProfile value.
//
// Scale: 2; Offset: -110; Units: mm;
func (m *BikeProfile) SetCrankLength(v uint8) *BikeProfile {
	m.CrankLength = v
	return m
}

// SetEnabled sets BikeProfile value.
func (m *BikeProfile) SetEnabled(v bool) *BikeProfile {
	m.Enabled = v
	return m
}

// SetBikeSpdAntIdTransType sets BikeProfile value.
func (m *BikeProfile) SetBikeSpdAntIdTransType(v uint8) *BikeProfile {
	m.BikeSpdAntIdTransType = v
	return m
}

// SetBikeCadAntIdTransType sets BikeProfile value.
func (m *BikeProfile) SetBikeCadAntIdTransType(v uint8) *BikeProfile {
	m.BikeCadAntIdTransType = v
	return m
}

// SetBikeSpdcadAntIdTransType sets BikeProfile value.
func (m *BikeProfile) SetBikeSpdcadAntIdTransType(v uint8) *BikeProfile {
	m.BikeSpdcadAntIdTransType = v
	return m
}

// SetBikePowerAntIdTransType sets BikeProfile value.
func (m *BikeProfile) SetBikePowerAntIdTransType(v uint8) *BikeProfile {
	m.BikePowerAntIdTransType = v
	return m
}

// SetOdometerRollover sets BikeProfile value.
//
// Rollover counter that can be used to extend the odometer
func (m *BikeProfile) SetOdometerRollover(v uint8) *BikeProfile {
	m.OdometerRollover = v
	return m
}

// SetFrontGearNum sets BikeProfile value.
//
// Number of front gears
func (m *BikeProfile) SetFrontGearNum(v uint8) *BikeProfile {
	m.FrontGearNum = v
	return m
}

// SetFrontGear sets BikeProfile value.
//
// Array: [N]; Number of teeth on each gear 0 is innermost
func (m *BikeProfile) SetFrontGear(v []uint8) *BikeProfile {
	m.FrontGear = v
	return m
}

// SetRearGearNum sets BikeProfile value.
//
// Number of rear gears
func (m *BikeProfile) SetRearGearNum(v uint8) *BikeProfile {
	m.RearGearNum = v
	return m
}

// SetRearGear sets BikeProfile value.
//
// Array: [N]; Number of teeth on each gear 0 is innermost
func (m *BikeProfile) SetRearGear(v []uint8) *BikeProfile {
	m.RearGear = v
	return m
}

// SetShimanoDi2Enabled sets BikeProfile value.
func (m *BikeProfile) SetShimanoDi2Enabled(v bool) *BikeProfile {
	m.ShimanoDi2Enabled = v
	return m
}

// SetDeveloperFields BikeProfile's DeveloperFields.
func (m *BikeProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *BikeProfile {
	m.DeveloperFields = developerFields
	return m
}
