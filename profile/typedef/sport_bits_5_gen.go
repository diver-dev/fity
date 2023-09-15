// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits5 uint8

const (
	SportBits5WaterSkiing SportBits5 = 0x01
	SportBits5Kayaking    SportBits5 = 0x02
	SportBits5Rafting     SportBits5 = 0x04
	SportBits5Windsurfing SportBits5 = 0x08
	SportBits5Kitesurfing SportBits5 = 0x10
	SportBits5Tactical    SportBits5 = 0x20
	SportBits5Jumpmaster  SportBits5 = 0x40
	SportBits5Boxing      SportBits5 = 0x80
	SportBits5Invalid     SportBits5 = 0x0 // INVALID
)

var sportbits5tostrs = map[SportBits5]string{
	SportBits5WaterSkiing: "water_skiing",
	SportBits5Kayaking:    "kayaking",
	SportBits5Rafting:     "rafting",
	SportBits5Windsurfing: "windsurfing",
	SportBits5Kitesurfing: "kitesurfing",
	SportBits5Tactical:    "tactical",
	SportBits5Jumpmaster:  "jumpmaster",
	SportBits5Boxing:      "boxing",
	SportBits5Invalid:     "invalid",
}

func (s SportBits5) String() string {
	val, ok := sportbits5tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits5 = func() map[string]SportBits5 {
	m := make(map[string]SportBits5)
	for t, str := range sportbits5tostrs {
		m[str] = SportBits5(t)
	}
	return m
}()

// FromString parse string into SportBits5 constant it's represent, return SportBits5Invalid if not found.
func SportBits5FromString(s string) SportBits5 {
	val, ok := strtosportbits5[s]
	if !ok {
		return strtosportbits5["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits5() []SportBits5 {
	vs := make([]SportBits5, 0, len(sportbits5tostrs))
	for i := range sportbits5tostrs {
		vs = append(vs, SportBits5(i))
	}
	return vs
}