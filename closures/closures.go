package closures

func Fibonacci() func() int {
	x, y := -1, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}
