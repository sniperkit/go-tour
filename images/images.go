package images

import (
	"image"
	"image/color"
)

// Represents an image with Width and Height
type Image struct {
	Width, Height int
}

// Always returns color.RGBAModel
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Returns the rectangular bounds of Image
// (0,0,i.Width,i.Height)
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// Returns color.RGBA{R: 0x91, G: 0xd9, B: 0xd2, A: 0xff}
// for all values of (x,y)
func (Image) At(x, y int) color.Color {
	return color.RGBA{R: 0x91, G: 0xd9, B: 0xd2, A: 0xff}
}
