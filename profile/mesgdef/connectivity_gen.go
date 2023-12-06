// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/kit/typeconv"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Connectivity is a Connectivity message.
type Connectivity struct {
	BluetoothEnabled            bool // Use Bluetooth for connectivity features
	BluetoothLeEnabled          bool // Use Bluetooth Low Energy for connectivity features
	AntEnabled                  bool // Use ANT for connectivity features
	Name                        string
	LiveTrackingEnabled         bool
	WeatherConditionsEnabled    bool
	WeatherAlertsEnabled        bool
	AutoActivityUploadEnabled   bool
	CourseDownloadEnabled       bool
	WorkoutDownloadEnabled      bool
	GpsEphemerisDownloadEnabled bool
	IncidentDetectionEnabled    bool
	GrouptrackEnabled           bool

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewConnectivity creates new Connectivity struct based on given mesg. If mesg is nil or mesg.Num is not equal to Connectivity mesg number, it will return nil.
func NewConnectivity(mesg proto.Message) *Connectivity {
	if mesg.Num != typedef.MesgNumConnectivity {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0:  false,                  /* BluetoothEnabled */
		1:  false,                  /* BluetoothLeEnabled */
		2:  false,                  /* AntEnabled */
		3:  basetype.StringInvalid, /* Name */
		4:  false,                  /* LiveTrackingEnabled */
		5:  false,                  /* WeatherConditionsEnabled */
		6:  false,                  /* WeatherAlertsEnabled */
		7:  false,                  /* AutoActivityUploadEnabled */
		8:  false,                  /* CourseDownloadEnabled */
		9:  false,                  /* WorkoutDownloadEnabled */
		10: false,                  /* GpsEphemerisDownloadEnabled */
		11: false,                  /* IncidentDetectionEnabled */
		12: false,                  /* GrouptrackEnabled */
	}

	for i := range mesg.Fields {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &Connectivity{
		BluetoothEnabled:            typeconv.ToBool[bool](vals[0]),
		BluetoothLeEnabled:          typeconv.ToBool[bool](vals[1]),
		AntEnabled:                  typeconv.ToBool[bool](vals[2]),
		Name:                        typeconv.ToString[string](vals[3]),
		LiveTrackingEnabled:         typeconv.ToBool[bool](vals[4]),
		WeatherConditionsEnabled:    typeconv.ToBool[bool](vals[5]),
		WeatherAlertsEnabled:        typeconv.ToBool[bool](vals[6]),
		AutoActivityUploadEnabled:   typeconv.ToBool[bool](vals[7]),
		CourseDownloadEnabled:       typeconv.ToBool[bool](vals[8]),
		WorkoutDownloadEnabled:      typeconv.ToBool[bool](vals[9]),
		GpsEphemerisDownloadEnabled: typeconv.ToBool[bool](vals[10]),
		IncidentDetectionEnabled:    typeconv.ToBool[bool](vals[11]),
		GrouptrackEnabled:           typeconv.ToBool[bool](vals[12]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to Connectivity mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumConnectivity)
func (m Connectivity) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumConnectivity {
		return
	}

	vals := [256]any{
		0:  m.BluetoothEnabled,
		1:  m.BluetoothLeEnabled,
		2:  m.AntEnabled,
		3:  m.Name,
		4:  m.LiveTrackingEnabled,
		5:  m.WeatherConditionsEnabled,
		6:  m.WeatherAlertsEnabled,
		7:  m.AutoActivityUploadEnabled,
		8:  m.CourseDownloadEnabled,
		9:  m.WorkoutDownloadEnabled,
		10: m.GpsEphemerisDownloadEnabled,
		11: m.IncidentDetectionEnabled,
		12: m.GrouptrackEnabled,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
