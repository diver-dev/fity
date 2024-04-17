// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// TankUpdate is a TankUpdate message.
type TankUpdate struct {
	Timestamp time.Time // Units: s
	Sensor    typedef.AntChannelId
	Pressure  uint16 // Scale: 100; Units: bar

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewTankUpdate creates new TankUpdate struct based on given mesg.
// If mesg is nil, it will return TankUpdate with all fields being set to its corresponding invalid value.
func NewTankUpdate(mesg *proto.Message) *TankUpdate {
	vals := [254]proto.Value{}

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

	return &TankUpdate{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		Sensor:    typedef.AntChannelId(vals[0].Uint32z()),
		Pressure:  vals[1].Uint16(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts TankUpdate into proto.Message. If options is nil, default options will be used.
func (m *TankUpdate) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumTankUpdate}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if uint32(m.Sensor) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint32(uint32(m.Sensor))
		fields = append(fields, field)
	}
	if m.Pressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.Pressure)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *TankUpdate) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// PressureScaled return Pressure in its scaled value [Scale: 100; Units: bar].
//
// If Pressure value is invalid, float64 invalid value will be returned.
func (m *TankUpdate) PressureScaled() float64 {
	if m.Pressure == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.Pressure, 100, 0)
}

// SetTimestamp sets TankUpdate value.
//
// Units: s
func (m *TankUpdate) SetTimestamp(v time.Time) *TankUpdate {
	m.Timestamp = v
	return m
}

// SetSensor sets TankUpdate value.
func (m *TankUpdate) SetSensor(v typedef.AntChannelId) *TankUpdate {
	m.Sensor = v
	return m
}

// SetPressure sets TankUpdate value.
//
// Scale: 100; Units: bar
func (m *TankUpdate) SetPressure(v uint16) *TankUpdate {
	m.Pressure = v
	return m
}

// SetDeveloperFields TankUpdate's DeveloperFields.
func (m *TankUpdate) SetDeveloperFields(developerFields ...proto.DeveloperField) *TankUpdate {
	m.DeveloperFields = developerFields
	return m
}
