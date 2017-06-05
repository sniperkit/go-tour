package exercise1

import "math"

func Sqrt(x float64) (float64, int) {
	const minDelta = 1e-8
	z := 1.0
	delta := 1.0
	iterations := 0
	for ; delta > minDelta; iterations++ {
		next := z - (z*z-x)/(2*z)
		delta = math.Abs(z - next)
		z = next
	}
	return z, iterations
}
