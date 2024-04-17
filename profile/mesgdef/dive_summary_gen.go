// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/kit/scaleoffset"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"math"
	"time"
)

// DiveSummary is a DiveSummary message.
type DiveSummary struct {
	Timestamp       time.Time // Units: s
	AvgDepth        uint32    // Scale: 1000; Units: m; 0 if above water
	MaxDepth        uint32    // Scale: 1000; Units: m; 0 if above water
	SurfaceInterval uint32    // Units: s; Time since end of last dive
	DiveNumber      uint32
	BottomTime      uint32 // Scale: 1000; Units: s
	DescentTime     uint32 // Scale: 1000; Units: s; Time to reach deepest level stop
	AscentTime      uint32 // Scale: 1000; Units: s; Time after leaving bottom until reaching surface
	AvgAscentRate   int32  // Scale: 1000; Units: m/s; Average ascent rate, not including descents or stops
	AvgDescentRate  uint32 // Scale: 1000; Units: m/s; Average descent rate, not including ascents or stops
	MaxAscentRate   uint32 // Scale: 1000; Units: m/s; Maximum ascent rate
	MaxDescentRate  uint32 // Scale: 1000; Units: m/s; Maximum descent rate
	HangTime        uint32 // Scale: 1000; Units: s; Time spent neither ascending nor descending
	ReferenceMesg   typedef.MesgNum
	ReferenceIndex  typedef.MessageIndex
	StartN2         uint16 // Units: percent
	EndN2           uint16 // Units: percent
	O2Toxicity      uint16 // Units: OTUs
	AvgPressureSac  uint16 // Scale: 100; Units: bar/min; Average pressure-based surface air consumption
	AvgVolumeSac    uint16 // Scale: 100; Units: L/min; Average volumetric surface air consumption
	AvgRmv          uint16 // Scale: 100; Units: L/min; Average respiratory minute volume
	StartCns        uint8  // Units: percent
	EndCns          uint8  // Units: percent

	// Developer Fields are dynamic, can't be mapped as struct's fields.
	// [Added since protocol version 2.0]
	DeveloperFields []proto.DeveloperField
}

// NewDiveSummary creates new DiveSummary struct based on given mesg.
// If mesg is nil, it will return DiveSummary with all fields being set to its corresponding invalid value.
func NewDiveSummary(mesg *proto.Message) *DiveSummary {
	vals := [254]proto.Value{}

	var developerFields []proto.DeveloperField
	if mesg != nil {
		for i := range mesg.Fields {
			if mesg.Fields[i].Num >= byte(len(vals)) {
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		developerFields = mesg.DeveloperFields
	}

	return &DiveSummary{
		Timestamp:       datetime.ToTime(vals[253].Uint32()),
		AvgDepth:        vals[2].Uint32(),
		MaxDepth:        vals[3].Uint32(),
		SurfaceInterval: vals[4].Uint32(),
		DiveNumber:      vals[10].Uint32(),
		BottomTime:      vals[11].Uint32(),
		DescentTime:     vals[15].Uint32(),
		AscentTime:      vals[16].Uint32(),
		AvgAscentRate:   vals[17].Int32(),
		AvgDescentRate:  vals[22].Uint32(),
		MaxAscentRate:   vals[23].Uint32(),
		MaxDescentRate:  vals[24].Uint32(),
		HangTime:        vals[25].Uint32(),
		ReferenceMesg:   typedef.MesgNum(vals[0].Uint16()),
		ReferenceIndex:  typedef.MessageIndex(vals[1].Uint16()),
		StartN2:         vals[7].Uint16(),
		EndN2:           vals[8].Uint16(),
		O2Toxicity:      vals[9].Uint16(),
		AvgPressureSac:  vals[12].Uint16(),
		AvgVolumeSac:    vals[13].Uint16(),
		AvgRmv:          vals[14].Uint16(),
		StartCns:        vals[5].Uint8(),
		EndCns:          vals[6].Uint8(),

		DeveloperFields: developerFields,
	}
}

// ToMesg converts DiveSummary into proto.Message. If options is nil, default options will be used.
func (m *DiveSummary) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[256]proto.Field)
	defer pool.Put(arr)

	fields := arr[:0] // Create slice from array with zero len.
	mesg := proto.Message{Num: typedef.MesgNumDiveSummary}

	if datetime.ToUint32(m.Timestamp) != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(datetime.ToUint32(m.Timestamp))
		fields = append(fields, field)
	}
	if m.AvgDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 2)
		field.Value = proto.Uint32(m.AvgDepth)
		fields = append(fields, field)
	}
	if m.MaxDepth != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 3)
		field.Value = proto.Uint32(m.MaxDepth)
		fields = append(fields, field)
	}
	if m.SurfaceInterval != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 4)
		field.Value = proto.Uint32(m.SurfaceInterval)
		fields = append(fields, field)
	}
	if m.DiveNumber != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 10)
		field.Value = proto.Uint32(m.DiveNumber)
		fields = append(fields, field)
	}
	if m.BottomTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 11)
		field.Value = proto.Uint32(m.BottomTime)
		fields = append(fields, field)
	}
	if m.DescentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 15)
		field.Value = proto.Uint32(m.DescentTime)
		fields = append(fields, field)
	}
	if m.AscentTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 16)
		field.Value = proto.Uint32(m.AscentTime)
		fields = append(fields, field)
	}
	if m.AvgAscentRate != basetype.Sint32Invalid {
		field := fac.CreateField(mesg.Num, 17)
		field.Value = proto.Int32(m.AvgAscentRate)
		fields = append(fields, field)
	}
	if m.AvgDescentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 22)
		field.Value = proto.Uint32(m.AvgDescentRate)
		fields = append(fields, field)
	}
	if m.MaxAscentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 23)
		field.Value = proto.Uint32(m.MaxAscentRate)
		fields = append(fields, field)
	}
	if m.MaxDescentRate != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 24)
		field.Value = proto.Uint32(m.MaxDescentRate)
		fields = append(fields, field)
	}
	if m.HangTime != basetype.Uint32Invalid {
		field := fac.CreateField(mesg.Num, 25)
		field.Value = proto.Uint32(m.HangTime)
		fields = append(fields, field)
	}
	if uint16(m.ReferenceMesg) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(uint16(m.ReferenceMesg))
		fields = append(fields, field)
	}
	if uint16(m.ReferenceIndex) != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.Uint16(uint16(m.ReferenceIndex))
		fields = append(fields, field)
	}
	if m.StartN2 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 7)
		field.Value = proto.Uint16(m.StartN2)
		fields = append(fields, field)
	}
	if m.EndN2 != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 8)
		field.Value = proto.Uint16(m.EndN2)
		fields = append(fields, field)
	}
	if m.O2Toxicity != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 9)
		field.Value = proto.Uint16(m.O2Toxicity)
		fields = append(fields, field)
	}
	if m.AvgPressureSac != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 12)
		field.Value = proto.Uint16(m.AvgPressureSac)
		fields = append(fields, field)
	}
	if m.AvgVolumeSac != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 13)
		field.Value = proto.Uint16(m.AvgVolumeSac)
		fields = append(fields, field)
	}
	if m.AvgRmv != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 14)
		field.Value = proto.Uint16(m.AvgRmv)
		fields = append(fields, field)
	}
	if m.StartCns != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 5)
		field.Value = proto.Uint8(m.StartCns)
		fields = append(fields, field)
	}
	if m.EndCns != basetype.Uint8Invalid {
		field := fac.CreateField(mesg.Num, 6)
		field.Value = proto.Uint8(m.EndCns)
		fields = append(fields, field)
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *DiveSummary) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// AvgDepthScaled return AvgDepth in its scaled value [Scale: 1000; Units: m; 0 if above water].
//
// If AvgDepth value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgDepthScaled() float64 {
	if m.AvgDepth == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgDepth, 1000, 0)
}

// MaxDepthScaled return MaxDepth in its scaled value [Scale: 1000; Units: m; 0 if above water].
//
// If MaxDepth value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) MaxDepthScaled() float64 {
	if m.MaxDepth == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.MaxDepth, 1000, 0)
}

// BottomTimeScaled return BottomTime in its scaled value [Scale: 1000; Units: s].
//
// If BottomTime value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) BottomTimeScaled() float64 {
	if m.BottomTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.BottomTime, 1000, 0)
}

// DescentTimeScaled return DescentTime in its scaled value [Scale: 1000; Units: s; Time to reach deepest level stop].
//
// If DescentTime value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) DescentTimeScaled() float64 {
	if m.DescentTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.DescentTime, 1000, 0)
}

// AscentTimeScaled return AscentTime in its scaled value [Scale: 1000; Units: s; Time after leaving bottom until reaching surface].
//
// If AscentTime value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AscentTimeScaled() float64 {
	if m.AscentTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AscentTime, 1000, 0)
}

// AvgAscentRateScaled return AvgAscentRate in its scaled value [Scale: 1000; Units: m/s; Average ascent rate, not including descents or stops].
//
// If AvgAscentRate value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgAscentRateScaled() float64 {
	if m.AvgAscentRate == basetype.Sint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgAscentRate, 1000, 0)
}

// AvgDescentRateScaled return AvgDescentRate in its scaled value [Scale: 1000; Units: m/s; Average descent rate, not including ascents or stops].
//
// If AvgDescentRate value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgDescentRateScaled() float64 {
	if m.AvgDescentRate == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgDescentRate, 1000, 0)
}

// MaxAscentRateScaled return MaxAscentRate in its scaled value [Scale: 1000; Units: m/s; Maximum ascent rate].
//
// If MaxAscentRate value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) MaxAscentRateScaled() float64 {
	if m.MaxAscentRate == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.MaxAscentRate, 1000, 0)
}

// MaxDescentRateScaled return MaxDescentRate in its scaled value [Scale: 1000; Units: m/s; Maximum descent rate].
//
// If MaxDescentRate value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) MaxDescentRateScaled() float64 {
	if m.MaxDescentRate == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.MaxDescentRate, 1000, 0)
}

// HangTimeScaled return HangTime in its scaled value [Scale: 1000; Units: s; Time spent neither ascending nor descending].
//
// If HangTime value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) HangTimeScaled() float64 {
	if m.HangTime == basetype.Uint32Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.HangTime, 1000, 0)
}

// AvgPressureSacScaled return AvgPressureSac in its scaled value [Scale: 100; Units: bar/min; Average pressure-based surface air consumption].
//
// If AvgPressureSac value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgPressureSacScaled() float64 {
	if m.AvgPressureSac == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgPressureSac, 100, 0)
}

// AvgVolumeSacScaled return AvgVolumeSac in its scaled value [Scale: 100; Units: L/min; Average volumetric surface air consumption].
//
// If AvgVolumeSac value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgVolumeSacScaled() float64 {
	if m.AvgVolumeSac == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgVolumeSac, 100, 0)
}

// AvgRmvScaled return AvgRmv in its scaled value [Scale: 100; Units: L/min; Average respiratory minute volume].
//
// If AvgRmv value is invalid, float64 invalid value will be returned.
func (m *DiveSummary) AvgRmvScaled() float64 {
	if m.AvgRmv == basetype.Uint16Invalid {
		return math.Float64frombits(basetype.Float64Invalid)
	}
	return scaleoffset.Apply(m.AvgRmv, 100, 0)
}

// SetTimestamp sets DiveSummary value.
//
// Units: s
func (m *DiveSummary) SetTimestamp(v time.Time) *DiveSummary {
	m.Timestamp = v
	return m
}

// SetAvgDepth sets DiveSummary value.
//
// Scale: 1000; Units: m; 0 if above water
func (m *DiveSummary) SetAvgDepth(v uint32) *DiveSummary {
	m.AvgDepth = v
	return m
}

// SetMaxDepth sets DiveSummary value.
//
// Scale: 1000; Units: m; 0 if above water
func (m *DiveSummary) SetMaxDepth(v uint32) *DiveSummary {
	m.MaxDepth = v
	return m
}

// SetSurfaceInterval sets DiveSummary value.
//
// Units: s; Time since end of last dive
func (m *DiveSummary) SetSurfaceInterval(v uint32) *DiveSummary {
	m.SurfaceInterval = v
	return m
}

// SetDiveNumber sets DiveSummary value.
func (m *DiveSummary) SetDiveNumber(v uint32) *DiveSummary {
	m.DiveNumber = v
	return m
}

// SetBottomTime sets DiveSummary value.
//
// Scale: 1000; Units: s
func (m *DiveSummary) SetBottomTime(v uint32) *DiveSummary {
	m.BottomTime = v
	return m
}

// SetDescentTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time to reach deepest level stop
func (m *DiveSummary) SetDescentTime(v uint32) *DiveSummary {
	m.DescentTime = v
	return m
}

// SetAscentTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time after leaving bottom until reaching surface
func (m *DiveSummary) SetAscentTime(v uint32) *DiveSummary {
	m.AscentTime = v
	return m
}

// SetAvgAscentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Average ascent rate, not including descents or stops
func (m *DiveSummary) SetAvgAscentRate(v int32) *DiveSummary {
	m.AvgAscentRate = v
	return m
}

// SetAvgDescentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Average descent rate, not including ascents or stops
func (m *DiveSummary) SetAvgDescentRate(v uint32) *DiveSummary {
	m.AvgDescentRate = v
	return m
}

// SetMaxAscentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Maximum ascent rate
func (m *DiveSummary) SetMaxAscentRate(v uint32) *DiveSummary {
	m.MaxAscentRate = v
	return m
}

// SetMaxDescentRate sets DiveSummary value.
//
// Scale: 1000; Units: m/s; Maximum descent rate
func (m *DiveSummary) SetMaxDescentRate(v uint32) *DiveSummary {
	m.MaxDescentRate = v
	return m
}

// SetHangTime sets DiveSummary value.
//
// Scale: 1000; Units: s; Time spent neither ascending nor descending
func (m *DiveSummary) SetHangTime(v uint32) *DiveSummary {
	m.HangTime = v
	return m
}

// SetReferenceMesg sets DiveSummary value.
func (m *DiveSummary) SetReferenceMesg(v typedef.MesgNum) *DiveSummary {
	m.ReferenceMesg = v
	return m
}

// SetReferenceIndex sets DiveSummary value.
func (m *DiveSummary) SetReferenceIndex(v typedef.MessageIndex) *DiveSummary {
	m.ReferenceIndex = v
	return m
}

// SetStartN2 sets DiveSummary value.
//
// Units: percent
func (m *DiveSummary) SetStartN2(v uint16) *DiveSummary {
	m.StartN2 = v
	return m
}

// SetEndN2 sets DiveSummary value.
//
// Units: percent
func (m *DiveSummary) SetEndN2(v uint16) *DiveSummary {
	m.EndN2 = v
	return m
}

// SetO2Toxicity sets DiveSummary value.
//
// Units: OTUs
func (m *DiveSummary) SetO2Toxicity(v uint16) *DiveSummary {
	m.O2Toxicity = v
	return m
}

// SetAvgPressureSac sets DiveSummary value.
//
// Scale: 100; Units: bar/min; Average pressure-based surface air consumption
func (m *DiveSummary) SetAvgPressureSac(v uint16) *DiveSummary {
	m.AvgPressureSac = v
	return m
}

// SetAvgVolumeSac sets DiveSummary value.
//
// Scale: 100; Units: L/min; Average volumetric surface air consumption
func (m *DiveSummary) SetAvgVolumeSac(v uint16) *DiveSummary {
	m.AvgVolumeSac = v
	return m
}

// SetAvgRmv sets DiveSummary value.
//
// Scale: 100; Units: L/min; Average respiratory minute volume
func (m *DiveSummary) SetAvgRmv(v uint16) *DiveSummary {
	m.AvgRmv = v
	return m
}

// SetStartCns sets DiveSummary value.
//
// Units: percent
func (m *DiveSummary) SetStartCns(v uint8) *DiveSummary {
	m.StartCns = v
	return m
}

// SetEndCns sets DiveSummary value.
//
// Units: percent
func (m *DiveSummary) SetEndCns(v uint8) *DiveSummary {
	m.EndCns = v
	return m
}

// SetDeveloperFields DiveSummary's DeveloperFields.
func (m *DiveSummary) SetDeveloperFields(developerFields ...proto.DeveloperField) *DiveSummary {
	m.DeveloperFields = developerFields
	return m
}
