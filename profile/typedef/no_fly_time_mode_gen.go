// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type NoFlyTimeMode byte

const (
	NoFlyTimeModeStandard    NoFlyTimeMode = 0    // Standard Diver Alert Network no-fly guidance
	NoFlyTimeModeFlat24Hours NoFlyTimeMode = 1    // Flat 24 hour no-fly guidance
	NoFlyTimeModeInvalid     NoFlyTimeMode = 0xFF // INVALID
)

var noflytimemodetostrs = map[NoFlyTimeMode]string{
	NoFlyTimeModeStandard:    "standard",
	NoFlyTimeModeFlat24Hours: "flat_24_hours",
	NoFlyTimeModeInvalid:     "invalid",
}

func (n NoFlyTimeMode) String() string {
	val, ok := noflytimemodetostrs[n]
	if !ok {
		return strconv.Itoa(int(n))
	}
	return val
}

var strtonoflytimemode = func() map[string]NoFlyTimeMode {
	m := make(map[string]NoFlyTimeMode)
	for t, str := range noflytimemodetostrs {
		m[str] = NoFlyTimeMode(t)
	}
	return m
}()

// FromString parse string into NoFlyTimeMode constant it's represent, return NoFlyTimeModeInvalid if not found.
func NoFlyTimeModeFromString(s string) NoFlyTimeMode {
	val, ok := strtonoflytimemode[s]
	if !ok {
		return strtonoflytimemode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListNoFlyTimeMode() []NoFlyTimeMode {
	vs := make([]NoFlyTimeMode, 0, len(noflytimemodetostrs))
	for i := range noflytimemodetostrs {
		vs = append(vs, NoFlyTimeMode(i))
	}
	return vs
}
