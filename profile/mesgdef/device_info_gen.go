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

// DeviceInfo is a DeviceInfo message.
type DeviceInfo struct {
	Timestamp           time.Time // Units: s
	Descriptor          string    // Used to describe the sensor or location
	ProductName         string    // Optional free form string to indicate the devices name or model
	SerialNumber        uint32
	CumOperatingTime    uint32 // Units: s; Reset by new battery or charge.
	Manufacturer        typedef.Manufacturer
	Product             uint16
	SoftwareVersion     uint16 // Scale: 100
	BatteryVoltage      uint16 // Scale: 256; Units: V
	AntDeviceNumber     uint16
	DeviceIndex         typedef.DeviceIndex
	DeviceType          uint8
	HardwareVersion     uint8
	BatteryStatus       typedef.BatteryStatus
	SensorPosition      typedef.BodyLocation // Indicates the location of the sensor
	AntTransmissionType uint8
	AntNetwork          typedef.AntNetwork
	SourceType          typedef.SourceType
	BatteryLevel        uint8 // Units: %

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDeviceInfo creates new DeviceInfo struct based on given mesg.
// If mesg is nil, it will return DeviceInfo with all fields being set to its corresponding invalid value.
func NewDeviceInfo(mesg *proto.Message) *DeviceInfo {
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

	return &DeviceInfo{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		Descriptor:          vals[19].String(),
		ProductName:         vals[27].String(),
		SerialNumber:        vals[3].Uint32z(),
		CumOperatingTime:    vals[7].Uint32(),
		Manufacturer:        typedef.Manufacturer(vals[2].Uint16()),
		Product:             vals[4].Uint16(),
		SoftwareVersion:     vals[5].Uint16(),
		BatteryVoltage:      vals[10].Uint16(),
		AntDeviceNumber:     vals[21].Uint16z(),
		DeviceIndex:         typedef.DeviceIndex(vals[0].Uint8()),
		DeviceType:          vals[1].Uint8(),
		HardwareVersion:     vals[6].Uint8(),
		BatteryStatus:       typedef.BatteryStatus(vals[11].Uint8()),
		SensorPosition:      typedef.BodyLocation(vals[18].Uint8()),
		AntTransmissionType: vals[20].Uint8z(),
		AntNetwork:          typedef.AntNetwork(vals[22].Uint8()),
		SourceType:          typedef.SourceType(vals[25].Uint8()),
		BatteryLevel:        vals[32].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DeviceInfo into proto.Message. If options is nil, default options will be used.
func (m *DeviceInfo) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumDeviceInfo}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.Descriptor != basetype.StringInvalid && m.Descriptor != "" {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.String(m.Descriptor)
		fields = append(fields, field)
	}
	if m.ProductName != basetype.StringInvalid && m.ProductName != "" {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = proto.String(m.ProductName)
		fields = append(fields, field)
	}
	if uint32(m.SerialNumber) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if m.CumOperatingTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint32(m.CumOperatingTime)
		fields = append(fields, field)
	}
	if uint16(m.Manufacturer) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(uint16(m.Manufacturer))
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.Product)
		fields = append(fields, field)
	}
	if m.SoftwareVersion != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.SoftwareVersion)
		fields = append(fields, field)
	}
	if m.BatteryVoltage != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(m.BatteryVoltage)
		fields = append(fields, field)
	}
	if uint16(m.AntDeviceNumber) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = proto.Uint16(m.AntDeviceNumber)
		fields = append(fields, field)
	}
	if uint8(m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if m.DeviceType != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.DeviceType)
		fields = append(fields, field)
	}
	if m.HardwareVersion != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.HardwareVersion)
		fields = append(fields, field)
	}
	if uint8(m.BatteryStatus) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint8(uint8(m.BatteryStatus))
		fields = append(fields, field)
	}
	if byte(m.SensorPosition) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 18)
		field.Value = proto.Uint8(byte(m.SensorPosition))
		fields = append(fields, field)
	}
	if uint8(m.AntTransmissionType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 20)
		field.Value = proto.Uint8(m.AntTransmissionType)
		fields = append(fields, field)
	}
	if byte(m.AntNetwork) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint8(byte(m.AntNetwork))
		fields = append(fields, field)
	}
	if byte(m.SourceType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint8(byte(m.SourceType))
		fields = append(fields, field)
	}
	if m.BatteryLevel != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = proto.Uint8(m.BatteryLevel)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SoftwareVersionScaled return SoftwareVersion in its scaled value [Scale: 100].
//
// If SoftwareVersion value is invalid, float64 invalid value will be returned.
func (m *DeviceInfo) SoftwareVersionScaled() float64 {
	if m.SoftwareVersion == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.SoftwareVersion, 100, 0)
}

// BatteryVoltageScaled return BatteryVoltage in its scaled value [Scale: 256; Units: V].
//
// If BatteryVoltage value is invalid, float64 invalid value will be returned.
func (m *DeviceInfo) BatteryVoltageScaled() float64 {
	if m.BatteryVoltage == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BatteryVoltage, 256, 0)
}

// SetTimestamp sets DeviceInfo value.
//
// Units: s
func (m *DeviceInfo) SetTimestamp(v time.Time) *DeviceInfo {
	m.Timestamp = v
	return m
}

// SetDescriptor sets DeviceInfo value.
//
// Used to describe the sensor or location
func (m *DeviceInfo) SetDescriptor(v string) *DeviceInfo {
	m.Descriptor = v
	return m
}

// SetProductName sets DeviceInfo value.
//
// Optional free form string to indicate the devices name or model
func (m *DeviceInfo) SetProductName(v string) *DeviceInfo {
	m.ProductName = v
	return m
}

// SetSerialNumber sets DeviceInfo value.
func (m *DeviceInfo) SetSerialNumber(v uint32) *DeviceInfo {
	m.SerialNumber = v
	return m
}

// SetCumOperatingTime sets DeviceInfo value.
//
// Units: s; Reset by new battery or charge.
func (m *DeviceInfo) SetCumOperatingTime(v uint32) *DeviceInfo {
	m.CumOperatingTime = v
	return m
}

// SetManufacturer sets DeviceInfo value.
func (m *DeviceInfo) SetManufacturer(v typedef.Manufacturer) *DeviceInfo {
	m.Manufacturer = v
	return m
}

// SetProduct sets DeviceInfo value.
func (m *DeviceInfo) SetProduct(v uint16) *DeviceInfo {
	m.Product = v
	return m
}

// SetSoftwareVersion sets DeviceInfo value.
//
// Scale: 100
func (m *DeviceInfo) SetSoftwareVersion(v uint16) *DeviceInfo {
	m.SoftwareVersion = v
	return m
}

// SetBatteryVoltage sets DeviceInfo value.
//
// Scale: 256; Units: V
func (m *DeviceInfo) SetBatteryVoltage(v uint16) *DeviceInfo {
	m.BatteryVoltage = v
	return m
}

// SetAntDeviceNumber sets DeviceInfo value.
func (m *DeviceInfo) SetAntDeviceNumber(v uint16) *DeviceInfo {
	m.AntDeviceNumber = v
	return m
}

// SetDeviceIndex sets DeviceInfo value.
func (m *DeviceInfo) SetDeviceIndex(v typedef.DeviceIndex) *DeviceInfo {
	m.DeviceIndex = v
	return m
}

// SetDeviceType sets DeviceInfo value.
func (m *DeviceInfo) SetDeviceType(v uint8) *DeviceInfo {
	m.DeviceType = v
	return m
}

// SetHardwareVersion sets DeviceInfo value.
func (m *DeviceInfo) SetHardwareVersion(v uint8) *DeviceInfo {
	m.HardwareVersion = v
	return m
}

// SetBatteryStatus sets DeviceInfo value.
func (m *DeviceInfo) SetBatteryStatus(v typedef.BatteryStatus) *DeviceInfo {
	m.BatteryStatus = v
	return m
}

// SetSensorPosition sets DeviceInfo value.
//
// Indicates the location of the sensor
func (m *DeviceInfo) SetSensorPosition(v typedef.BodyLocation) *DeviceInfo {
	m.SensorPosition = v
	return m
}

// SetAntTransmissionType sets DeviceInfo value.
func (m *DeviceInfo) SetAntTransmissionType(v uint8) *DeviceInfo {
	m.AntTransmissionType = v
	return m
}

// SetAntNetwork sets DeviceInfo value.
func (m *DeviceInfo) SetAntNetwork(v typedef.AntNetwork) *DeviceInfo {
	m.AntNetwork = v
	return m
}

// SetSourceType sets DeviceInfo value.
func (m *DeviceInfo) SetSourceType(v typedef.SourceType) *DeviceInfo {
	m.SourceType = v
	return m
}

// SetBatteryLevel sets DeviceInfo value.
//
// Units: %
func (m *DeviceInfo) SetBatteryLevel(v uint8) *DeviceInfo {
	m.BatteryLevel = v
	return m
}

// SetDeveloperFields DeviceInfo's DeveloperFields.
func (m *DeviceInfo) SetDeveloperFields(developerFields ...proto.DeveloperField) *DeviceInfo {
	m.DeveloperFields = developerFields
	return m
}
