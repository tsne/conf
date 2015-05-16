package conf

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	maxUint = uint64(^uint(0))
	maxInt  = int64(maxUint >> 1)
	minInt  = -maxInt - 1
)

func decode(output, input reflect.Value) error {
	if input.Kind() == reflect.Interface && !input.IsNil() {
		input = input.Elem()
	}

	switch output.Kind() {
	case reflect.Bool:
		return decodeBool(output, input)

	case reflect.Int:
		return decodeInt(output, input, minInt, maxInt)

	case reflect.Int8:
		return decodeInt(output, input, math.MinInt8, math.MaxInt8)

	case reflect.Int16:
		return decodeInt(output, input, math.MinInt16, math.MaxInt16)

	case reflect.Int32:
		return decodeInt(output, input, math.MinInt32, math.MaxInt32)

	case reflect.Int64:
		outputType := output.Type()
		if outputType.PkgPath() == "time" && outputType.Name() == "Duration" {
			return decodeDuration(output, input)
		}
		return decodeInt(output, input, math.MinInt64, math.MaxInt64)

	case reflect.Uint:
		return decodeUint(output, input, maxUint)

	case reflect.Uint8:
		return decodeUint(output, input, math.MaxUint8)

	case reflect.Uint16:
		return decodeUint(output, input, math.MaxUint16)

	case reflect.Uint32:
		return decodeUint(output, input, math.MaxUint32)

	case reflect.Uint64:
		return decodeUint(output, input, math.MaxUint64)

	case reflect.Float32:
		return decodeFloat(output, input, math.MaxFloat32)

	case reflect.Float64:
		return decodeFloat(output, input, math.MaxFloat64)

	case reflect.String:
		return decodeString(output, input)

	case reflect.Array:
		return decodeArray(output, input)

	case reflect.Slice:
		return decodeSlice(output, input)

	case reflect.Map:
		return decodeMap(output, input)

	case reflect.Interface:
		return decodeInterface(output, input)

	case reflect.Struct:
		return decodeStruct(output, input)

	case reflect.Ptr:
		return decodePtr(output, input)

	default:
		return fmt.Errorf("type '%s' is not supported", output.Kind())
	}
}

func decodeBool(output, input reflect.Value) error {
	input = convertNumericString(input)

	switch input.Kind() {
	case reflect.Bool:
		output.SetBool(input.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		output.SetBool(input.Int() != 0)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		output.SetBool(input.Uint() != 0)

	case reflect.Float32, reflect.Float64:
		output.SetBool(input.Float() != 0)

	default:
		return fmt.Errorf("'%s' could not be converted to '%s'", input.Type(), output.Type())
	}

	return nil
}

func decodeInt(output, input reflect.Value, min, max int64) error {
	input = convertNumericString(input)

	switch input.Kind() {
	case reflect.Bool:
		if input.Bool() {
			output.SetInt(1)
		} else {
			output.SetInt(0)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := input.Int()
		if i < min || i > max {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetInt(i)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i := input.Uint()
		if i > uint64(max) {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetInt(int64(i))

	case reflect.Float32, reflect.Float64:
		f := input.Float()
		if f < float64(min) || f > float64(max) {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetInt(int64(f))

	default:
		return fmt.Errorf("'%s' could not be converted to '%s'", input.Type(), output.Type())
	}

	return nil
}

func decodeUint(output, input reflect.Value, max uint64) error {
	input = convertNumericString(input)

	switch input.Kind() {
	case reflect.Bool:
		if input.Bool() {
			output.SetUint(1)
		} else {
			output.SetUint(0)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := input.Int()
		if i < 0 || uint64(i) > max {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetUint(uint64(i))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i := input.Uint()
		if i > max {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetUint(i)

	case reflect.Float32, reflect.Float64:
		f := input.Float()
		if f < 0 || f > float64(max) {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetUint(uint64(f))

	default:
		return fmt.Errorf("'%s' could not be converted to '%s'", input.Type(), output.Type())
	}

	return nil
}

func decodeDuration(output, input reflect.Value) error {
	input = convertNumericString(input)

	switch input.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		output.SetInt(input.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i := input.Uint()
		if i > math.MaxInt64 {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetInt(int64(i))

	case reflect.Float32, reflect.Float64:
		f := input.Float()
		if f < math.MinInt64 || f > math.MaxInt64 {
			return fmt.Errorf("value out of range ('%s' expected)", output.Type())
		}
		output.SetInt(int64(f))

	case reflect.String:
		d, err := time.ParseDuration(input.String())
		if err != nil {
			return fmt.Errorf("'%s' is not a valid duration", input.String())
		}
		output.SetInt(int64(d))

	default:
		return fmt.Errorf("'%s' could not be converted to 'time.Duration'", input.Type())
	}

	return nil
}

func decodeFloat(output, input reflect.Value, max float64) error {
	input = convertNumericString(input)

	switch input.Kind() {
	case reflect.Bool:
		if input.Bool() {
			output.SetFloat(1)
		} else {
			output.SetFloat(0)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i := input.Int()
		if float64(i) < -max || float64(i) > max {
			return fmt.Errorf("value out of range ('%s' expected)", input.Type())
		}
		output.SetFloat(float64(i))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i := input.Uint()
		if float64(i) > max {
			return fmt.Errorf("value out of range ('%s' expected)", input.Type())
		}
		output.SetFloat(float64(i))

	case reflect.Float32, reflect.Float64:
		f := input.Float()
		if f < -max || f > max {
			return fmt.Errorf("value out of range ('%s' expected)", input.Type())
		}
		output.SetFloat(f)

	default:
		return fmt.Errorf("'%s' could not be converted to '%s'", input.Type(), output.Type())
	}

	return nil
}

func decodeString(output, input reflect.Value) error {
	switch input.Kind() {
	case reflect.Bool:
		output.SetString(strconv.FormatBool(input.Bool()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		output.SetString(strconv.FormatInt(input.Int(), 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		output.SetString(strconv.FormatUint(input.Uint(), 10))

	case reflect.Float32, reflect.Float64:
		output.SetString(strconv.FormatFloat(input.Float(), 'f', -1, 64))

	case reflect.String:
		output.SetString(input.String())

	default:
		output.SetString(fmt.Sprintf("%v", input.Interface()))
	}

	return nil
}

func decodeArray(output, input reflect.Value) error {
	switch input.Kind() {
	case reflect.Array, reflect.Slice:
		n := input.Len()
		if n != output.Len() {
			return fmt.Errorf("'[%d]%s' could not be converted to '[%d]%s'", input.Len(), input.Type().Elem(), output.Len(), output.Type().Elem())
		}

		for i := 0; i < n; i++ {
			if err := decode(output.Index(i), input.Index(i)); err != nil {
				return err
			}
		}

	default:
		if output.Len() != 1 {
			return fmt.Errorf("'[1]%s' could not be converted to '[%d]%s'", input.Type(), output.Len(), output.Type().Elem())
		}
		return decode(output.Index(0), input)
	}

	return nil
}

func decodeSlice(output, input reflect.Value) error {
	switch input.Kind() {
	case reflect.Array, reflect.Slice:
		n := input.Len()
		sliceVal := reflect.MakeSlice(reflect.SliceOf(output.Type().Elem()), n, n)
		for i := 0; i < n; i++ {
			if err := decode(sliceVal.Index(i), input.Index(i)); err != nil {
				return err
			}
		}
		output.Set(sliceVal)

	default:
		sliceVal := reflect.MakeSlice(reflect.SliceOf(output.Type().Elem()), 1, 1)
		if err := decode(sliceVal.Index(0), input); err != nil {
			return err
		}
		output.Set(sliceVal)
	}

	return nil
}

func decodeMap(output, input reflect.Value) error {
	if input.Kind() != reflect.Map {
		return fmt.Errorf("'%s' could not be converted to 'map'", input.Type())
	}

	outputType := output.Type()
	mapType := reflect.MapOf(outputType.Key(), outputType.Elem())
	mapVal := reflect.MakeMap(mapType)

	for _, key := range input.MapKeys() {
		k := reflect.Indirect(reflect.New(mapType.Key()))
		if err := decode(k, key); err != nil {
			return err
		}

		v := reflect.Indirect(reflect.New(mapType.Elem()))
		if err := decode(v, input.MapIndex(key)); err != nil {
			return err
		}

		mapVal.SetMapIndex(k, v)
	}

	output.Set(mapVal)
	return nil
}

func decodeInterface(output, input reflect.Value) error {
	if !output.CanSet() {
		return fmt.Errorf("'%s' cannot be set", output.Type())
	}

	output.Set(input)
	return nil
}

func decodeStruct(output, input reflect.Value) error {
	if input.Kind() != reflect.Map {
		return fmt.Errorf("'%s' could not be converted to 'map'", input.Type())
	}

	stringType := reflect.TypeOf("")
	keyType := input.Type().Key()
	if !stringType.AssignableTo(keyType) {
		return fmt.Errorf("map[%s]' could not be converted to 'map[string]'", keyType)
	}

	for _, field := range fieldsOf(output) {
		if field.ignore {
			continue
		}

		key := reflect.ValueOf(field.mapkey())
		val := input.MapIndex(key)
		if !val.IsValid() && len(field.key) == 0 {
			// map key not found and no key specified (search case-insensitive)
			for _, k := range input.MapKeys() {
				if s, ok := k.Interface().(string); ok {
					if strings.EqualFold(s, field.name) {
						key = k
						val = input.MapIndex(k)
						break
					}
				}
			}
		}

		if !val.IsValid() {
			// map key not found
			if field.required {
				return fmt.Errorf("required field '%s' not found", field.name)
			}
			continue
		}

		if err := decode(field.value, val); err != nil {
			return fmt.Errorf("[struct field '%s'] %s", field.name, err)
		}
	}

	return nil
}

func decodePtr(output, input reflect.Value) error {
	if input.Kind() == reflect.Ptr {
		input = input.Elem()
	}
	return decode(output.Elem(), input)
}

func convertNumericString(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.String {
		s := v.String()
		if b, err := strconv.ParseBool(s); err == nil {
			return reflect.ValueOf(b)
		}
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return reflect.ValueOf(i)
		}
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			return reflect.ValueOf(f)
		}
	}
	return v
}

type field struct {
	name     string
	value    reflect.Value
	key      string
	required bool
	ignore   bool
}

func fieldsOf(v reflect.Value) []*field {
	t := v.Type()
	n := v.NumField()
	fields := make([]*field, 0, n)
	for i := 0; i < n; i++ {
		structField := t.Field(i)
		f := &field{
			name:  structField.Name,
			value: v.Field(i),
		}

		tag := structField.Tag.Get("config")
		switch tag {
		case "":
			// no 'config' tag
		case "-":
			f.ignore = true
		default:
			tagParts := strings.Split(tag, ",")
			f.key = tagParts[0]
			if len(tagParts) > 1 {
				switch tagParts[1] {
				case "required":
					f.required = true
				default:
					panic(fmt.Sprintf("'%s.%s' contains an invalid 'config' tag (%s)", t, f.name, tag))
				}
			}
		}

		fields = append(fields, f)
	}
	return fields
}

func (f *field) mapkey() string {
	if len(f.key) == 0 {
		return f.name
	}
	return f.key
}
