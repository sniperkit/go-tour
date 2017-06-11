package sqrt_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/sahilm/go-tour/sqrt"
)

func TestSqrt(t *testing.T) {
	cases := []float64{1, 2, 3.3, 4, 1000, 1024}
	for _, c := range cases {
		got, iterations := sqrt.Sqrt(c)
		want := math.Sqrt(c)
		if got != want {
			t.Errorf("Sqrt(%f) == %f, want %f", c, got, want)
		} else {
			fmt.Printf("sqrt(%.3f) == %.14f in %d iterations\n", c, got, iterations)
		}
	}
}