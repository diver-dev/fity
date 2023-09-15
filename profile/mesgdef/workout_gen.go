// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// Workout is a Workout message.
type Workout struct {
	MessageIndex   typedef.MessageIndex
	Sport          typedef.Sport
	Capabilities   typedef.WorkoutCapabilities
	NumValidSteps  uint16 // number of valid steps
	WktName        string
	SubSport       typedef.SubSport
	PoolLength     uint16 // Scale: 100; Units: m;
	PoolLengthUnit typedef.DisplayMeasure

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewWorkout creates new Workout struct based on given mesg. If mesg is nil or mesg.Num is not equal to Workout mesg number, it will return nil.
func NewWorkout(mesg proto.Message) *Workout {
	if mesg.Num != typedef.MesgNumWorkout {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid,  /* MessageIndex */
		4:   basetype.EnumInvalid,    /* Sport */
		5:   basetype.Uint32zInvalid, /* Capabilities */
		6:   basetype.Uint16Invalid,  /* NumValidSteps */
		8:   basetype.StringInvalid,  /* WktName */
		11:  basetype.EnumInvalid,    /* SubSport */
		14:  basetype.Uint16Invalid,  /* PoolLength */
		15:  basetype.EnumInvalid,    /* PoolLengthUnit */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Workout{
		MessageIndex:   typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Sport:          typeconv.ToEnum[typedef.Sport](vals[4]),
		Capabilities:   typeconv.ToUint32z[typedef.WorkoutCapabilities](vals[5]),
		NumValidSteps:  typeconv.ToUint16[uint16](vals[6]),
		WktName:        typeconv.ToString[string](vals[8]),
		SubSport:       typeconv.ToEnum[typedef.SubSport](vals[11]),
		PoolLength:     typeconv.ToUint16[uint16](vals[14]),
		PoolLengthUnit: typeconv.ToEnum[typedef.DisplayMeasure](vals[15]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Workout mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumWorkout)
func (m Workout) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumWorkout {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		4:   m.Sport,
		5:   m.Capabilities,
		6:   m.NumValidSteps,
		8:   m.WktName,
		11:  m.SubSport,
		14:  m.PoolLength,
		15:  m.PoolLengthUnit,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}