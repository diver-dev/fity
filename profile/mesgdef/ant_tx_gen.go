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

// AntTx is a AntTx message.
type AntTx struct {
	Timestamp           time.Time // Units: s;
	FractionalTimestamp uint16    // Scale: 32768; Units: s;
	MesgId              byte
	MesgData            []byte // Array: [N];
	ChannelNumber       uint8
	Data                []byte // Array: [N];

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField

	IsExpandedFields [5]bool // Used for tracking expanded fields, field.Num as index.
}

// NewAntTx creates new AntTx struct based on given mesg.
// If mesg is nil, it will return AntTx with all fields being set to its corresponding invalid value.
func NewAntTx(mesg *proto.Message) *AntTx {
	vals := [254]any{}
	isExpandedFields := [5]bool{}

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

	return &AntTx{
		Timestamp:           datetime.ToTime(vals[253]),
		FractionalTimestamp: typeconv.ToUint16[uint16](vals[0]),
		MesgId:              typeconv.ToByte[byte](vals[1]),
		MesgData:            typeconv.ToSliceByte[byte](vals[2]),
		ChannelNumber:       typeconv.ToUint8[uint8](vals[3]),
		Data:                typeconv.ToSliceByte[byte](vals[4]),

		DeveloperFields: developerFields,

		IsExpandedFields: isExpandedFields,
	}
}

// ToMesg converts AntTx into proto.Message.
func (m *AntTx) ToMesg(fac Factory) proto.Message {
	fieldsPtr := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsPtr)

	fields := (*fieldsPtr)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumAntTx)

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = datetime.ToUint32(m.Timestamp)
		fields = append(fields, field)
	}
	if m.FractionalTimestamp != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.FractionalTimestamp
		fields = append(fields, field)
	}
	if m.MesgId != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.MesgId
		fields = append(fields, field)
	}
	if m.MesgData != nil {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.MesgData
		fields = append(fields, field)
	}
	if m.ChannelNumber != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.ChannelNumber
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if m.Data != nil {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.Data
		field.IsExpandedField = m.IsExpandedFields[4]
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTimestamp sets AntTx value.
//
// Units: s;
func (m *AntTx) SetTimestamp(v time.Time) *AntTx {
	m.Timestamp = v
	return m
}

// SetFractionalTimestamp sets AntTx value.
//
// Scale: 32768; Units: s;
func (m *AntTx) SetFractionalTimestamp(v uint16) *AntTx {
	m.FractionalTimestamp = v
	return m
}

// SetMesgId sets AntTx value.
func (m *AntTx) SetMesgId(v byte) *AntTx {
	m.MesgId = v
	return m
}

// SetMesgData sets AntTx value.
//
// Array: [N];
func (m *AntTx) SetMesgData(v []byte) *AntTx {
	m.MesgData = v
	return m
}

// SetChannelNumber sets AntTx value.
func (m *AntTx) SetChannelNumber(v uint8) *AntTx {
	m.ChannelNumber = v
	return m
}

// SetData sets AntTx value.
//
// Array: [N];
func (m *AntTx) SetData(v []byte) *AntTx {
	m.Data = v
	return m
}

// SetDeveloperFields AntTx's DeveloperFields.
func (m *AntTx) SetDeveloperFields(developerFields ...proto.DeveloperField) *AntTx {
	m.DeveloperFields = developerFields
	return m
}
