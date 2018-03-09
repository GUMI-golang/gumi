package media

import (
	"image"
	"image/draw"
)

type Fixed struct {
	src *image.RGBA
}

func NewFixed(img image.Image) *Fixed {
	var src *image.RGBA
	var ok bool
	if src, ok = img.(*image.RGBA); !ok {
		src = image.NewRGBA(img.Bounds())
		draw.Draw(src, src.Rect, img, image.ZP, draw.Src)
	}
	return &Fixed{
		src: src,
	}
}

func (s Fixed) Bound() image.Rectangle {
	return s.src.Bounds()
}
func (s Fixed) Draw(dst draw.Image) {
	draw.Draw(dst, dst.Bounds(), s.src, s.src.Rect.Min, draw.Over)
}
