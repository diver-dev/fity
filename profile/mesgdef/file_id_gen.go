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

// FileId is a FileId message.
type FileId struct {
	TimeCreated  time.Time // Only set for files that are can be created/erased.
	ProductName  string    // Optional free form string to indicate the devices name or model
	SerialNumber uint32
	Manufacturer typedef.Manufacturer
	Product      uint16
	Number       uint16 // Only set for files that are not created/erased.
	Type         typedef.File
}

// NewFileId creates new FileId struct based on given mesg.
// If mesg is nil, it will return FileId with all fields being set to its corresponding invalid value.
func NewFileId(mesg *proto.Message) *FileId {
	vals := [9]any{}

	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
	}

	return &FileId{
		TimeCreated:  datetime.ToTime(vals[4]),
		ProductName:  typeconv.ToString[string](vals[8]),
		SerialNumber: typeconv.ToUint32z[uint32](vals[3]),
		Manufacturer: typeconv.ToUint16[typedef.Manufacturer](vals[1]),
		Product:      typeconv.ToUint16[uint16](vals[2]),
		Number:       typeconv.ToUint16[uint16](vals[5]),
		Type:         typeconv.ToEnum[typedef.File](vals[0]),
	}
}

// ToMesg converts FileId into proto.Message.
func (m *FileId) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumFileId)

	if datetime.ToUint32(m.TimeCreated) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = datetime.ToUint32(m.TimeCreated)
		fields = append(fields, field)
	}
	if m.ProductName != basetype.StringInvalid && m.ProductName != "" {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = m.ProductName
		fields = append(fields, field)
	}
	if uint32(m.SerialNumber) != basetype.Uint32zInvalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = uint32(m.SerialNumber)
		fields = append(fields, field)
	}
	if uint16(m.Manufacturer) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = uint16(m.Manufacturer)
		fields = append(fields, field)
	}
	if m.Product != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.Product
		fields = append(fields, field)
	}
	if m.Number != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = m.Number
		fields = append(fields, field)
	}
	if byte(m.Type) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.Type)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	return mesg
}

// SetTimeCreated sets FileId value.
//
// Only set for files that are can be created/erased.
func (m *FileId) SetTimeCreated(v time.Time) *FileId {
	m.TimeCreated = v
	return m
}

// SetProductName sets FileId value.
//
// Optional free form string to indicate the devices name or model
func (m *FileId) SetProductName(v string) *FileId {
	m.ProductName = v
	return m
}

// SetSerialNumber sets FileId value.
func (m *FileId) SetSerialNumber(v uint32) *FileId {
	m.SerialNumber = v
	return m
}

// SetManufacturer sets FileId value.
func (m *FileId) SetManufacturer(v typedef.Manufacturer) *FileId {
	m.Manufacturer = v
	return m
}

// SetProduct sets FileId value.
func (m *FileId) SetProduct(v uint16) *FileId {
	m.Product = v
	return m
}

// SetNumber sets FileId value.
//
// Only set for files that are not created/erased.
func (m *FileId) SetNumber(v uint16) *FileId {
	m.Number = v
	return m
}

// SetType sets FileId value.
func (m *FileId) SetType(v typedef.File) *FileId {
	m.Type = v
	return m
}
