// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SportBits0 uint8

const (
	SportBits0Generic          SportBits0 = 0x01
	SportBits0Running          SportBits0 = 0x02
	SportBits0Cycling          SportBits0 = 0x04
	SportBits0Transition       SportBits0 = 0x08 // Mulitsport transition
	SportBits0FitnessEquipment SportBits0 = 0x10
	SportBits0Swimming         SportBits0 = 0x20
	SportBits0Basketball       SportBits0 = 0x40
	SportBits0Soccer           SportBits0 = 0x80
	SportBits0Invalid          SportBits0 = 0x0 // INVALID
)

var sportbits0tostrs = map[SportBits0]string{
	SportBits0Generic:          "generic",
	SportBits0Running:          "running",
	SportBits0Cycling:          "cycling",
	SportBits0Transition:       "transition",
	SportBits0FitnessEquipment: "fitness_equipment",
	SportBits0Swimming:         "swimming",
	SportBits0Basketball:       "basketball",
	SportBits0Soccer:           "soccer",
	SportBits0Invalid:          "invalid",
}

func (s SportBits0) String() string {
	val, ok := sportbits0tostrs[s]
	if !ok {
		return strconv.FormatUint(uint64(s), 10)
	}
	return val
}

var strtosportbits0 = func() map[string]SportBits0 {
	m := make(map[string]SportBits0)
	for t, str := range sportbits0tostrs {
		m[str] = SportBits0(t)
	}
	return m
}()

// FromString parse string into SportBits0 constant it's represent, return SportBits0Invalid if not found.
func SportBits0FromString(s string) SportBits0 {
	val, ok := strtosportbits0[s]
	if !ok {
		return strtosportbits0["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSportBits0() []SportBits0 {
	vs := make([]SportBits0, 0, len(sportbits0tostrs))
	for i := range sportbits0tostrs {
		vs = append(vs, SportBits0(i))
	}
	return vs
}
