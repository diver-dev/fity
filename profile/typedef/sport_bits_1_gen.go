// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits1 uint8

const (
	SportBits1Tennis             SportBits1 = 0x01
	SportBits1AmericanFootball   SportBits1 = 0x02
	SportBits1Training           SportBits1 = 0x04
	SportBits1Walking            SportBits1 = 0x08
	SportBits1CrossCountrySkiing SportBits1 = 0x10
	SportBits1AlpineSkiing       SportBits1 = 0x20
	SportBits1Snowboarding       SportBits1 = 0x40
	SportBits1Rowing             SportBits1 = 0x80
	SportBits1Invalid            SportBits1 = 0x0 // INVALID
)

var sportbits1tostrs = map[SportBits1]string{
	SportBits1Tennis:             "tennis",
	SportBits1AmericanFootball:   "american_football",
	SportBits1Training:           "training",
	SportBits1Walking:            "walking",
	SportBits1CrossCountrySkiing: "cross_country_skiing",
	SportBits1AlpineSkiing:       "alpine_skiing",
	SportBits1Snowboarding:       "snowboarding",
	SportBits1Rowing:             "rowing",
	SportBits1Invalid:            "invalid",
}

func (s SportBits1) String() string {
	val, ok := sportbits1tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits1 = func() map[string]SportBits1 {
	m := make(map[string]SportBits1)
	for t, str := range sportbits1tostrs {
		m[str] = SportBits1(t)
	}
	return m
}()

// FromString parse string into SportBits1 constant it's represent, return SportBits1Invalid if not found.
func SportBits1FromString(s string) SportBits1 {
	val, ok := strtosportbits1[s]
	if !ok {
		return strtosportbits1["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits1() []SportBits1 {
	vs := make([]SportBits1, 0, len(sportbits1tostrs))
	for i := range sportbits1tostrs {
		vs = append(vs, SportBits1(i))
	}
	return vs
}