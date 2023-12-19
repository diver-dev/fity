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

// VideoTitle is a VideoTitle message.
type VideoTitle struct {
	MessageIndex typedef.MessageIndex // Long titles will be split into multiple parts
	MessageCount uint16               // Total number of title parts
	Text         string

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewVideoTitle creates new VideoTitle struct based on given mesg.
// If mesg is nil, it will return VideoTitle with all fields being set to its corresponding invalid value.
func NewVideoTitle(mesg *proto.Message) *VideoTitle {
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

	return &VideoTitle{
		MessageIndex: typeconv.ToUint16[typedef.MessageIndex](vals[254]),
		MessageCount: typeconv.ToUint16[uint16](vals[0]),
		Text:         typeconv.ToString[string](vals[1]),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts VideoTitle into proto.Message.
func (m *VideoTitle) ToMesg(fac Factory) proto.Message {
	mesg := fac.CreateMesgOnly(typedef.MesgNumVideoTitle)
	mesg.Fields = make([]proto.Field, 0, m.size())

	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 254)
		field.Value = typeconv.ToUint16[uint16](m.MessageIndex)
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.MessageCount != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = m.MessageCount
		mesg.Fields = append(mesg.Fields, field)
	}
	if m.Text != basetype.StringInvalid && m.Text != "" {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = m.Text
		mesg.Fields = append(mesg.Fields, field)
	}

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// size returns size of VideoTitle's valid fields.
func (m *VideoTitle) size() byte {
	var size byte
	if typeconv.ToUint16[uint16](m.MessageIndex) != basetype.Uint16Invalid {
		size++
	}
	if m.MessageCount != basetype.Uint16Invalid {
		size++
	}
	if m.Text != basetype.StringInvalid && m.Text != "" {
		size++
	}
	return size
}

// SetMessageIndex sets VideoTitle value.
//
// Long titles will be split into multiple parts
func (m *VideoTitle) SetMessageIndex(v typedef.MessageIndex) *VideoTitle {
	m.MessageIndex = v
	return m
}

// SetMessageCount sets VideoTitle value.
//
// Total number of title parts
func (m *VideoTitle) SetMessageCount(v uint16) *VideoTitle {
	m.MessageCount = v
	return m
}

// SetText sets VideoTitle value.
func (m *VideoTitle) SetText(v string) *VideoTitle {
	m.Text = v
	return m
}

// SetDeveloperFields VideoTitle's DeveloperFields.
func (m *VideoTitle) SetDeveloperFields(developerFields ...proto.DeveloperField) *VideoTitle {
	m.DeveloperFields = developerFields
	return m
}
