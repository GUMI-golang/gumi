package gumi

import "image/color"

var Scale _Scale

type _Scale struct {
}

func (_Scale) Color(c1, c2 color.Color, zeroToOne float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	var h2 = zeroToOne
	var h1 = 1 - h2
	var res color.RGBA64
	res.R = uint16(float64(r1)*h1 + float64(r2)*h2)
	res.G = uint16(float64(g1)*h1 + float64(g2)*h2)
	res.B = uint16(float64(b1)*h1 + float64(b2)*h2)
	res.A = uint16(float64(a1)*h1 + float64(a2)*h2)
	return res

}
func (_Scale) Length(length float64, zeroToOne float64) float64 {
	return length * zeroToOne
}
