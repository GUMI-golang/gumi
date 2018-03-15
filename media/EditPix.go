package media

import (
	"image"
	"image/draw"
)

const (
	R           = 0
	G           = 1
	B           = 2
	A           = 3
	STRIDECOUNT = 4
)

const m = 1<<16 - 1

func startEdit(src image.Image) (rgba *image.RGBA) {
	if v, ok := src.(*image.RGBA); ok {
		return v
	}
	rgba = image.NewRGBA(src.Bounds())
	draw.Draw(rgba, rgba.Rect, src, src.Bounds().Min, draw.Src)
	return rgba
}
func endEdit(dst draw.Image, rgba *image.RGBA) {
	if _, ok := dst.(*image.RGBA); ok {
		return
	}
	draw.Draw(dst, dst.Bounds(), rgba, rgba.Rect.Min, draw.Src)
}

//

//
func sizeEidt(rgba *image.RGBA) (w, h int) {
	sz := rgba.Rect.Size()
	return sz.X, sz.Y
}

//
func paddingEmpty(src *image.RGBA, padW, padH int) (dst *image.RGBA) {
	var srcw, srch = sizeEidt(src)
	var dstW, dstH = srcw + padW*2, srch + padH*2
	dst = image.NewRGBA(image.Rect(0, 0, dstW, dstH))
	draw.Draw(dst, image.Rect(padW, padH, padW+srcw, padH+srch), src, image.ZP, draw.Src)
	return dst
}
func paddingExtend(src *image.RGBA, padW, padH int) (dst *image.RGBA) {
	dst = paddingEmpty(src, padW, padH)
	dstW, dstH := sizeEidt(dst)
	for y := 0; y < dstH; y++ {
		iy := y
		if iy < padH {
			iy = padH
		} else if iy >= dstH-padH {
			iy = dstH - padH - 1
		}
		for x := 0; x < dstW; x++ {
			ix := x
			if ix < padW {
				ix = padW
			} else if x >= dstW-padW {
				ix = dstW - padW - 1
			} else if iy == y {
				// This only enters if we are not in a y-padded area or
				// x-padded area, so nothing to extend here.
				// So simply jump to the next padded-x index.
				x = dstW - padW - 1
				continue
			}
			dstOffset := dst.PixOffset(x, y)
			edgeOffset := dst.PixOffset(ix, iy)
			dst.Pix[dstOffset+R] = dst.Pix[edgeOffset+R]
			dst.Pix[dstOffset+G] = dst.Pix[edgeOffset+G]
			dst.Pix[dstOffset+B] = dst.Pix[edgeOffset+B]
			dst.Pix[dstOffset+A] = dst.Pix[edgeOffset+A]
		}
	}
	return dst
}
