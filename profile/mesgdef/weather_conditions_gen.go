// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/semicircles"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// WeatherConditions is a WeatherConditions message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type WeatherConditions struct {
	Timestamp                time.Time // time of update for current conditions, else forecast time
	Location                 string    // string corresponding to GCS response location string
	ObservedAtTime           time.Time
	ObservedLocationLat      int32                 // Units: semicircles
	ObservedLocationLong     int32                 // Units: semicircles
	WindDirection            uint16                // Units: degrees
	WindSpeed                uint16                // Scale: 1000; Units: m/s
	WeatherReport            typedef.WeatherReport // Current or forecast
	Temperature              int8                  // Units: C
	Condition                typedef.WeatherStatus // Corresponds to GSC Response weatherIcon field
	PrecipitationProbability uint8                 // range 0-100
	TemperatureFeelsLike     int8                  // Units: C; Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
	RelativeHumidity         uint8
	DayOfWeek                typedef.DayOfWeek
	HighTemperature          int8 // Units: C
	LowTemperature           int8 // Units: C

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewWeatherConditions creates new WeatherConditions struct based on given mesg.
// If mesg is nil, it will return WeatherConditions with all fields being set to its corresponding invalid value.
func NewWeatherConditions(mesg *proto.Message) *WeatherConditions {
	vals := [254]proto.Value{}

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
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		if len(unknownFields) == 0 {
			unknownFields = nil
		}
		unknownFields = append(unknownFields[:0:0], unknownFields...)
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &WeatherConditions{
		Timestamp:                datetime.ToTime(vals[253].Uint32()),
		WeatherReport:            typedef.WeatherReport(vals[0].Uint8()),
		Temperature:              vals[1].Int8(),
		Condition:                typedef.WeatherStatus(vals[2].Uint8()),
		WindDirection:            vals[3].Uint16(),
		WindSpeed:                vals[4].Uint16(),
		PrecipitationProbability: vals[5].Uint8(),
		TemperatureFeelsLike:     vals[6].Int8(),
		RelativeHumidity:         vals[7].Uint8(),
		Location:                 vals[8].String(),
		ObservedAtTime:           datetime.ToTime(vals[9].Uint32()),
		ObservedLocationLat:      vals[10].Int32(),
		ObservedLocationLong:     vals[11].Int32(),
		DayOfWeek:                typedef.DayOfWeek(vals[12].Uint8()),
		HighTemperature:          vals[13].Int8(),
		LowTemperature:           vals[14].Int8(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts WeatherConditions into proto.Message. If options is nil, default options will be used.
func (m *WeatherConditions) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumWeatherConditions}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.WeatherReport != typedef.WeatherReportInvalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint8(byte(m.WeatherReport))
		fields = append(fields, field)
	}
	if m.Temperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Int8(m.Temperature)
		fields = append(fields, field)
	}
	if m.Condition != typedef.WeatherStatusInvalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint8(byte(m.Condition))
		fields = append(fields, field)
	}
	if m.WindDirection != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint16(m.WindDirection)
		fields = append(fields, field)
	}
	if m.WindSpeed != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint16(m.WindSpeed)
		fields = append(fields, field)
	}
	if m.PrecipitationProbability != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.PrecipitationProbability)
		fields = append(fields, field)
	}
	if m.TemperatureFeelsLike != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Int8(m.TemperatureFeelsLike)
		fields = append(fields, field)
	}
	if m.RelativeHumidity != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint8(m.RelativeHumidity)
		fields = append(fields, field)
	}
	if m.Location != basetype.StringInvalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.String(m.Location)
		fields = append(fields, field)
	}
	if !m.ObservedAtTime.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint32(uint32(m.ObservedAtTime.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ObservedLocationLat != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Int32(m.ObservedLocationLat)
		fields = append(fields, field)
	}
	if m.ObservedLocationLong != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Int32(m.ObservedLocationLong)
		fields = append(fields, field)
	}
	if m.DayOfWeek != typedef.DayOfWeekInvalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint8(byte(m.DayOfWeek))
		fields = append(fields, field)
	}
	if m.HighTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Int8(m.HighTemperature)
		fields = append(fields, field)
	}
	if m.LowTemperature != basetype.Sint8Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Int8(m.LowTemperature)
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

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeatherConditions) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// ObservedAtTimeUint32 returns ObservedAtTime in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *WeatherConditions) ObservedAtTimeUint32() uint32 { return datetime.ToUint32(m.ObservedAtTime) }

// WindSpeedScaled return WindSpeed in its scaled value.
// If WindSpeed value is invalid, float64 invalid value will be returned.
//
// Scale: 1000; Units: m/s
func (m *WeatherConditions) WindSpeedScaled() float64 {
	if m.WindSpeed == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return float64(m.WindSpeed)/1000 - 0
}

// ObservedLocationLatDegrees returns ObservedLocationLat in degrees instead of semicircles.
// If ObservedLocationLat value is invalid, float64 invalid value will be returned.
func (m *WeatherConditions) ObservedLocationLatDegrees() float64 {
	return semicircles.ToDegrees(m.ObservedLocationLat)
}

// ObservedLocationLongDegrees returns ObservedLocationLong in degrees instead of semicircles.
// If ObservedLocationLong value is invalid, float64 invalid value will be returned.
func (m *WeatherConditions) ObservedLocationLongDegrees() float64 {
	return semicircles.ToDegrees(m.ObservedLocationLong)
}

// SetTimestamp sets Timestamp value.
//
// time of update for current conditions, else forecast time
func (m *WeatherConditions) SetTimestamp(v time.Time) *WeatherConditions {
	m.Timestamp = v
	return m
}

// SetWeatherReport sets WeatherReport value.
//
// Current or forecast
func (m *WeatherConditions) SetWeatherReport(v typedef.WeatherReport) *WeatherConditions {
	m.WeatherReport = v
	return m
}

// SetTemperature sets Temperature value.
//
// Units: C
func (m *WeatherConditions) SetTemperature(v int8) *WeatherConditions {
	m.Temperature = v
	return m
}

// SetCondition sets Condition value.
//
// Corresponds to GSC Response weatherIcon field
func (m *WeatherConditions) SetCondition(v typedef.WeatherStatus) *WeatherConditions {
	m.Condition = v
	return m
}

// SetWindDirection sets WindDirection value.
//
// Units: degrees
func (m *WeatherConditions) SetWindDirection(v uint16) *WeatherConditions {
	m.WindDirection = v
	return m
}

// SetWindSpeed sets WindSpeed value.
//
// Scale: 1000; Units: m/s
func (m *WeatherConditions) SetWindSpeed(v uint16) *WeatherConditions {
	m.WindSpeed = v
	return m
}

// SetWindSpeedScaled is similar to SetWindSpeed except it accepts a scaled value.
// This method automatically converts the given value to its uint16 form, discarding any applied scale and offset.
//
// Scale: 1000; Units: m/s
func (m *WeatherConditions) SetWindSpeedScaled(v float64) *WeatherConditions {
	unscaled := (v + 0) * 1000
	if math.IsNaN(unscaled) || math.IsInf(unscaled, 0) || unscaled > float64(basetype.Uint16Invalid) {
		m.WindSpeed = uint16(basetype.Uint16Invalid)
		return m
	}
	m.WindSpeed = uint16(unscaled)
	return m
}

// SetPrecipitationProbability sets PrecipitationProbability value.
//
// range 0-100
func (m *WeatherConditions) SetPrecipitationProbability(v uint8) *WeatherConditions {
	m.PrecipitationProbability = v
	return m
}

// SetTemperatureFeelsLike sets TemperatureFeelsLike value.
//
// Units: C; Heat Index if GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
func (m *WeatherConditions) SetTemperatureFeelsLike(v int8) *WeatherConditions {
	m.TemperatureFeelsLike = v
	return m
}

// SetRelativeHumidity sets RelativeHumidity value.
func (m *WeatherConditions) SetRelativeHumidity(v uint8) *WeatherConditions {
	m.RelativeHumidity = v
	return m
}

// SetLocation sets Location value.
//
// string corresponding to GCS response location string
func (m *WeatherConditions) SetLocation(v string) *WeatherConditions {
	m.Location = v
	return m
}

// SetObservedAtTime sets ObservedAtTime value.
func (m *WeatherConditions) SetObservedAtTime(v time.Time) *WeatherConditions {
	m.ObservedAtTime = v
	return m
}

// SetObservedLocationLat sets ObservedLocationLat value.
//
// Units: semicircles
func (m *WeatherConditions) SetObservedLocationLat(v int32) *WeatherConditions {
	m.ObservedLocationLat = v
	return m
}

// SetObservedLocationLatDegrees is similar to SetObservedLocationLat except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *WeatherConditions) SetObservedLocationLatDegrees(degrees float64) *WeatherConditions {
	m.ObservedLocationLat = semicircles.ToSemicircles(degrees)
	return m
}

// SetObservedLocationLong sets ObservedLocationLong value.
//
// Units: semicircles
func (m *WeatherConditions) SetObservedLocationLong(v int32) *WeatherConditions {
	m.ObservedLocationLong = v
	return m
}

// SetObservedLocationLongDegrees is similar to SetObservedLocationLong except it accepts a value in degrees.
// This method will automatically convert given degrees value to semicircles (int32) form.
func (m *WeatherConditions) SetObservedLocationLongDegrees(degrees float64) *WeatherConditions {
	m.ObservedLocationLong = semicircles.ToSemicircles(degrees)
	return m
}

// SetDayOfWeek sets DayOfWeek value.
func (m *WeatherConditions) SetDayOfWeek(v typedef.DayOfWeek) *WeatherConditions {
	m.DayOfWeek = v
	return m
}

// SetHighTemperature sets HighTemperature value.
//
// Units: C
func (m *WeatherConditions) SetHighTemperature(v int8) *WeatherConditions {
	m.HighTemperature = v
	return m
}

// SetLowTemperature sets LowTemperature value.
//
// Units: C
func (m *WeatherConditions) SetLowTemperature(v int8) *WeatherConditions {
	m.LowTemperature = v
	return m
}

// SetDeveloperFields WeatherConditions's UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *WeatherConditions) SetUnknownFields(unknownFields ...proto.Field) *WeatherConditions {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields WeatherConditions's DeveloperFields.
func (m *WeatherConditions) SetDeveloperFields(developerFields ...proto.DeveloperField) *WeatherConditions {
	m.DeveloperFields = developerFields
	return m
}
