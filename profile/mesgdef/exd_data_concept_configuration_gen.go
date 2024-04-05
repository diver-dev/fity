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

// ExdDataConceptConfiguration is a ExdDataConceptConfiguration message.
type ExdDataConceptConfiguration struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptIndex uint8
	DataPage     uint8
	ConceptKey   uint8
	Scaling      uint8
	DataUnits    typedef.ExdDataUnits
	Qualifier    typedef.ExdQualifiers
	Descriptor   typedef.ExdDescriptors
	IsSigned     bool

	IsExpandedFields [4]bool // Used for tracking expanded fields, field.Num as index.

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewExdDataConceptConfiguration creates new ExdDataConceptConfiguration struct based on given mesg.
// If mesg is nil, it will return ExdDataConceptConfiguration with all fields being set to its corresponding invalid value.
func NewExdDataConceptConfiguration(mesg *proto.Message) *ExdDataConceptConfiguration {
	vals := [12]proto.Value{}
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

	return &ExdDataConceptConfiguration{
		ScreenIndex:  vals[0].Uint8(),
		ConceptField: vals[1].Uint8(),
		FieldId:      vals[2].Uint8(),
		ConceptIndex: vals[3].Uint8(),
		DataPage:     vals[4].Uint8(),
		ConceptKey:   vals[5].Uint8(),
		Scaling:      vals[6].Uint8(),
		DataUnits:    typedef.ExdDataUnits(vals[8].Uint8()),
		Qualifier:    typedef.ExdQualifiers(vals[9].Uint8()),
		Descriptor:   typedef.ExdDescriptors(vals[10].Uint8()),
		IsSigned:     vals[11].Bool(),

		IsExpandedFields: isExpandedFields,

		DeveloperFields: developerFields,
	}
}

// ToMesg converts ExdDataConceptConfiguration into proto.Message. If options is nil, default options will be used.
func (m *ExdDataConceptConfiguration) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumExdDataConceptConfiguration}

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
	if m.ConceptIndex != basetype.Uint8Invalid && ((m.IsExpandedFields[3] && options.IncludeExpandedFields) || !m.IsExpandedFields[3]) {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint8(m.ConceptIndex)
		field.IsExpandedField = m.IsExpandedFields[3]
		fields = append(fields, field)
	}
	if m.DataPage != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint8(m.DataPage)
		fields = append(fields, field)
	}
	if m.ConceptKey != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.ConceptKey)
		fields = append(fields, field)
	}
	if m.Scaling != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.Scaling)
		fields = append(fields, field)
	}
	if byte(m.DataUnits) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint8(byte(m.DataUnits))
		fields = append(fields, field)
	}
	if byte(m.Qualifier) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint8(byte(m.Qualifier))
		fields = append(fields, field)
	}
	if byte(m.Descriptor) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint8(byte(m.Descriptor))
		fields = append(fields, field)
	}
	if m.IsSigned != false {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Bool(m.IsSigned)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// SetScreenIndex sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetScreenIndex(v uint8) *ExdDataConceptConfiguration {
	m.ScreenIndex = v
	return m
}

// SetConceptField sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptField(v byte) *ExdDataConceptConfiguration {
	m.ConceptField = v
	return m
}

// SetFieldId sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetFieldId(v uint8) *ExdDataConceptConfiguration {
	m.FieldId = v
	return m
}

// SetConceptIndex sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptIndex(v uint8) *ExdDataConceptConfiguration {
	m.ConceptIndex = v
	return m
}

// SetDataPage sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDataPage(v uint8) *ExdDataConceptConfiguration {
	m.DataPage = v
	return m
}

// SetConceptKey sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetConceptKey(v uint8) *ExdDataConceptConfiguration {
	m.ConceptKey = v
	return m
}

// SetScaling sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetScaling(v uint8) *ExdDataConceptConfiguration {
	m.Scaling = v
	return m
}

// SetDataUnits sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDataUnits(v typedef.ExdDataUnits) *ExdDataConceptConfiguration {
	m.DataUnits = v
	return m
}

// SetQualifier sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetQualifier(v typedef.ExdQualifiers) *ExdDataConceptConfiguration {
	m.Qualifier = v
	return m
}

// SetDescriptor sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetDescriptor(v typedef.ExdDescriptors) *ExdDataConceptConfiguration {
	m.Descriptor = v
	return m
}

// SetIsSigned sets ExdDataConceptConfiguration value.
func (m *ExdDataConceptConfiguration) SetIsSigned(v bool) *ExdDataConceptConfiguration {
	m.IsSigned = v
	return m
}

// SetDeveloperFields ExdDataConceptConfiguration's DeveloperFields.
func (m *ExdDataConceptConfiguration) SetDeveloperFields(developerFields ...proto.DeveloperField) *ExdDataConceptConfiguration {
	m.DeveloperFields = developerFields
	return m
}
