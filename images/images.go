package images

import (
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{R: 0x91, G: 0xd9, B: 0xd2, A: 0xff}
}
