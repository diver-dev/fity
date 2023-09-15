// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type EventType byte

const (
	EventTypeStart                  EventType = 0
	EventTypeStop                   EventType = 1
	EventTypeConsecutiveDepreciated EventType = 2
	EventTypeMarker                 EventType = 3
	EventTypeStopAll                EventType = 4
	EventTypeBeginDepreciated       EventType = 5
	EventTypeEndDepreciated         EventType = 6
	EventTypeEndAllDepreciated      EventType = 7
	EventTypeStopDisable            EventType = 8
	EventTypeStopDisableAll         EventType = 9
	EventTypeInvalid                EventType = 0xFF // INVALID
)

var eventtypetostrs = map[EventType]string{
	EventTypeStart:                  "start",
	EventTypeStop:                   "stop",
	EventTypeConsecutiveDepreciated: "consecutive_depreciated",
	EventTypeMarker:                 "marker",
	EventTypeStopAll:                "stop_all",
	EventTypeBeginDepreciated:       "begin_depreciated",
	EventTypeEndDepreciated:         "end_depreciated",
	EventTypeEndAllDepreciated:      "end_all_depreciated",
	EventTypeStopDisable:            "stop_disable",
	EventTypeStopDisableAll:         "stop_disable_all",
	EventTypeInvalid:                "invalid",
}

func (e EventType) String() string {
	val, ok := eventtypetostrs[e]
	if !ok {
		return strconv.Itoa(int(e))
	}
	return val
}

var strtoeventtype = func() map[string]EventType {
	m := make(map[string]EventType)
	for t, str := range eventtypetostrs {
		m[str] = EventType(t)
	}
	return m
}()

// FromString parse string into EventType constant it's represent, return EventTypeInvalid if not found.
func EventTypeFromString(s string) EventType {
	val, ok := strtoeventtype[s]
	if !ok {
		return strtoeventtype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListEventType() []EventType {
	vs := make([]EventType, 0, len(eventtypetostrs))
	for i := range eventtypetostrs {
		vs = append(vs, EventType(i))
	}
	return vs
}
