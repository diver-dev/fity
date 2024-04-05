// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"fmt"
	"strconv"
)

type MesgNum uint16

const (
	MesgNumFileId                      MesgNum = 0
	MesgNumCapabilities                MesgNum = 1
	MesgNumDeviceSettings              MesgNum = 2
	MesgNumUserProfile                 MesgNum = 3
	MesgNumHrmProfile                  MesgNum = 4
	MesgNumSdmProfile                  MesgNum = 5
	MesgNumBikeProfile                 MesgNum = 6
	MesgNumZonesTarget                 MesgNum = 7
	MesgNumHrZone                      MesgNum = 8
	MesgNumPowerZone                   MesgNum = 9
	MesgNumMetZone                     MesgNum = 10
	MesgNumSport                       MesgNum = 12
	MesgNumGoal                        MesgNum = 15
	MesgNumSession                     MesgNum = 18
	MesgNumLap                         MesgNum = 19
	MesgNumRecord                      MesgNum = 20
	MesgNumEvent                       MesgNum = 21
	MesgNumDeviceInfo                  MesgNum = 23
	MesgNumWorkout                     MesgNum = 26
	MesgNumWorkoutStep                 MesgNum = 27
	MesgNumSchedule                    MesgNum = 28
	MesgNumWeightScale                 MesgNum = 30
	MesgNumCourse                      MesgNum = 31
	MesgNumCoursePoint                 MesgNum = 32
	MesgNumTotals                      MesgNum = 33
	MesgNumActivity                    MesgNum = 34
	MesgNumSoftware                    MesgNum = 35
	MesgNumFileCapabilities            MesgNum = 37
	MesgNumMesgCapabilities            MesgNum = 38
	MesgNumFieldCapabilities           MesgNum = 39
	MesgNumFileCreator                 MesgNum = 49
	MesgNumBloodPressure               MesgNum = 51
	MesgNumSpeedZone                   MesgNum = 53
	MesgNumMonitoring                  MesgNum = 55
	MesgNumTrainingFile                MesgNum = 72
	MesgNumHrv                         MesgNum = 78
	MesgNumAntRx                       MesgNum = 80
	MesgNumAntTx                       MesgNum = 81
	MesgNumAntChannelId                MesgNum = 82
	MesgNumLength                      MesgNum = 101
	MesgNumMonitoringInfo              MesgNum = 103
	MesgNumPad                         MesgNum = 105
	MesgNumSlaveDevice                 MesgNum = 106
	MesgNumConnectivity                MesgNum = 127
	MesgNumWeatherConditions           MesgNum = 128
	MesgNumWeatherAlert                MesgNum = 129
	MesgNumCadenceZone                 MesgNum = 131
	MesgNumHr                          MesgNum = 132
	MesgNumSegmentLap                  MesgNum = 142
	MesgNumMemoGlob                    MesgNum = 145
	MesgNumSegmentId                   MesgNum = 148
	MesgNumSegmentLeaderboardEntry     MesgNum = 149
	MesgNumSegmentPoint                MesgNum = 150
	MesgNumSegmentFile                 MesgNum = 151
	MesgNumWorkoutSession              MesgNum = 158
	MesgNumWatchfaceSettings           MesgNum = 159
	MesgNumGpsMetadata                 MesgNum = 160
	MesgNumCameraEvent                 MesgNum = 161
	MesgNumTimestampCorrelation        MesgNum = 162
	MesgNumGyroscopeData               MesgNum = 164
	MesgNumAccelerometerData           MesgNum = 165
	MesgNumThreeDSensorCalibration     MesgNum = 167
	MesgNumVideoFrame                  MesgNum = 169
	MesgNumObdiiData                   MesgNum = 174
	MesgNumNmeaSentence                MesgNum = 177
	MesgNumAviationAttitude            MesgNum = 178
	MesgNumVideo                       MesgNum = 184
	MesgNumVideoTitle                  MesgNum = 185
	MesgNumVideoDescription            MesgNum = 186
	MesgNumVideoClip                   MesgNum = 187
	MesgNumOhrSettings                 MesgNum = 188
	MesgNumExdScreenConfiguration      MesgNum = 200
	MesgNumExdDataFieldConfiguration   MesgNum = 201
	MesgNumExdDataConceptConfiguration MesgNum = 202
	MesgNumFieldDescription            MesgNum = 206
	MesgNumDeveloperDataId             MesgNum = 207
	MesgNumMagnetometerData            MesgNum = 208
	MesgNumBarometerData               MesgNum = 209
	MesgNumOneDSensorCalibration       MesgNum = 210
	MesgNumMonitoringHrData            MesgNum = 211
	MesgNumTimeInZone                  MesgNum = 216
	MesgNumSet                         MesgNum = 225
	MesgNumStressLevel                 MesgNum = 227
	MesgNumMaxMetData                  MesgNum = 229
	MesgNumDiveSettings                MesgNum = 258
	MesgNumDiveGas                     MesgNum = 259
	MesgNumDiveAlarm                   MesgNum = 262
	MesgNumExerciseTitle               MesgNum = 264
	MesgNumDiveSummary                 MesgNum = 268
	MesgNumSpo2Data                    MesgNum = 269
	MesgNumSleepLevel                  MesgNum = 275
	MesgNumJump                        MesgNum = 285
	MesgNumAadAccelFeatures            MesgNum = 289
	MesgNumBeatIntervals               MesgNum = 290
	MesgNumRespirationRate             MesgNum = 297
	MesgNumHsaAccelerometerData        MesgNum = 302
	MesgNumHsaStepData                 MesgNum = 304
	MesgNumHsaSpo2Data                 MesgNum = 305
	MesgNumHsaStressData               MesgNum = 306
	MesgNumHsaRespirationData          MesgNum = 307
	MesgNumHsaHeartRateData            MesgNum = 308
	MesgNumSplit                       MesgNum = 312
	MesgNumSplitSummary                MesgNum = 313
	MesgNumHsaBodyBatteryData          MesgNum = 314
	MesgNumHsaEvent                    MesgNum = 315
	MesgNumClimbPro                    MesgNum = 317
	MesgNumTankUpdate                  MesgNum = 319
	MesgNumTankSummary                 MesgNum = 323
	MesgNumSleepAssessment             MesgNum = 346
	MesgNumHrvStatusSummary            MesgNum = 370
	MesgNumHrvValue                    MesgNum = 371
	MesgNumRawBbi                      MesgNum = 372
	MesgNumDeviceAuxBatteryInfo        MesgNum = 375
	MesgNumHsaGyroscopeData            MesgNum = 376
	MesgNumChronoShotSession           MesgNum = 387
	MesgNumChronoShotData              MesgNum = 388
	MesgNumHsaConfigurationData        MesgNum = 389
	MesgNumDiveApneaAlarm              MesgNum = 393
	MesgNumHsaWristTemperatureData     MesgNum = 409    // Message number for the HSA wrist temperature data message
	MesgNumMfgRangeMin                 MesgNum = 0xFF00 // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumMfgRangeMax                 MesgNum = 0xFFFE // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumInvalid                     MesgNum = 0xFFFF
)

func (m MesgNum) Uint16() uint16 { return uint16(m) }

var mesgnumToString = map[MesgNum]string{}
var stringToMesgNum = map[string]MesgNum{}

func (m MesgNum) String() string {
	switch m {
	case MesgNumFileId:
		return "file_id"
	case MesgNumCapabilities:
		return "capabilities"
	case MesgNumDeviceSettings:
		return "device_settings"
	case MesgNumUserProfile:
		return "user_profile"
	case MesgNumHrmProfile:
		return "hrm_profile"
	case MesgNumSdmProfile:
		return "sdm_profile"
	case MesgNumBikeProfile:
		return "bike_profile"
	case MesgNumZonesTarget:
		return "zones_target"
	case MesgNumHrZone:
		return "hr_zone"
	case MesgNumPowerZone:
		return "power_zone"
	case MesgNumMetZone:
		return "met_zone"
	case MesgNumSport:
		return "sport"
	case MesgNumGoal:
		return "goal"
	case MesgNumSession:
		return "session"
	case MesgNumLap:
		return "lap"
	case MesgNumRecord:
		return "record"
	case MesgNumEvent:
		return "event"
	case MesgNumDeviceInfo:
		return "device_info"
	case MesgNumWorkout:
		return "workout"
	case MesgNumWorkoutStep:
		return "workout_step"
	case MesgNumSchedule:
		return "schedule"
	case MesgNumWeightScale:
		return "weight_scale"
	case MesgNumCourse:
		return "course"
	case MesgNumCoursePoint:
		return "course_point"
	case MesgNumTotals:
		return "totals"
	case MesgNumActivity:
		return "activity"
	case MesgNumSoftware:
		return "software"
	case MesgNumFileCapabilities:
		return "file_capabilities"
	case MesgNumMesgCapabilities:
		return "mesg_capabilities"
	case MesgNumFieldCapabilities:
		return "field_capabilities"
	case MesgNumFileCreator:
		return "file_creator"
	case MesgNumBloodPressure:
		return "blood_pressure"
	case MesgNumSpeedZone:
		return "speed_zone"
	case MesgNumMonitoring:
		return "monitoring"
	case MesgNumTrainingFile:
		return "training_file"
	case MesgNumHrv:
		return "hrv"
	case MesgNumAntRx:
		return "ant_rx"
	case MesgNumAntTx:
		return "ant_tx"
	case MesgNumAntChannelId:
		return "ant_channel_id"
	case MesgNumLength:
		return "length"
	case MesgNumMonitoringInfo:
		return "monitoring_info"
	case MesgNumPad:
		return "pad"
	case MesgNumSlaveDevice:
		return "slave_device"
	case MesgNumConnectivity:
		return "connectivity"
	case MesgNumWeatherConditions:
		return "weather_conditions"
	case MesgNumWeatherAlert:
		return "weather_alert"
	case MesgNumCadenceZone:
		return "cadence_zone"
	case MesgNumHr:
		return "hr"
	case MesgNumSegmentLap:
		return "segment_lap"
	case MesgNumMemoGlob:
		return "memo_glob"
	case MesgNumSegmentId:
		return "segment_id"
	case MesgNumSegmentLeaderboardEntry:
		return "segment_leaderboard_entry"
	case MesgNumSegmentPoint:
		return "segment_point"
	case MesgNumSegmentFile:
		return "segment_file"
	case MesgNumWorkoutSession:
		return "workout_session"
	case MesgNumWatchfaceSettings:
		return "watchface_settings"
	case MesgNumGpsMetadata:
		return "gps_metadata"
	case MesgNumCameraEvent:
		return "camera_event"
	case MesgNumTimestampCorrelation:
		return "timestamp_correlation"
	case MesgNumGyroscopeData:
		return "gyroscope_data"
	case MesgNumAccelerometerData:
		return "accelerometer_data"
	case MesgNumThreeDSensorCalibration:
		return "three_d_sensor_calibration"
	case MesgNumVideoFrame:
		return "video_frame"
	case MesgNumObdiiData:
		return "obdii_data"
	case MesgNumNmeaSentence:
		return "nmea_sentence"
	case MesgNumAviationAttitude:
		return "aviation_attitude"
	case MesgNumVideo:
		return "video"
	case MesgNumVideoTitle:
		return "video_title"
	case MesgNumVideoDescription:
		return "video_description"
	case MesgNumVideoClip:
		return "video_clip"
	case MesgNumOhrSettings:
		return "ohr_settings"
	case MesgNumExdScreenConfiguration:
		return "exd_screen_configuration"
	case MesgNumExdDataFieldConfiguration:
		return "exd_data_field_configuration"
	case MesgNumExdDataConceptConfiguration:
		return "exd_data_concept_configuration"
	case MesgNumFieldDescription:
		return "field_description"
	case MesgNumDeveloperDataId:
		return "developer_data_id"
	case MesgNumMagnetometerData:
		return "magnetometer_data"
	case MesgNumBarometerData:
		return "barometer_data"
	case MesgNumOneDSensorCalibration:
		return "one_d_sensor_calibration"
	case MesgNumMonitoringHrData:
		return "monitoring_hr_data"
	case MesgNumTimeInZone:
		return "time_in_zone"
	case MesgNumSet:
		return "set"
	case MesgNumStressLevel:
		return "stress_level"
	case MesgNumMaxMetData:
		return "max_met_data"
	case MesgNumDiveSettings:
		return "dive_settings"
	case MesgNumDiveGas:
		return "dive_gas"
	case MesgNumDiveAlarm:
		return "dive_alarm"
	case MesgNumExerciseTitle:
		return "exercise_title"
	case MesgNumDiveSummary:
		return "dive_summary"
	case MesgNumSpo2Data:
		return "spo2_data"
	case MesgNumSleepLevel:
		return "sleep_level"
	case MesgNumJump:
		return "jump"
	case MesgNumAadAccelFeatures:
		return "aad_accel_features"
	case MesgNumBeatIntervals:
		return "beat_intervals"
	case MesgNumRespirationRate:
		return "respiration_rate"
	case MesgNumHsaAccelerometerData:
		return "hsa_accelerometer_data"
	case MesgNumHsaStepData:
		return "hsa_step_data"
	case MesgNumHsaSpo2Data:
		return "hsa_spo2_data"
	case MesgNumHsaStressData:
		return "hsa_stress_data"
	case MesgNumHsaRespirationData:
		return "hsa_respiration_data"
	case MesgNumHsaHeartRateData:
		return "hsa_heart_rate_data"
	case MesgNumSplit:
		return "split"
	case MesgNumSplitSummary:
		return "split_summary"
	case MesgNumHsaBodyBatteryData:
		return "hsa_body_battery_data"
	case MesgNumHsaEvent:
		return "hsa_event"
	case MesgNumClimbPro:
		return "climb_pro"
	case MesgNumTankUpdate:
		return "tank_update"
	case MesgNumTankSummary:
		return "tank_summary"
	case MesgNumSleepAssessment:
		return "sleep_assessment"
	case MesgNumHrvStatusSummary:
		return "hrv_status_summary"
	case MesgNumHrvValue:
		return "hrv_value"
	case MesgNumRawBbi:
		return "raw_bbi"
	case MesgNumDeviceAuxBatteryInfo:
		return "device_aux_battery_info"
	case MesgNumHsaGyroscopeData:
		return "hsa_gyroscope_data"
	case MesgNumChronoShotSession:
		return "chrono_shot_session"
	case MesgNumChronoShotData:
		return "chrono_shot_data"
	case MesgNumHsaConfigurationData:
		return "hsa_configuration_data"
	case MesgNumDiveApneaAlarm:
		return "dive_apnea_alarm"
	case MesgNumHsaWristTemperatureData:
		return "hsa_wrist_temperature_data"
	case MesgNumMfgRangeMin:
		return "mfg_range_min"
	case MesgNumMfgRangeMax:
		return "mfg_range_max"
	default:
		if val, ok := mesgnumToString[m]; ok {
			return val
		}
		return "MesgNumInvalid(" + strconv.FormatUint(uint64(m), 10) + ")"
	}
}

// FromString parse string into MesgNum constant it's represent, return MesgNumInvalid if not found.
func MesgNumFromString(s string) MesgNum {
	switch s {
	case "file_id":
		return MesgNumFileId
	case "capabilities":
		return MesgNumCapabilities
	case "device_settings":
		return MesgNumDeviceSettings
	case "user_profile":
		return MesgNumUserProfile
	case "hrm_profile":
		return MesgNumHrmProfile
	case "sdm_profile":
		return MesgNumSdmProfile
	case "bike_profile":
		return MesgNumBikeProfile
	case "zones_target":
		return MesgNumZonesTarget
	case "hr_zone":
		return MesgNumHrZone
	case "power_zone":
		return MesgNumPowerZone
	case "met_zone":
		return MesgNumMetZone
	case "sport":
		return MesgNumSport
	case "goal":
		return MesgNumGoal
	case "session":
		return MesgNumSession
	case "lap":
		return MesgNumLap
	case "record":
		return MesgNumRecord
	case "event":
		return MesgNumEvent
	case "device_info":
		return MesgNumDeviceInfo
	case "workout":
		return MesgNumWorkout
	case "workout_step":
		return MesgNumWorkoutStep
	case "schedule":
		return MesgNumSchedule
	case "weight_scale":
		return MesgNumWeightScale
	case "course":
		return MesgNumCourse
	case "course_point":
		return MesgNumCoursePoint
	case "totals":
		return MesgNumTotals
	case "activity":
		return MesgNumActivity
	case "software":
		return MesgNumSoftware
	case "file_capabilities":
		return MesgNumFileCapabilities
	case "mesg_capabilities":
		return MesgNumMesgCapabilities
	case "field_capabilities":
		return MesgNumFieldCapabilities
	case "file_creator":
		return MesgNumFileCreator
	case "blood_pressure":
		return MesgNumBloodPressure
	case "speed_zone":
		return MesgNumSpeedZone
	case "monitoring":
		return MesgNumMonitoring
	case "training_file":
		return MesgNumTrainingFile
	case "hrv":
		return MesgNumHrv
	case "ant_rx":
		return MesgNumAntRx
	case "ant_tx":
		return MesgNumAntTx
	case "ant_channel_id":
		return MesgNumAntChannelId
	case "length":
		return MesgNumLength
	case "monitoring_info":
		return MesgNumMonitoringInfo
	case "pad":
		return MesgNumPad
	case "slave_device":
		return MesgNumSlaveDevice
	case "connectivity":
		return MesgNumConnectivity
	case "weather_conditions":
		return MesgNumWeatherConditions
	case "weather_alert":
		return MesgNumWeatherAlert
	case "cadence_zone":
		return MesgNumCadenceZone
	case "hr":
		return MesgNumHr
	case "segment_lap":
		return MesgNumSegmentLap
	case "memo_glob":
		return MesgNumMemoGlob
	case "segment_id":
		return MesgNumSegmentId
	case "segment_leaderboard_entry":
		return MesgNumSegmentLeaderboardEntry
	case "segment_point":
		return MesgNumSegmentPoint
	case "segment_file":
		return MesgNumSegmentFile
	case "workout_session":
		return MesgNumWorkoutSession
	case "watchface_settings":
		return MesgNumWatchfaceSettings
	case "gps_metadata":
		return MesgNumGpsMetadata
	case "camera_event":
		return MesgNumCameraEvent
	case "timestamp_correlation":
		return MesgNumTimestampCorrelation
	case "gyroscope_data":
		return MesgNumGyroscopeData
	case "accelerometer_data":
		return MesgNumAccelerometerData
	case "three_d_sensor_calibration":
		return MesgNumThreeDSensorCalibration
	case "video_frame":
		return MesgNumVideoFrame
	case "obdii_data":
		return MesgNumObdiiData
	case "nmea_sentence":
		return MesgNumNmeaSentence
	case "aviation_attitude":
		return MesgNumAviationAttitude
	case "video":
		return MesgNumVideo
	case "video_title":
		return MesgNumVideoTitle
	case "video_description":
		return MesgNumVideoDescription
	case "video_clip":
		return MesgNumVideoClip
	case "ohr_settings":
		return MesgNumOhrSettings
	case "exd_screen_configuration":
		return MesgNumExdScreenConfiguration
	case "exd_data_field_configuration":
		return MesgNumExdDataFieldConfiguration
	case "exd_data_concept_configuration":
		return MesgNumExdDataConceptConfiguration
	case "field_description":
		return MesgNumFieldDescription
	case "developer_data_id":
		return MesgNumDeveloperDataId
	case "magnetometer_data":
		return MesgNumMagnetometerData
	case "barometer_data":
		return MesgNumBarometerData
	case "one_d_sensor_calibration":
		return MesgNumOneDSensorCalibration
	case "monitoring_hr_data":
		return MesgNumMonitoringHrData
	case "time_in_zone":
		return MesgNumTimeInZone
	case "set":
		return MesgNumSet
	case "stress_level":
		return MesgNumStressLevel
	case "max_met_data":
		return MesgNumMaxMetData
	case "dive_settings":
		return MesgNumDiveSettings
	case "dive_gas":
		return MesgNumDiveGas
	case "dive_alarm":
		return MesgNumDiveAlarm
	case "exercise_title":
		return MesgNumExerciseTitle
	case "dive_summary":
		return MesgNumDiveSummary
	case "spo2_data":
		return MesgNumSpo2Data
	case "sleep_level":
		return MesgNumSleepLevel
	case "jump":
		return MesgNumJump
	case "aad_accel_features":
		return MesgNumAadAccelFeatures
	case "beat_intervals":
		return MesgNumBeatIntervals
	case "respiration_rate":
		return MesgNumRespirationRate
	case "hsa_accelerometer_data":
		return MesgNumHsaAccelerometerData
	case "hsa_step_data":
		return MesgNumHsaStepData
	case "hsa_spo2_data":
		return MesgNumHsaSpo2Data
	case "hsa_stress_data":
		return MesgNumHsaStressData
	case "hsa_respiration_data":
		return MesgNumHsaRespirationData
	case "hsa_heart_rate_data":
		return MesgNumHsaHeartRateData
	case "split":
		return MesgNumSplit
	case "split_summary":
		return MesgNumSplitSummary
	case "hsa_body_battery_data":
		return MesgNumHsaBodyBatteryData
	case "hsa_event":
		return MesgNumHsaEvent
	case "climb_pro":
		return MesgNumClimbPro
	case "tank_update":
		return MesgNumTankUpdate
	case "tank_summary":
		return MesgNumTankSummary
	case "sleep_assessment":
		return MesgNumSleepAssessment
	case "hrv_status_summary":
		return MesgNumHrvStatusSummary
	case "hrv_value":
		return MesgNumHrvValue
	case "raw_bbi":
		return MesgNumRawBbi
	case "device_aux_battery_info":
		return MesgNumDeviceAuxBatteryInfo
	case "hsa_gyroscope_data":
		return MesgNumHsaGyroscopeData
	case "chrono_shot_session":
		return MesgNumChronoShotSession
	case "chrono_shot_data":
		return MesgNumChronoShotData
	case "hsa_configuration_data":
		return MesgNumHsaConfigurationData
	case "dive_apnea_alarm":
		return MesgNumDiveApneaAlarm
	case "hsa_wrist_temperature_data":
		return MesgNumHsaWristTemperatureData
	case "mfg_range_min":
		return MesgNumMfgRangeMin
	case "mfg_range_max":
		return MesgNumMfgRangeMax
	default:
		if val, ok := stringToMesgNum[s]; ok {
			return val
		}
		return MesgNumInvalid
	}
}

// List returns all constants.
func ListMesgNum() []MesgNum {
	list := []MesgNum{
		MesgNumFileId,
		MesgNumCapabilities,
		MesgNumDeviceSettings,
		MesgNumUserProfile,
		MesgNumHrmProfile,
		MesgNumSdmProfile,
		MesgNumBikeProfile,
		MesgNumZonesTarget,
		MesgNumHrZone,
		MesgNumPowerZone,
		MesgNumMetZone,
		MesgNumSport,
		MesgNumGoal,
		MesgNumSession,
		MesgNumLap,
		MesgNumRecord,
		MesgNumEvent,
		MesgNumDeviceInfo,
		MesgNumWorkout,
		MesgNumWorkoutStep,
		MesgNumSchedule,
		MesgNumWeightScale,
		MesgNumCourse,
		MesgNumCoursePoint,
		MesgNumTotals,
		MesgNumActivity,
		MesgNumSoftware,
		MesgNumFileCapabilities,
		MesgNumMesgCapabilities,
		MesgNumFieldCapabilities,
		MesgNumFileCreator,
		MesgNumBloodPressure,
		MesgNumSpeedZone,
		MesgNumMonitoring,
		MesgNumTrainingFile,
		MesgNumHrv,
		MesgNumAntRx,
		MesgNumAntTx,
		MesgNumAntChannelId,
		MesgNumLength,
		MesgNumMonitoringInfo,
		MesgNumPad,
		MesgNumSlaveDevice,
		MesgNumConnectivity,
		MesgNumWeatherConditions,
		MesgNumWeatherAlert,
		MesgNumCadenceZone,
		MesgNumHr,
		MesgNumSegmentLap,
		MesgNumMemoGlob,
		MesgNumSegmentId,
		MesgNumSegmentLeaderboardEntry,
		MesgNumSegmentPoint,
		MesgNumSegmentFile,
		MesgNumWorkoutSession,
		MesgNumWatchfaceSettings,
		MesgNumGpsMetadata,
		MesgNumCameraEvent,
		MesgNumTimestampCorrelation,
		MesgNumGyroscopeData,
		MesgNumAccelerometerData,
		MesgNumThreeDSensorCalibration,
		MesgNumVideoFrame,
		MesgNumObdiiData,
		MesgNumNmeaSentence,
		MesgNumAviationAttitude,
		MesgNumVideo,
		MesgNumVideoTitle,
		MesgNumVideoDescription,
		MesgNumVideoClip,
		MesgNumOhrSettings,
		MesgNumExdScreenConfiguration,
		MesgNumExdDataFieldConfiguration,
		MesgNumExdDataConceptConfiguration,
		MesgNumFieldDescription,
		MesgNumDeveloperDataId,
		MesgNumMagnetometerData,
		MesgNumBarometerData,
		MesgNumOneDSensorCalibration,
		MesgNumMonitoringHrData,
		MesgNumTimeInZone,
		MesgNumSet,
		MesgNumStressLevel,
		MesgNumMaxMetData,
		MesgNumDiveSettings,
		MesgNumDiveGas,
		MesgNumDiveAlarm,
		MesgNumExerciseTitle,
		MesgNumDiveSummary,
		MesgNumSpo2Data,
		MesgNumSleepLevel,
		MesgNumJump,
		MesgNumAadAccelFeatures,
		MesgNumBeatIntervals,
		MesgNumRespirationRate,
		MesgNumHsaAccelerometerData,
		MesgNumHsaStepData,
		MesgNumHsaSpo2Data,
		MesgNumHsaStressData,
		MesgNumHsaRespirationData,
		MesgNumHsaHeartRateData,
		MesgNumSplit,
		MesgNumSplitSummary,
		MesgNumHsaBodyBatteryData,
		MesgNumHsaEvent,
		MesgNumClimbPro,
		MesgNumTankUpdate,
		MesgNumTankSummary,
		MesgNumSleepAssessment,
		MesgNumHrvStatusSummary,
		MesgNumHrvValue,
		MesgNumRawBbi,
		MesgNumDeviceAuxBatteryInfo,
		MesgNumHsaGyroscopeData,
		MesgNumChronoShotSession,
		MesgNumChronoShotData,
		MesgNumHsaConfigurationData,
		MesgNumDiveApneaAlarm,
		MesgNumHsaWristTemperatureData,
		MesgNumMfgRangeMin,
		MesgNumMfgRangeMax,
	}
	for k := range mesgnumToString {
		list = append(list, k)
	}
	return list
}

// MesgNumRegister registers a manufacturer specific MesgNum so that the value can be recognized.
// It is recommended to define the constants somewhere else to track your own specifications.
//
// This is intended for those who prefer using this SDK as it is without the need to generate custom SDK using cmd/fitgen.
func MesgNumRegister(v MesgNum, s string) error {
	if v >= MesgNumInvalid {
		return fmt.Errorf("could not register outside max range: %d", MesgNumInvalid)
	}

	switch v {
	case MesgNumFileId:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumFileId", v)
	case MesgNumCapabilities:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumCapabilities", v)
	case MesgNumDeviceSettings:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDeviceSettings", v)
	case MesgNumUserProfile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumUserProfile", v)
	case MesgNumHrmProfile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHrmProfile", v)
	case MesgNumSdmProfile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSdmProfile", v)
	case MesgNumBikeProfile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumBikeProfile", v)
	case MesgNumZonesTarget:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumZonesTarget", v)
	case MesgNumHrZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHrZone", v)
	case MesgNumPowerZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumPowerZone", v)
	case MesgNumMetZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMetZone", v)
	case MesgNumSport:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSport", v)
	case MesgNumGoal:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumGoal", v)
	case MesgNumSession:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSession", v)
	case MesgNumLap:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumLap", v)
	case MesgNumRecord:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumRecord", v)
	case MesgNumEvent:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumEvent", v)
	case MesgNumDeviceInfo:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDeviceInfo", v)
	case MesgNumWorkout:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWorkout", v)
	case MesgNumWorkoutStep:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWorkoutStep", v)
	case MesgNumSchedule:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSchedule", v)
	case MesgNumWeightScale:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWeightScale", v)
	case MesgNumCourse:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumCourse", v)
	case MesgNumCoursePoint:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumCoursePoint", v)
	case MesgNumTotals:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTotals", v)
	case MesgNumActivity:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumActivity", v)
	case MesgNumSoftware:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSoftware", v)
	case MesgNumFileCapabilities:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumFileCapabilities", v)
	case MesgNumMesgCapabilities:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMesgCapabilities", v)
	case MesgNumFieldCapabilities:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumFieldCapabilities", v)
	case MesgNumFileCreator:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumFileCreator", v)
	case MesgNumBloodPressure:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumBloodPressure", v)
	case MesgNumSpeedZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSpeedZone", v)
	case MesgNumMonitoring:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMonitoring", v)
	case MesgNumTrainingFile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTrainingFile", v)
	case MesgNumHrv:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHrv", v)
	case MesgNumAntRx:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAntRx", v)
	case MesgNumAntTx:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAntTx", v)
	case MesgNumAntChannelId:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAntChannelId", v)
	case MesgNumLength:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumLength", v)
	case MesgNumMonitoringInfo:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMonitoringInfo", v)
	case MesgNumPad:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumPad", v)
	case MesgNumSlaveDevice:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSlaveDevice", v)
	case MesgNumConnectivity:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumConnectivity", v)
	case MesgNumWeatherConditions:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWeatherConditions", v)
	case MesgNumWeatherAlert:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWeatherAlert", v)
	case MesgNumCadenceZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumCadenceZone", v)
	case MesgNumHr:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHr", v)
	case MesgNumSegmentLap:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSegmentLap", v)
	case MesgNumMemoGlob:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMemoGlob", v)
	case MesgNumSegmentId:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSegmentId", v)
	case MesgNumSegmentLeaderboardEntry:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSegmentLeaderboardEntry", v)
	case MesgNumSegmentPoint:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSegmentPoint", v)
	case MesgNumSegmentFile:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSegmentFile", v)
	case MesgNumWorkoutSession:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWorkoutSession", v)
	case MesgNumWatchfaceSettings:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumWatchfaceSettings", v)
	case MesgNumGpsMetadata:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumGpsMetadata", v)
	case MesgNumCameraEvent:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumCameraEvent", v)
	case MesgNumTimestampCorrelation:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTimestampCorrelation", v)
	case MesgNumGyroscopeData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumGyroscopeData", v)
	case MesgNumAccelerometerData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAccelerometerData", v)
	case MesgNumThreeDSensorCalibration:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumThreeDSensorCalibration", v)
	case MesgNumVideoFrame:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumVideoFrame", v)
	case MesgNumObdiiData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumObdiiData", v)
	case MesgNumNmeaSentence:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumNmeaSentence", v)
	case MesgNumAviationAttitude:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAviationAttitude", v)
	case MesgNumVideo:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumVideo", v)
	case MesgNumVideoTitle:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumVideoTitle", v)
	case MesgNumVideoDescription:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumVideoDescription", v)
	case MesgNumVideoClip:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumVideoClip", v)
	case MesgNumOhrSettings:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumOhrSettings", v)
	case MesgNumExdScreenConfiguration:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumExdScreenConfiguration", v)
	case MesgNumExdDataFieldConfiguration:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumExdDataFieldConfiguration", v)
	case MesgNumExdDataConceptConfiguration:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumExdDataConceptConfiguration", v)
	case MesgNumFieldDescription:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumFieldDescription", v)
	case MesgNumDeveloperDataId:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDeveloperDataId", v)
	case MesgNumMagnetometerData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMagnetometerData", v)
	case MesgNumBarometerData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumBarometerData", v)
	case MesgNumOneDSensorCalibration:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumOneDSensorCalibration", v)
	case MesgNumMonitoringHrData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMonitoringHrData", v)
	case MesgNumTimeInZone:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTimeInZone", v)
	case MesgNumSet:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSet", v)
	case MesgNumStressLevel:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumStressLevel", v)
	case MesgNumMaxMetData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMaxMetData", v)
	case MesgNumDiveSettings:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDiveSettings", v)
	case MesgNumDiveGas:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDiveGas", v)
	case MesgNumDiveAlarm:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDiveAlarm", v)
	case MesgNumExerciseTitle:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumExerciseTitle", v)
	case MesgNumDiveSummary:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDiveSummary", v)
	case MesgNumSpo2Data:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSpo2Data", v)
	case MesgNumSleepLevel:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSleepLevel", v)
	case MesgNumJump:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumJump", v)
	case MesgNumAadAccelFeatures:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumAadAccelFeatures", v)
	case MesgNumBeatIntervals:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumBeatIntervals", v)
	case MesgNumRespirationRate:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumRespirationRate", v)
	case MesgNumHsaAccelerometerData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaAccelerometerData", v)
	case MesgNumHsaStepData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaStepData", v)
	case MesgNumHsaSpo2Data:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaSpo2Data", v)
	case MesgNumHsaStressData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaStressData", v)
	case MesgNumHsaRespirationData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaRespirationData", v)
	case MesgNumHsaHeartRateData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaHeartRateData", v)
	case MesgNumSplit:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSplit", v)
	case MesgNumSplitSummary:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSplitSummary", v)
	case MesgNumHsaBodyBatteryData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaBodyBatteryData", v)
	case MesgNumHsaEvent:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaEvent", v)
	case MesgNumClimbPro:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumClimbPro", v)
	case MesgNumTankUpdate:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTankUpdate", v)
	case MesgNumTankSummary:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumTankSummary", v)
	case MesgNumSleepAssessment:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumSleepAssessment", v)
	case MesgNumHrvStatusSummary:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHrvStatusSummary", v)
	case MesgNumHrvValue:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHrvValue", v)
	case MesgNumRawBbi:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumRawBbi", v)
	case MesgNumDeviceAuxBatteryInfo:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDeviceAuxBatteryInfo", v)
	case MesgNumHsaGyroscopeData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaGyroscopeData", v)
	case MesgNumChronoShotSession:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumChronoShotSession", v)
	case MesgNumChronoShotData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumChronoShotData", v)
	case MesgNumHsaConfigurationData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaConfigurationData", v)
	case MesgNumDiveApneaAlarm:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumDiveApneaAlarm", v)
	case MesgNumHsaWristTemperatureData:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumHsaWristTemperatureData", v)
	case MesgNumMfgRangeMin:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMfgRangeMin", v)
	case MesgNumMfgRangeMax:
		return fmt.Errorf("duplicate: %d is already exist for MesgNumMfgRangeMax", v)
	}

	mesgnumToString[v] = s
	stringToMesgNum[s] = v

	return nil
}
