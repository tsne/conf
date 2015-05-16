package conf

import (
	"reflect"
	"testing"
)

func TestValueBool(t *testing.T) {
	v := newTestValue(true)

	b, err := v.Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 1 {
		t.Fatalf("unexpected int value: %v", i)
	}

	ui, err := v.Uint()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ui != 1 {
		t.Fatalf("unexpected uint value: %v", ui)
	}

	f, err := v.Float()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != 1 {
		t.Fatalf("unexpected float value: %v", f)
	}

	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "true" {
		t.Fatalf("unexpected string value: %v", s)
	}
}

func TestValueInt(t *testing.T) {
	v := newTestValue(7)

	b, err := v.Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 7 {
		t.Fatalf("unexpected int value: %v", i)
	}

	ui, err := v.Uint()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ui != 7 {
		t.Fatalf("unexpected uint value: %v", ui)
	}

	f, err := v.Float()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != 7 {
		t.Fatalf("unexpected float value: %v", f)
	}

	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "7" {
		t.Fatalf("unexpected string value: %v", s)
	}
}

func TestValueUint(t *testing.T) {
	v := newTestValue(uint(7))

	b, err := v.Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 7 {
		t.Fatalf("unexpected int value: %v", i)
	}

	ui, err := v.Uint()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ui != 7 {
		t.Fatalf("unexpected uint value: %v", ui)
	}

	f, err := v.Float()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != 7 {
		t.Fatalf("unexpected float value: %v", f)
	}

	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "7" {
		t.Fatalf("unexpected string value: %v", s)
	}
}

func TestValueFloat(t *testing.T) {
	v := newTestValue(1.2)

	b, err := v.Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 1 {
		t.Fatalf("unexpected int value: %v", i)
	}

	ui, err := v.Uint()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ui != 1 {
		t.Fatalf("unexpected uint value: %v", ui)
	}

	f, err := v.Float()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != 1.2 {
		t.Fatalf("unexpected float value: %v", f)
	}

	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "1.2" {
		t.Fatalf("unexpected string value: %v", s)
	}
}

func TestValueString(t *testing.T) {
	v := newTestValue("3")

	b, err := v.Bool()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !b {
		t.Fatalf("unexpected bool value: %v", b)
	}

	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 3 {
		t.Fatalf("unexpected int value: %v", i)
	}

	ui, err := v.Uint()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ui != 3 {
		t.Fatalf("unexpected uint value: %v", ui)
	}

	f, err := v.Float()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f != 3 {
		t.Fatalf("unexpected float value: %v", f)
	}

	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "3" {
		t.Fatalf("unexpected string value: %v", s)
	}
}

func newTestValue(v interface{}) Value {
	return Value(reflect.ValueOf(v))
}
