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
	"math"
)

// ClimbPro is a ClimbPro message.
type ClimbPro struct {
	Timestamp     typedef.DateTime // Units: s;
	PositionLat   int32            // Units: semicircles;
	PositionLong  int32            // Units: semicircles;
	ClimbProEvent typedef.ClimbProEvent
	ClimbNumber   uint16
	ClimbCategory uint8
	CurrentDist   float32 // Units: m;

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewClimbPro creates new ClimbPro struct based on given mesg. If mesg is nil or mesg.Num is not equal to ClimbPro mesg number, it will return nil.
func NewClimbPro(mesg proto.Message) *ClimbPro {
	if mesg.Num != typedef.MesgNumClimbPro {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid,                        /* Timestamp */
		0:   basetype.Sint32Invalid,                        /* PositionLat */
		1:   basetype.Sint32Invalid,                        /* PositionLong */
		2:   basetype.EnumInvalid,                          /* ClimbProEvent */
		3:   basetype.Uint16Invalid,                        /* ClimbNumber */
		4:   basetype.Uint8Invalid,                         /* ClimbCategory */
		5:   math.Float32frombits(basetype.Float32Invalid), /* CurrentDist */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &ClimbPro{
		Timestamp:     typeconv.ToUint32[typedef.DateTime](vals[253]),
		PositionLat:   typeconv.ToSint32[int32](vals[0]),
		PositionLong:  typeconv.ToSint32[int32](vals[1]),
		ClimbProEvent: typeconv.ToEnum[typedef.ClimbProEvent](vals[2]),
		ClimbNumber:   typeconv.ToUint16[uint16](vals[3]),
		ClimbCategory: typeconv.ToUint8[uint8](vals[4]),
		CurrentDist:   typeconv.ToFloat32[float32](vals[5]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to ClimbPro mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumClimbPro)
func (m ClimbPro) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumClimbPro {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.PositionLat,
		1:   m.PositionLong,
		2:   m.ClimbProEvent,
		3:   m.ClimbNumber,
		4:   m.ClimbCategory,
		5:   m.CurrentDist,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
