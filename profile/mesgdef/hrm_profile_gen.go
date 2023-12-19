// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// HrmProfile is a HrmProfile message.
type HrmProfile struct {
	MessageIndex      typedef.MessageIndex
	Enabled           bool
	HrmAntId          uint16
	LogHrv            bool
	HrmAntIdTransType uint8

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewHrmProfile creates new HrmProfile struct based on given mesg.
// If mesg is nil, it will return HrmProfile with all fields being set to its corresponding invalid value.
func NewHrmProfile(mesg *proto.Message) *HrmProfile {
	vals := [255]any{}

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

	return &HrmProfile{
		MessageIndex:      typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Enabled:           typeconv.ToBool[bool](vals[0]),
		HrmAntId:          typeconv.ToUint16z[uint16](vals[1]),
		LogHrv:            typeconv.ToBool[bool](vals[2]),
		HrmAntIdTransType: typeconv.ToUint8z[uint8](vals[3]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts HrmProfile into proto.Message.
func (m *HrmProfile) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumHrmProfile)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Enabled != false {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.Enabled
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint16z[uint16](m.HrmAntId) != basetype.Uint16zInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = typeconv.ToUint16z[uint16](m.HrmAntId)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.LogHrv != false {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.LogHrv
		mesg.Fields = append(mesg.Fields, field)
	}
	if typeconv.ToUint8z[uint8](m.HrmAntIdTransType) != basetype.Uint8zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = typeconv.ToUint8z[uint8](m.HrmAntIdTransType)
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of HrmProfile's valid fields.
func (m *HrmProfile) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if m.Enabled != false {
		size++
	}
	if typeconv.ToUint16z[uint16](m.HrmAntId) != basetype.Uint16zInvalid {
		size++
	}
	if m.LogHrv != false {
		size++
	}
	if typeconv.ToUint8z[uint8](m.HrmAntIdTransType) != basetype.Uint8zInvalid {
		size++
	}
	return size
}

// SetMessageIndex sets HrmProfile value.
func (m *HrmProfile) SetMessageIndex(v typedef.MessageIndex) *HrmProfile {
	m.MessageIndex = v
	return m
}

// SetEnabled sets HrmProfile value.
func (m *HrmProfile) SetEnabled(v bool) *HrmProfile {
	m.Enabled = v
	return m
}

// SetHrmAntId sets HrmProfile value.
func (m *HrmProfile) SetHrmAntId(v uint16) *HrmProfile {
	m.HrmAntId = v
	return m
}

// SetLogHrv sets HrmProfile value.
func (m *HrmProfile) SetLogHrv(v bool) *HrmProfile {
	m.LogHrv = v
	return m
}

// SetHrmAntIdTransType sets HrmProfile value.
func (m *HrmProfile) SetHrmAntIdTransType(v uint8) *HrmProfile {
	m.HrmAntIdTransType = v
	return m
}

// SetDeveloperFields HrmProfile's DeveloperFields.
func (m *HrmProfile) SetDeveloperFields(developerFields ...proto.DeveloperField) *HrmProfile {
	m.DeveloperFields = developerFields
	return m
}
