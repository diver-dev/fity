// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type CommTimeoutType uint16

const (
	CommTimeoutTypeWildcardPairingTimeout CommTimeoutType = 0      // Timeout pairing to any device
	CommTimeoutTypePairingTimeout         CommTimeoutType = 1      // Timeout pairing to previously paired device
	CommTimeoutTypeConnectionLost         CommTimeoutType = 2      // Temporary loss of communications
	CommTimeoutTypeConnectionTimeout      CommTimeoutType = 3      // Connection closed due to extended bad communications
	CommTimeoutTypeInvalid                CommTimeoutType = 0xFFFF // INVALID
)

var commtimeouttypetostrs = map[CommTimeoutType]string{
	CommTimeoutTypeWildcardPairingTimeout: "wildcard_pairing_timeout",
	CommTimeoutTypePairingTimeout:         "pairing_timeout",
	CommTimeoutTypeConnectionLost:         "connection_lost",
	CommTimeoutTypeConnectionTimeout:      "connection_timeout",
	CommTimeoutTypeInvalid:                "invalid",
}

func (c CommTimeoutType) String() string {
	val, ok := commtimeouttypetostrs[c]
	if !ok {
		return strconv.FormatUint(uint64(c), 10)
	}
	return val
}

var strtocommtimeouttype = func() map[string]CommTimeoutType {
	m := make(map[string]CommTimeoutType)
	for t, str := range commtimeouttypetostrs {
		m[str] = CommTimeoutType(t)
	}
	return m
}()

// FromString parse string into CommTimeoutType constant it's represent, return CommTimeoutTypeInvalid if not found.
func CommTimeoutTypeFromString(s string) CommTimeoutType {
	val, ok := strtocommtimeouttype[s]
	if !ok {
		return strtocommtimeouttype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListCommTimeoutType() []CommTimeoutType {
	vs := make([]CommTimeoutType, 0, len(commtimeouttypetostrs))
	for i := range commtimeouttypetostrs {
		vs = append(vs, CommTimeoutType(i))
	}
	return vs
}
