package media

import (
	"image"
	"image/draw"
	"math"
	"sync"
)

const (
	FillupNearestNeighbor FillupMode = iota
	FillupNearest         FillupMode = iota
	FillupGausian         FillupMode = iota
)

type FillupMode uint8
type Fillup struct {
	src  *image.RGBA
	Mode FillupMode
	//

}

func NewFillup(img image.Image, mode FillupMode) *Fillup {
	var src *image.RGBA
	var ok bool
	if src, ok = img.(*image.RGBA); !ok {
		src = image.NewRGBA(img.Bounds())
		draw.Draw(src, src.Rect, img, image.ZP, draw.Src)
	}
	return &Fillup{
		src:  src,
		Mode: mode,
	}
}
func (s Fillup) Bound() image.Rectangle {
	return s.src.Bounds()
}
func (s Fillup) Draw(dst draw.Image) {
	var pix = startEdit(dst)
	defer endEdit(dst, pix)
	//
	switch s.Mode {
	case FillupNearestNeighbor:
		resizeNearestNeighbor(pix, s.src)
	case FillupNearest:
		resize(pix, s.src, 1., nearest)
	default:
		fallthrough
	case FillupGausian:
		resize(pix, s.src, 1., gausian)
		//resizeminGoroutine(pix, stride, s.src.Pix, s.src.Stride, 1., gausian)
	}
}
func gausian(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		exp := 2.0
		x *= 2.0
		y := math.Pow(0.5, math.Pow(x, exp))
		base := math.Pow(0.5, math.Pow(2, exp))
		return (y - base) / (1 - base)
	}
	return 0
}
func nearest(x float64) float64 {
	x = math.Abs(x)
	if x < 1.0 {
		return 1.0 - x
	}
	return 0
}
func filterRange(v, length int, delta, radius float64) (val float64, start, end int) {
	val = float64(v) * delta
	if val < 0 {
		val = 0
	} else if val >= float64(length) {
		val = float64(length)
	}
	start, end = int(val-radius+0.5), int(val+radius)
	if start < 0 {
		start = 0
	}
	if end >= length {
		end = length
	}
	return
}
func resize(dst *image.RGBA, src *image.RGBA, support float64, fn func(x float64) float64) {
	var dstw, dsth = sizeEidt(dst)
	var srcw, srch = sizeEidt(src)
	var deltaH = float64(srcw) / float64(dstw)
	var scaleH = math.Max(deltaH, 1.0)
	var radiusH = math.Ceil(scaleH * support)
	var deltaV = float64(srch) / float64(dsth)
	var scaleV = math.Max(deltaV, 1.0)
	var radiusV = math.Ceil(scaleV * support)
	var wg = new(sync.WaitGroup)
	//
	var memorizationW = make([]struct {
		val        float64
		start, end int
	}, dstw)
	for x := 0; x < dst.Rect.Dx(); x++ {
		xsrc, hstart, hend := filterRange(x, srcw, deltaH, radiusH)
		memorizationW[x] = struct {
			val        float64
			start, end int
		}{
			val:   xsrc + float64(src.Rect.Min.X),
			start: hstart + src.Rect.Min.X,
			end:   hend + src.Rect.Min.X,
		}
	}
	var memorizationH = make([]struct {
		val        float64
		start, end int
	}, dstw)
	for y := 0; y < dst.Rect.Dy(); y++ {
		ysrc, vstart, vend := filterRange(y, srch, deltaV, radiusV)
		memorizationH[y] = struct {
			val        float64
			start, end int
		}{
			val:   ysrc + float64(src.Rect.Min.Y),
			start: vstart + src.Rect.Min.Y,
			end:   vend + src.Rect.Min.Y,
		}
	}
	//
	for x := dst.Rect.Min.X; x < dst.Rect.Max.X; x++ {
		wg.Add(1)
		go func(x int) {
			for y := dst.Rect.Min.Y; y < dst.Rect.Max.Y; y++ {
				var r, g, b, a float64
				var sum float64
				//
				h, v := memorizationW[x-dst.Rect.Min.X], memorizationH[y-dst.Rect.Min.Y]
				// H : pixel evaluate
				for kx := h.start; kx < h.end; kx++ {
					srcoffset := src.PixOffset(kx, int(v.val))
					normal := (float64(kx) - h.val) / scaleH
					res := fn(normal)
					// normalized r, g, b, a and sum
					r += float64(src.Pix[srcoffset+R]) * res
					g += float64(src.Pix[srcoffset+G]) * res
					b += float64(src.Pix[srcoffset+B]) * res
					a += float64(src.Pix[srcoffset+A]) * res
					sum += res
				}
				// V : pixel evaluate
				for ky := v.start; ky < v.end; ky++ {
					srcoffset := src.PixOffset(int(h.val), ky)
					normal := (float64(ky) - v.val) / scaleV
					res := fn(normal)
					// normalized r, g, b, a and sum
					r += float64(src.Pix[srcoffset+R]) * res
					g += float64(src.Pix[srcoffset+G]) * res
					b += float64(src.Pix[srcoffset+B]) * res
					a += float64(src.Pix[srcoffset+A]) * res
					sum += res
				}
				// pixel set
				dstoffset := dst.PixOffset(x, y)
				sr := uint32(clamp((r/sum)+0.5, 0, 255)) * 0x101
				sg := uint32(clamp((g/sum)+0.5, 0, 255)) * 0x101
				sb := uint32(clamp((b/sum)+0.5, 0, 255)) * 0x101
				sa := uint32(clamp((a/sum)+0.5, 0, 255)) * 0x101
				tempa := (math.MaxUint16 - sa) * 0x101
				// R, G, B, A
				dst.Pix[dstoffset+R] = uint8((uint32(dst.Pix[dstoffset+R])*tempa/math.MaxUint16 + sr) >> 8)
				dst.Pix[dstoffset+G] = uint8((uint32(dst.Pix[dstoffset+G])*tempa/math.MaxUint16 + sg) >> 8)
				dst.Pix[dstoffset+B] = uint8((uint32(dst.Pix[dstoffset+B])*tempa/math.MaxUint16 + sb) >> 8)
				dst.Pix[dstoffset+A] = uint8((uint32(dst.Pix[dstoffset+A])*tempa/math.MaxUint16 + sa) >> 8)
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
}
func resizeNearestNeighbor(dst *image.RGBA, src *image.RGBA) {
	var dstw, dsth = sizeEidt(dst)
	var srcw, srch = sizeEidt(src)
	var dx = float32(srcw) / float32(dstw)
	var dy = float32(srch) / float32(dsth)
	var wg = new(sync.WaitGroup)

	for y := dst.Rect.Min.Y; y < dst.Rect.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			for x := dst.Rect.Min.X; x < dst.Rect.Max.X; x++ {
				dstIdx := dst.PixOffset(x, y)
				srcIdx := src.PixOffset(
					int((float32(x-dst.Rect.Min.X)+0.5)*dx)+src.Rect.Min.X,
					int((float32(y-dst.Rect.Min.Y)+0.5)*dy)+src.Rect.Min.Y,
				)
				sr := uint32(src.Pix[srcIdx+R]) * 0x101
				sg := uint32(src.Pix[srcIdx+G]) * 0x101
				sb := uint32(src.Pix[srcIdx+B]) * 0x101
				sa := uint32(src.Pix[srcIdx+A]) * 0x101
				a := (math.MaxUint16 - sa) * 0x101
				// R, G, B, A
				dst.Pix[dstIdx+R] = uint8((uint32(dst.Pix[dstIdx+R])*a/math.MaxUint16 + sr) >> 8)
				dst.Pix[dstIdx+G] = uint8((uint32(dst.Pix[dstIdx+G])*a/math.MaxUint16 + sg) >> 8)
				dst.Pix[dstIdx+B] = uint8((uint32(dst.Pix[dstIdx+B])*a/math.MaxUint16 + sb) >> 8)
				dst.Pix[dstIdx+A] = uint8((uint32(dst.Pix[dstIdx+A])*a/math.MaxUint16 + sa) >> 8)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
