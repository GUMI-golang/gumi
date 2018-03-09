package media

import (
	"image"
	"image/draw"
)

const (
	RepeatNormal     RepeatMode = iota
	RepeatHorizontal RepeatMode = iota
	RepeatVertical   RepeatMode = iota
	//RepeatRound      RepeatMode = iota
	//RepeatSpace      RepeatMode = iota
)

type RepeatMode uint8
type Repeat struct {
	src  *image.RGBA
	Mode RepeatMode
}

func NewRepeat(img image.Image, mode RepeatMode) *Repeat {
	var src *image.RGBA
	var ok bool
	if src, ok = img.(*image.RGBA); !ok {
		src = image.NewRGBA(img.Bounds())
		draw.Draw(src, src.Rect, img, image.ZP, draw.Src)
	}
	return &Repeat{
		src:  src,
		Mode: mode,
	}
}
func (s Repeat) Bound() image.Rectangle {
	return s.src.Bounds()
}
func (s Repeat) Draw(dst draw.Image) {
	var dstsz = dst.Bounds()
	var dstw, dsth = dstsz.Dx(), dstsz.Dy()
	var srcw, srch = s.src.Rect.Dx(), s.src.Rect.Dy()
	var rcnth, rcntv = dstw/srcw + 1, dsth/srch + 1
	switch s.Mode {
	default:
		fallthrough
	case RepeatNormal:
		for hidx := 0; hidx < rcnth; hidx++ {
			for vidx := 0; vidx < rcntv; vidx++ {
				draw.Draw(
					dst,
					image.Rect(
						dstsz.Min.X + hidx*srcw,
						dstsz.Min.Y + vidx*srch,
						dstsz.Min.X + (hidx + 1)*srcw,
						dstsz.Min.Y + (vidx + 1)*srch,
					),
					s.src,
					s.src.Rect.Min,
					draw.Over,
				)
			}
		}
	case RepeatHorizontal:
		for hidx := 0; hidx < rcnth; hidx++ {
			draw.Draw(
				dst,
				image.Rect(
					dstsz.Min.X + hidx*srcw,
					dstsz.Min.Y,
					dstsz.Min.X + (hidx + 1)*srcw,
					dstsz.Min.Y + srch,
				),
				s.src,
				s.src.Rect.Min,
				draw.Over,
			)
		}
	case RepeatVertical:
		for vidx := 0; vidx < rcntv; vidx++ {
			draw.Draw(
				dst,
				image.Rect(
					dstsz.Min.X,
					dstsz.Min.Y + vidx*srch,
					dstsz.Min.X + srcw,
					dstsz.Min.Y + (vidx + 1)*srch,
				),
				s.src,
				s.src.Rect.Min,
				draw.Over,
			)
		}
	}
}
