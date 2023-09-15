// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DiveGasMode byte

const (
	DiveGasModeOpenCircuit          DiveGasMode = 0
	DiveGasModeClosedCircuitDiluent DiveGasMode = 1
	DiveGasModeInvalid              DiveGasMode = 0xFF // INVALID
)

var divegasmodetostrs = map[DiveGasMode]string{
	DiveGasModeOpenCircuit:          "open_circuit",
	DiveGasModeClosedCircuitDiluent: "closed_circuit_diluent",
	DiveGasModeInvalid:              "invalid",
}

func (d DiveGasMode) String() string {
	val, ok := divegasmodetostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodivegasmode = func() map[string]DiveGasMode {
	m := make(map[string]DiveGasMode)
	for t, str := range divegasmodetostrs {
		m[str] = DiveGasMode(t)
	}
	return m
}()

// FromString parse string into DiveGasMode constant it's represent, return DiveGasModeInvalid if not found.
func DiveGasModeFromString(s string) DiveGasMode {
	val, ok := strtodivegasmode[s]
	if !ok {
		return strtodivegasmode["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDiveGasMode() []DiveGasMode {
	vs := make([]DiveGasMode, 0, len(divegasmodetostrs))
	for i := range divegasmodetostrs {
		vs = append(vs, DiveGasMode(i))
	}
	return vs
}
