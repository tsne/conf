package conf

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	buf := bytes.NewBufferString(`{ "foo":1, "bar":"abc" }`)

	c, err := Load(buf, json.Unmarshal)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(c) != 2 {
		t.Fatalf("invalid number of configuration values: %d", len(c))
	}

	foo, err := c.Value("foo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	i, err := foo.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 1 {
		t.Fatalf("unexpected configuration value: %d", i)
	}

	bar, err := c.Value("bar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	s, err := bar.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "abc" {
		t.Fatalf("unexpected configuration value: %s", s)
	}

	// invalid configuration
	buf = bytes.NewBufferString(`{ "foo"=1 }`)

	_, err = Load(buf, json.Unmarshal)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	// erroneous reader
	readErr := errors.New("read error")
	_, err = Load(&erroneousReader{err: readErr}, json.Unmarshal)
	if err != readErr {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConfigMustLoad(t *testing.T) {
	buf := bytes.NewBufferString(`{ "foo":1, "bar":"abc" }`)
	p, b := panicked(func() {
		MustLoad(buf, json.Unmarshal)
	})
	if b {
		t.Fatalf("unexpected panic: %v", p)
	}

	readErr := errors.New("read error")
	p, b = panicked(func() {
		MustLoad(&erroneousReader{err: readErr}, json.Unmarshal)
	})
	if !b {
		t.Fatalf("expected panic, got none")
	}
	if p != readErr {
		t.Fatalf("unexpected panic error: %v", p)
	}
}

func TestConfigValue(t *testing.T) {
	c := Config{
		"foo": map[string]interface{}{
			"one": 1,
			"two": 2,
			"others": map[interface{}]interface{}{
				"3": "three",
				"4": "four",
			},
		},
	}

	v, err := c.Value("foo.one")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	i, err := v.Int()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 1 {
		t.Fatalf("unexpected value: %v", i)
	}

	v, err = c.Value("foo.others.3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	s, err := v.String()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s != "three" {
		t.Fatalf("unexpected value: '%v'", s)
	}

	_, err = c.Value("bar")
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	_, err = c.Value("foo.three")
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	_, err = c.Value("foo.others.10")
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestConfigDecode(t *testing.T) {
	c := Config{
		"foo": map[string]interface{}{
			"one": 1,
			"two": "a",
		},
	}

	var i int
	err := c.Decode("foo.one", &i)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i != 1 {
		t.Fatalf("unexpected value: %v", i)
	}

	err = c.Decode("foo.three", &i)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	err = c.Decode("foo.two", &i)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func panicked(f func()) (interface{}, bool) {
	var (
		p interface{}
		b bool
	)
	func() {
		defer func() {
			p = recover()
			b = p != nil
		}()
		f()
	}()

	return p, b
}

type erroneousReader struct {
	err error
}

func (r *erroneousReader) Read(p []byte) (int, error) {
	return 0, r.err
}
