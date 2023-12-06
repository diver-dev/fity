// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type LocaltimeIntoDay uint32

const (
	LocaltimeIntoDayInvalid LocaltimeIntoDay = 0xFFFFFFFF // INVALID
)

var localtimeintodaytostrs = map[LocaltimeIntoDay]string{
	LocaltimeIntoDayInvalid: "invalid",
}

func (l LocaltimeIntoDay) String() string {
	val, ok := localtimeintodaytostrs[l]
	if !ok {
		return strconv.FormatUint(uint64(l), 10)
	}
	return val
}

var strtolocaltimeintoday = func() map[string]LocaltimeIntoDay {
	m := make(map[string]LocaltimeIntoDay)
	for t, str := range localtimeintodaytostrs {
		m[str] = LocaltimeIntoDay(t)
	}
	return m
}()

// FromString parse string into LocaltimeIntoDay constant it's represent, return LocaltimeIntoDayInvalid if not found.
func LocaltimeIntoDayFromString(s string) LocaltimeIntoDay {
	val, ok := strtolocaltimeintoday[s]
	if !ok {
		return strtolocaltimeintoday["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListLocaltimeIntoDay() []LocaltimeIntoDay {
	vs := make([]LocaltimeIntoDay, 0, len(localtimeintodaytostrs))
	for i := range localtimeintodaytostrs {
		vs = append(vs, LocaltimeIntoDay(i))
	}
	return vs
}
