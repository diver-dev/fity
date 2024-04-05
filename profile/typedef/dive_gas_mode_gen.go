// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
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
	DiveGasModeInvalid              DiveGasMode = 0xFF
)

func (d DiveGasMode) Byte() byte { return byte(d) }

func (d DiveGasMode) String() string {
	switch d {
	case DiveGasModeOpenCircuit:
		return "open_circuit"
	case DiveGasModeClosedCircuitDiluent:
		return "closed_circuit_diluent"
	default:
		return "DiveGasModeInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DiveGasMode constant it's represent, return DiveGasModeInvalid if not found.
func DiveGasModeFromString(s string) DiveGasMode {
	switch s {
	case "open_circuit":
		return DiveGasModeOpenCircuit
	case "closed_circuit_diluent":
		return DiveGasModeClosedCircuitDiluent
	default:
		return DiveGasModeInvalid
	}
}

// List returns all constants.
func ListDiveGasMode() []DiveGasMode {
	return []DiveGasMode{
		DiveGasModeOpenCircuit,
		DiveGasModeClosedCircuitDiluent,
	}
}
