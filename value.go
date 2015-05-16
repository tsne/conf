package conf

import (
	"fmt"
	"reflect"
)

// Value represents a specific configuration value which can be converted
// to several types.
type Value reflect.Value

// Bool tries to convert the configuration value to a boolean value.
// If the conversion fails an error is returned.
func (v *Value) Bool() (bool, error) {
	var b bool
	err := v.populate(&b)
	return b, err
}

// Int tries to convert the configuration value to an integer value.
// If the conversion fails an error is returned.
func (v *Value) Int() (int64, error) {
	var i int64
	err := v.populate(&i)
	return i, err
}

// Uint tries to convert the configuration value to an unsigned integer value.
// If the conversion fails an error is returned.
func (v *Value) Uint() (uint64, error) {
	var i uint64
	err := v.populate(&i)
	return i, err
}

// Float tries to convert the configuration value to a floating point value.
// If the conversion fails an error is returned.
func (v *Value) Float() (float64, error) {
	var f float64
	err := v.populate(&f)
	return f, err
}

// Float tries to convert the configuration value to a string value.
// If the conversion fails an error is returned.
func (v *Value) String() (string, error) {
	var s string
	err := v.populate(&s)
	return s, err
}

func (v *Value) populate(value interface{}) error {
	input := *(*reflect.Value)(v)
	output := reflect.ValueOf(value)
	if output.Kind() != reflect.Ptr {
		return fmt.Errorf("'%T' is not a pointer type", value)
	}
	return decode(output, input)
}
