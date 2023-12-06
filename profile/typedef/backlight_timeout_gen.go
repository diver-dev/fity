// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type BacklightTimeout uint8

const (
	BacklightTimeoutInfinite BacklightTimeout = 0    // Backlight stays on forever.
	BacklightTimeoutInvalid  BacklightTimeout = 0xFF // INVALID
)

var backlighttimeouttostrs = map[BacklightTimeout]string{
	BacklightTimeoutInfinite: "infinite",
	BacklightTimeoutInvalid:  "invalid",
}

func (b BacklightTimeout) String() string {
	val, ok := backlighttimeouttostrs[b]
	if !ok {
		return strconv.FormatUint(uint64(b), 10)
	}
	return val
}

var strtobacklighttimeout = func() map[string]BacklightTimeout {
	m := make(map[string]BacklightTimeout)
	for t, str := range backlighttimeouttostrs {
		m[str] = BacklightTimeout(t)
	}
	return m
}()

// FromString parse string into BacklightTimeout constant it's represent, return BacklightTimeoutInvalid if not found.
func BacklightTimeoutFromString(s string) BacklightTimeout {
	val, ok := strtobacklighttimeout[s]
	if !ok {
		return strtobacklighttimeout["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListBacklightTimeout() []BacklightTimeout {
	vs := make([]BacklightTimeout, 0, len(backlighttimeouttostrs))
	for i := range backlighttimeouttostrs {
		vs = append(vs, BacklightTimeout(i))
	}
	return vs
}
