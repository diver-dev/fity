// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type SensorType byte

const (
	SensorTypeAccelerometer SensorType = 0
	SensorTypeGyroscope     SensorType = 1
	SensorTypeCompass       SensorType = 2 // Magnetometer
	SensorTypeBarometer     SensorType = 3
	SensorTypeInvalid       SensorType = 0xFF // INVALID
)

var sensortypetostrs = map[SensorType]string{
	SensorTypeAccelerometer: "accelerometer",
	SensorTypeGyroscope:     "gyroscope",
	SensorTypeCompass:       "compass",
	SensorTypeBarometer:     "barometer",
	SensorTypeInvalid:       "invalid",
}

func (s SensorType) String() string {
	val, ok := sensortypetostrs[s]
	if !ok {
		return strconv.Itoa(int(s))
	}
	return val
}

var strtosensortype = func() map[string]SensorType {
	m := make(map[string]SensorType)
	for t, str := range sensortypetostrs {
		m[str] = SensorType(t)
	}
	return m
}()

// FromString parse string into SensorType constant it's represent, return SensorTypeInvalid if not found.
func SensorTypeFromString(s string) SensorType {
	val, ok := strtosensortype[s]
	if !ok {
		return strtosensortype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListSensorType() []SensorType {
	vs := make([]SensorType, 0, len(sensortypetostrs))
	for i := range sensortypetostrs {
		vs = append(vs, SensorType(i))
	}
	return vs
}