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
	"math"
	"time"
)

// MaxMetData is a MaxMetData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type MaxMetData struct {
	UpdateTime     time.Time // Time maxMET and vo2 were calculated
	Vo2Max         uint16    // Scale: 10; Units: mL/kg/min
	Sport          typedef.Sport
	SubSport       typedef.SubSport
	MaxMetCategory typedef.MaxMetCategory
	CalibratedData bool                          // Indicates if calibrated data was used in the calculation
	HrSource       typedef.MaxMetHeartRateSource // Indicates if the estimate was obtained using a chest strap or wrist heart rate
	SpeedSource    typedef.MaxMetSpeedSource     // Indidcates if the estimate was obtained using onboard GPS or connected GPS

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewMaxMetData creates new MaxMetData struct based on given mesg.
// If mesg is nil, it will return MaxMetData with all fields being set to its corresponding invalid value.
func NewMaxMetData(mesg *proto.Message) *MaxMetData {
	vals := [14]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 13 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &MaxMetData{
		UpdateTime:     datetime.ToTime(vals[0].Uint32()),
		Vo2Max:         vals[2].Uint16(),
		Sport:          typedef.Sport(vals[5].Uint8()),
		SubSport:       typedef.SubSport(vals[6].Uint8()),
		MaxMetCategory: typedef.MaxMetCategory(vals[8].Uint8()),
		CalibratedData: vals[9].Bool(),
		HrSource:       typedef.MaxMetHeartRateSource(vals[12].Uint8()),
		SpeedSource:    typedef.MaxMetSpeedSource(vals[13].Uint8()),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts MaxMetData into proto.Message. If options is nil, default options will be used.
func (m *MaxMetData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumMaxMetData}

	if datetime.ToUint32(m.UpdateTime) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(datetime.ToUint32(m.UpdateTime))
		fields = append(fields, field)
	}
	if m.Vo2Max != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.Vo2Max)
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(byte(m.Sport))
		fields = append(fields, field)
	}
	if byte(m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(byte(m.SubSport))
		fields = append(fields, field)
	}
	if byte(m.MaxMetCategory) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(byte(m.MaxMetCategory))
		fields = append(fields, field)
	}
	if m.CalibratedData != false {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Bool(m.CalibratedData)
		fields = append(fields, field)
	}
	if byte(m.HrSource) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(byte(m.HrSource))
		fields = append(fields, field)
	}
	if byte(m.SpeedSource) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint8(byte(m.SpeedSource))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// UpdateTimeUint32 returns UpdateTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *MaxMetData) UpdateTimeUint32() uint32 { return datetime.ToUint32(m.UpdateTime) }

// Vo2MaxScaled return Vo2Max in its scaled value.
// If Vo2Max value is invalid, float64 invalid value will be returned.
//
// Scale: 10; Units: mL/kg/min
func (m *MaxMetData) Vo2MaxScaled() float64 {
	if m.Vo2Max == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Vo2Max)/10 - 0
}

// SetUpdateTime sets UpdateTime value.
//
// Time maxMET and vo2 were calculated
func (m *MaxMetData) SetUpdateTime(v time.Time) *MaxMetData {
	m.UpdateTime = v
	return m
}

// SetVo2Max sets Vo2Max value.
//
// Scale: 10; Units: mL/kg/min
func (m *MaxMetData) SetVo2Max(v uint16) *MaxMetData {
	m.Vo2Max = v
	return m
}

// SetVo2MaxScaled is similar to SetVo2Max except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 10; Units: mL/kg/min
func (m *MaxMetData) SetVo2MaxScaled(v float64) *MaxMetData {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.Vo2Max = uint16(basetype.Uint16Invalid)
		return m
	}
	m.Vo2Max = uint16(unscaled)
	return m
}

// SetSport sets Sport value.
func (m *MaxMetData) SetSport(v typedef.Sport) *MaxMetData {
	m.Sport = v
	return m
}

// SetSubSport sets SubSport value.
func (m *MaxMetData) SetSubSport(v typedef.SubSport) *MaxMetData {
	m.SubSport = v
	return m
}

// SetMaxMetCategory sets MaxMetCategory value.
func (m *MaxMetData) SetMaxMetCategory(v typedef.MaxMetCategory) *MaxMetData {
	m.MaxMetCategory = v
	return m
}

// SetCalibratedData sets CalibratedData value.
//
// Indicates if calibrated data was used in the calculation
func (m *MaxMetData) SetCalibratedData(v bool) *MaxMetData {
	m.CalibratedData = v
	return m
}

// SetHrSource sets HrSource value.
//
// Indicates if the estimate was obtained using a chest strap or wrist heart rate
func (m *MaxMetData) SetHrSource(v typedef.MaxMetHeartRateSource) *MaxMetData {
	m.HrSource = v
	return m
}

// SetSpeedSource sets SpeedSource value.
//
// Indidcates if the estimate was obtained using onboard GPS or connected GPS
func (m *MaxMetData) SetSpeedSource(v typedef.MaxMetSpeedSource) *MaxMetData {
	m.SpeedSource = v
	return m
}

// SetDeveloperFields MaxMetData's DeveloperFields.
func (m *MaxMetData) SetDeveloperFields(developerFields ...proto.DeveloperField) *MaxMetData {
	m.DeveloperFields = developerFields
	return m
}
