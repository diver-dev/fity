// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.
// SDK Version: 21.126

// Copyright 2023 The Fit SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package profile

import (
	"strconv"

	"github.com/muktihari/fit/profile/basetype"
)

type ProfileType uint16

const (
	Enum ProfileType = iota
	Sint8
	Uint8
	Sint16
	Uint16
	Sint32
	Uint32
	String
	Float32
	Float64
	Uint8z
	Uint16z
	Uint32z
	Byte
	Sint64
	Uint64
	Uint64z
	Bool
	File
	MesgNum
	Checksum
	FileFlags
	MesgCount
	DateTime
	LocalDateTime
	MessageIndex
	DeviceIndex
	Gender
	Language
	LanguageBits0
	LanguageBits1
	LanguageBits2
	LanguageBits3
	LanguageBits4
	TimeZone
	DisplayMeasure
	DisplayHeart
	DisplayPower
	DisplayPosition
	Switch
	Sport
	SportBits0
	SportBits1
	SportBits2
	SportBits3
	SportBits4
	SportBits5
	SportBits6
	SubSport
	SportEvent
	Activity
	Intensity
	SessionTrigger
	AutolapTrigger
	LapTrigger
	TimeMode
	BacklightMode
	DateMode
	BacklightTimeout
	Event
	EventType
	TimerTrigger
	FitnessEquipmentState
	Tone
	Autoscroll
	ActivityClass
	HrZoneCalc
	PwrZoneCalc
	WktStepDuration
	WktStepTarget
	Goal
	GoalRecurrence
	GoalSource
	Schedule
	CoursePoint
	Manufacturer
	GarminProduct
	AntplusDeviceType
	AntNetwork
	WorkoutCapabilities
	BatteryStatus
	HrType
	CourseCapabilities
	Weight
	WorkoutHr
	WorkoutPower
	BpStatus
	UserLocalId
	SwimStroke
	ActivityType
	ActivitySubtype
	ActivityLevel
	Side
	LeftRightBalance
	LeftRightBalance100
	LengthType
	DayOfWeek
	ConnectivityCapabilities
	WeatherReport
	WeatherStatus
	WeatherSeverity
	WeatherSevereType
	TimeIntoDay
	LocaltimeIntoDay
	StrokeType
	BodyLocation
	SegmentLapStatus
	SegmentLeaderboardType
	SegmentDeleteStatus
	SegmentSelectionType
	SourceType
	LocalDeviceType
	BleDeviceType
	AntChannelId
	DisplayOrientation
	WorkoutEquipment
	WatchfaceMode
	DigitalWatchfaceLayout
	AnalogWatchfaceLayout
	RiderPositionType
	PowerPhaseType
	CameraEventType
	SensorType
	BikeLightNetworkConfigType
	CommTimeoutType
	CameraOrientationType
	AttitudeStage
	AttitudeValidity
	AutoSyncFrequency
	ExdLayout
	ExdDisplayType
	ExdDataUnits
	ExdQualifiers
	ExdDescriptors
	AutoActivityDetect
	SupportedExdScreenLayouts
	FitBaseType
	TurnType
	BikeLightBeamAngleMode
	FitBaseUnit
	SetType
	MaxMetCategory
	ExerciseCategory
	BenchPressExerciseName
	CalfRaiseExerciseName
	CardioExerciseName
	CarryExerciseName
	ChopExerciseName
	CoreExerciseName
	CrunchExerciseName
	CurlExerciseName
	DeadliftExerciseName
	FlyeExerciseName
	HipRaiseExerciseName
	HipStabilityExerciseName
	HipSwingExerciseName
	HyperextensionExerciseName
	LateralRaiseExerciseName
	LegCurlExerciseName
	LegRaiseExerciseName
	LungeExerciseName
	OlympicLiftExerciseName
	PlankExerciseName
	PlyoExerciseName
	PullUpExerciseName
	PushUpExerciseName
	RowExerciseName
	ShoulderPressExerciseName
	ShoulderStabilityExerciseName
	ShrugExerciseName
	SitUpExerciseName
	SquatExerciseName
	TotalBodyExerciseName
	TricepsExtensionExerciseName
	WarmUpExerciseName
	RunExerciseName
	WaterType
	TissueModelType
	DiveGasStatus
	DiveAlert
	DiveAlarmType
	DiveBacklightMode
	SleepLevel
	Spo2MeasurementType
	CcrSetpointSwitchMode
	DiveGasMode
	FaveroProduct
	SplitType
	ClimbProEvent
	GasConsumptionRateType
	TapSensitivity
	RadarThreatLevelType
	MaxMetSpeedSource
	MaxMetHeartRateSource
	HrvStatus
	NoFlyTimeMode
	Invalid // INVALID
)

var profiletypetostrs = [...]string{
	Enum:                          "enum",
	Sint8:                         "sint8",
	Uint8:                         "uint8",
	Sint16:                        "sint16",
	Uint16:                        "uint16",
	Sint32:                        "sint32",
	Uint32:                        "uint32",
	String:                        "string",
	Float32:                       "float32",
	Float64:                       "float64",
	Uint8z:                        "uint8z",
	Uint16z:                       "uint16z",
	Uint32z:                       "uint32z",
	Byte:                          "byte",
	Sint64:                        "sint64",
	Uint64:                        "uint64",
	Uint64z:                       "uint64z",
	Bool:                          "bool",
	File:                          "file",
	MesgNum:                       "mesg_num",
	Checksum:                      "checksum",
	FileFlags:                     "file_flags",
	MesgCount:                     "mesg_count",
	DateTime:                      "date_time",
	LocalDateTime:                 "local_date_time",
	MessageIndex:                  "message_index",
	DeviceIndex:                   "device_index",
	Gender:                        "gender",
	Language:                      "language",
	LanguageBits0:                 "language_bits_0",
	LanguageBits1:                 "language_bits_1",
	LanguageBits2:                 "language_bits_2",
	LanguageBits3:                 "language_bits_3",
	LanguageBits4:                 "language_bits_4",
	TimeZone:                      "time_zone",
	DisplayMeasure:                "display_measure",
	DisplayHeart:                  "display_heart",
	DisplayPower:                  "display_power",
	DisplayPosition:               "display_position",
	Switch:                        "switch",
	Sport:                         "sport",
	SportBits0:                    "sport_bits_0",
	SportBits1:                    "sport_bits_1",
	SportBits2:                    "sport_bits_2",
	SportBits3:                    "sport_bits_3",
	SportBits4:                    "sport_bits_4",
	SportBits5:                    "sport_bits_5",
	SportBits6:                    "sport_bits_6",
	SubSport:                      "sub_sport",
	SportEvent:                    "sport_event",
	Activity:                      "activity",
	Intensity:                     "intensity",
	SessionTrigger:                "session_trigger",
	AutolapTrigger:                "autolap_trigger",
	LapTrigger:                    "lap_trigger",
	TimeMode:                      "time_mode",
	BacklightMode:                 "backlight_mode",
	DateMode:                      "date_mode",
	BacklightTimeout:              "backlight_timeout",
	Event:                         "event",
	EventType:                     "event_type",
	TimerTrigger:                  "timer_trigger",
	FitnessEquipmentState:         "fitness_equipment_state",
	Tone:                          "tone",
	Autoscroll:                    "autoscroll",
	ActivityClass:                 "activity_class",
	HrZoneCalc:                    "hr_zone_calc",
	PwrZoneCalc:                   "pwr_zone_calc",
	WktStepDuration:               "wkt_step_duration",
	WktStepTarget:                 "wkt_step_target",
	Goal:                          "goal",
	GoalRecurrence:                "goal_recurrence",
	GoalSource:                    "goal_source",
	Schedule:                      "schedule",
	CoursePoint:                   "course_point",
	Manufacturer:                  "manufacturer",
	GarminProduct:                 "garmin_product",
	AntplusDeviceType:             "antplus_device_type",
	AntNetwork:                    "ant_network",
	WorkoutCapabilities:           "workout_capabilities",
	BatteryStatus:                 "battery_status",
	HrType:                        "hr_type",
	CourseCapabilities:            "course_capabilities",
	Weight:                        "weight",
	WorkoutHr:                     "workout_hr",
	WorkoutPower:                  "workout_power",
	BpStatus:                      "bp_status",
	UserLocalId:                   "user_local_id",
	SwimStroke:                    "swim_stroke",
	ActivityType:                  "activity_type",
	ActivitySubtype:               "activity_subtype",
	ActivityLevel:                 "activity_level",
	Side:                          "side",
	LeftRightBalance:              "left_right_balance",
	LeftRightBalance100:           "left_right_balance_100",
	LengthType:                    "length_type",
	DayOfWeek:                     "day_of_week",
	ConnectivityCapabilities:      "connectivity_capabilities",
	WeatherReport:                 "weather_report",
	WeatherStatus:                 "weather_status",
	WeatherSeverity:               "weather_severity",
	WeatherSevereType:             "weather_severe_type",
	TimeIntoDay:                   "time_into_day",
	LocaltimeIntoDay:              "localtime_into_day",
	StrokeType:                    "stroke_type",
	BodyLocation:                  "body_location",
	SegmentLapStatus:              "segment_lap_status",
	SegmentLeaderboardType:        "segment_leaderboard_type",
	SegmentDeleteStatus:           "segment_delete_status",
	SegmentSelectionType:          "segment_selection_type",
	SourceType:                    "source_type",
	LocalDeviceType:               "local_device_type",
	BleDeviceType:                 "ble_device_type",
	AntChannelId:                  "ant_channel_id",
	DisplayOrientation:            "display_orientation",
	WorkoutEquipment:              "workout_equipment",
	WatchfaceMode:                 "watchface_mode",
	DigitalWatchfaceLayout:        "digital_watchface_layout",
	AnalogWatchfaceLayout:         "analog_watchface_layout",
	RiderPositionType:             "rider_position_type",
	PowerPhaseType:                "power_phase_type",
	CameraEventType:               "camera_event_type",
	SensorType:                    "sensor_type",
	BikeLightNetworkConfigType:    "bike_light_network_config_type",
	CommTimeoutType:               "comm_timeout_type",
	CameraOrientationType:         "camera_orientation_type",
	AttitudeStage:                 "attitude_stage",
	AttitudeValidity:              "attitude_validity",
	AutoSyncFrequency:             "auto_sync_frequency",
	ExdLayout:                     "exd_layout",
	ExdDisplayType:                "exd_display_type",
	ExdDataUnits:                  "exd_data_units",
	ExdQualifiers:                 "exd_qualifiers",
	ExdDescriptors:                "exd_descriptors",
	AutoActivityDetect:            "auto_activity_detect",
	SupportedExdScreenLayouts:     "supported_exd_screen_layouts",
	FitBaseType:                   "fit_base_type",
	TurnType:                      "turn_type",
	BikeLightBeamAngleMode:        "bike_light_beam_angle_mode",
	FitBaseUnit:                   "fit_base_unit",
	SetType:                       "set_type",
	MaxMetCategory:                "max_met_category",
	ExerciseCategory:              "exercise_category",
	BenchPressExerciseName:        "bench_press_exercise_name",
	CalfRaiseExerciseName:         "calf_raise_exercise_name",
	CardioExerciseName:            "cardio_exercise_name",
	CarryExerciseName:             "carry_exercise_name",
	ChopExerciseName:              "chop_exercise_name",
	CoreExerciseName:              "core_exercise_name",
	CrunchExerciseName:            "crunch_exercise_name",
	CurlExerciseName:              "curl_exercise_name",
	DeadliftExerciseName:          "deadlift_exercise_name",
	FlyeExerciseName:              "flye_exercise_name",
	HipRaiseExerciseName:          "hip_raise_exercise_name",
	HipStabilityExerciseName:      "hip_stability_exercise_name",
	HipSwingExerciseName:          "hip_swing_exercise_name",
	HyperextensionExerciseName:    "hyperextension_exercise_name",
	LateralRaiseExerciseName:      "lateral_raise_exercise_name",
	LegCurlExerciseName:           "leg_curl_exercise_name",
	LegRaiseExerciseName:          "leg_raise_exercise_name",
	LungeExerciseName:             "lunge_exercise_name",
	OlympicLiftExerciseName:       "olympic_lift_exercise_name",
	PlankExerciseName:             "plank_exercise_name",
	PlyoExerciseName:              "plyo_exercise_name",
	PullUpExerciseName:            "pull_up_exercise_name",
	PushUpExerciseName:            "push_up_exercise_name",
	RowExerciseName:               "row_exercise_name",
	ShoulderPressExerciseName:     "shoulder_press_exercise_name",
	ShoulderStabilityExerciseName: "shoulder_stability_exercise_name",
	ShrugExerciseName:             "shrug_exercise_name",
	SitUpExerciseName:             "sit_up_exercise_name",
	SquatExerciseName:             "squat_exercise_name",
	TotalBodyExerciseName:         "total_body_exercise_name",
	TricepsExtensionExerciseName:  "triceps_extension_exercise_name",
	WarmUpExerciseName:            "warm_up_exercise_name",
	RunExerciseName:               "run_exercise_name",
	WaterType:                     "water_type",
	TissueModelType:               "tissue_model_type",
	DiveGasStatus:                 "dive_gas_status",
	DiveAlert:                     "dive_alert",
	DiveAlarmType:                 "dive_alarm_type",
	DiveBacklightMode:             "dive_backlight_mode",
	SleepLevel:                    "sleep_level",
	Spo2MeasurementType:           "spo2_measurement_type",
	CcrSetpointSwitchMode:         "ccr_setpoint_switch_mode",
	DiveGasMode:                   "dive_gas_mode",
	FaveroProduct:                 "favero_product",
	SplitType:                     "split_type",
	ClimbProEvent:                 "climb_pro_event",
	GasConsumptionRateType:        "gas_consumption_rate_type",
	TapSensitivity:                "tap_sensitivity",
	RadarThreatLevelType:          "radar_threat_level_type",
	MaxMetSpeedSource:             "max_met_speed_source",
	MaxMetHeartRateSource:         "max_met_heart_rate_source",
	HrvStatus:                     "hrv_status",
	NoFlyTimeMode:                 "no_fly_time_mode",
	Invalid:                       "invalid",
}

func (p ProfileType) String() string {
	if p >= ProfileType(len(profiletypetostrs)) {
		return strconv.FormatUint(uint64(p), 10)
	}
	return profiletypetostrs[p]
}

var strtoprofiletype = func() map[string]ProfileType {
	m := make(map[string]ProfileType)
	for t, str := range profiletypetostrs {
		m[str] = ProfileType(t)
	}
	return m
}()

// FromString parse string into ProfileType constant it's represent, return ProfileTypeInvalid if not found.
func ProfileTypeFromString(s string) ProfileType {
	val, ok := strtoprofiletype[s]
	if !ok {
		return strtoprofiletype["invalid"]
	}
	return val
}

// List returns all constants. The result might be unsorted (depend on stringer is in array or map), it's up to the caller to sort.
func ListProfileType() []ProfileType {
	vs := make([]ProfileType, 0, len(profiletypetostrs))
	for i := range profiletypetostrs {
		vs = append(vs, ProfileType(i))
	}
	return vs
}

var mappingtobasetypes = [...]basetype.BaseType{
	Enum:                          basetype.Enum,
	Sint8:                         basetype.Sint8,
	Uint8:                         basetype.Uint8,
	Sint16:                        basetype.Sint16,
	Uint16:                        basetype.Uint16,
	Sint32:                        basetype.Sint32,
	Uint32:                        basetype.Uint32,
	String:                        basetype.String,
	Float32:                       basetype.Float32,
	Float64:                       basetype.Float64,
	Uint8z:                        basetype.Uint8z,
	Uint16z:                       basetype.Uint16z,
	Uint32z:                       basetype.Uint32z,
	Byte:                          basetype.Byte,
	Sint64:                        basetype.Sint64,
	Uint64:                        basetype.Uint64,
	Uint64z:                       basetype.Uint64z,
	Bool:                          basetype.Enum,
	File:                          basetype.Enum,
	MesgNum:                       basetype.Uint16,
	Checksum:                      basetype.Uint8,
	FileFlags:                     basetype.Uint8z,
	MesgCount:                     basetype.Enum,
	DateTime:                      basetype.Uint32,
	LocalDateTime:                 basetype.Uint32,
	MessageIndex:                  basetype.Uint16,
	DeviceIndex:                   basetype.Uint8,
	Gender:                        basetype.Enum,
	Language:                      basetype.Enum,
	LanguageBits0:                 basetype.Uint8z,
	LanguageBits1:                 basetype.Uint8z,
	LanguageBits2:                 basetype.Uint8z,
	LanguageBits3:                 basetype.Uint8z,
	LanguageBits4:                 basetype.Uint8z,
	TimeZone:                      basetype.Enum,
	DisplayMeasure:                basetype.Enum,
	DisplayHeart:                  basetype.Enum,
	DisplayPower:                  basetype.Enum,
	DisplayPosition:               basetype.Enum,
	Switch:                        basetype.Enum,
	Sport:                         basetype.Enum,
	SportBits0:                    basetype.Uint8z,
	SportBits1:                    basetype.Uint8z,
	SportBits2:                    basetype.Uint8z,
	SportBits3:                    basetype.Uint8z,
	SportBits4:                    basetype.Uint8z,
	SportBits5:                    basetype.Uint8z,
	SportBits6:                    basetype.Uint8z,
	SubSport:                      basetype.Enum,
	SportEvent:                    basetype.Enum,
	Activity:                      basetype.Enum,
	Intensity:                     basetype.Enum,
	SessionTrigger:                basetype.Enum,
	AutolapTrigger:                basetype.Enum,
	LapTrigger:                    basetype.Enum,
	TimeMode:                      basetype.Enum,
	BacklightMode:                 basetype.Enum,
	DateMode:                      basetype.Enum,
	BacklightTimeout:              basetype.Uint8,
	Event:                         basetype.Enum,
	EventType:                     basetype.Enum,
	TimerTrigger:                  basetype.Enum,
	FitnessEquipmentState:         basetype.Enum,
	Tone:                          basetype.Enum,
	Autoscroll:                    basetype.Enum,
	ActivityClass:                 basetype.Enum,
	HrZoneCalc:                    basetype.Enum,
	PwrZoneCalc:                   basetype.Enum,
	WktStepDuration:               basetype.Enum,
	WktStepTarget:                 basetype.Enum,
	Goal:                          basetype.Enum,
	GoalRecurrence:                basetype.Enum,
	GoalSource:                    basetype.Enum,
	Schedule:                      basetype.Enum,
	CoursePoint:                   basetype.Enum,
	Manufacturer:                  basetype.Uint16,
	GarminProduct:                 basetype.Uint16,
	AntplusDeviceType:             basetype.Uint8,
	AntNetwork:                    basetype.Enum,
	WorkoutCapabilities:           basetype.Uint32z,
	BatteryStatus:                 basetype.Uint8,
	HrType:                        basetype.Enum,
	CourseCapabilities:            basetype.Uint32z,
	Weight:                        basetype.Uint16,
	WorkoutHr:                     basetype.Uint32,
	WorkoutPower:                  basetype.Uint32,
	BpStatus:                      basetype.Enum,
	UserLocalId:                   basetype.Uint16,
	SwimStroke:                    basetype.Enum,
	ActivityType:                  basetype.Enum,
	ActivitySubtype:               basetype.Enum,
	ActivityLevel:                 basetype.Enum,
	Side:                          basetype.Enum,
	LeftRightBalance:              basetype.Uint8,
	LeftRightBalance100:           basetype.Uint16,
	LengthType:                    basetype.Enum,
	DayOfWeek:                     basetype.Enum,
	ConnectivityCapabilities:      basetype.Uint32z,
	WeatherReport:                 basetype.Enum,
	WeatherStatus:                 basetype.Enum,
	WeatherSeverity:               basetype.Enum,
	WeatherSevereType:             basetype.Enum,
	TimeIntoDay:                   basetype.Uint32,
	LocaltimeIntoDay:              basetype.Uint32,
	StrokeType:                    basetype.Enum,
	BodyLocation:                  basetype.Enum,
	SegmentLapStatus:              basetype.Enum,
	SegmentLeaderboardType:        basetype.Enum,
	SegmentDeleteStatus:           basetype.Enum,
	SegmentSelectionType:          basetype.Enum,
	SourceType:                    basetype.Enum,
	LocalDeviceType:               basetype.Uint8,
	BleDeviceType:                 basetype.Uint8,
	AntChannelId:                  basetype.Uint32z,
	DisplayOrientation:            basetype.Enum,
	WorkoutEquipment:              basetype.Enum,
	WatchfaceMode:                 basetype.Enum,
	DigitalWatchfaceLayout:        basetype.Enum,
	AnalogWatchfaceLayout:         basetype.Enum,
	RiderPositionType:             basetype.Enum,
	PowerPhaseType:                basetype.Enum,
	CameraEventType:               basetype.Enum,
	SensorType:                    basetype.Enum,
	BikeLightNetworkConfigType:    basetype.Enum,
	CommTimeoutType:               basetype.Uint16,
	CameraOrientationType:         basetype.Enum,
	AttitudeStage:                 basetype.Enum,
	AttitudeValidity:              basetype.Uint16,
	AutoSyncFrequency:             basetype.Enum,
	ExdLayout:                     basetype.Enum,
	ExdDisplayType:                basetype.Enum,
	ExdDataUnits:                  basetype.Enum,
	ExdQualifiers:                 basetype.Enum,
	ExdDescriptors:                basetype.Enum,
	AutoActivityDetect:            basetype.Uint32,
	SupportedExdScreenLayouts:     basetype.Uint32z,
	FitBaseType:                   basetype.Uint8,
	TurnType:                      basetype.Enum,
	BikeLightBeamAngleMode:        basetype.Uint8,
	FitBaseUnit:                   basetype.Uint16,
	SetType:                       basetype.Uint8,
	MaxMetCategory:                basetype.Enum,
	ExerciseCategory:              basetype.Uint16,
	BenchPressExerciseName:        basetype.Uint16,
	CalfRaiseExerciseName:         basetype.Uint16,
	CardioExerciseName:            basetype.Uint16,
	CarryExerciseName:             basetype.Uint16,
	ChopExerciseName:              basetype.Uint16,
	CoreExerciseName:              basetype.Uint16,
	CrunchExerciseName:            basetype.Uint16,
	CurlExerciseName:              basetype.Uint16,
	DeadliftExerciseName:          basetype.Uint16,
	FlyeExerciseName:              basetype.Uint16,
	HipRaiseExerciseName:          basetype.Uint16,
	HipStabilityExerciseName:      basetype.Uint16,
	HipSwingExerciseName:          basetype.Uint16,
	HyperextensionExerciseName:    basetype.Uint16,
	LateralRaiseExerciseName:      basetype.Uint16,
	LegCurlExerciseName:           basetype.Uint16,
	LegRaiseExerciseName:          basetype.Uint16,
	LungeExerciseName:             basetype.Uint16,
	OlympicLiftExerciseName:       basetype.Uint16,
	PlankExerciseName:             basetype.Uint16,
	PlyoExerciseName:              basetype.Uint16,
	PullUpExerciseName:            basetype.Uint16,
	PushUpExerciseName:            basetype.Uint16,
	RowExerciseName:               basetype.Uint16,
	ShoulderPressExerciseName:     basetype.Uint16,
	ShoulderStabilityExerciseName: basetype.Uint16,
	ShrugExerciseName:             basetype.Uint16,
	SitUpExerciseName:             basetype.Uint16,
	SquatExerciseName:             basetype.Uint16,
	TotalBodyExerciseName:         basetype.Uint16,
	TricepsExtensionExerciseName:  basetype.Uint16,
	WarmUpExerciseName:            basetype.Uint16,
	RunExerciseName:               basetype.Uint16,
	WaterType:                     basetype.Enum,
	TissueModelType:               basetype.Enum,
	DiveGasStatus:                 basetype.Enum,
	DiveAlert:                     basetype.Enum,
	DiveAlarmType:                 basetype.Enum,
	DiveBacklightMode:             basetype.Enum,
	SleepLevel:                    basetype.Enum,
	Spo2MeasurementType:           basetype.Enum,
	CcrSetpointSwitchMode:         basetype.Enum,
	DiveGasMode:                   basetype.Enum,
	FaveroProduct:                 basetype.Uint16,
	SplitType:                     basetype.Enum,
	ClimbProEvent:                 basetype.Enum,
	GasConsumptionRateType:        basetype.Enum,
	TapSensitivity:                basetype.Enum,
	RadarThreatLevelType:          basetype.Enum,
	MaxMetSpeedSource:             basetype.Enum,
	MaxMetHeartRateSource:         basetype.Enum,
	HrvStatus:                     basetype.Enum,
	NoFlyTimeMode:                 basetype.Enum,
	Invalid:                       255,
}

func (p *ProfileType) BaseType() basetype.BaseType {
	if *p > ProfileType(len(mappingtobasetypes)) {
		return mappingtobasetypes[Invalid]
	}

	return mappingtobasetypes[*p]
}
