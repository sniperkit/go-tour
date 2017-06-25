package concurrent_test

import (
	"testing"

	"github.com/sahilm/go-tour/concurrent"
)

func TestPutIfAbsent(t *testing.T) {
	m := concurrent.NewMap()
	v, ok := m.PutIfAbsent("foo", "bar")

	if v != "bar" {
		t.Errorf("got %v, want: %v", v, "bar")
	}

	if !ok {
		t.Errorf("expected to be ok")
	}

	v, ok = m.PutIfAbsent("foo", 1)

	if v != "bar" {
		t.Errorf("got %v, want: %v", v, "bar")
	}

	if ok {
		t.Errorf("expected to be not ok")
	}
}

func TestPutAndGet(t *testing.T) {
	m := concurrent.NewMap()

	cases := []struct {
		k string
		v interface{}
	}{
		{"foo", "bar"},
		{"1", 2},
	}

	for _, c := range cases {
		m.Put(c.k, c.v)
	}

	for _, c := range cases {
		got, ok := m.Get(c.k)
		if got != c.v {
			t.Errorf("got %v, want %v", got, c.k)
		}
		if !ok {
			t.Errorf("inserted key: %v indicated as notpresent", c.k)
		}
	}

	_, ok := m.Get("missing-key")
	if ok {
		t.Errorf("missing key indicated as present")
	}
}
