// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

func TestHeaderMarshaler(t *testing.T) {
	tt := []struct {
		name       string
		fileHeader *proto.FileHeader
		b          []byte
		err        error
	}{
		{
			name: "correct header",
			fileHeader: &proto.FileHeader{
				Size:            14,
				ProtocolVersion: 32,
				ProfileVersion:  2132,
				DataSize:        642262,
				DataType:        ".FIT",
				CRC:             12856,
			},
			b: []byte{
				14,
				32,
				84, 8,
				214, 204, 9, 0,
				46, 70, 73, 84,
				56, 50,
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.fileHeader.MarshalBinary()
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected err: %v, got: %v", tc.err, err)
			}
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageDefinitionMarshaler(t *testing.T) {
	tt := []struct {
		name    string
		mesgdef *proto.MessageDefinition
		b       []byte
	}{
		{
			name: "mesg def fields only",
			mesgdef: &proto.MessageDefinition{
				Header:       64,
				LocalMesgNum: 0,
				Reserved:     0,
				Architecture: 0,
				MesgNum:      typedef.MesgNumFileId,
				FieldDefinitions: []proto.FieldDefinition{
					{Num: 0, Size: 1, BaseType: basetype.Enum},
					{Num: 1, Size: 2, BaseType: basetype.Uint16},
					{Num: 2, Size: 2, BaseType: basetype.Uint16},
					{Num: 3, Size: 4, BaseType: basetype.Uint32z},
					{Num: 8, Size: 13, BaseType: basetype.String},
					{Num: 5, Size: 2, BaseType: basetype.Uint16},
				},
				DeveloperFieldDefinitions: nil},
			b: []byte{
				64, // Header
				0,  // LocalMesgNum
				0,  // Reserved
				0,  // Architecture
				0,  // MesgNum
				6,  // len(FieldDefinitions)
				0, 1, 0,
				1, 2, 132,
				2, 2, 132,
				3, 4, 140,
				8, 13, 7,
				5, 2, 132,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, _ := tc.mesgdef.MarshalBinary()
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestMessageMarshaler(t *testing.T) {
	tt := []struct {
		name string
		mesg *proto.Message
		b    []byte
	}{
		{
			name: "file_id mesg",
			mesg: &proto.Message{
				Header:   0,
				Name:     "file_id",
				Num:      typedef.MesgNumFileId,
				LocalNum: 0,
				Fields: []proto.Field{
					{Value: uint8(4)},
					{Value: uint16(292)},
					{Value: uint16(100)},
					{Value: uint32(120188)},
					{Value: string("XOSS iOS APP")},
					{Value: uint16(1873)},
				}, DeveloperFields: nil,
			},
			b: []byte{
				0, // Header
				4, // Field[0] ...
				36, 1,
				100, 0,
				124, 213, 1, 0,
				88, 79, 83, 83, 32, 105, 79, 83, 32, 65, 80, 80, 00,
				81, 7,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.mesg.MarshalBinary()
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(b, tc.b); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
