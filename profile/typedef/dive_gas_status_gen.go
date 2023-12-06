// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DiveGasStatus byte

const (
	DiveGasStatusDisabled   DiveGasStatus = 0
	DiveGasStatusEnabled    DiveGasStatus = 1
	DiveGasStatusBackupOnly DiveGasStatus = 2
	DiveGasStatusInvalid    DiveGasStatus = 0xFF // INVALID
)

var divegasstatustostrs = map[DiveGasStatus]string{
	DiveGasStatusDisabled:   "disabled",
	DiveGasStatusEnabled:    "enabled",
	DiveGasStatusBackupOnly: "backup_only",
	DiveGasStatusInvalid:    "invalid",
}

func (d DiveGasStatus) String() string {
	val, ok := divegasstatustostrs[d]
	if !ok {
		return strconv.Itoa(int(d))
	}
	return val
}

var strtodivegasstatus = func() map[string]DiveGasStatus {
	m := make(map[string]DiveGasStatus)
	for t, str := range divegasstatustostrs {
		m[str] = DiveGasStatus(t)
	}
	return m
}()

// FromString parse string into DiveGasStatus constant it's represent, return DiveGasStatusInvalid if not found.
func DiveGasStatusFromString(s string) DiveGasStatus {
	val, ok := strtodivegasstatus[s]
	if !ok {
		return strtodivegasstatus["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListDiveGasStatus() []DiveGasStatus {
	vs := make([]DiveGasStatus, 0, len(divegasstatustostrs))
	for i := range divegasstatustostrs {
		vs = append(vs, DiveGasStatus(i))
	}
	return vs
}
