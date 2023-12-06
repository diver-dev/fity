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

// SlaveDevice is a SlaveDevice message.
type SlaveDevice struct {
	Manufacturer typedef.Manufacturer
	Product      uint16

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSlaveDevice creates new SlaveDevice struct based on given mesg. If mesg is nil or mesg.Num is not equal to SlaveDevice mesg number, it will return nil.
func NewSlaveDevice(mesg proto.Message) *SlaveDevice {
	if mesg.Num != typedef.MesgNumSlaveDevice {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0: basetype.Uint16Invalid, /* Manufacturer */
		1: basetype.Uint16Invalid, /* Product */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &SlaveDevice{
		Manufacturer: typeconv.ToUint16[typedef.Manufacturer](vals[0]),
		Product:      typeconv.ToUint16[uint16](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SlaveDevice mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSlaveDevice)
func (m SlaveDevice) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSlaveDevice {
		return
	}

	vals := [256]any{
		0: m.Manufacturer,
		1: m.Product,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
