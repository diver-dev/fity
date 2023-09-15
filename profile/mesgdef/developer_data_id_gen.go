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

// DeveloperDataId is a DeveloperDataId message.
type DeveloperDataId struct {
	DeveloperId        []byte // Array: [N];
	ApplicationId      []byte // Array: [N];
	ManufacturerId     typedef.Manufacturer
	DeveloperDataIndex uint8
	ApplicationVersion uint32
}

// NewDeveloperDataId creates new DeveloperDataId struct based on given mesg. If mesg is nil or mesg.Num is not equal to DeveloperDataId mesg number, it will return nil.
func NewDeveloperDataId(mesg proto.Message) *DeveloperDataId {
	if mesg.Num != typedef.MesgNumDeveloperDataId {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0: nil,                    /* DeveloperId */
		1: nil,                    /* ApplicationId */
		2: basetype.Uint16Invalid, /* ManufacturerId */
		3: basetype.Uint8Invalid,  /* DeveloperDataIndex */
		4: basetype.Uint32Invalid, /* ApplicationVersion */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &DeveloperDataId{
		DeveloperId:        typeconv.ToSliceByte[byte](vals[0]),
		ApplicationId:      typeconv.ToSliceByte[byte](vals[1]),
		ManufacturerId:     typeconv.ToUint16[typedef.Manufacturer](vals[2]),
		DeveloperDataIndex: typeconv.ToUint8[uint8](vals[3]),
		ApplicationVersion: typeconv.ToUint32[uint32](vals[4]),
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to DeveloperDataId mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumDeveloperDataId)
func (m DeveloperDataId) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumDeveloperDataId {
		return
	}

	vals := [256]any{
		0: m.DeveloperId,
		1: m.ApplicationId,
		2: m.ManufacturerId,
		3: m.DeveloperDataIndex,
		4: m.ApplicationVersion,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
}