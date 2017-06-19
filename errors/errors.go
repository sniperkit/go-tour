package errors

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, int, error) {
	if x < 0 {
		return 0, 0, ErrNegativeSqrt(x)
	}
	const minDelta = 1e-8
	z := 1.0
	delta := 1.0
	iterations := 0
	for ; delta > minDelta; iterations++ {
		next := z - (z*z-x)/(2*z)
		delta = math.Abs(z - next)
		z = next
	}
	return z, iterations, nil
}
