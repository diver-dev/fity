// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type ExdDisplayType byte

const (
	ExdDisplayTypeNumerical         ExdDisplayType = 0
	ExdDisplayTypeSimple            ExdDisplayType = 1
	ExdDisplayTypeGraph             ExdDisplayType = 2
	ExdDisplayTypeBar               ExdDisplayType = 3
	ExdDisplayTypeCircleGraph       ExdDisplayType = 4
	ExdDisplayTypeVirtualPartner    ExdDisplayType = 5
	ExdDisplayTypeBalance           ExdDisplayType = 6
	ExdDisplayTypeStringList        ExdDisplayType = 7
	ExdDisplayTypeString            ExdDisplayType = 8
	ExdDisplayTypeSimpleDynamicIcon ExdDisplayType = 9
	ExdDisplayTypeGauge             ExdDisplayType = 10
	ExdDisplayTypeInvalid           ExdDisplayType = 0xFF // INVALID
)

var exddisplaytypetostrs = map[ExdDisplayType]string{
	ExdDisplayTypeNumerical:         "numerical",
	ExdDisplayTypeSimple:            "simple",
	ExdDisplayTypeGraph:             "graph",
	ExdDisplayTypeBar:               "bar",
	ExdDisplayTypeCircleGraph:       "circle_graph",
	ExdDisplayTypeVirtualPartner:    "virtual_partner",
	ExdDisplayTypeBalance:           "balance",
	ExdDisplayTypeStringList:        "string_list",
	ExdDisplayTypeString:            "string",
	ExdDisplayTypeSimpleDynamicIcon: "simple_dynamic_icon",
	ExdDisplayTypeGauge:             "gauge",
	ExdDisplayTypeInvalid:           "invalid",
}

func (e ExdDisplayType) String() string {
	val, ok := exddisplaytypetostrs[e]
	if !ok {
		return strconv.Itoa(int(e))
	}
	return val
}

var strtoexddisplaytype = func() map[string]ExdDisplayType {
	m := make(map[string]ExdDisplayType)
	for t, str := range exddisplaytypetostrs {
		m[str] = ExdDisplayType(t)
	}
	return m
}()

// FromString parse string into ExdDisplayType constant it's represent, return ExdDisplayTypeInvalid if not found.
func ExdDisplayTypeFromString(s string) ExdDisplayType {
	val, ok := strtoexddisplaytype[s]
	if !ok {
		return strtoexddisplaytype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListExdDisplayType() []ExdDisplayType {
	vs := make([]ExdDisplayType, 0, len(exddisplaytypetostrs))
	for i := range exddisplaytypetostrs {
		vs = append(vs, ExdDisplayType(i))
	}
	return vs
}
