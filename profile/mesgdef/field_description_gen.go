// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// FieldDescription is a FieldDescription message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type FieldDescription struct {
	FieldName             []string // Array: [N]
	Components            string
	Units                 []string // Array: [N]
	Bits                  string
	Accumulate            string
	FitBaseUnitId         typedef.FitBaseUnit
	NativeMesgNum         typedef.MesgNum
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	FitBaseTypeId         basetype.BaseType
	Array                 uint8
	Scale                 uint8
	Offset                int8
	NativeFieldNum        uint8

	UnknownFields []proto.Field // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
}

// NewFieldDescription creates new FieldDescription struct based on given mesg.
// If mesg is nil, it will return FieldDescription with all fields being set to its corresponding invalid value.
func NewFieldDescription(mesg *proto.Message) *FieldDescription {
	m := new(FieldDescription)
	m.Reset(mesg)
	return m
}

// Reset resets all FieldDescription's fields based on given mesg.
// If mesg is nil, all fields will be set to its corresponding invalid value.
func (m *FieldDescription) Reset(mesg *proto.Message) {
	var (
		vals          [16]proto.Value
		unknownFields []proto.Field
	)

	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 15 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
	}

	*m = FieldDescription{
		DeveloperDataIndex:    vals[0].Uint8(),
		FieldDefinitionNumber: vals[1].Uint8(),
		FitBaseTypeId:         basetype.BaseType((vals[2]).Uint8()),
		FieldName:             vals[3].SliceString(),
		Array:                 vals[4].Uint8(),
		Components:            vals[5].String(),
		Scale:                 vals[6].Uint8(),
		Offset:                vals[7].Int8(),
		Units:                 vals[8].SliceString(),
		Bits:                  vals[9].String(),
		Accumulate:            vals[10].String(),
		FitBaseUnitId:         typedef.FitBaseUnit(vals[13].Uint16()),
		NativeMesgNum:         typedef.MesgNum(vals[14].Uint16()),
		NativeFieldNum:        vals[15].Uint8(),

		UnknownFields: unknownFields,
	}
}

// ToMesg converts FieldDescription into proto.Message. If options is nil, default options will be used.
func (m *FieldDescription) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumFieldDescription}

	if m.DeveloperDataIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.DeveloperDataIndex)
		fields = append(fields, field)
	}
	if m.FieldDefinitionNumber != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.FieldDefinitionNumber)
		fields = append(fields, field)
	}
	if m.FitBaseTypeId != 255 {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(uint8(m.FitBaseTypeId))
		fields = append(fields, field)
	}
	if m.FieldName != nil {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.SliceString(m.FieldName)
		fields = append(fields, field)
	}
	if m.Array != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(m.Array)
		fields = append(fields, field)
	}
	if m.Components != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.String(m.Components)
		fields = append(fields, field)
	}
	if m.Scale != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.Scale)
		fields = append(fields, field)
	}
	if m.Offset != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Int8(m.Offset)
		fields = append(fields, field)
	}
	if m.Units != nil {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.SliceString(m.Units)
		fields = append(fields, field)
	}
	if m.Bits != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.String(m.Bits)
		fields = append(fields, field)
	}
	if m.Accumulate != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.String(m.Accumulate)
		fields = append(fields, field)
	}
	if m.FitBaseUnitId != typedef.FitBaseUnitInvalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint16(uint16(m.FitBaseUnitId))
		fields = append(fields, field)
	}
	if m.NativeMesgNum != typedef.MesgNumInvalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint16(uint16(m.NativeMesgNum))
		fields = append(fields, field)
	}
	if m.NativeFieldNum != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint8(m.NativeFieldNum)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	return mesg
}

// SetDeveloperDataIndex sets DeveloperDataIndex value.
func (m *FieldDescription) SetDeveloperDataIndex(v uint8) *FieldDescription {
	m.DeveloperDataIndex = v
	return m
}

// SetFieldDefinitionNumber sets FieldDefinitionNumber value.
func (m *FieldDescription) SetFieldDefinitionNumber(v uint8) *FieldDescription {
	m.FieldDefinitionNumber = v
	return m
}

// SetFitBaseTypeId sets FitBaseTypeId value.
func (m *FieldDescription) SetFitBaseTypeId(v basetype.BaseType) *FieldDescription {
	m.FitBaseTypeId = v
	return m
}

// SetFieldName sets FieldName value.
//
// Array: [N]
func (m *FieldDescription) SetFieldName(v []string) *FieldDescription {
	m.FieldName = v
	return m
}

// SetArray sets Array value.
func (m *FieldDescription) SetArray(v uint8) *FieldDescription {
	m.Array = v
	return m
}

// SetComponents sets Components value.
func (m *FieldDescription) SetComponents(v string) *FieldDescription {
	m.Components = v
	return m
}

// SetScale sets Scale value.
func (m *FieldDescription) SetScale(v uint8) *FieldDescription {
	m.Scale = v
	return m
}

// SetOffset sets Offset value.
func (m *FieldDescription) SetOffset(v int8) *FieldDescription {
	m.Offset = v
	return m
}

// SetUnits sets Units value.
//
// Array: [N]
func (m *FieldDescription) SetUnits(v []string) *FieldDescription {
	m.Units = v
	return m
}

// SetBits sets Bits value.
func (m *FieldDescription) SetBits(v string) *FieldDescription {
	m.Bits = v
	return m
}

// SetAccumulate sets Accumulate value.
func (m *FieldDescription) SetAccumulate(v string) *FieldDescription {
	m.Accumulate = v
	return m
}

// SetFitBaseUnitId sets FitBaseUnitId value.
func (m *FieldDescription) SetFitBaseUnitId(v typedef.FitBaseUnit) *FieldDescription {
	m.FitBaseUnitId = v
	return m
}

// SetNativeMesgNum sets NativeMesgNum value.
func (m *FieldDescription) SetNativeMesgNum(v typedef.MesgNum) *FieldDescription {
	m.NativeMesgNum = v
	return m
}

// SetNativeFieldNum sets NativeFieldNum value.
func (m *FieldDescription) SetNativeFieldNum(v uint8) *FieldDescription {
	m.NativeFieldNum = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *FieldDescription) SetUnknownFields(unknownFields ...proto.Field) *FieldDescription {
	m.UnknownFields = unknownFields
	return m
}
