// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SleepLevel byte

const (
	SleepLevelUnmeasurable SleepLevel = 0
	SleepLevelAwake        SleepLevel = 1
	SleepLevelLight        SleepLevel = 2
	SleepLevelDeep         SleepLevel = 3
	SleepLevelRem          SleepLevel = 4
	SleepLevelInvalid      SleepLevel = 0xFF // INVALID
)

var sleepleveltostrs = map[SleepLevel]string{
	SleepLevelUnmeasurable: "unmeasurable",
	SleepLevelAwake:        "awake",
	SleepLevelLight:        "light",
	SleepLevelDeep:         "deep",
	SleepLevelRem:          "rem",
	SleepLevelInvalid:      "invalid",
}

func (s SleepLevel) String() string {
	val, ok := sleepleveltostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosleeplevel = func() map[string]SleepLevel {
	m := make(map[string]SleepLevel)
	for t, str := range sleepleveltostrs {
		m[str] = SleepLevel(t)
	}
	return m
}()

// FromString parse string into SleepLevel constant it's represent, return SleepLevelInvalid if not found.
func SleepLevelFromString(s string) SleepLevel {
	val, ok := strtosleeplevel[s]
	if !ok {
		return strtosleeplevel["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSleepLevel() []SleepLevel {
	vs := make([]SleepLevel, 0, len(sleepleveltostrs))
	for i := range sleepleveltostrs {
		vs = append(vs, SleepLevel(i))
	}
	return vs
}