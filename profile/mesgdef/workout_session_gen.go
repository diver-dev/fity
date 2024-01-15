// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// WorkoutSession is a WorkoutSession message.
type WorkoutSession struct {
	MessageIndex   typedef.MessageIndex
	NumValidSteps  uint16
	FirstStepIndex uint16
	PoolLength     uint16 // Scale: 100; Units: m
	Sport          typedef.Sport
	SubSport       typedef.SubSport
	PoolLengthUnit typedef.DisplayMeasure

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkoutSession creates new WorkoutSession struct based on given mesg.
// If mesg is nil, it will return WorkoutSession with all fields being set to its corresponding invalid value.
func NewWorkoutSession(mesg *proto.Message) *WorkoutSession {
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

	return &WorkoutSession{
		MessageIndex:   typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		NumValidSteps:  typeconv.ToUint16[uint16](vals[2]),
		FirstStepIndex: typeconv.ToUint16[uint16](vals[3]),
		PoolLength:     typeconv.ToUint16[uint16](vals[4]),
		Sport:          typeconv.ToEnum[typedef.Sport](vals[0]),
		SubSport:       typeconv.ToEnum[typedef.SubSport](vals[1]),
		PoolLengthUnit: typeconv.ToEnum[typedef.DisplayMeasure](vals[5]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts WorkoutSession into proto.Message.
func (m *WorkoutSession) ToMesg(fac Factory) proto.Message {
	fieldsArray := fieldsPool.Get().(*[256]proto.Field)
	defer fieldsPool.Put(fieldsArray)

	fields := (*fieldsArray)[:0] // Create slice from array with zero len.
	mesg := fac.CreateMesgOnly(typedef.MesgNumWorkoutSession)

	if uint16(m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = uint16(m.MessageIndex)
		fields = append(fields, field)
	}
	if m.NumValidSteps != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = m.NumValidSteps
		fields = append(fields, field)
	}
	if m.FirstStepIndex != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = m.FirstStepIndex
		fields = append(fields, field)
	}
	if m.PoolLength != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = m.PoolLength
		fields = append(fields, field)
	}
	if byte(m.Sport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = byte(m.Sport)
		fields = append(fields, field)
	}
	if byte(m.SubSport) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = byte(m.SubSport)
		fields = append(fields, field)
	}
	if byte(m.PoolLengthUnit) != basetype.EnumInvalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = byte(m.PoolLengthUnit)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// PoolLengthScaled return PoolLength in its scaled value [Scale: 100; Units: m].
//
// If PoolLength value is invalid, float64 invalid value will be returned.
func (m *WorkoutSession) PoolLengthScaled() float64 {
	if m.PoolLength == basetype.Uint16Invalid {
		return basetype.Float64InvalidInFloatForm()
	}
	return scaleoffset.Apply(m.PoolLength, 100, 0)
}

// SetMessageIndex sets WorkoutSession value.
func (m *WorkoutSession) SetMessageIndex(v typedef.MessageIndex) *WorkoutSession {
	m.MessageIndex = v
	return m
}

// SetNumValidSteps sets WorkoutSession value.
func (m *WorkoutSession) SetNumValidSteps(v uint16) *WorkoutSession {
	m.NumValidSteps = v
	return m
}

// SetFirstStepIndex sets WorkoutSession value.
func (m *WorkoutSession) SetFirstStepIndex(v uint16) *WorkoutSession {
	m.FirstStepIndex = v
	return m
}

// SetPoolLength sets WorkoutSession value.
//
// Scale: 100; Units: m
func (m *WorkoutSession) SetPoolLength(v uint16) *WorkoutSession {
	m.PoolLength = v
	return m
}

// SetSport sets WorkoutSession value.
func (m *WorkoutSession) SetSport(v typedef.Sport) *WorkoutSession {
	m.Sport = v
	return m
}

// SetSubSport sets WorkoutSession value.
func (m *WorkoutSession) SetSubSport(v typedef.SubSport) *WorkoutSession {
	m.SubSport = v
	return m
}

// SetPoolLengthUnit sets WorkoutSession value.
func (m *WorkoutSession) SetPoolLengthUnit(v typedef.DisplayMeasure) *WorkoutSession {
	m.PoolLengthUnit = v
	return m
}

// SetDeveloperFields WorkoutSession's DeveloperFields.
func (m *WorkoutSession) SetDeveloperFields(developerFields ...proto.DeveloperField) *WorkoutSession {
	m.DeveloperFields = developerFields
	return m
}
