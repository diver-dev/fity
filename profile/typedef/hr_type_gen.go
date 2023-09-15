// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type HrType byte

const (
	HrTypeNormal    HrType = 0
	HrTypeIrregular HrType = 1
	HrTypeInvalid   HrType = 0xFF // INVALID
)

var hrtypetostrs = map[HrType]string{
	HrTypeNormal:    "normal",
	HrTypeIrregular: "irregular",
	HrTypeInvalid:   "invalid",
}

func (h HrType) String() string {
	val, ok := hrtypetostrs[h]
	if !ok {
		return strconv.Itoa(int(h))
	}
	return val
}

var strtohrtype = func() map[string]HrType {
	m := make(map[string]HrType)
	for t, str := range hrtypetostrs {
		m[str] = HrType(t)
	}
	return m
}()

// FromString parse string into HrType constant it's represent, return HrTypeInvalid if not found.
func HrTypeFromString(s string) HrType {
	val, ok := strtohrtype[s]
	if !ok {
		return strtohrtype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListHrType() []HrType {
	vs := make([]HrType, 0, len(hrtypetostrs))
	for i := range hrtypetostrs {
		vs = append(vs, HrType(i))
	}
	return vs
}
