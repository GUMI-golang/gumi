package media

import (
	"image"
	"image/color"
	"image/draw"
)

type Uniform struct{
	c color.Color
	uni *image.Uniform
}

func NewUniform(c color.Color) *Uniform {
	return &Uniform{c:c,}
}

func (Uniform) Bound() image.Rectangle {
	return image.Rect(0,0,0,0)
}
func (s Uniform) Draw(dst draw.Image) {
	bd := dst.Bounds()
	switch v := dst.(type) {
	case *image.RGBA:
		r, g, b, a := RGBA32ToRGBA8(s.c.RGBA())
		for x := bd.Min.X; x < bd.Max.X ; x ++{
			for y := bd.Min.Y; y < bd.Max.Y ; y ++{
				off := v.PixOffset(x, y)
				v.Pix[off + R] = r
				v.Pix[off + G] = g
				v.Pix[off + B] = b
				v.Pix[off + A] = a
			}
		}
	default:
		for x := bd.Min.X; x < bd.Max.X ; x ++{
			for y := bd.Min.Y; y < bd.Max.Y ; y ++{
				dst.Set(x, y, s.c)
			}
		}
	}

}
func RGBA32ToRGBA8(r,g,b,a uint32) (uint8,uint8,uint8,uint8) {
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}