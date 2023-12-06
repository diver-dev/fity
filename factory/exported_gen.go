// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

var std = New()

// StandardFactory returns standard factory.
func StandardFactory() *Factory { return std }

// CreateMesg creates new message based on defined messages in the factory. If not found, it returns new message with "unknown" name.
func CreateMesg(num typedef.MesgNum) proto.Message {
	return std.CreateMesg(num)
}

// CreateMesgOnly is similar to CreateMesg, but it sets Fields and DeveloperFields to nil. This is useful when we plan to fill these values ourselves
// to avoid unnecessary malloc when cloning them, as they will be removed anyway. For example, the decoding process will populate them with decoded data.
func CreateMesgOnly(num typedef.MesgNum) proto.Message {
	return std.CreateMesgOnly(num)
}

// CreateField creates new field based on defined messages in the factory. If not found, it returns new field with "unknown" name.
//
// Returned Field's FieldBase is a pointer, referencing a value in this factory to reduce unnecessary malloc or runtime duffcopy
// since the content should not be changed.
func CreateField(mesgNum typedef.MesgNum, num byte) proto.Field {
	return std.CreateField(mesgNum, num)
}

// RegisterMesg registers a new message that is not defined in the profile.xlsx.
// You can not edit or replace existing message in the factory, including the messages you have registered.
// If you intend to edit your own messages, create a new factory instance using New() and define the new message definitions on it.
//
// By registering, any Fit file containing these messages can be recognized instead of returning "unknown" message.
func RegisterMesg(mesg proto.Message) error {
	return std.RegisterMesg(mesg)
}
