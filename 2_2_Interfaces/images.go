package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width  int
	height int
}

// ColorModel() color.Model
// Bounds() Rectangle
// At(x, y int) color.Color
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	image_func := func(x, y int) uint8 {
		return uint8(x * y)
	}
	value := image_func(x, y)
	return color.RGBA{value, value, 255, 255}
}

func main() {
	m := Image{300, 200}
	pic.ShowImage(m)
}
