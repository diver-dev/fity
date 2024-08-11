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
	"time"
)

// HsaConfigurationData is a HsaConfigurationData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaConfigurationData struct {
	Timestamp time.Time // Units: s; Encoded configuration data
	Data      []byte    // Array: [N]
	DataSize  uint8     // Size in bytes of data field

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHsaConfigurationData creates new HsaConfigurationData struct based on given mesg.
// If mesg is nil, it will return HsaConfigurationData with all fields being set to its corresponding invalid value.
func NewHsaConfigurationData(mesg *proto.Message) *HsaConfigurationData {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &HsaConfigurationData{
		Timestamp: datetime.ToTime(vals[253].Uint32()),
		Data:      vals[0].SliceUint8(),
		DataSize:  vals[1].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaConfigurationData into proto.Message. If options is nil, default options will be used.
func (m *HsaConfigurationData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[255]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHsaConfigurationData}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.SliceUint8(m.Data)
		fields = append(fields, field)
	}
	if m.DataSize != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.DataSize)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaConfigurationData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s; Encoded configuration data
func (m *HsaConfigurationData) SetTimestamp(v time.Time) *HsaConfigurationData {
	m.Timestamp = v
	return m
}

// SetData sets Data value.
//
// Array: [N]
func (m *HsaConfigurationData) SetData(v []byte) *HsaConfigurationData {
	m.Data = v
	return m
}

// SetDataSize sets DataSize value.
//
// Size in bytes of data field
func (m *HsaConfigurationData) SetDataSize(v uint8) *HsaConfigurationData {
	m.DataSize = v
	return m
}

// SetDeveloperFields HsaConfigurationData's DeveloperFields.
func (m *HsaConfigurationData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaConfigurationData {
	m.DeveloperFields = developerFields
	return m
}
