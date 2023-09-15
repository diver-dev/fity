// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LanguageBits4 uint8

const (
	LanguageBits4BrazilianPortuguese LanguageBits4 = 0x01
	LanguageBits4Indonesian          LanguageBits4 = 0x02
	LanguageBits4Malaysian           LanguageBits4 = 0x04
	LanguageBits4Vietnamese          LanguageBits4 = 0x08
	LanguageBits4Burmese             LanguageBits4 = 0x10
	LanguageBits4Mongolian           LanguageBits4 = 0x20
	LanguageBits4Invalid             LanguageBits4 = 0x0 // INVALID
)

var languagebits4tostrs = map[LanguageBits4]string{
	LanguageBits4BrazilianPortuguese: "brazilian_portuguese",
	LanguageBits4Indonesian:          "indonesian",
	LanguageBits4Malaysian:           "malaysian",
	LanguageBits4Vietnamese:          "vietnamese",
	LanguageBits4Burmese:             "burmese",
	LanguageBits4Mongolian:           "mongolian",
	LanguageBits4Invalid:             "invalid",
}

func (l LanguageBits4) String() string {
	val, ok := languagebits4tostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolanguagebits4 = func() map[string]LanguageBits4 {
	m := make(map[string]LanguageBits4)
	for t, str := range languagebits4tostrs {
		m[str] = LanguageBits4(t)
	}
	return m
}()

// FromString parse string into LanguageBits4 constant it's represent, return LanguageBits4Invalid if not found.
func LanguageBits4FromString(s string) LanguageBits4 {
	val, ok := strtolanguagebits4[s]
	if !ok {
		return strtolanguagebits4["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLanguageBits4() []LanguageBits4 {
	vs := make([]LanguageBits4, 0, len(languagebits4tostrs))
	for i := range languagebits4tostrs {
		vs = append(vs, LanguageBits4(i))
	}
	return vs
}