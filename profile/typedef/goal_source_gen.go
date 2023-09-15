// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type GoalSource byte

const (
	GoalSourceAuto      GoalSource = 0    // Device generated
	GoalSourceCommunity GoalSource = 1    // Social network sourced goal
	GoalSourceUser      GoalSource = 2    // Manually generated
	GoalSourceInvalid   GoalSource = 0xFF // INVALID
)

var goalsourcetostrs = map[GoalSource]string{
	GoalSourceAuto:      "auto",
	GoalSourceCommunity: "community",
	GoalSourceUser:      "user",
	GoalSourceInvalid:   "invalid",
}

func (g GoalSource) String() string {
	val, ok := goalsourcetostrs[g]
	if !ok {
		return strconv.Itoa(int(g))
	}
	return val
}

var strtogoalsource = func() map[string]GoalSource {
	m := make(map[string]GoalSource)
	for t, str := range goalsourcetostrs {
		m[str] = GoalSource(t)
	}
	return m
}()

// FromString parse string into GoalSource constant it's represent, return GoalSourceInvalid if not found.
func GoalSourceFromString(s string) GoalSource {
	val, ok := strtogoalsource[s]
	if !ok {
		return strtogoalsource["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListGoalSource() []GoalSource {
	vs := make([]GoalSource, 0, len(goalsourcetostrs))
	for i := range goalsourcetostrs {
		vs = append(vs, GoalSource(i))
	}
	return vs
}
