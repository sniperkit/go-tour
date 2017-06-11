package slices_test

import (
	"testing"

	"github.com/sahilm/go-tour/slices"
)

func TestPic(t *testing.T) {
	got := slices.Pic(4, 4)
	actualLength := len(got)
	t.Log(got)
	expectedLength := 4
	if actualLength != expectedLength {
		t.Errorf("Expected len(pic(4,4)): "+
			"%d, got: %d",
			expectedLength,
			actualLength)
	}
}
