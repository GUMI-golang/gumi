package renderline

import (
	"image"
	"sync"
	"image/draw"
)

const m = 1<<16 - 1

var wgpool = &sync.Pool{
	New: func() interface{} {
		return new(sync.WaitGroup)
	},
}
func clip(dst image.Rectangle, r *image.Rectangle, src image.Rectangle, sp *image.Point) {
	orig := r.Min
	*r = r.Intersect(dst)
	*r = r.Intersect(src.Add(orig.Sub(*sp)))
	dx := r.Min.X - orig.X
	dy := r.Min.Y - orig.Y
	if dx == 0 && dy == 0 {
		return
	}
	sp.X += dx
	sp.Y += dy
}
func ParallelDraw(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point, method draw.Op){
	switch method {
	case draw.Src:
		ParallelDrawSrc(dst, r, src, sp)
	case draw.Over:
		ParallelDrawOver(dst, r, src, sp)
	}
}
func ParallelDrawSrc(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point) {
	wg := wgpool.Get().(*sync.WaitGroup)
	defer wgpool.Put(wg)
	//
	n, dy := 4*r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var ddelta, sdelta int
	//
	clip(dst.Rect, &r, src.Rect, &sp)
	//
	if r.Min.Y <= sp.Y {
		ddelta = dst.Stride
		sdelta = src.Stride
	} else {
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
	}
	//
	wg.Add(dy)
	for ; dy > 0; dy-- {
		go func(d, s int) {
			copy(dst.Pix[d:d+n], src.Pix[s:s+n])
			wg.Done()
		}(d0, s0)
		d0 += ddelta
		s0 += sdelta
	}
	wg.Wait()
}
func ParallelDrawOver(dst *image.RGBA, r image.Rectangle, src *image.RGBA, sp image.Point) {
	wg := wgpool.Get().(*sync.WaitGroup)
	defer wgpool.Put(wg)
	//
	dx, dy := r.Dx(), r.Dy()
	d0 := dst.PixOffset(r.Min.X, r.Min.Y)
	s0 := src.PixOffset(sp.X, sp.Y)
	var (
		ddelta, sdelta int
		i0, i1, idelta int
	)
	if r.Min.Y < sp.Y || r.Min.Y == sp.Y && r.Min.X <= sp.X {
		ddelta = dst.Stride
		sdelta = src.Stride
		i0, i1, idelta = 0, dx*4, +4
	} else {
		// If the source start point is higher than the destination start point, or equal height but to the left,
		// then we compose the rows in right-to-left, bottom-up order instead of left-to-right, top-down.
		d0 += (dy - 1) * dst.Stride
		s0 += (dy - 1) * src.Stride
		ddelta = -dst.Stride
		sdelta = -src.Stride
		i0, i1, idelta = (dx-1)*4, -4, -4
	}


	wg.Add(dy)
	for ; dy > 0; dy-- {
		go func(d, s int) {
			dpix := dst.Pix[d:]
			spix := src.Pix[s:]
			for i := i0; i != i1; i += idelta {
				sr := uint32(spix[i+0]) * 0x101
				sg := uint32(spix[i+1]) * 0x101
				sb := uint32(spix[i+2]) * 0x101
				sa := uint32(spix[i+3]) * 0x101

				dr := &dpix[i+0]
				dg := &dpix[i+1]
				db := &dpix[i+2]
				da := &dpix[i+3]

				// The 0x101 is here for the same reason as in drawRGBA.
				a := (m - sa) * 0x101

				*dr = uint8((uint32(*dr)*a/m + sr) >> 8)
				*dg = uint8((uint32(*dg)*a/m + sg) >> 8)
				*db = uint8((uint32(*db)*a/m + sb) >> 8)
				*da = uint8((uint32(*da)*a/m + sa) >> 8)
			}
			wg.Done()
		}(d0, s0)
		d0 += ddelta
		s0 += sdelta
	}
	wg.Wait()
}

