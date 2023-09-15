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

// VideoFrame is a VideoFrame message.
type VideoFrame struct {
	Timestamp   typedef.DateTime // Units: s; Whole second part of the timestamp
	TimestampMs uint16           // Units: ms; Millisecond part of the timestamp.
	FrameNumber uint32           // Number of the frame that the timestamp and timestamp_ms correlate to

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideoFrame creates new VideoFrame struct based on given mesg. If mesg is nil or mesg.Num is not equal to VideoFrame mesg number, it will return nil.
func NewVideoFrame(mesg proto.Message) *VideoFrame {
	if mesg.Num != typedef.MesgNumVideoFrame {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		253: basetype.Uint32Invalid, /* Timestamp */
		0:   basetype.Uint16Invalid, /* TimestampMs */
		1:   basetype.Uint32Invalid, /* FrameNumber */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &VideoFrame{
		Timestamp:   typeconv.ToUint32[typedef.DateTime](vals[253]),
		TimestampMs: typeconv.ToUint16[uint16](vals[0]),
		FrameNumber: typeconv.ToUint32[uint32](vals[1]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to VideoFrame mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumVideoFrame)
func (m VideoFrame) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumVideoFrame {
		return
	}

	vals := [256]any{
		253: m.Timestamp,
		0:   m.TimestampMs,
		1:   m.FrameNumber,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}