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
	"time"
)

// BloodPressure is a BloodPressure message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type BloodPressure struct {
	Timestamp            time.Time            // Units: s
	SystolicPressure     uint16               // Units: mmHg
	DiastolicPressure    uint16               // Units: mmHg
	MeanArterialPressure uint16               // Units: mmHg
	Map3SampleMean       uint16               // Units: mmHg
	MapMorningValues     uint16               // Units: mmHg
	MapEveningValues     uint16               // Units: mmHg
	UserProfileIndex     typedef.MessageIndex // Associates this blood pressure message to a user. This corresponds to the index of the user profile message in the blood pressure file.
	HeartRate            uint8                // Units: bpm
	HeartRateType        typedef.HrType
	Status               typedef.BpStatus

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewBloodPressure creates new BloodPressure struct based on given mesg.
// If mesg is nil, it will return BloodPressure with all fields being set to its corresponding invalid value.
func NewBloodPressure(mesg *proto.Message) *BloodPressure {
	m := new(BloodPressure)
	m.Reset(mesg)
	return m
}

// Reset resets all BloodPressure's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *BloodPressure) Reset(mesg *proto.Message) {
	var (
		vals            [254]proto.Value
		unknownFields   []proto.Field
		developerFields []proto.DeveloperField
	)

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

	*m = BloodPressure{
		Timestamp:            datetime.ToTime(vals[253].Uint32()),
		SystolicPressure:     vals[0].Uint16(),
		DiastolicPressure:    vals[1].Uint16(),
		MeanArterialPressure: vals[2].Uint16(),
		Map3SampleMean:       vals[3].Uint16(),
		MapMorningValues:     vals[4].Uint16(),
		MapEveningValues:     vals[5].Uint16(),
		HeartRate:            vals[6].Uint8(),
		HeartRateType:        typedef.HrType(vals[7].Uint8()),
		Status:               typedef.BpStatus(vals[8].Uint8()),
		UserProfileIndex:     typedef.MessageIndex(vals[9].Uint16()),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts BloodPressure into proto.Message. If options is nil, default options will be used.
func (m *BloodPressure) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumBloodPressure}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.SystolicPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.SystolicPressure)
		fields = append(fields, field)
	}
	if m.DiastolicPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.DiastolicPressure)
		fields = append(fields, field)
	}
	if m.MeanArterialPressure != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint16(m.MeanArterialPressure)
		fields = append(fields, field)
	}
	if m.Map3SampleMean != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.Map3SampleMean)
		fields = append(fields, field)
	}
	if m.MapMorningValues != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.MapMorningValues)
		fields = append(fields, field)
	}
	if m.MapEveningValues != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint16(m.MapEveningValues)
		fields = append(fields, field)
	}
	if m.HeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.HeartRate)
		fields = append(fields, field)
	}
	if m.HeartRateType != typedef.HrTypeInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.HeartRateType))
		fields = append(fields, field)
	}
	if m.Status != typedef.BpStatusInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(byte(m.Status))
		fields = append(fields, field)
	}
	if m.UserProfileIndex != typedef.MessageIndexInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(uint16(m.UserProfileIndex))
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
func (m *BloodPressure) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *BloodPressure) SetTimestamp(v time.Time) *BloodPressure {
	m.Timestamp = v
	return m
}

// SetSystolicPressure sets SystolicPressure value.
//
// Units: mmHg
func (m *BloodPressure) SetSystolicPressure(v uint16) *BloodPressure {
	m.SystolicPressure = v
	return m
}

// SetDiastolicPressure sets DiastolicPressure value.
//
// Units: mmHg
func (m *BloodPressure) SetDiastolicPressure(v uint16) *BloodPressure {
	m.DiastolicPressure = v
	return m
}

// SetMeanArterialPressure sets MeanArterialPressure value.
//
// Units: mmHg
func (m *BloodPressure) SetMeanArterialPressure(v uint16) *BloodPressure {
	m.MeanArterialPressure = v
	return m
}

// SetMap3SampleMean sets Map3SampleMean value.
//
// Units: mmHg
func (m *BloodPressure) SetMap3SampleMean(v uint16) *BloodPressure {
	m.Map3SampleMean = v
	return m
}

// SetMapMorningValues sets MapMorningValues value.
//
// Units: mmHg
func (m *BloodPressure) SetMapMorningValues(v uint16) *BloodPressure {
	m.MapMorningValues = v
	return m
}

// SetMapEveningValues sets MapEveningValues value.
//
// Units: mmHg
func (m *BloodPressure) SetMapEveningValues(v uint16) *BloodPressure {
	m.MapEveningValues = v
	return m
}

// SetHeartRate sets HeartRate value.
//
// Units: bpm
func (m *BloodPressure) SetHeartRate(v uint8) *BloodPressure {
	m.HeartRate = v
	return m
}

// SetHeartRateType sets HeartRateType value.
func (m *BloodPressure) SetHeartRateType(v typedef.HrType) *BloodPressure {
	m.HeartRateType = v
	return m
}

// SetStatus sets Status value.
func (m *BloodPressure) SetStatus(v typedef.BpStatus) *BloodPressure {
	m.Status = v
	return m
}

// SetUserProfileIndex sets UserProfileIndex value.
//
// Associates this blood pressure message to a user. This corresponds to the index of the user profile message in the blood pressure file.
func (m *BloodPressure) SetUserProfileIndex(v typedef.MessageIndex) *BloodPressure {
	m.UserProfileIndex = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *BloodPressure) SetUnknownFields(unknownFields ...proto.Field) *BloodPressure {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *BloodPressure) SetDeveloperFields(developerFields ...proto.DeveloperField) *BloodPressure {
	m.DeveloperFields = developerFields
	return m
}
