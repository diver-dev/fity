// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func BenchmarkNewRecord(b *testing.B) {
	mesg := factory.CreateMesg(mesgnum.Record)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NewRecord(&mesg)
	}
}

func BenchmarkRecordToMesg(b *testing.B) {
	mesg := factory.CreateMesg(mesgnum.Record)
	for i := range mesg.Fields {
		if mesg.Fields[i].Num == fieldnum.RecordTimestamp {
			mesg.Fields[i].Value = proto.Uint32(datetime.ToUint32(time.Now()))
			break
		}
	}
	record := NewRecord(&mesg)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = record.ToMesg(nil)
	}
}

func TestRecordFieldExpansionCorrectness(t *testing.T) {
	t.Run("NewRecord -> IsExpandedField()", func(t *testing.T) {
		r := NewRecord(&proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
			{FieldBase: factory.CreateField(mesgnum.Record, fieldnum.RecordEnhancedRespirationRate).FieldBase, IsExpandedField: true},
		}})
		if !r.IsExpandedField(fieldnum.RecordEnhancedRespirationRate) {
			t.Fatalf("expected expanded: true, got: false")
		}
	})

	r := NewRecord(nil)

	type numExpand struct {
		num        byte
		isExpanded bool
	}

	state := []numExpand{
		{5, false},
		{6, false},
		{19, false},
		{29, false},
		{73, false},
		{78, false},
		{108, false},
		{255, false}, // test outside range
	}

	// Check: init should all false.
	for i, v := range state {
		if res := r.IsExpandedField(v.num); res != v.isExpanded {
			t.Errorf("[%d] expected[%d]: %t, got: %t",
				i, v.num, v.isExpanded, res)
		}
	}

	// Set all to true and check
	for i := range state {
		r.MarkAsExpandedField(state[i].num, true)
		state[i].isExpanded = true
		if state[i].num == 255 { // invalid
			state[i].isExpanded = false
		}

		if res := r.IsExpandedField(state[i].num); res != state[i].isExpanded {
			t.Errorf("[%d] expected[%d]: %t, got: %t",
				i, state[i].num, state[i].isExpanded, res)
		}
	}

	// Set all to false and check
	for i := range state {
		r.MarkAsExpandedField(state[i].num, false)
		state[i].isExpanded = false

		if res := r.IsExpandedField(state[i].num); res != state[i].isExpanded {
			t.Errorf("[%d] expected[%d]: %t, got: %t",
				i, state[i].num, state[i].isExpanded, res)
		}
	}

	// Test partial
	for i := range state {
		switch state[i].num {
		case 19, 78, 255:
		default:
			r.MarkAsExpandedField(state[i].num, true)
			state[i].isExpanded = true
		}
	}

	// Check partial
	for i, v := range state {
		if res := r.IsExpandedField(v.num); res != v.isExpanded {
			t.Errorf("[%d] expected[%d]: %t, got: %t",
				i, v.num, v.isExpanded, res)
		}
	}
}

func TestRecordScaled(t *testing.T) {
	// const unscaleExpected = 37304
	const scaleExpected = 6960.8

	r := NewRecord(nil)
	r.Altitude = 37304
	scaled := r.AltitudeScaled()
	if scaled != scaleExpected {
		t.Fatalf("expected: %v, got: %v", scaleExpected, scaled)
	}
}

func TestRecordMarkAsScaled(t *testing.T) {
	tt := []struct {
		name      string
		initValue uint16
		input     float64
		expected  uint16
	}{
		{
			name:      "happy flow",
			initValue: basetype.Uint16Invalid,
			input:     6960.8,
			expected:  37304,
		},
		{
			name:      "float64 invalid",
			initValue: 0,
			input:     math.Float64frombits(basetype.Float64Invalid),
			expected:  basetype.Uint16Invalid,
		},
		{
			name:      "value higher than uint16",
			initValue: 0,
			input:     float64(math.MaxFloat32),
			expected:  basetype.Uint16Invalid,
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			r := NewRecord(nil).SetAltitude(tc.initValue)
			r.SetAltitudeScaled(tc.input)
			if r.Altitude != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, r.Altitude)
			}
		})
	}
}

func BenchmarkRecordIsExpandedField(b *testing.B) {
	r := NewRecord(nil)
	r.MarkAsExpandedField(19, true)
	if !r.IsExpandedField(19) {
		b.Fail()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.IsExpandedField(19)
	}
}

func BenchmarkRecordMarkAsExpandedField(b *testing.B) {
	r := NewRecord(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = r.MarkAsExpandedField(19, true)
	}
}

func TestRecordToMesgTimestampCorrectness(t *testing.T) {
	now := time.Now()

	r := NewRecord(nil)
	r.Timestamp = now
	mesg := r.ToMesg(nil)
	field := mesg.FieldByNum(proto.FieldNumTimestamp)

	if expected := datetime.ToUint32(now); field.Value.Uint32() != expected {
		t.Fatalf("expected: %d, got: %d", expected, field.Value.Uint32())
	}

	r.Timestamp = time.Time{}
	mesg = r.ToMesg(nil)
	field = mesg.FieldByNum(proto.FieldNumTimestamp)
	if field != nil {
		t.Fatalf("field should be nil, got: fieldNum: %d, value: %d",
			field.Num, field.Value.Uint32())
	}
}

func TestRecordUnknownFields(t *testing.T) {
	now := time.Now()

	t.Run("with unknown fields", func(t *testing.T) {
		r := NewRecord(&proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
			{FieldBase: &proto.FieldBase{Num: 252, Name: factory.NameUnknown}},
		}})
		mesg := r.ToMesg(nil)
		field := mesg.FieldByNum(252)
		if field == nil {
			t.Fatal("field should not be nil")
		}
		if field.Num != 252 && field.Name != factory.NameUnknown {
			t.Fatalf("expected num: %d, name: %s, got num: %d, name: %s",
				252, factory.NameUnknown, field.Num, field.Name)
		}
	})

	t.Run("without unknown fields", func(t *testing.T) {
		r := NewRecord(nil)
		if r.UnknownFields != nil {
			t.Fatalf("UnknownFields should be nil, got: %+v", r.UnknownFields)
		}
	})
}

func TestRecordReset(t *testing.T) {
	now := time.Now()
	distance := uint32(1000)
	mesg := &proto.Message{Num: mesgnum.Record, Fields: []proto.Field{
		factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
		factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(distance)),
	}}

	r1 := NewRecord(mesg)
	r2 := new(Record)
	r2.Reset(mesg)

	if diff := cmp.Diff(r1, r2,
		cmp.Transformer("float32", func(v float32) uint32 { return math.Float32bits(v) }),
		cmp.Exporter(func(_ reflect.Type) bool { return true }),
	); diff != "" {
		t.Fatal(diff)
	}

	if r1.Distance != distance && r1.Distance != r2.Distance {
		t.Fatalf("expected: %d, got %d != %d",
			distance, r1.Distance, r2.Distance)
	}

	r1.Reset(nil)
	r2.Reset(nil)

	if r1.Distance != basetype.Uint32Invalid && r1.Distance != r2.Distance {
		t.Fatalf("expected:%d, got: %d != %d",
			basetype.Uint32Invalid, r1.Distance, r2.Distance)
	}
}
