// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type Autoscroll byte

const (
	AutoscrollNone    Autoscroll = 0
	AutoscrollSlow    Autoscroll = 1
	AutoscrollMedium  Autoscroll = 2
	AutoscrollFast    Autoscroll = 3
	AutoscrollInvalid Autoscroll = 0xFF // INVALID
)

var autoscrolltostrs = map[Autoscroll]string{
	AutoscrollNone:    "none",
	AutoscrollSlow:    "slow",
	AutoscrollMedium:  "medium",
	AutoscrollFast:    "fast",
	AutoscrollInvalid: "invalid",
}

func (a Autoscroll) String() string {
	val, ok := autoscrolltostrs[a]
	if !ok {
		return strconv.Itoa(int(a))
	}
	return val
}

var strtoautoscroll = func() map[string]Autoscroll {
	m := make(map[string]Autoscroll)
	for t, str := range autoscrolltostrs {
		m[str] = Autoscroll(t)
	}
	return m
}()

// FromString parse string into Autoscroll constant it's represent, return AutoscrollInvalid if not found.
func AutoscrollFromString(s string) Autoscroll {
	val, ok := strtoautoscroll[s]
	if !ok {
		return strtoautoscroll["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListAutoscroll() []Autoscroll {
	vs := make([]Autoscroll, 0, len(autoscrolltostrs))
	for i := range autoscrolltostrs {
		vs = append(vs, Autoscroll(i))
	}
	return vs
}