package gcore

import (
	"image/color"
	"math"
)

func Clamp(i float64, min, max float64) float64 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

// h [...1], s [...1], v [...1]
func HSVToRGB(h, s, v float64) color.RGBA {
	var i, f, p, q, t float64
	if s == 0 {
		outV := uint8(Clamp(v*255+0.5, 0, 255))
		return color.RGBA{outV, outV, outV, 0xFF}
	}

	i = math.Floor(h * 6)
	f = h*6 - i
	p = v * (1 - s)
	q = v * (1 - s*f)
	t = v * (1 - s*(1-f))

	var r, g, b float64
	switch i {
	case 0:
		r = v
		g = t
		b = p
	case 1:
		r = q
		g = v
		b = p
	case 2:
		r = p
		g = v
		b = t
	case 3:
		r = p
		g = q
		b = v
	case 4:
		r = t
		g = p
		b = v
	default:
		r = v
		g = p
		b = q
	}

	outR := uint8(Clamp(r*255+0.5, 0, 255))
	outG := uint8(Clamp(g*255+0.5, 0, 255))
	outB := uint8(Clamp(b*255+0.5, 0, 255))
	return color.RGBA{outR, outG, outB, 0xFF}
}
