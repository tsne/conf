package conf

import (
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
)

// Unmarshaler represents a function which is able to parse the configuration
// data and store the parsed values into the given value.
type Unmarshaler func(data []byte, value interface{}) error

// Config represents a container for all parsed configuration values.
type Config map[string]interface{}

// Load loads the configuration from the source r using unmarshal to parse the
// input stream. The returned Config object holds all parsed values.
func Load(r io.Reader, unmarshal Unmarshaler) (Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	c := Config{}
	if err = unmarshal(data, &c); err != nil {
		return nil, err
	}
	return c, nil
}

// MustLoad ensures the loading of the configuration from the r using
// unmarshal to parse the input stream. This function calls Load and panics
// on error.
func MustLoad(r io.Reader, unmarshal Unmarshaler) Config {
	c, err := Load(r, unmarshal)
	if err != nil {
		panic(err)
	}
	return c
}

// Value returns the configuration value with the given key. If there is
// a hierarchy of configuration values a dot can be used to separate the
// different levels (e.g. "foo.bar" gets the value of the key 'bar' which lies
// under the key 'foo').
func (c Config) Value(key string) (*Value, error) {
	val := c.value(key)
	if !val.IsValid() {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	v := Value(val)
	return &v, nil
}

// Decode stores the configuration value with the given key in value.
// If value has an invalid type an error is returned. Decode tries to
// convert the configuration value to value's type (e.g. the string "7"
// is converted to the integer 7). The given value has to be a pointer,
// otherwise an error is returned.
//
// When decoding a struct each field could provide a 'config' tag. If the
// tag is "-" the field will be ignored. Otherwise it consists of a custom
// configuration key, followed by an optional comma and options. If the
// custom configuration key is empty the field name will be used.
// Currently only "required" is a valid option. If the "required" option
// will be set and the key does not exist an error is returned.
//
// Tag examples:
//   // Field will be ignored
//   Field string `config:"-"`
//
//   // Field is required, the key is "Field"
//   Field string `config:",required"`
//
//   // Field is optional, the key is "foobar"
//   Field string `config:"foobar"`
//
//   // Field is required, the key is "foobar"
//   Field string `config:"foobar,required"`
func (c Config) Decode(key string, value interface{}) error {
	val, err := c.Value(key)
	if err != nil {
		return err
	}
	if err = val.decode(value); err != nil {
		return fmt.Errorf("cannot decode key '%s': %v", key, err)
	}
	return nil
}

func (c Config) value(configKey string) reflect.Value {
	val := reflect.ValueOf(c)
	keys := strings.Split(configKey, ".")
	nKeys := len(keys)
	for i := 0; i < nKeys-1; i++ {
		key := reflect.ValueOf(keys[i])
		if !key.Type().AssignableTo(val.Type().Key()) {
			return reflect.Value{}
		}

		val = val.MapIndex(key)
		if val.Kind() == reflect.Interface && !val.IsNil() {
			val = val.Elem()
		}
		if !val.IsValid() || val.Kind() != reflect.Map {
			return reflect.Value{}
		}
	}

	key := reflect.ValueOf(keys[nKeys-1])
	return val.MapIndex(key)
}
