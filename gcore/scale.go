package gcore

import (
	"image/color"
)

type scale float64

func Scale(f float64) scale {
	return scale(Clamp(f, 0, 1))
}
func (s scale) Color(c1, c2 color.Color) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	var h2 = float64(s)
	var h1 = 1 - h2
	var res color.RGBA64
	res.R = uint16(float64(r1)*h1 + float64(r2)*h2)
	res.G = uint16(float64(g1)*h1 + float64(g2)*h2)
	res.B = uint16(float64(b1)*h1 + float64(b2)*h2)
	res.A = uint16(float64(a1)*h1 + float64(a2)*h2)
	return res

}
func (s scale) Length(length float64) float64 {
	return length * float64(s)
}
