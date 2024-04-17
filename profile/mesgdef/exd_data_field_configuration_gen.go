// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// ExdDataFieldConfiguration is a ExdDataFieldConfiguration message.
type ExdDataFieldConfiguration struct {
	Title        []string // Array: [32]
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptCount uint8
	DisplayType  typedef.ExdDisplayType

	IsExpandedFields [4]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExdDataFieldConfiguration creates new ExdDataFieldConfiguration struct based on given mesg.
// If mesg is nil, it will return ExdDataFieldConfiguration with all fields being set to its corresponding invalid value.
func NewExdDataFieldConfiguration(mesg *proto.Message) *ExdDataFieldConfiguration {
	vals := [6]proto.Value{}
	isExpandedFields := [4]bool{}

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

	return &ExdDataFieldConfiguration{
		Title:        vals[5].SliceString(),
		ScreenIndex:  vals[0].Uint8(),
		ConceptField: vals[1].Uint8(),
		FieldId:      vals[2].Uint8(),
		ConceptCount: vals[3].Uint8(),
		DisplayType:  typedef.ExdDisplayType(vals[4].Uint8()),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExdDataFieldConfiguration into proto.Message. If options is nil, default options will be used.
func (m *ExdDataFieldConfiguration) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumExdDataFieldConfiguration}

	if m.Title != nil {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.SliceString(m.Title)
		fields = append(fields, field)
	}
	if m.ScreenIndex != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(m.ScreenIndex)
		fields = append(fields, field)
	}
	if m.ConceptField != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint8(m.ConceptField)
		fields = append(fields, field)
	}
	if m.FieldId != basetype.Uint8Invalid && ((m.IsExpandedFields[2] && options.IncludeExpandedFields) || !m.IsExpandedFields[2]) {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(m.FieldId)
		field.IsExpandedField = m.IsExpandedFields[2]
		fields = append(fields, field)
	}
	if m.ConceptCount != basetype.Uint8Invalid && ((m.IsExpandedFields[3] && options.IncludeExpandedFields) || !m.IsExpandedFields[3]) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.ConceptCount)
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if byte(m.DisplayType) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(byte(m.DisplayType))
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetTitle sets ExdDataFieldConfiguration value.
//
// Array: [32]
func (m *ExdDataFieldConfiguration) SetTitle(v []string) *ExdDataFieldConfiguration {
	m.Title = v
	return m
}

// SetScreenIndex sets ExdDataFieldConfiguration value.
func (m *ExdDataFieldConfiguration) SetScreenIndex(v uint8) *ExdDataFieldConfiguration {
	m.ScreenIndex = v
	return m
}

// SetConceptField sets ExdDataFieldConfiguration value.
func (m *ExdDataFieldConfiguration) SetConceptField(v byte) *ExdDataFieldConfiguration {
	m.ConceptField = v
	return m
}

// SetFieldId sets ExdDataFieldConfiguration value.
func (m *ExdDataFieldConfiguration) SetFieldId(v uint8) *ExdDataFieldConfiguration {
	m.FieldId = v
	return m
}

// SetConceptCount sets ExdDataFieldConfiguration value.
func (m *ExdDataFieldConfiguration) SetConceptCount(v uint8) *ExdDataFieldConfiguration {
	m.ConceptCount = v
	return m
}

// SetDisplayType sets ExdDataFieldConfiguration value.
func (m *ExdDataFieldConfiguration) SetDisplayType(v typedef.ExdDisplayType) *ExdDataFieldConfiguration {
	m.DisplayType = v
	return m
}

// SetDeveloperFields ExdDataFieldConfiguration's DeveloperFields.
func (m *ExdDataFieldConfiguration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExdDataFieldConfiguration {
	m.DeveloperFields = developerFields
	return m
}
