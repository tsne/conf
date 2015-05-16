package conf

import (
	"math"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestDecodeBool(t *testing.T) {
	var b bool

	validDecode(t, &b, true)
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, false)
	if b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, int(3))
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, int(0))
	if b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, uint(3))
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, uint(0))
	if b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, float64(3))
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, float64(0))
	if b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	validDecode(t, &b, "true")
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	invalidDecode(t, &b, "foobar")
	invalidDecode(t, &b, struct{}{})
}

func TestDecodeInt(t *testing.T) {
	var i int
	validDecode(t, &i, -7)
	if i != -7 {
		t.Fatalf("unexpected int value: %d", i)
	}

	validDecode(t, &i, true)
	if i != 1 {
		t.Fatalf("unexpected int value: %d", i)
	}

	validDecode(t, &i, false)
	if i != 0 {
		t.Fatalf("unexpected int value: %d", i)
	}

	validDecode(t, &i, float64(17.1))
	if i != 17 {
		t.Fatalf("unexpected int value: %d", i)
	}

	var i8 int8
	validDecode(t, &i8, math.MinInt8)
	if i8 != math.MinInt8 {
		t.Fatalf("unexpected int8 value: %d", i8)
	}
	validDecode(t, &i8, uint8(math.MaxInt8))
	if i8 != math.MaxInt8 {
		t.Fatalf("unexpected int8 value: %d", i8)
	}

	var i16 int16
	validDecode(t, &i16, math.MinInt16)
	if i16 != math.MinInt16 {
		t.Fatalf("unexpected int16 value: %d", i16)
	}
	validDecode(t, &i16, uint16(math.MaxInt16))
	if i16 != math.MaxInt16 {
		t.Fatalf("unexpected int16 value: %d", i16)
	}

	var i32 int32
	validDecode(t, &i32, math.MinInt32)
	if i32 != math.MinInt32 {
		t.Fatalf("unexpected int32 value: %d", i32)
	}
	validDecode(t, &i32, uint32(math.MaxInt32))
	if i32 != math.MaxInt32 {
		t.Fatalf("unexpected int32 value: %d", i32)
	}

	var i64 int64
	validDecode(t, &i64, math.MinInt64)
	if i64 != math.MinInt64 {
		t.Fatalf("unexpected int64 value: %d", i64)
	}
	validDecode(t, &i64, uint64(math.MaxInt64))
	if i64 != math.MaxInt64 {
		t.Fatalf("unexpected int64 value: %d", i64)
	}

	// underflow
	invalidDecode(t, &i8, math.MinInt8-1)
	invalidDecode(t, &i16, math.MinInt16-1)
	invalidDecode(t, &i32, float64(math.MinInt32-1))

	// overflow
	invalidDecode(t, &i8, math.MaxInt8+1)
	invalidDecode(t, &i16, math.MaxInt16+1)
	invalidDecode(t, &i32, uint64(math.MaxInt32+1))

	// invalid value
	invalidDecode(t, &i8, "a")
	invalidDecode(t, &i16, "b")
	invalidDecode(t, &i32, struct{}{})
	invalidDecode(t, &i64, struct{}{})
}

func TestDecodeUint(t *testing.T) {
	var ui uint
	validDecode(t, &ui, 7)
	if ui != 7 {
		t.Fatalf("unexpected uint value: %d", ui)
	}

	validDecode(t, &ui, true)
	if ui != 1 {
		t.Fatalf("unexpected int value: %d", ui)
	}

	validDecode(t, &ui, false)
	if ui != 0 {
		t.Fatalf("unexpected int value: %d", ui)
	}

	validDecode(t, &ui, float64(17.1))
	if ui != 17 {
		t.Fatalf("unexpected int value: %d", ui)
	}

	var ui8 uint8
	validDecode(t, &ui8, math.MaxUint8)
	if ui8 != math.MaxUint8 {
		t.Fatalf("unexpected uint8 value: %d", ui8)
	}
	validDecode(t, &ui8, 0)
	if ui8 != 0 {
		t.Fatalf("unexpected uint8 value: %d", ui8)
	}

	var ui16 uint16
	validDecode(t, &ui16, math.MaxUint16)
	if ui16 != math.MaxUint16 {
		t.Fatalf("unexpected uint16 value: %d", ui16)
	}
	validDecode(t, &ui16, 0)
	if ui16 != 0 {
		t.Fatalf("unexpected uint16 value: %d", ui16)
	}

	var ui32 uint32
	validDecode(t, &ui32, math.MaxUint32)
	if ui32 != math.MaxUint32 {
		t.Fatalf("unexpected uint32 value: %d", ui32)
	}
	validDecode(t, &ui32, 0)
	if ui32 != 0 {
		t.Fatalf("unexpected uint32 value: %d", ui32)
	}

	var ui64 uint64
	validDecode(t, &ui64, uint64(math.MaxUint64))
	if ui64 != math.MaxUint64 {
		t.Fatalf("unexpected uint64 value: %d", ui64)
	}
	validDecode(t, &ui64, 0)
	if ui64 != 0 {
		t.Fatalf("unexpected uint64 value: %d", ui64)
	}

	// underflow
	invalidDecode(t, &ui, -1)
	invalidDecode(t, &ui8, -1)
	invalidDecode(t, &ui16, -1)
	invalidDecode(t, &ui32, int(-1))
	invalidDecode(t, &ui64, float64(-1))

	// overflow
	invalidDecode(t, &ui8, math.MaxUint8+1)
	invalidDecode(t, &ui16, uint64(math.MaxUint16+1))
	invalidDecode(t, &ui32, float64(math.MaxUint32+1))

	// invalid value
	invalidDecode(t, &ui8, "a")
	invalidDecode(t, &ui16, "b")
	invalidDecode(t, &ui32, struct{}{})
	invalidDecode(t, &ui64, struct{}{})
}

func TestDecodeDuration(t *testing.T) {
	var d time.Duration

	validDecode(t, &d, "5s")
	if d != 5*time.Second {
		t.Fatalf("unexpected duration value: %s", d)
	}

	validDecode(t, &d, int64(12))
	if int64(d) != 12 {
		t.Fatalf("unexpected duration value: %s", d)
	}

	validDecode(t, &d, uint8(100))
	if int64(d) != 100 {
		t.Fatalf("unexpected duration value: %s", d)
	}

	validDecode(t, &d, float32(73))
	if int64(d) != 73 {
		t.Fatalf("unexpected duration value: %s", d)
	}

	invalidDecode(t, &d, "4r")
	invalidDecode(t, &d, uint64(math.MaxInt64)+1)
	invalidDecode(t, &d, 2*float64(math.MaxInt64))
	invalidDecode(t, &d, 2*float64(math.MinInt64))
	invalidDecode(t, &d, struct{}{})
}

func TestDecodeFloat(t *testing.T) {
	var f32 float32

	validDecode(t, &f32, true)
	if f32 != 1 {
		t.Fatalf("unexpected float32 value: %f", f32)
	}

	validDecode(t, &f32, uint16(12))
	if f32 != 12 {
		t.Fatalf("unexpected float64 value: %f", f32)
	}

	validDecode(t, &f32, -math.MaxFloat32)
	if f32 != -math.MaxFloat32 {
		t.Fatalf("unexpected float32 value: %f", f32)
	}

	validDecode(t, &f32, math.MaxFloat32)
	if f32 != math.MaxFloat32 {
		t.Fatalf("unexpected float32 value: %f", f32)
	}

	var f64 float64

	validDecode(t, &f64, false)
	if f64 != 0 {
		t.Fatalf("unexpected float64 value: %f", f64)
	}

	validDecode(t, &f64, uint8(10))
	if f64 != 10 {
		t.Fatalf("unexpected float64 value: %f", f64)
	}

	validDecode(t, &f64, -math.MaxFloat64)
	if f64 != -math.MaxFloat64 {
		t.Fatalf("unexpected float64 value: %f", f64)
	}

	validDecode(t, &f64, math.MaxFloat64)
	if f64 != math.MaxFloat64 {
		t.Fatalf("unexpected float64 value: %f", f64)
	}

	// underflow
	invalidDecode(t, &f32, -2*math.MaxFloat32)

	// overflow
	invalidDecode(t, &f32, 2*math.MaxFloat32)
}

func TestDecodeString(t *testing.T) {
	var s string

	validDecode(t, &s, "foo")
	if s != "foo" {
		t.Fatalf("unexpected string value: %s", s)
	}

	validDecode(t, &s, true)
	if s != "true" {
		t.Fatalf("unexpected string value: %s", s)
	}

	validDecode(t, &s, int(-7))
	if s != "-7" {
		t.Fatalf("unexpected string value: %s", s)
	}

	validDecode(t, &s, uint(7))
	if s != "7" {
		t.Fatalf("unexpected string value: %s", s)
	}

	validDecode(t, &s, float64(7))
	if s != "7" {
		t.Fatalf("unexpected string value: %s", s)
	}
}

func TestDecodeArray(t *testing.T) {
	var ai [3]int

	validDecode(t, &ai, []int{1, 2, 3})
	if ai[0] != 1 || ai[1] != 2 || ai[2] != 3 {
		t.Fatalf("unexpected array value: %v", ai)
	}

	validDecode(t, &ai, []string{"1", "2", "3"})
	if ai[0] != 1 || ai[1] != 2 || ai[2] != 3 {
		t.Fatalf("unexpected array value: %v", ai)
	}

	var as [3]string

	validDecode(t, &as, []string{"a", "b", "c"})
	if as[0] != "a" || as[1] != "b" || as[2] != "c" {
		t.Fatalf("unexpected array value: %v", as)
	}

	validDecode(t, &as, []int{1, 2, 3})
	if as[0] != "1" || as[1] != "2" || as[2] != "3" {
		t.Fatalf("unexpected array value: %v", as)
	}

	var af [1]float64

	validDecode(t, &af, 0.1)
	if af[0] != 0.1 {
		t.Fatalf("unexpected array value: %v", af)
	}

	invalidDecode(t, &ai, []string{"a", "2", "3"})
	invalidDecode(t, &as, []string{"a", "b"})
	invalidDecode(t, &as, "foo")
}

func TestDecodeSlice(t *testing.T) {
	var si []int

	validDecode(t, &si, []int{1, 2})
	if len(si) != 2 || si[0] != 1 || si[1] != 2 {
		t.Fatalf("unexpected slice value: %v", si)
	}

	validDecode(t, &si, []string{"1", "2", "3"})
	if len(si) != 3 || si[0] != 1 || si[1] != 2 || si[2] != 3 {
		t.Fatalf("unexpected slice value: %v", si)
	}

	var ss []string

	validDecode(t, &ss, []string{"a", "b", "c", "d"})
	if len(ss) != 4 || ss[0] != "a" || ss[1] != "b" || ss[2] != "c" || ss[3] != "d" {
		t.Fatalf("unexpected slice value: %v", ss)
	}

	validDecode(t, &ss, []int{1, 2, 3})
	if len(ss) != 3 || ss[0] != "1" || ss[1] != "2" || ss[2] != "3" {
		t.Fatalf("unexpected slice value: %v", ss)
	}

	validDecode(t, &ss, "foo")
	if len(ss) != 1 || ss[0] != "foo" {
		t.Fatalf("unexpected slice value: %v", ss)
	}

	invalidDecode(t, &si, []string{"a", "2", "3"})
	invalidDecode(t, &si, "foo")
}

func TestDecodeMap(t *testing.T) {
	msi := make(map[string]int)

	validDecode(t, &msi, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	if len(msi) != 3 || msi["one"] != 1 || msi["two"] != 2 || msi["three"] != 3 {
		t.Fatalf("unexpected map value: %v", msi)
	}

	validDecode(t, &msi, map[string]string{
		"four": "4",
		"five": "5",
		"six":  "6",
	})
	if len(msi) != 3 || msi["four"] != 4 || msi["five"] != 5 || msi["six"] != 6 {
		t.Fatalf("unexpected map value: %v", msi)
	}

	mii := make(map[int]float64)

	validDecode(t, &mii, map[float32]int{
		2.1: 4,
		3:   9,
		4:   16,
	})
	if len(mii) != 3 || mii[2] != 4 || mii[3] != 9 || mii[4] != 16 {
		t.Fatalf("unexpected map value: %v", mii)
	}

	validDecode(t, &mii, map[string]string{
		"2": "4",
		"3": "9",
		"4": "16",
	})
	if len(mii) != 3 || mii[2] != 4 || mii[3] != 9 || mii[4] != 16 {
		t.Fatalf("unexpected map value: %v", mii)
	}

	invalidDecode(t, &msi, "not a map")
	invalidDecode(t, &msi, map[string]string{
		"four": "a",
		"five": "5",
		"six":  "6",
	})
	invalidDecode(t, &mii, map[int]string{
		1: "a",
		2: "5",
	})
	invalidDecode(t, &mii, map[string]string{
		"b": "a",
		"2": "5",
	})
}

func TestDecodeInterface(t *testing.T) {
	var v interface{}

	validDecode(t, &v, int(7))
	i, ok := v.(int)
	switch {
	case !ok:
		t.Fatalf("unexpected interface type: %T", v)
	case i != 7:
		t.Fatalf("unexpected interface value: %v", v)
	}

	validDecode(t, &v, "foo")
	s, ok := v.(string)
	switch {
	case !ok:
		t.Fatalf("unexpected interface type: %T", v)
	case s != "foo":
		t.Fatalf("unexpected interface value: %v", v)
	}
}

func TestDecodeSimpleStruct(t *testing.T) {
	simple := struct {
		Foo    int
		Fu     string
		Bar    []int
		Baz    []string
		Foobar map[string]string
	}{}

	validDecode(t, &simple, map[string]interface{}{
		"Foo": 7,
		"fu":  "seven",
		"Bar": []int{1, 1, 2, 3, 5, 8},
		"baz": []string{"one", "one", "two", "three", "five", "eight"},
		"Foobar": map[string]string{
			"user": "0815",
			"meta": "my-meta-data",
		},
		"unused-key": "unused",
	})
	if simple.Foo != 7 ||
		simple.Fu != "seven" ||
		len(simple.Bar) != 6 ||
		simple.Bar[0] != 1 ||
		simple.Bar[1] != 1 ||
		simple.Bar[2] != 2 ||
		simple.Bar[3] != 3 ||
		simple.Bar[4] != 5 ||
		simple.Bar[5] != 8 ||
		len(simple.Baz) != 6 ||
		simple.Baz[0] != "one" ||
		simple.Baz[1] != "one" ||
		simple.Baz[2] != "two" ||
		simple.Baz[3] != "three" ||
		simple.Baz[4] != "five" ||
		simple.Baz[5] != "eight" ||
		len(simple.Foobar) != 2 ||
		simple.Foobar["user"] != "0815" ||
		simple.Foobar["meta"] != "my-meta-data" {

		t.Fatalf("unexpected struct value: %+v", simple)
	}

	partial := struct {
		Foo    uint8
		Bar    string
		Foobar []float32
	}{}

	validDecode(t, &partial, map[string]interface{}{
		"foo":        "7",
		"unused-key": "unused",
	})
	if partial.Foo != 7 ||
		partial.Bar != "" ||
		len(partial.Foobar) != 0 {

		t.Fatalf("unexpected struct value: %+v", partial)
	}

	invalidDecode(t, &simple, "no a map")
	invalidDecode(t, &partial, map[string]interface{}{
		"foo": "a",
	})
}

func TestDecodeNestedStruct(t *testing.T) {
	nested := struct {
		Foo string
		Bar struct {
			Sub1 int
			Sub2 float32
		}
	}{}

	validDecode(t, &nested, map[string]interface{}{
		"Foo": "foo",
		"bar": map[string]interface{}{
			"Sub1": 7,
			"sub2": 1.5,
		},
		"unused-key": "unused",
	})
	if nested.Foo != "foo" ||
		nested.Bar.Sub1 != 7 ||
		nested.Bar.Sub2 != 1.5 {

		t.Fatalf("unexpected struct value: %+v", nested)
	}

	partial := struct {
		Foo struct {
			Sub int
		}
		Bar struct {
			Sub string
		}
	}{}

	validDecode(t, &partial, map[string]interface{}{
		"foo": map[string]interface{}{
			"sub":        7,
			"unused-key": "unused",
		},
		"another-unused-key": "useless",
	})
	if partial.Foo.Sub != 7 ||
		partial.Bar.Sub != "" {

		t.Fatalf("unexpected struct value: %+v", nested)
	}
}

func TestDecodeStructTags(t *testing.T) {
	tagged := struct {
		Renamed   float64 `config:"ren"`
		Required1 int     `config:",required"`
		Required2 int     `config:"req2,required"`
		Omitted   string  `config:"-"`
		Foobar    []string
	}{}

	validDecode(t, &tagged, map[string]interface{}{
		"ren":       7,
		"required1": "13",
		"req2":      23,
		"omitted":   "foo",
	})
	if tagged.Renamed != 7 ||
		tagged.Required1 != 13 ||
		tagged.Required2 != 23 ||
		tagged.Omitted != "" ||
		len(tagged.Foobar) != 0 {

		t.Fatalf("unexpected struct value: %+v", tagged)
	}

	invalidDecode(t, &tagged, map[string]interface{}{
		"ren":    "1",
		"req2":   "2",
		"Foobar": []string{"foo", "bar"},
	})

	invalidDecode(t, &tagged, map[int]interface{}{
		7:  "foo",
		13: "bar",
	})
}

func validDecode(t *testing.T, out, in interface{}) {
	err := decode(reflect.ValueOf(out), reflect.ValueOf(in))
	if err != nil {
		_, _, line, _ := runtime.Caller(1)
		t.Fatalf("unexpected error decoding %T: %v (line %d)", out, err, line)
	}
}

func invalidDecode(t *testing.T, out, in interface{}) {
	err := decode(reflect.ValueOf(out), reflect.ValueOf(in))
	if err == nil {
		_, _, line, _ := runtime.Caller(1)
		t.Fatalf("error expected decoding %T, got none (line %d)", out, line)
	}
}
