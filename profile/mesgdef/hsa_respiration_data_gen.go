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

// HsaRespirationData is a HsaRespirationData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaRespirationData struct {
	Timestamp          time.Time // Units: s
	RespirationRate    []int16   // Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHsaRespirationData creates new HsaRespirationData struct based on given mesg.
// If mesg is nil, it will return HsaRespirationData with all fields being set to its corresponding invalid value.
func NewHsaRespirationData(mesg *proto.Message) *HsaRespirationData {
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
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &HsaRespirationData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		ProcessingInterval: vals[0].Uint16(),
		RespirationRate:    vals[1].SliceInt16(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaRespirationData into proto.Message. If options is nil, default options will be used.
func (m *HsaRespirationData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHsaRespirationData}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}
	if m.RespirationRate != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceInt16(m.RespirationRate)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaRespirationData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// RespirationRateScaled return RespirationRate in its scaled value.
// If RespirationRate value is invalid, nil will be returned.
//
// Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
func (m *HsaRespirationData) RespirationRateScaled() []float64 {
	if m.RespirationRate == nil {
		return nil
	}
	var vals = make([]float64, len(m.RespirationRate))
	for i := range m.RespirationRate {
		if m.RespirationRate[i] == basetype.Sint16Invalid {
			vals[i] = math.Float64frombits(basetype.Float64Invalid)
			continue
		}
		vals[i] = float64(m.RespirationRate[i])/100 - 0
	}
	return vals
}

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaRespirationData) SetTimestamp(v time.Time) *HsaRespirationData {
	m.Timestamp = v
	return m
}

// SetProcessingInterval sets ProcessingInterval value.
//
// Units: s; Processing interval length in seconds
func (m *HsaRespirationData) SetProcessingInterval(v uint16) *HsaRespirationData {
	m.ProcessingInterval = v
	return m
}

// SetRespirationRate sets RespirationRate value.
//
// Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
func (m *HsaRespirationData) SetRespirationRate(v []int16) *HsaRespirationData {
	m.RespirationRate = v
	return m
}

// SetRespirationRateScaled is similar to SetRespirationRate except it accepts a scaled value.
// This method automatically converts the given value to its []int16 form, discarding any applied scale and offset.
//
// Array: [N]; Scale: 100; Units: breaths/min; Breaths * 100 /min -300 indicates invalid -200 indicates large motion -100 indicates off wrist
func (m *HsaRespirationData) SetRespirationRateScaled(vs []float64) *HsaRespirationData {
	if vs == nil {
		m.RespirationRate = nil
		return m
	}
	m.RespirationRate = make([]int16, len(vs))
	for i := range vs {
		unscaled := (vs[i] + 0) * 100
		if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
			m.RespirationRate[i] = int16(basetype.Sint16Invalid)
			continue
		}
		m.RespirationRate[i] = int16(unscaled)
	}
	return m
}

// SetDeveloperFields HsaRespirationData's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *HsaRespirationData) SetUnknownFields(unknownFields ...proto.Field) *HsaRespirationData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields HsaRespirationData's DeveloperFields.
func (m *HsaRespirationData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaRespirationData {
	m.DeveloperFields = developerFields
	return m
}
