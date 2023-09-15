// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.115

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type TricepsExtensionExerciseName uint16

const (
	TricepsExtensionExerciseNameBenchDip                                     TricepsExtensionExerciseName = 0
	TricepsExtensionExerciseNameWeightedBenchDip                             TricepsExtensionExerciseName = 1
	TricepsExtensionExerciseNameBodyWeightDip                                TricepsExtensionExerciseName = 2
	TricepsExtensionExerciseNameCableKickback                                TricepsExtensionExerciseName = 3
	TricepsExtensionExerciseNameCableLyingTricepsExtension                   TricepsExtensionExerciseName = 4
	TricepsExtensionExerciseNameCableOverheadTricepsExtension                TricepsExtensionExerciseName = 5
	TricepsExtensionExerciseNameDumbbellKickback                             TricepsExtensionExerciseName = 6
	TricepsExtensionExerciseNameDumbbellLyingTricepsExtension                TricepsExtensionExerciseName = 7
	TricepsExtensionExerciseNameEzBarOverheadTricepsExtension                TricepsExtensionExerciseName = 8
	TricepsExtensionExerciseNameInclineDip                                   TricepsExtensionExerciseName = 9
	TricepsExtensionExerciseNameWeightedInclineDip                           TricepsExtensionExerciseName = 10
	TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension            TricepsExtensionExerciseName = 11
	TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension             TricepsExtensionExerciseName = 12
	TricepsExtensionExerciseNameLyingEzBarTricepsExtension                   TricepsExtensionExerciseName = 13
	TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress   TricepsExtensionExerciseName = 14
	TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension             TricepsExtensionExerciseName = 15
	TricepsExtensionExerciseNameRecliningTricepsPress                        TricepsExtensionExerciseName = 16
	TricepsExtensionExerciseNameReverseGripPressdown                         TricepsExtensionExerciseName = 17
	TricepsExtensionExerciseNameReverseGripTricepsPressdown                  TricepsExtensionExerciseName = 18
	TricepsExtensionExerciseNameRopePressdown                                TricepsExtensionExerciseName = 19
	TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension        TricepsExtensionExerciseName = 20
	TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension       TricepsExtensionExerciseName = 21
	TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension          TricepsExtensionExerciseName = 22
	TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension     TricepsExtensionExerciseName = 23
	TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension    TricepsExtensionExerciseName = 24
	TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension TricepsExtensionExerciseName = 25
	TricepsExtensionExerciseNameSingleLegBenchDipAndKick                     TricepsExtensionExerciseName = 26
	TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick             TricepsExtensionExerciseName = 27
	TricepsExtensionExerciseNameSingleLegDip                                 TricepsExtensionExerciseName = 28
	TricepsExtensionExerciseNameWeightedSingleLegDip                         TricepsExtensionExerciseName = 29
	TricepsExtensionExerciseNameStaticLyingTricepsExtension                  TricepsExtensionExerciseName = 30
	TricepsExtensionExerciseNameSuspendedDip                                 TricepsExtensionExerciseName = 31
	TricepsExtensionExerciseNameWeightedSuspendedDip                         TricepsExtensionExerciseName = 32
	TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension       TricepsExtensionExerciseName = 33
	TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension          TricepsExtensionExerciseName = 34
	TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension       TricepsExtensionExerciseName = 35
	TricepsExtensionExerciseNameTabletopDip                                  TricepsExtensionExerciseName = 36
	TricepsExtensionExerciseNameWeightedTabletopDip                          TricepsExtensionExerciseName = 37
	TricepsExtensionExerciseNameTricepsExtensionOnFloor                      TricepsExtensionExerciseName = 38
	TricepsExtensionExerciseNameTricepsPressdown                             TricepsExtensionExerciseName = 39
	TricepsExtensionExerciseNameWeightedDip                                  TricepsExtensionExerciseName = 40
	TricepsExtensionExerciseNameInvalid                                      TricepsExtensionExerciseName = 0xFFFF // INVALID
)

var tricepsextensionexercisenametostrs = map[TricepsExtensionExerciseName]string{
	TricepsExtensionExerciseNameBenchDip:                                     "bench_dip",
	TricepsExtensionExerciseNameWeightedBenchDip:                             "weighted_bench_dip",
	TricepsExtensionExerciseNameBodyWeightDip:                                "body_weight_dip",
	TricepsExtensionExerciseNameCableKickback:                                "cable_kickback",
	TricepsExtensionExerciseNameCableLyingTricepsExtension:                   "cable_lying_triceps_extension",
	TricepsExtensionExerciseNameCableOverheadTricepsExtension:                "cable_overhead_triceps_extension",
	TricepsExtensionExerciseNameDumbbellKickback:                             "dumbbell_kickback",
	TricepsExtensionExerciseNameDumbbellLyingTricepsExtension:                "dumbbell_lying_triceps_extension",
	TricepsExtensionExerciseNameEzBarOverheadTricepsExtension:                "ez_bar_overhead_triceps_extension",
	TricepsExtensionExerciseNameInclineDip:                                   "incline_dip",
	TricepsExtensionExerciseNameWeightedInclineDip:                           "weighted_incline_dip",
	TricepsExtensionExerciseNameInclineEzBarLyingTricepsExtension:            "incline_ez_bar_lying_triceps_extension",
	TricepsExtensionExerciseNameLyingDumbbellPulloverToExtension:             "lying_dumbbell_pullover_to_extension",
	TricepsExtensionExerciseNameLyingEzBarTricepsExtension:                   "lying_ez_bar_triceps_extension",
	TricepsExtensionExerciseNameLyingTricepsExtensionToCloseGripBenchPress:   "lying_triceps_extension_to_close_grip_bench_press",
	TricepsExtensionExerciseNameOverheadDumbbellTricepsExtension:             "overhead_dumbbell_triceps_extension",
	TricepsExtensionExerciseNameRecliningTricepsPress:                        "reclining_triceps_press",
	TricepsExtensionExerciseNameReverseGripPressdown:                         "reverse_grip_pressdown",
	TricepsExtensionExerciseNameReverseGripTricepsPressdown:                  "reverse_grip_triceps_pressdown",
	TricepsExtensionExerciseNameRopePressdown:                                "rope_pressdown",
	TricepsExtensionExerciseNameSeatedBarbellOverheadTricepsExtension:        "seated_barbell_overhead_triceps_extension",
	TricepsExtensionExerciseNameSeatedDumbbellOverheadTricepsExtension:       "seated_dumbbell_overhead_triceps_extension",
	TricepsExtensionExerciseNameSeatedEzBarOverheadTricepsExtension:          "seated_ez_bar_overhead_triceps_extension",
	TricepsExtensionExerciseNameSeatedSingleArmOverheadDumbbellExtension:     "seated_single_arm_overhead_dumbbell_extension",
	TricepsExtensionExerciseNameSingleArmDumbbellOverheadTricepsExtension:    "single_arm_dumbbell_overhead_triceps_extension",
	TricepsExtensionExerciseNameSingleDumbbellSeatedOverheadTricepsExtension: "single_dumbbell_seated_overhead_triceps_extension",
	TricepsExtensionExerciseNameSingleLegBenchDipAndKick:                     "single_leg_bench_dip_and_kick",
	TricepsExtensionExerciseNameWeightedSingleLegBenchDipAndKick:             "weighted_single_leg_bench_dip_and_kick",
	TricepsExtensionExerciseNameSingleLegDip:                                 "single_leg_dip",
	TricepsExtensionExerciseNameWeightedSingleLegDip:                         "weighted_single_leg_dip",
	TricepsExtensionExerciseNameStaticLyingTricepsExtension:                  "static_lying_triceps_extension",
	TricepsExtensionExerciseNameSuspendedDip:                                 "suspended_dip",
	TricepsExtensionExerciseNameWeightedSuspendedDip:                         "weighted_suspended_dip",
	TricepsExtensionExerciseNameSwissBallDumbbellLyingTricepsExtension:       "swiss_ball_dumbbell_lying_triceps_extension",
	TricepsExtensionExerciseNameSwissBallEzBarLyingTricepsExtension:          "swiss_ball_ez_bar_lying_triceps_extension",
	TricepsExtensionExerciseNameSwissBallEzBarOverheadTricepsExtension:       "swiss_ball_ez_bar_overhead_triceps_extension",
	TricepsExtensionExerciseNameTabletopDip:                                  "tabletop_dip",
	TricepsExtensionExerciseNameWeightedTabletopDip:                          "weighted_tabletop_dip",
	TricepsExtensionExerciseNameTricepsExtensionOnFloor:                      "triceps_extension_on_floor",
	TricepsExtensionExerciseNameTricepsPressdown:                             "triceps_pressdown",
	TricepsExtensionExerciseNameWeightedDip:                                  "weighted_dip",
	TricepsExtensionExerciseNameInvalid:                                      "invalid",
}

func (t TricepsExtensionExerciseName) String() string {
	val, ok := tricepsextensionexercisenametostrs[t]
	if !ok {
		return strconv.FormatUint(uint64(t), 10)
	}
	return val
}

var strtotricepsextensionexercisename = func() map[string]TricepsExtensionExerciseName {
	m := make(map[string]TricepsExtensionExerciseName)
	for t, str := range tricepsextensionexercisenametostrs {
		m[str] = TricepsExtensionExerciseName(t)
	}
	return m
}()

// FromString parse string into TricepsExtensionExerciseName constant it's represent, return TricepsExtensionExerciseNameInvalid if not found.
func TricepsExtensionExerciseNameFromString(s string) TricepsExtensionExerciseName {
	val, ok := strtotricepsextensionexercisename[s]
	if !ok {
		return strtotricepsextensionexercisename["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListTricepsExtensionExerciseName() []TricepsExtensionExerciseName {
	vs := make([]TricepsExtensionExerciseName, 0, len(tricepsextensionexercisenametostrs))
	for i := range tricepsextensionexercisenametostrs {
		vs = append(vs, TricepsExtensionExerciseName(i))
	}
	return vs
}