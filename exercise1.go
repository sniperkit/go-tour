package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	z := 1.0
	iterate := true
	i := 0
	for ; iterate; i++ {
		next := z - (z*z-x)/(2*z)
		if math.Abs(z-next) < 1e-7 {
			iterate = false
		} else {
			z = next
		}
	}
	return z, i
}

func compareAndPrintDelta(x float64) {
	z, i := sqrt(x)
	mathSqrt := math.Sqrt(x)
	fmt.Printf(
		"Newton's method (iterations = %d): %f, "+
			"Sqrt: %f, "+
			"Delta: %f\n",
		i, z, mathSqrt, math.Abs(mathSqrt-z))
}

func main() {
	compareAndPrintDelta(1)
	compareAndPrintDelta(2)
	compareAndPrintDelta(3)
	compareAndPrintDelta(4)
	compareAndPrintDelta(5)
	compareAndPrintDelta(1000)
	compareAndPrintDelta(1234566)
}
