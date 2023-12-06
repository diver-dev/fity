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

// Schedule is a Schedule message.
type Schedule struct {
	Manufacturer  typedef.Manufacturer // Corresponds to file_id of scheduled workout / course.
	Product       uint16               // Corresponds to file_id of scheduled workout / course.
	SerialNumber  uint32               // Corresponds to file_id of scheduled workout / course.
	TimeCreated   typedef.DateTime     // Corresponds to file_id of scheduled workout / course.
	Completed     bool                 // TRUE if this activity has been started
	Type          typedef.Schedule
	ScheduledTime typedef.LocalDateTime

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSchedule creates new Schedule struct based on given mesg. If mesg is nil or mesg.Num is not equal to Schedule mesg number, it will return nil.
func NewSchedule(mesg proto.Message) *Schedule {
	if mesg.Num != typedef.MesgNumSchedule {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0: basetype.Uint16Invalid,  /* Manufacturer */
		1: basetype.Uint16Invalid,  /* Product */
		2: basetype.Uint32zInvalid, /* SerialNumber */
		3: basetype.Uint32Invalid,  /* TimeCreated */
		4: false,                   /* Completed */
		5: basetype.EnumInvalid,    /* Type */
		6: basetype.Uint32Invalid,  /* ScheduledTime */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Schedule{
		Manufacturer:  typeconv.ToUint16[typedef.Manufacturer](vals[0]),
		Product:       typeconv.ToUint16[uint16](vals[1]),
		SerialNumber:  typeconv.ToUint32z[uint32](vals[2]),
		TimeCreated:   typeconv.ToUint32[typedef.DateTime](vals[3]),
		Completed:     typeconv.ToBool[bool](vals[4]),
		Type:          typeconv.ToEnum[typedef.Schedule](vals[5]),
		ScheduledTime: typeconv.ToUint32[typedef.LocalDateTime](vals[6]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Schedule mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSchedule)
func (m Schedule) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSchedule {
		return
	}

	vals := [256]any{
		0: m.Manufacturer,
		1: m.Product,
		2: m.SerialNumber,
		3: m.TimeCreated,
		4: m.Completed,
		5: m.Type,
		6: m.ScheduledTime,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
