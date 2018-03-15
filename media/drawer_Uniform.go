package media

import (
	"image"
	"image/color"
	"image/draw"
)

type Uniform struct {
	c   color.Color
	uni *image.Uniform
}

func NewUniform(c color.Color) *Uniform {
	return &Uniform{c: c}
}

func (Uniform) Bound() image.Rectangle {
	return image.Rect(0, 0, 0, 0)
}
func (s Uniform) Draw(dst draw.Image) {
	//draw.Draw()
	bd := dst.Bounds()
	switch v := dst.(type) {
	case *image.RGBA:
		sr, sg, sb, sa := s.c.RGBA()
		for x := bd.Min.X; x < bd.Max.X; x++ {
			for y := bd.Min.Y; y < bd.Max.Y; y++ {
				off := v.PixOffset(x, y)
				dr := &v.Pix[off+R]
				dg := &v.Pix[off+G]
				db := &v.Pix[off+B]
				da := &v.Pix[off+A]
				//
				a := (m - sa) * 0x101
				//
				*dr = uint8((uint32(*dr)*a/m + sr) >> 8)
				*dg = uint8((uint32(*dg)*a/m + sg) >> 8)
				*db = uint8((uint32(*db)*a/m + sb) >> 8)
				*da = uint8((uint32(*da)*a/m + sa) >> 8)
			}
		}
	default:
		draw.Draw(dst, dst.Bounds(), image.NewUniform(s.c), image.ZP, draw.Over)
	}

}
func RGBA32ToRGBA8(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}
