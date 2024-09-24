// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// Monitoring is a Monitoring message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type Monitoring struct {
	Timestamp                    time.Time           // Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
	LocalTimestamp               time.Time           // Must align to logging interval, for example, time must be 00:00:00 for daily log.
	ActivityTime                 [8]uint16           // Array: [8]; Units: minutes; Indexed using minute_activity_level enum
	Distance                     uint32              // Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	Cycles                       uint32              // Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
	ActiveTime                   uint32              // Scale: 1000; Units: s
	Duration                     uint32              // Units: s
	Ascent                       uint32              // Scale: 1000; Units: m
	Descent                      uint32              // Scale: 1000; Units: m
	Calories                     uint16              // Units: kcal; Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation
	Distance16                   uint16              // Units: 100 * m
	Cycles16                     uint16              // Units: 2 * cycles (steps)
	ActiveTime16                 uint16              // Units: s
	Temperature                  int16               // Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
	TemperatureMin               int16               // Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
	TemperatureMax               int16               // Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
	ActiveCalories               uint16              // Units: kcal
	Timestamp16                  uint16              // Units: s
	DurationMin                  uint16              // Units: min
	ModerateActivityMinutes      uint16              // Units: minutes
	VigorousActivityMinutes      uint16              // Units: minutes
	DeviceIndex                  typedef.DeviceIndex // Associates this data to device_info message. Not required for file with single device (sensor).
	ActivityType                 typedef.ActivityType
	ActivitySubtype              typedef.ActivitySubtype
	ActivityLevel                typedef.ActivityLevel
	CurrentActivityTypeIntensity byte  // Indicates single type / intensity for duration since last monitoring message.
	TimestampMin8                uint8 // Units: min
	HeartRate                    uint8 // Units: bpm
	Intensity                    uint8 // Scale: 10

	state [4]uint8 // Used for tracking expanded fields.

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewMonitoring creates new Monitoring struct based on given mesg.
// If mesg is nil, it will return Monitoring with all fields being set to its corresponding invalid value.
func NewMonitoring(mesg *proto.Message) *Monitoring {
	vals := [254]proto.Value{}

	var state [4]uint8
	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			if mesg.Fields[i].Num < 29 && mesg.Fields[i].IsExpandedField {
				pos := mesg.Fields[i].Num / 8
				state[pos] |= 1 << (mesg.Fields[i].Num - (8 * pos))
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &Monitoring{
		Timestamp:       datetime.ToTime(vals[253].Uint32()),
		DeviceIndex:     typedef.DeviceIndex(vals[0].Uint8()),
		Calories:        vals[1].Uint16(),
		Distance:        vals[2].Uint32(),
		Cycles:          vals[3].Uint32(),
		ActiveTime:      vals[4].Uint32(),
		ActivityType:    typedef.ActivityType(vals[5].Uint8()),
		ActivitySubtype: typedef.ActivitySubtype(vals[6].Uint8()),
		ActivityLevel:   typedef.ActivityLevel(vals[7].Uint8()),
		Distance16:      vals[8].Uint16(),
		Cycles16:        vals[9].Uint16(),
		ActiveTime16:    vals[10].Uint16(),
		LocalTimestamp:  datetime.ToTime(vals[11].Uint32()),
		Temperature:     vals[12].Int16(),
		TemperatureMin:  vals[14].Int16(),
		TemperatureMax:  vals[15].Int16(),
		ActivityTime: func() (arr [8]uint16) {
			arr = [8]uint16{
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
				basetype.Uint16Invalid,
			}
			copy(arr[:], vals[16].SliceUint16())
			return arr
		}(),
		ActiveCalories:               vals[19].Uint16(),
		CurrentActivityTypeIntensity: vals[24].Uint8(),
		TimestampMin8:                vals[25].Uint8(),
		Timestamp16:                  vals[26].Uint16(),
		HeartRate:                    vals[27].Uint8(),
		Intensity:                    vals[28].Uint8(),
		DurationMin:                  vals[29].Uint16(),
		Duration:                     vals[30].Uint32(),
		Ascent:                       vals[31].Uint32(),
		Descent:                      vals[32].Uint32(),
		ModerateActivityMinutes:      vals[33].Uint16(),
		VigorousActivityMinutes:      vals[34].Uint16(),

		state: state,

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts Monitoring into proto.Message. If options is nil, default options will be used.
func (m *Monitoring) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumMonitoring}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.DeviceIndex != typedef.DeviceIndexInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(uint8(m.DeviceIndex))
		fields = append(fields, field)
	}
	if m.Calories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(m.Calories)
		fields = append(fields, field)
	}
	if m.Distance != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.Distance)
		fields = append(fields, field)
	}
	if m.Cycles != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.Cycles)
		fields = append(fields, field)
	}
	if m.ActiveTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.ActiveTime)
		fields = append(fields, field)
	}
	if m.ActivityType != typedef.ActivityTypeInvalid {
		if expanded := m.IsExpandedField(5); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 5)
			field.Value = proto.Uint8(byte(m.ActivityType))
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.ActivitySubtype != typedef.ActivitySubtypeInvalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(byte(m.ActivitySubtype))
		fields = append(fields, field)
	}
	if m.ActivityLevel != typedef.ActivityLevelInvalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(byte(m.ActivityLevel))
		fields = append(fields, field)
	}
	if m.Distance16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint16(m.Distance16)
		fields = append(fields, field)
	}
	if m.Cycles16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.Cycles16)
		fields = append(fields, field)
	}
	if m.ActiveTime16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint16(m.ActiveTime16)
		fields = append(fields, field)
	}
	if !m.LocalTimestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint32(uint32(m.LocalTimestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.Temperature != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Int16(m.Temperature)
		fields = append(fields, field)
	}
	if m.TemperatureMin != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Int16(m.TemperatureMin)
		fields = append(fields, field)
	}
	if m.TemperatureMax != basetype.Sint16Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Int16(m.TemperatureMax)
		fields = append(fields, field)
	}
	if m.ActivityTime != [8]uint16{
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
		basetype.Uint16Invalid,
	} {
		field := fac.CreateField(mesg.Num, 16)
		copied := m.ActivityTime
		field.Value = proto.SliceUint16(copied[:])
		fields = append(fields, field)
	}
	if m.ActiveCalories != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 19)
		field.Value = proto.Uint16(m.ActiveCalories)
		fields = append(fields, field)
	}
	if m.CurrentActivityTypeIntensity != basetype.ByteInvalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Uint8(m.CurrentActivityTypeIntensity)
		fields = append(fields, field)
	}
	if m.TimestampMin8 != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint8(m.TimestampMin8)
		fields = append(fields, field)
	}
	if m.Timestamp16 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 26)
		field.Value = proto.Uint16(m.Timestamp16)
		fields = append(fields, field)
	}
	if m.HeartRate != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 27)
		field.Value = proto.Uint8(m.HeartRate)
		fields = append(fields, field)
	}
	if m.Intensity != basetype.Uint8Invalid {
		if expanded := m.IsExpandedField(28); !expanded || (expanded && options.IncludeExpandedFields) {
			field := fac.CreateField(mesg.Num, 28)
			field.Value = proto.Uint8(m.Intensity)
			field.IsExpandedField = expanded
			fields = append(fields, field)
		}
	}
	if m.DurationMin != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 29)
		field.Value = proto.Uint16(m.DurationMin)
		fields = append(fields, field)
	}
	if m.Duration != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 30)
		field.Value = proto.Uint32(m.Duration)
		fields = append(fields, field)
	}
	if m.Ascent != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 31)
		field.Value = proto.Uint32(m.Ascent)
		fields = append(fields, field)
	}
	if m.Descent != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 32)
		field.Value = proto.Uint32(m.Descent)
		fields = append(fields, field)
	}
	if m.ModerateActivityMinutes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 33)
		field.Value = proto.Uint16(m.ModerateActivityMinutes)
		fields = append(fields, field)
	}
	if m.VigorousActivityMinutes != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 34)
		field.Value = proto.Uint16(m.VigorousActivityMinutes)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// GetCycles returns Dynamic Field interpretation of Cycles. Otherwise, returns the original value of Cycles.
//
// Based on m.ActivityType:
//   - name: "steps", units: "steps" , value: uint32(m.Cycles)
//   - name: "strokes", units: "strokes" , value: (float64(m.Cycles) * 2) - 0
//
// Otherwise:
//   - name: "cycles", units: "cycles" , value: m.Cycles
func (m *Monitoring) GetCycles() (name string, value any) {
	switch m.ActivityType {
	case typedef.ActivityTypeWalking, typedef.ActivityTypeRunning:
		return "steps", uint32(m.Cycles)
	case typedef.ActivityTypeCycling, typedef.ActivityTypeSwimming:
		return "strokes", (float64(m.Cycles) * 2) - 0
	}
	return "cycles", m.Cycles
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Monitoring) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// LocalTimestampUint32 returns LocalTimestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *Monitoring) LocalTimestampUint32() uint32 { return datetime.ToUint32(m.LocalTimestamp) }

// DistanceScaled return Distance in its scaled value.
// If Distance value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) DistanceScaled() float64 {
	if m.Distance == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Distance)/100 - 0
}

// CyclesScaled return Cycles in its scaled value.
// If Cycles value is invalid, float64 invalid value will be returned.
//
// Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) CyclesScaled() float64 {
	if m.Cycles == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Cycles)/2 - 0
}

// ActiveTimeScaled return ActiveTime in its scaled value.
// If ActiveTime value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: s
func (m *Monitoring) ActiveTimeScaled() float64 {
	if m.ActiveTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.ActiveTime)/1000 - 0
}

// TemperatureScaled return Temperature in its scaled value.
// If Temperature value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
func (m *Monitoring) TemperatureScaled() float64 {
	if m.Temperature == basetype.Sint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Temperature)/100 - 0
}

// TemperatureMinScaled return TemperatureMin in its scaled value.
// If TemperatureMin value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
func (m *Monitoring) TemperatureMinScaled() float64 {
	if m.TemperatureMin == basetype.Sint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TemperatureMin)/100 - 0
}

// TemperatureMaxScaled return TemperatureMax in its scaled value.
// If TemperatureMax value is invalid, float64 invalid value will be returned.
//
// Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
func (m *Monitoring) TemperatureMaxScaled() float64 {
	if m.TemperatureMax == basetype.Sint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.TemperatureMax)/100 - 0
}

// IntensityScaled return Intensity in its scaled value.
// If Intensity value is invalid, float64 invalid value will be returned.
//
// Scale: 10
func (m *Monitoring) IntensityScaled() float64 {
	if m.Intensity == basetype.Uint8Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Intensity)/10 - 0
}

// AscentScaled return Ascent in its scaled value.
// If Ascent value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m
func (m *Monitoring) AscentScaled() float64 {
	if m.Ascent == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Ascent)/1000 - 0
}

// DescentScaled return Descent in its scaled value.
// If Descent value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m
func (m *Monitoring) DescentScaled() float64 {
	if m.Descent == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.Descent)/1000 - 0
}

// SetTimestamp sets Timestamp value.
//
// Units: s; Must align to logging interval, for example, time must be 00:00:00 for daily log.
func (m *Monitoring) SetTimestamp(v time.Time) *Monitoring {
	m.Timestamp = v
	return m
}

// SetDeviceIndex sets DeviceIndex value.
//
// Associates this data to device_info message. Not required for file with single device (sensor).
func (m *Monitoring) SetDeviceIndex(v typedef.DeviceIndex) *Monitoring {
	m.DeviceIndex = v
	return m
}

// SetCalories sets Calories value.
//
// Units: kcal; Accumulated total calories. Maintained by MonitoringReader for each activity_type. See SDK documentation
func (m *Monitoring) SetCalories(v uint16) *Monitoring {
	m.Calories = v
	return m
}

// SetDistance sets Distance value.
//
// Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetDistance(v uint32) *Monitoring {
	m.Distance = v
	return m
}

// SetDistanceScaled is similar to SetDistance except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 100; Units: m; Accumulated distance. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetDistanceScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Distance = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Distance = uint32(unscaled)
	return m
}

// SetCycles sets Cycles value.
//
// Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetCycles(v uint32) *Monitoring {
	m.Cycles = v
	return m
}

// SetCyclesScaled is similar to SetCycles except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 2; Units: cycles; Accumulated cycles. Maintained by MonitoringReader for each activity_type. See SDK documentation.
func (m *Monitoring) SetCyclesScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 2
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Cycles = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Cycles = uint32(unscaled)
	return m
}

// SetActiveTime sets ActiveTime value.
//
// Scale: 1000; Units: s
func (m *Monitoring) SetActiveTime(v uint32) *Monitoring {
	m.ActiveTime = v
	return m
}

// SetActiveTimeScaled is similar to SetActiveTime except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: s
func (m *Monitoring) SetActiveTimeScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.ActiveTime = uint32(basetype.Uint32Invalid)
		return m
	}
	m.ActiveTime = uint32(unscaled)
	return m
}

// SetActivityType sets ActivityType value.
func (m *Monitoring) SetActivityType(v typedef.ActivityType) *Monitoring {
	m.ActivityType = v
	return m
}

// SetActivitySubtype sets ActivitySubtype value.
func (m *Monitoring) SetActivitySubtype(v typedef.ActivitySubtype) *Monitoring {
	m.ActivitySubtype = v
	return m
}

// SetActivityLevel sets ActivityLevel value.
func (m *Monitoring) SetActivityLevel(v typedef.ActivityLevel) *Monitoring {
	m.ActivityLevel = v
	return m
}

// SetDistance16 sets Distance16 value.
//
// Units: 100 * m
func (m *Monitoring) SetDistance16(v uint16) *Monitoring {
	m.Distance16 = v
	return m
}

// SetCycles16 sets Cycles16 value.
//
// Units: 2 * cycles (steps)
func (m *Monitoring) SetCycles16(v uint16) *Monitoring {
	m.Cycles16 = v
	return m
}

// SetActiveTime16 sets ActiveTime16 value.
//
// Units: s
func (m *Monitoring) SetActiveTime16(v uint16) *Monitoring {
	m.ActiveTime16 = v
	return m
}

// SetLocalTimestamp sets LocalTimestamp value.
//
// Must align to logging interval, for example, time must be 00:00:00 for daily log.
func (m *Monitoring) SetLocalTimestamp(v time.Time) *Monitoring {
	m.LocalTimestamp = v
	return m
}

// SetTemperature sets Temperature value.
//
// Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperature(v int16) *Monitoring {
	m.Temperature = v
	return m
}

// SetTemperatureScaled is similar to SetTemperature except it accepts a scaled value.
// This method automatically converts the given value to its int16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: C; Avg temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
		m.Temperature = int16(basetype.Sint16Invalid)
		return m
	}
	m.Temperature = int16(unscaled)
	return m
}

// SetTemperatureMin sets TemperatureMin value.
//
// Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMin(v int16) *Monitoring {
	m.TemperatureMin = v
	return m
}

// SetTemperatureMinScaled is similar to SetTemperatureMin except it accepts a scaled value.
// This method automatically converts the given value to its int16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: C; Min temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMinScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
		m.TemperatureMin = int16(basetype.Sint16Invalid)
		return m
	}
	m.TemperatureMin = int16(unscaled)
	return m
}

// SetTemperatureMax sets TemperatureMax value.
//
// Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMax(v int16) *Monitoring {
	m.TemperatureMax = v
	return m
}

// SetTemperatureMaxScaled is similar to SetTemperatureMax except it accepts a scaled value.
// This method automatically converts the given value to its int16 form, discarding any applied scale and offset.
//
// Scale: 100; Units: C; Max temperature during the logging interval ended at timestamp
func (m *Monitoring) SetTemperatureMaxScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 100
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Sint16Invalid) {
		m.TemperatureMax = int16(basetype.Sint16Invalid)
		return m
	}
	m.TemperatureMax = int16(unscaled)
	return m
}

// SetActivityTime sets ActivityTime value.
//
// Array: [8]; Units: minutes; Indexed using minute_activity_level enum
func (m *Monitoring) SetActivityTime(v [8]uint16) *Monitoring {
	m.ActivityTime = v
	return m
}

// SetActiveCalories sets ActiveCalories value.
//
// Units: kcal
func (m *Monitoring) SetActiveCalories(v uint16) *Monitoring {
	m.ActiveCalories = v
	return m
}

// SetCurrentActivityTypeIntensity sets CurrentActivityTypeIntensity value.
//
// Indicates single type / intensity for duration since last monitoring message.
func (m *Monitoring) SetCurrentActivityTypeIntensity(v byte) *Monitoring {
	m.CurrentActivityTypeIntensity = v
	return m
}

// SetTimestampMin8 sets TimestampMin8 value.
//
// Units: min
func (m *Monitoring) SetTimestampMin8(v uint8) *Monitoring {
	m.TimestampMin8 = v
	return m
}

// SetTimestamp16 sets Timestamp16 value.
//
// Units: s
func (m *Monitoring) SetTimestamp16(v uint16) *Monitoring {
	m.Timestamp16 = v
	return m
}

// SetHeartRate sets HeartRate value.
//
// Units: bpm
func (m *Monitoring) SetHeartRate(v uint8) *Monitoring {
	m.HeartRate = v
	return m
}

// SetIntensity sets Intensity value.
//
// Scale: 10
func (m *Monitoring) SetIntensity(v uint8) *Monitoring {
	m.Intensity = v
	return m
}

// SetIntensityScaled is similar to SetIntensity except it accepts a scaled value.
// This method automatically converts the given value to its uint8 form, discarding any applied scale and offset.
//
// Scale: 10
func (m *Monitoring) SetIntensityScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 10
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint8Invalid) {
		m.Intensity = uint8(basetype.Uint8Invalid)
		return m
	}
	m.Intensity = uint8(unscaled)
	return m
}

// SetDurationMin sets DurationMin value.
//
// Units: min
func (m *Monitoring) SetDurationMin(v uint16) *Monitoring {
	m.DurationMin = v
	return m
}

// SetDuration sets Duration value.
//
// Units: s
func (m *Monitoring) SetDuration(v uint32) *Monitoring {
	m.Duration = v
	return m
}

// SetAscent sets Ascent value.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetAscent(v uint32) *Monitoring {
	m.Ascent = v
	return m
}

// SetAscentScaled is similar to SetAscent except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetAscentScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Ascent = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Ascent = uint32(unscaled)
	return m
}

// SetDescent sets Descent value.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetDescent(v uint32) *Monitoring {
	m.Descent = v
	return m
}

// SetDescentScaled is similar to SetDescent except it accepts a scaled value.
// This method automatically converts the given value to its uint32 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m
func (m *Monitoring) SetDescentScaled(v float64) *Monitoring {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint32Invalid) {
		m.Descent = uint32(basetype.Uint32Invalid)
		return m
	}
	m.Descent = uint32(unscaled)
	return m
}

// SetModerateActivityMinutes sets ModerateActivityMinutes value.
//
// Units: minutes
func (m *Monitoring) SetModerateActivityMinutes(v uint16) *Monitoring {
	m.ModerateActivityMinutes = v
	return m
}

// SetVigorousActivityMinutes sets VigorousActivityMinutes value.
//
// Units: minutes
func (m *Monitoring) SetVigorousActivityMinutes(v uint16) *Monitoring {
	m.VigorousActivityMinutes = v
	return m
}

// SetDeveloperFields Monitoring's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *Monitoring) SetUnknownFields(unknownFields ...proto.Field) *Monitoring {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields Monitoring's DeveloperFields.
func (m *Monitoring) SetDeveloperFields(developerFields ...proto.DeveloperField) *Monitoring {
	m.DeveloperFields = developerFields
	return m
}

// MarkAsExpandedField marks whether given fieldNum is an expanded field (field that being
// generated through a component expansion). Eligible for field number: 5, 28.
func (m *Monitoring) MarkAsExpandedField(fieldNum byte, flag bool) (ok bool) {
	switch fieldNum {
	case 5, 28:
	default:
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	m.state[pos] &^= bit
	if flag {
		m.state[pos] |= bit
	}
	return true
}

// IsExpandedField checks whether given fieldNum is a field generated through
// a component expansion. Eligible for field number: 5, 28.
func (m *Monitoring) IsExpandedField(fieldNum byte) bool {
	if fieldNum >= 29 {
		return false
	}
	pos := fieldNum / 8
	bit := uint8(1) << (fieldNum - (8 * pos))
	return m.state[pos]&bit == bit
}
