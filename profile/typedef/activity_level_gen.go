// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ActivityLevel byte

const (
	ActivityLevelLow     ActivityLevel = 0
	ActivityLevelMedium  ActivityLevel = 1
	ActivityLevelHigh    ActivityLevel = 2
	ActivityLevelInvalid ActivityLevel = 0xFF // INVALID
)

var activityleveltostrs = map[ActivityLevel]string{
	ActivityLevelLow:     "low",
	ActivityLevelMedium:  "medium",
	ActivityLevelHigh:    "high",
	ActivityLevelInvalid: "invalid",
}

func (a ActivityLevel) String() string {
	val, ok := activityleveltostrs[a]
	if !ok {
		return strconv.Itoa(int(a))
	}
	return val
}

var strtoactivitylevel = func() map[string]ActivityLevel {
	m := make(map[string]ActivityLevel)
	for t, str := range activityleveltostrs {
		m[str] = ActivityLevel(t)
	}
	return m
}()

// FromString parse string into ActivityLevel constant it's represent, return ActivityLevelInvalid if not found.
func ActivityLevelFromString(s string) ActivityLevel {
	val, ok := strtoactivitylevel[s]
	if !ok {
		return strtoactivitylevel["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListActivityLevel() []ActivityLevel {
	vs := make([]ActivityLevel, 0, len(activityleveltostrs))
	for i := range activityleveltostrs {
		vs = append(vs, ActivityLevel(i))
	}
	return vs
}
