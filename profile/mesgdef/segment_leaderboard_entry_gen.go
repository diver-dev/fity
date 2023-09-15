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

// SegmentLeaderboardEntry is a SegmentLeaderboardEntry message.
type SegmentLeaderboardEntry struct {
	MessageIndex     typedef.MessageIndex
	Name             string                         // Friendly name assigned to leader
	Type             typedef.SegmentLeaderboardType // Leader classification
	GroupPrimaryKey  uint32                         // Primary user ID of this leader
	ActivityId       uint32                         // ID of the activity associated with this leader time
	SegmentTime      uint32                         // Scale: 1000; Units: s; Segment Time (includes pauses)
	ActivityIdString string                         // String version of the activity_id. 21 characters long, express in decimal

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSegmentLeaderboardEntry creates new SegmentLeaderboardEntry struct based on given mesg. If mesg is nil or mesg.Num is not equal to SegmentLeaderboardEntry mesg number, it will return nil.
func NewSegmentLeaderboardEntry(mesg proto.Message) *SegmentLeaderboardEntry {
	if mesg.Num != typedef.MesgNumSegmentLeaderboardEntry {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		254: basetype.Uint16Invalid, /* MessageIndex */
		0:   basetype.StringInvalid, /* Name */
		1:   basetype.EnumInvalid,   /* Type */
		2:   basetype.Uint32Invalid, /* GroupPrimaryKey */
		3:   basetype.Uint32Invalid, /* ActivityId */
		4:   basetype.Uint32Invalid, /* SegmentTime */
		5:   basetype.StringInvalid, /* ActivityIdString */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &SegmentLeaderboardEntry{
		MessageIndex:     typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		Name:             typeconv.ToString[string](vals[0]),
		Type:             typeconv.ToEnum[typedef.SegmentLeaderboardType](vals[1]),
		GroupPrimaryKey:  typeconv.ToUint32[uint32](vals[2]),
		ActivityId:       typeconv.ToUint32[uint32](vals[3]),
		SegmentTime:      typeconv.ToUint32[uint32](vals[4]),
		ActivityIdString: typeconv.ToString[string](vals[5]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SegmentLeaderboardEntry mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSegmentLeaderboardEntry)
func (m SegmentLeaderboardEntry) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSegmentLeaderboardEntry {
		return
	}

	vals := [256]any{
		254: m.MessageIndex,
		0:   m.Name,
		1:   m.Type,
		2:   m.GroupPrimaryKey,
		3:   m.ActivityId,
		4:   m.SegmentTime,
		5:   m.ActivityIdString,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
