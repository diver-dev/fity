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

// DeviceAuxBatteryInfo is a DeviceAuxBatteryInfo message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type DeviceAuxBatteryInfo struct {
	Timestamp         time.Time
	BatteryVoltage    uint16 // Scale: 256; Units: V
	DeviceIndex       typedef.DeviceIndex
	BatteryStatus     typedef.BatteryStatus
	BatteryIdentifier uint8

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewDeviceAuxBatteryInfo creates new DeviceAuxBatteryInfo struct based on given mesg.
// If mesg is nil, it will return DeviceAuxBatteryInfo with all fields being set to its corresponding invalid value.
func NewDeviceAuxBatteryInfo(mesg *proto.Message) *DeviceAuxBatteryInfo {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
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

	return &DeviceAuxBatteryInfo{
		Timestamp:         datetime.ToTime(vals[253].Uint32()),
		DeviceIndex:       typedef.DeviceIndex(vals[0].Uint8()),
		BatteryVoltage:    vals[1].Uint16(),
		BatteryStatus:     typedef.BatteryStatus(vals[2].Uint8()),
		BatteryIdentifier: vals[3].Uint8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts DeviceAuxBatteryInfo into proto.Message. If options is nil, default options will be used.
func (m *DeviceAuxBatteryInfo) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumDeviceAuxBatteryInfo}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.DeviceIndex != typedef.DeviceIndexInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if m.BatteryVoltage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.BatteryVoltage)
		fields = append(fields, field)
	}
	if m.BatteryStatus != typedef.BatteryStatusInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(uint8(m.BatteryStatus))
		fields = append(fields, field)
	}
	if m.BatteryIdentifier != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.BatteryIdentifier)
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
func (m *DeviceAuxBatteryInfo) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// BatteryVoltageScaled return BatteryVoltage in its scaled value.
// If BatteryVoltage value is invalid, float64 invalid value will be returned.
//
// Scale: 256; Units: V
func (m *DeviceAuxBatteryInfo) BatteryVoltageScaled() float64 {
	if m.BatteryVoltage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.BatteryVoltage)/256 - 0
}

// SetTimestamp sets Timestamp value.
func (m *DeviceAuxBatteryInfo) SetTimestamp(v time.Time) *DeviceAuxBatteryInfo {
	m.Timestamp = v
	return m
}

// SetDeviceIndex sets DeviceIndex value.
func (m *DeviceAuxBatteryInfo) SetDeviceIndex(v typedef.DeviceIndex) *DeviceAuxBatteryInfo {
	m.DeviceIndex = v
	return m
}

// SetBatteryVoltage sets BatteryVoltage value.
//
// Scale: 256; Units: V
func (m *DeviceAuxBatteryInfo) SetBatteryVoltage(v uint16) *DeviceAuxBatteryInfo {
	m.BatteryVoltage = v
	return m
}

// SetBatteryVoltageScaled is similar to SetBatteryVoltage except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 256; Units: V
func (m *DeviceAuxBatteryInfo) SetBatteryVoltageScaled(v float64) *DeviceAuxBatteryInfo {
	unscaled := (v + 0) * 256
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.BatteryVoltage = uint16(basetype.Uint16Invalid)
		return m
	}
	m.BatteryVoltage = uint16(unscaled)
	return m
}

// SetBatteryStatus sets BatteryStatus value.
func (m *DeviceAuxBatteryInfo) SetBatteryStatus(v typedef.BatteryStatus) *DeviceAuxBatteryInfo {
	m.BatteryStatus = v
	return m
}

// SetBatteryIdentifier sets BatteryIdentifier value.
func (m *DeviceAuxBatteryInfo) SetBatteryIdentifier(v uint8) *DeviceAuxBatteryInfo {
	m.BatteryIdentifier = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *DeviceAuxBatteryInfo) SetUnknownFields(unknownFields ...proto.Field) *DeviceAuxBatteryInfo {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *DeviceAuxBatteryInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *DeviceAuxBatteryInfo {
	m.DeveloperFields = developerFields
	return m
}
