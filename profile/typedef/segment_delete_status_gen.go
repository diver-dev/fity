// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SegmentDeleteStatus byte

const (
	SegmentDeleteStatusDoNotDelete SegmentDeleteStatus = 0
	SegmentDeleteStatusDeleteOne   SegmentDeleteStatus = 1
	SegmentDeleteStatusDeleteAll   SegmentDeleteStatus = 2
	SegmentDeleteStatusInvalid     SegmentDeleteStatus = 0xFF // INVALID
)

var segmentdeletestatustostrs = map[SegmentDeleteStatus]string{
	SegmentDeleteStatusDoNotDelete: "do_not_delete",
	SegmentDeleteStatusDeleteOne:   "delete_one",
	SegmentDeleteStatusDeleteAll:   "delete_all",
	SegmentDeleteStatusInvalid:     "invalid",
}

func (s SegmentDeleteStatus) String() string {
	val, ok := segmentdeletestatustostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosegmentdeletestatus = func() map[string]SegmentDeleteStatus {
	m := make(map[string]SegmentDeleteStatus)
	for t, str := range segmentdeletestatustostrs {
		m[str] = SegmentDeleteStatus(t)
	}
	return m
}()

// FromString parse string into SegmentDeleteStatus constant it's represent, return SegmentDeleteStatusInvalid if not found.
func SegmentDeleteStatusFromString(s string) SegmentDeleteStatus {
	val, ok := strtosegmentdeletestatus[s]
	if !ok {
		return strtosegmentdeletestatus["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSegmentDeleteStatus() []SegmentDeleteStatus {
	vs := make([]SegmentDeleteStatus, 0, len(segmentdeletestatustostrs))
	for i := range segmentdeletestatustostrs {
		vs = append(vs, SegmentDeleteStatus(i))
	}
	return vs
}