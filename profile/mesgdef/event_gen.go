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

// Event is a Event message.
type Event struct {
	Timestamp                   time.Time // Units: s;
	Event                       typedef.Event
	EventType                   typedef.EventType
	Data16                      uint16
	Data                        uint32
	EventGroup                  uint8
	Score                       uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	OpponentScore               uint16 // Do not populate directly. Autogenerated by decoder for sport_point subfield components
	FrontGearNum                uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Front gear number. 1 is innermost.
	FrontGear                   uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of front teeth.
	RearGearNum                 uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Rear gear number. 1 is innermost.
	RearGear                    uint8  // Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of rear teeth.
	DeviceIndex                 typedef.DeviceIndex
	ActivityType                typedef.ActivityType         // Activity Type associated with an auto_activity_detect event
	StartTimestamp              time.Time                    // Units: s; Timestamp of when the event started
	RadarThreatLevelMax         typedef.RadarThreatLevelType // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatCount            uint8                        // Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
	RadarThreatAvgApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
	RadarThreatMaxApproachSpeed uint8                        // Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField

	IsExpandedFields [25]bool // Used for tracking expanded fields, field.Num as index.
}

// NewEvent creates new Event struct based on given mesg.
// If mesg is nil, it will return Event with all fields being set to its corresponding invalid value.
func NewEvent(mesg *proto.Message) *Event {
	vals := [254]any{}
	isExpandedFields := [25]bool{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			if mesg.Fields[i].Num < byte(len(isExpandedFields)) {
				isExpandedFields[mesg.Fields[i].Num] = mesg.Fields[i].IsExpandedField
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &Event{
		Timestamp:                   datetime.ToTime(vals[253]),
		Event:                       typeconv.ToEnum[typedef.Event](vals[0]),
		EventType:                   typeconv.ToEnum[typedef.EventType](vals[1]),
		Data16:                      typeconv.ToUint16[uint16](vals[2]),
		Data:                        typeconv.ToUint32[uint32](vals[3]),
		EventGroup:                  typeconv.ToUint8[uint8](vals[4]),
		Score:                       typeconv.ToUint16[uint16](vals[7]),
		OpponentScore:               typeconv.ToUint16[uint16](vals[8]),
		FrontGearNum:                typeconv.ToUint8z[uint8](vals[9]),
		FrontGear:                   typeconv.ToUint8z[uint8](vals[10]),
		RearGearNum:                 typeconv.ToUint8z[uint8](vals[11]),
		RearGear:                    typeconv.ToUint8z[uint8](vals[12]),
		DeviceIndex:                 typeconv.ToUint8[typedef.DeviceIndex](vals[13]),
		ActivityType:                typeconv.ToEnum[typedef.ActivityType](vals[14]),
		StartTimestamp:              datetime.ToTime(vals[15]),
		RadarThreatLevelMax:         typeconv.ToEnum[typedef.RadarThreatLevelType](vals[21]),
		RadarThreatCount:            typeconv.ToUint8[uint8](vals[22]),
		RadarThreatAvgApproachSpeed: typeconv.ToUint8[uint8](vals[23]),
		RadarThreatMaxApproachSpeed: typeconv.ToUint8[uint8](vals[24]),

		DeveloperFields: developerFields,

		IsExpandedFields: isExpandedFields,
	}
}

// ToMesg converts Event into proto.Message.
func (m *Event) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumEvent)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.Event) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.Event)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.EventType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToEnum[byte](m.EventType)
		fields = append(fields, field)
	}
	if m.Data16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Data16
		fields = append(fields, field)
	}
	if m.Data != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.Data
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if m.EventGroup != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.EventGroup
		fields = append(fields, field)
	}
	if m.Score != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = m.Score
		field.IsExpandedField = m.IsExpandedFields[7]
		fields = append(fields, field)
	}
	if m.OpponentScore != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.OpponentScore
		field.IsExpandedField = m.IsExpandedFields[8]
		fields = append(fields, field)
	}
	if typeconv.ToUint8z[uint8](m.FrontGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = typeconv.ToUint8z[uint8](m.FrontGearNum)
		field.IsExpandedField = m.IsExpandedFields[9]
		fields = append(fields, field)
	}
	if typeconv.ToUint8z[uint8](m.FrontGear) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = typeconv.ToUint8z[uint8](m.FrontGear)
		field.IsExpandedField = m.IsExpandedFields[10]
		fields = append(fields, field)
	}
	if typeconv.ToUint8z[uint8](m.RearGearNum) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = typeconv.ToUint8z[uint8](m.RearGearNum)
		field.IsExpandedField = m.IsExpandedFields[11]
		fields = append(fields, field)
	}
	if typeconv.ToUint8z[uint8](m.RearGear) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = typeconv.ToUint8z[uint8](m.RearGear)
		field.IsExpandedField = m.IsExpandedFields[12]
		fields = append(fields, field)
	}
	if typeconv.ToUint8[uint8](m.DeviceIndex) != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = typeconv.ToUint8[uint8](m.DeviceIndex)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.ActivityType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = typeconv.ToEnum[byte](m.ActivityType)
		fields = append(fields, field)
	}
	if datetime.ToUint32(m.StartTimestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = datetime.ToUint32(m.StartTimestamp)
		fields = append(fields, field)
	}
	if typeconv.ToEnum[byte](m.RadarThreatLevelMax) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 21)
		field.Value = typeconv.ToEnum[byte](m.RadarThreatLevelMax)
		field.IsExpandedField = m.IsExpandedFields[21]
		fields = append(fields, field)
	}
	if m.RadarThreatCount != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = m.RadarThreatCount
		field.IsExpandedField = m.IsExpandedFields[22]
		fields = append(fields, field)
	}
	if m.RadarThreatAvgApproachSpeed != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = m.RadarThreatAvgApproachSpeed
		field.IsExpandedField = m.IsExpandedFields[23]
		fields = append(fields, field)
	}
	if m.RadarThreatMaxApproachSpeed != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = m.RadarThreatMaxApproachSpeed
		field.IsExpandedField = m.IsExpandedFields[24]
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets Event value.
//
// Units: s;
func (m *Event) SetTimestamp(v time.Time) *Event {
	m.Timestamp = v
	return m
}

// SetEvent sets Event value.
func (m *Event) SetEvent(v typedef.Event) *Event {
	m.Event = v
	return m
}

// SetEventType sets Event value.
func (m *Event) SetEventType(v typedef.EventType) *Event {
	m.EventType = v
	return m
}

// SetData16 sets Event value.
func (m *Event) SetData16(v uint16) *Event {
	m.Data16 = v
	return m
}

// SetData sets Event value.
func (m *Event) SetData(v uint32) *Event {
	m.Data = v
	return m
}

// SetEventGroup sets Event value.
func (m *Event) SetEventGroup(v uint8) *Event {
	m.EventGroup = v
	return m
}

// SetScore sets Event value.
//
// Do not populate directly. Autogenerated by decoder for sport_point subfield components
func (m *Event) SetScore(v uint16) *Event {
	m.Score = v
	return m
}

// SetOpponentScore sets Event value.
//
// Do not populate directly. Autogenerated by decoder for sport_point subfield components
func (m *Event) SetOpponentScore(v uint16) *Event {
	m.OpponentScore = v
	return m
}

// SetFrontGearNum sets Event value.
//
// Do not populate directly. Autogenerated by decoder for gear_change subfield components. Front gear number. 1 is innermost.
func (m *Event) SetFrontGearNum(v uint8) *Event {
	m.FrontGearNum = v
	return m
}

// SetFrontGear sets Event value.
//
// Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of front teeth.
func (m *Event) SetFrontGear(v uint8) *Event {
	m.FrontGear = v
	return m
}

// SetRearGearNum sets Event value.
//
// Do not populate directly. Autogenerated by decoder for gear_change subfield components. Rear gear number. 1 is innermost.
func (m *Event) SetRearGearNum(v uint8) *Event {
	m.RearGearNum = v
	return m
}

// SetRearGear sets Event value.
//
// Do not populate directly. Autogenerated by decoder for gear_change subfield components. Number of rear teeth.
func (m *Event) SetRearGear(v uint8) *Event {
	m.RearGear = v
	return m
}

// SetDeviceIndex sets Event value.
func (m *Event) SetDeviceIndex(v typedef.DeviceIndex) *Event {
	m.DeviceIndex = v
	return m
}

// SetActivityType sets Event value.
//
// Activity Type associated with an auto_activity_detect event
func (m *Event) SetActivityType(v typedef.ActivityType) *Event {
	m.ActivityType = v
	return m
}

// SetStartTimestamp sets Event value.
//
// Units: s; Timestamp of when the event started
func (m *Event) SetStartTimestamp(v time.Time) *Event {
	m.StartTimestamp = v
	return m
}

// SetRadarThreatLevelMax sets Event value.
//
// Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
func (m *Event) SetRadarThreatLevelMax(v typedef.RadarThreatLevelType) *Event {
	m.RadarThreatLevelMax = v
	return m
}

// SetRadarThreatCount sets Event value.
//
// Do not populate directly. Autogenerated by decoder for threat_alert subfield components.
func (m *Event) SetRadarThreatCount(v uint8) *Event {
	m.RadarThreatCount = v
	return m
}

// SetRadarThreatAvgApproachSpeed sets Event value.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatAvgApproachSpeed(v uint8) *Event {
	m.RadarThreatAvgApproachSpeed = v
	return m
}

// SetRadarThreatMaxApproachSpeed sets Event value.
//
// Scale: 10; Units: m/s; Do not populate directly. Autogenerated by decoder for radar_threat_alert subfield components
func (m *Event) SetRadarThreatMaxApproachSpeed(v uint8) *Event {
	m.RadarThreatMaxApproachSpeed = v
	return m
}

// SetDeveloperFields Event's DeveloperFields.
func (m *Event) SetDeveloperFields(developerFields ...proto.DeveloperField) *Event {
	m.DeveloperFields = developerFields
	return m
}
