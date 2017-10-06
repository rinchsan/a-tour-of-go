package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{
	x, y, w, h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(i.x, i.y, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	ux, uy := uint8(x), uint8(y)
	r := ux + uy
	g := ux^uy
	b := ux * uy
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{0, 0, 100, 100}
	pic.ShowImage(m)
}
