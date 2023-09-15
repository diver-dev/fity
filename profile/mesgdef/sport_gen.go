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

// Sport is a Sport message.
type Sport struct {
	Sport    typedef.Sport
	SubSport typedef.SubSport
	Name     string

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSport creates new Sport struct based on given mesg. If mesg is nil or mesg.Num is not equal to Sport mesg number, it will return nil.
func NewSport(mesg proto.Message) *Sport {
	if mesg.Num != typedef.MesgNumSport {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0: basetype.EnumInvalid,   /* Sport */
		1: basetype.EnumInvalid,   /* SubSport */
		3: basetype.StringInvalid, /* Name */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Sport{
		Sport:    typeconv.ToEnum[typedef.Sport](vals[0]),
		SubSport: typeconv.ToEnum[typedef.SubSport](vals[1]),
		Name:     typeconv.ToString[string](vals[3]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Sport mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSport)
func (m Sport) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSport {
		return
	}

	vals := [256]any{
		0: m.Sport,
		1: m.SubSport,
		3: m.Name,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
