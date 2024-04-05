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

// Hr is a Hr message.
type Hr struct {
	Timestamp           time.Time
	FilteredBpm         []uint8  // Array: [N]; Units: bpm
	EventTimestamp      []uint32 // Array: [N]; Scale: 1024; Units: s
	EventTimestamp12    []byte   // Array: [N]; Units: s
	FractionalTimestamp uint16   // Scale: 32768; Units: s
	Time256             uint8    // Scale: 256; Units: s

	IsExpandedFields [10]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHr creates new Hr struct based on given mesg.
// If mesg is nil, it will return Hr with all fields being set to its corresponding invalid value.
func NewHr(mesg *proto.Message) *Hr {
	vals := [254]proto.Value{}
	isExpandedFields := [10]bool{}

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

	return &Hr{
		Timestamp:           datetime.ToTime(vals[253].Uint32()),
		FilteredBpm:         vals[6].SliceUint8(),
		EventTimestamp:      vals[9].SliceUint32(),
		EventTimestamp12:    vals[10].SliceUint8(),
		FractionalTimestamp: vals[0].Uint16(),
		Time256:             vals[1].Uint8(),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts Hr into proto.Message. If options is nil, default options will be used.
func (m *Hr) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumHr}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.FilteredBpm != nil {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.SliceUint8(m.FilteredBpm)
		fields = append(fields, field)
	}
	if m.EventTimestamp != nil && ((m.IsExpandedFields[9] && options.IncludeExpandedFields) || !m.IsExpandedFields[9]) {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.SliceUint32(m.EventTimestamp)
		field.IsExpandedField = m.IsExpandedFields[9]
		fields = append(fields, field)
	}
	if m.EventTimestamp12 != nil {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.SliceUint8(m.EventTimestamp12)
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid && ((m.IsExpandedFields[0] && options.IncludeExpandedFields) || !m.IsExpandedFields[0]) {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.FractionalTimestamp)
		field.IsExpandedField = m.IsExpandedFields[0]
		fields = append(fields, field)
	}
	if m.Time256 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.Time256)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// EventTimestampScaled return EventTimestamp in its scaled value [Array: [N]; Scale: 1024; Units: s].
//
// If EventTimestamp value is invalid, nil will be returned.
func (m *Hr) EventTimestampScaled() []float64 {
	if m.EventTimestamp == nil {
		return nil
	}
	return scaleoffset.ApplySlice(m.EventTimestamp, 1024, 0)
}

// FractionalTimestampScaled return FractionalTimestamp in its scaled value [Scale: 32768; Units: s].
//
// If FractionalTimestamp value is invalid, float64 invalid value will be returned.
func (m *Hr) FractionalTimestampScaled() float64 {
	if m.FractionalTimestamp == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.FractionalTimestamp, 32768, 0)
}

// Time256Scaled return Time256 in its scaled value [Scale: 256; Units: s].
//
// If Time256 value is invalid, float64 invalid value will be returned.
func (m *Hr) Time256Scaled() float64 {
	if m.Time256 == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.Time256, 256, 0)
}

// SetTimestamp sets Hr value.
func (m *Hr) SetTimestamp(v time.Time) *Hr {
	m.Timestamp = v
	return m
}

// SetFilteredBpm sets Hr value.
//
// Array: [N]; Units: bpm
func (m *Hr) SetFilteredBpm(v []uint8) *Hr {
	m.FilteredBpm = v
	return m
}

// SetEventTimestamp sets Hr value.
//
// Array: [N]; Scale: 1024; Units: s
func (m *Hr) SetEventTimestamp(v []uint32) *Hr {
	m.EventTimestamp = v
	return m
}

// SetEventTimestamp12 sets Hr value.
//
// Array: [N]; Units: s
func (m *Hr) SetEventTimestamp12(v []byte) *Hr {
	m.EventTimestamp12 = v
	return m
}

// SetFractionalTimestamp sets Hr value.
//
// Scale: 32768; Units: s
func (m *Hr) SetFractionalTimestamp(v uint16) *Hr {
	m.FractionalTimestamp = v
	return m
}

// SetTime256 sets Hr value.
//
// Scale: 256; Units: s
func (m *Hr) SetTime256(v uint8) *Hr {
	m.Time256 = v
	return m
}

// SetDeveloperFields Hr's DeveloperFields.
func (m *Hr) SetDeveloperFields(developerFields ...proto.DeveloperField) *Hr {
	m.DeveloperFields = developerFields
	return m
}
