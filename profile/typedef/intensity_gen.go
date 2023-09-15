// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Intensity byte

const (
	IntensityActive   Intensity = 0
	IntensityRest     Intensity = 1
	IntensityWarmup   Intensity = 2
	IntensityCooldown Intensity = 3
	IntensityRecovery Intensity = 4
	IntensityInterval Intensity = 5
	IntensityOther    Intensity = 6
	IntensityInvalid  Intensity = 0xFF // INVALID
)

var intensitytostrs = map[Intensity]string{
	IntensityActive:   "active",
	IntensityRest:     "rest",
	IntensityWarmup:   "warmup",
	IntensityCooldown: "cooldown",
	IntensityRecovery: "recovery",
	IntensityInterval: "interval",
	IntensityOther:    "other",
	IntensityInvalid:  "invalid",
}

func (i Intensity) String() string {
	val, ok := intensitytostrs[i]
	if !ok {
		return strconv.Itoa(int(i))
	}
	return val
}

var strtointensity = func() map[string]Intensity {
	m := make(map[string]Intensity)
	for t, str := range intensitytostrs {
		m[str] = Intensity(t)
	}
	return m
}()

// FromString parse string into Intensity constant it's represent, return IntensityInvalid if not found.
func IntensityFromString(s string) Intensity {
	val, ok := strtointensity[s]
	if !ok {
		return strtointensity["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListIntensity() []Intensity {
	vs := make([]Intensity, 0, len(intensitytostrs))
	for i := range intensitytostrs {
		vs = append(vs, Intensity(i))
	}
	return vs
}
