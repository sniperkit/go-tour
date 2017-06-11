package slices

func Pic(dx, dy int) [][]uint8 {
	pic := [][]uint8{}

	for y := 0; y < dy; y++ {
		xs := []uint8{}
		for x := 0; x < dx; x++ {
			xs = append(xs, uint8(x*y))
		}
		pic = append(pic, xs)
	}
	return pic
}
