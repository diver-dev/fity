// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

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

// SleepAssessment is a SleepAssessment message.
type SleepAssessment struct {
	CombinedAwakeScore       uint8  // Average of awake_time_score and awakenings_count_score. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeTimeScore           uint8  // Score that evaluates the total time spent awake between sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCountScore     uint8  // Score that evaluates the number of awakenings that interrupt sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	DeepSleepScore           uint8  // Score that evaluates the amount of deep sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepDurationScore       uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	LightSleepScore          uint8  // Score that evaluates the amount of light sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	OverallSleepScore        uint8  // Total score that summarizes the overall quality of sleep, combining sleep duration and quality. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepQualityScore        uint8  // Score that evaluates the quality of sleep based on sleep stages, heart-rate variability and possible awakenings during the night. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRecoveryScore       uint8  // Score that evaluates stress and recovery during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	RemSleepScore            uint8  // Score that evaluates the amount of REM sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	SleepRestlessnessScore   uint8  // Score that evaluates the amount of restlessness during sleep. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AwakeningsCount          uint8  // The number of awakenings during sleep.
	InterruptionsScore       uint8  // Score that evaluates the sleep interruptions. If valid: 0 (worst) to 100 (best). If unknown: FIT_UINT8_INVALID.
	AverageStressDuringSleep uint16 // Scale: 100; Excludes stress during awake periods in the sleep window

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewSleepAssessment creates new SleepAssessment struct based on given mesg. If mesg is nil or mesg.Num is not equal to SleepAssessment mesg number, it will return nil.
func NewSleepAssessment(mesg proto.Message) *SleepAssessment {
	if mesg.Num != typedef.MesgNumSleepAssessment {
		return nil
	}

	vals := [256]any{ // Mark all values as invalid, replace only when specified.
		0:  basetype.Uint8Invalid,  /* CombinedAwakeScore */
		1:  basetype.Uint8Invalid,  /* AwakeTimeScore */
		2:  basetype.Uint8Invalid,  /* AwakeningsCountScore */
		3:  basetype.Uint8Invalid,  /* DeepSleepScore */
		4:  basetype.Uint8Invalid,  /* SleepDurationScore */
		5:  basetype.Uint8Invalid,  /* LightSleepScore */
		6:  basetype.Uint8Invalid,  /* OverallSleepScore */
		7:  basetype.Uint8Invalid,  /* SleepQualityScore */
		8:  basetype.Uint8Invalid,  /* SleepRecoveryScore */
		9:  basetype.Uint8Invalid,  /* RemSleepScore */
		10: basetype.Uint8Invalid,  /* SleepRestlessnessScore */
		11: basetype.Uint8Invalid,  /* AwakeningsCount */
		14: basetype.Uint8Invalid,  /* InterruptionsScore */
		15: basetype.Uint16Invalid, /* AverageStressDuringSleep */
	}

	for i := 0; i < len(mesg.Fields); i++ {
		if mesg.Fields[i].Value == nil {
			continue // keep the invalid value
		}
		vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
	}

	return &SleepAssessment{
		CombinedAwakeScore:       typeconv.ToUint8[uint8](vals[0]),
		AwakeTimeScore:           typeconv.ToUint8[uint8](vals[1]),
		AwakeningsCountScore:     typeconv.ToUint8[uint8](vals[2]),
		DeepSleepScore:           typeconv.ToUint8[uint8](vals[3]),
		SleepDurationScore:       typeconv.ToUint8[uint8](vals[4]),
		LightSleepScore:          typeconv.ToUint8[uint8](vals[5]),
		OverallSleepScore:        typeconv.ToUint8[uint8](vals[6]),
		SleepQualityScore:        typeconv.ToUint8[uint8](vals[7]),
		SleepRecoveryScore:       typeconv.ToUint8[uint8](vals[8]),
		RemSleepScore:            typeconv.ToUint8[uint8](vals[9]),
		SleepRestlessnessScore:   typeconv.ToUint8[uint8](vals[10]),
		AwakeningsCount:          typeconv.ToUint8[uint8](vals[11]),
		InterruptionsScore:       typeconv.ToUint8[uint8](vals[14]),
		AverageStressDuringSleep: typeconv.ToUint16[uint16](vals[15]),

		DeveloperFields: mesg.DeveloperFields,
	}
}

// PutMessage puts fields's value into mesg. If mesg is nil or mesg.Num is not equal to SleepAssessment mesg number, it will return nil.
// It is the caller responsibility to provide the appropriate mesg, it's recommended to create mesg using factory:
//
//	factory.CreateMesg(typedef.MesgNumSleepAssessment)
func (m SleepAssessment) PutMessage(mesg *proto.Message) {
	if mesg == nil {
		return
	}

	if mesg.Num != typedef.MesgNumSleepAssessment {
		return
	}

	vals := [256]any{
		0:  m.CombinedAwakeScore,
		1:  m.AwakeTimeScore,
		2:  m.AwakeningsCountScore,
		3:  m.DeepSleepScore,
		4:  m.SleepDurationScore,
		5:  m.LightSleepScore,
		6:  m.OverallSleepScore,
		7:  m.SleepQualityScore,
		8:  m.SleepRecoveryScore,
		9:  m.RemSleepScore,
		10: m.SleepRestlessnessScore,
		11: m.AwakeningsCount,
		14: m.InterruptionsScore,
		15: m.AverageStressDuringSleep,
	}

	for i := range mesg.Fields {
		mesg.Fields[i].Value = vals[mesg.Fields[i].Num]
	}
	mesg.DeveloperFields = m.DeveloperFields

}
