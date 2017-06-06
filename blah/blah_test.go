package blah_test

import (
	"testing"

	"github.com/sahilm/go-tour/blah"
)

func TestAdd(t *testing.T) {
	want := 2
	got := blah.Add(1, 1)
	if got != want {
		t.Errorf("failed")
	}
}
