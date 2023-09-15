// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SupportedExdScreenLayouts uint32

const (
	SupportedExdScreenLayoutsFullScreen                SupportedExdScreenLayouts = 0x00000001
	SupportedExdScreenLayoutsHalfVertical              SupportedExdScreenLayouts = 0x00000002
	SupportedExdScreenLayoutsHalfHorizontal            SupportedExdScreenLayouts = 0x00000004
	SupportedExdScreenLayoutsHalfVerticalRightSplit    SupportedExdScreenLayouts = 0x00000008
	SupportedExdScreenLayoutsHalfHorizontalBottomSplit SupportedExdScreenLayouts = 0x00000010
	SupportedExdScreenLayoutsFullQuarterSplit          SupportedExdScreenLayouts = 0x00000020
	SupportedExdScreenLayoutsHalfVerticalLeftSplit     SupportedExdScreenLayouts = 0x00000040
	SupportedExdScreenLayoutsHalfHorizontalTopSplit    SupportedExdScreenLayouts = 0x00000080
	SupportedExdScreenLayoutsInvalid                   SupportedExdScreenLayouts = 0x0 // INVALID
)

var supportedexdscreenlayoutstostrs = map[SupportedExdScreenLayouts]string{
	SupportedExdScreenLayoutsFullScreen:                "full_screen",
	SupportedExdScreenLayoutsHalfVertical:              "half_vertical",
	SupportedExdScreenLayoutsHalfHorizontal:            "half_horizontal",
	SupportedExdScreenLayoutsHalfVerticalRightSplit:    "half_vertical_right_split",
	SupportedExdScreenLayoutsHalfHorizontalBottomSplit: "half_horizontal_bottom_split",
	SupportedExdScreenLayoutsFullQuarterSplit:          "full_quarter_split",
	SupportedExdScreenLayoutsHalfVerticalLeftSplit:     "half_vertical_left_split",
	SupportedExdScreenLayoutsHalfHorizontalTopSplit:    "half_horizontal_top_split",
	SupportedExdScreenLayoutsInvalid:                   "invalid",
}

func (s SupportedExdScreenLayouts) String() string {
	val, ok := supportedexdscreenlayoutstostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosupportedexdscreenlayouts = func() map[string]SupportedExdScreenLayouts {
	m := make(map[string]SupportedExdScreenLayouts)
	for t, str := range supportedexdscreenlayoutstostrs {
		m[str] = SupportedExdScreenLayouts(t)
	}
	return m
}()

// FromString parse string into SupportedExdScreenLayouts constant it's represent, return SupportedExdScreenLayoutsInvalid if not found.
func SupportedExdScreenLayoutsFromString(s string) SupportedExdScreenLayouts {
	val, ok := strtosupportedexdscreenlayouts[s]
	if !ok {
		return strtosupportedexdscreenlayouts["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSupportedExdScreenLayouts() []SupportedExdScreenLayouts {
	vs := make([]SupportedExdScreenLayouts, 0, len(supportedexdscreenlayoutstostrs))
	for i := range supportedexdscreenlayoutstostrs {
		vs = append(vs, SupportedExdScreenLayouts(i))
	}
	return vs
}
