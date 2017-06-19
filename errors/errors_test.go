package errors_test

import (
	"math"
	"testing"

	"github.com/sahilm/go-tour/errors"
)

func TestSqrtErrors(t *testing.T) {
	cases := []struct {
		in           float64
		want         float64
		errorMessage string
	}{
		{-2, 0, "cannot Sqrt negative number: -2"},
		{-100, 0, "cannot Sqrt negative number: -100"},
		{2, math.Sqrt(2), ""},
	}

	for _, c := range cases {
		got, _, err := errors.Sqrt(c.in)

		if got != c.want {
			t.Errorf("Got: %v, but want: %v", got, c.want)
		}

		switch err {
		case nil:
			if c.errorMessage != "" {
				t.Errorf("Expected error message: %v, but got no error.", c.errorMessage)
			}

		default:
			gotErrorMessage := err.Error()
			if gotErrorMessage != c.errorMessage {
				t.Errorf("Expected error message: %v, but got error: %v", c.errorMessage, gotErrorMessage)
			}
		}
	}
}
