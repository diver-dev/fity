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

// SleepLevel is a SleepLevel message.
type SleepLevel struct {
	Timestamp  time.Time // Units: s;
	SleepLevel typedef.SleepLevel

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSleepLevel creates new SleepLevel struct based on given mesg.
// If mesg is nil, it will return SleepLevel with all fields being set to its corresponding invalid value.
func NewSleepLevel(mesg *proto.Message) *SleepLevel {
	vals := [254]any{}

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

	return &SleepLevel{
		Timestamp:  datetime.ToTime(vals[253]),
		SleepLevel: typeconv.ToEnum[typedef.SleepLevel](vals[0]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts SleepLevel into proto.Message.
func (m *SleepLevel) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumSleepLevel)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToEnum[byte](m.SleepLevel) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = typeconv.ToEnum[byte](m.SleepLevel)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of SleepLevel's valid fields.
func (m *SleepLevel) size() byte {
	var size byte
	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		size++
	}
	if typeconv.ToEnum[byte](m.SleepLevel) != basetype.EnumInvalid {
		size++
	}
	return size
}

// SetTimestamp sets SleepLevel value.
//
// Units: s;
func (m *SleepLevel) SetTimestamp(v time.Time) *SleepLevel {
	m.Timestamp = v
	return m
}

// SetSleepLevel sets SleepLevel value.
func (m *SleepLevel) SetSleepLevel(v typedef.SleepLevel) *SleepLevel {
	m.SleepLevel = v
	return m
}

// SetDeveloperFields SleepLevel's DeveloperFields.
func (m *SleepLevel) SetDeveloperFields(developerFields ...proto.DeveloperField) *SleepLevel {
	m.DeveloperFields = developerFields
	return m
}
