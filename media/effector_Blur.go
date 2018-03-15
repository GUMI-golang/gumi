package media

import (
	"image"
	"image/draw"
	"math"
	"sync"
)

const (
	BlurBox      BlurMode = iota
	BlurGaussian BlurMode = iota
)

type BlurMode uint8

type Blur struct {
	Radius    float64
	KeepAlpha bool
	Mode      BlurMode
}

func NewBlur(Radius float64, KeepAlpha bool, Mode BlurMode) *Blur {

	return &Blur{
		Radius:    Radius,
		KeepAlpha: KeepAlpha,
		Mode:      Mode,
	}
}
func (s Blur) Draw(target draw.Image) {
	var pix = startEdit(target)
	defer endEdit(target, pix)
	switch s.Mode {
	default:
		fallthrough
	case BlurBox:
		if s.KeepAlpha {
			boxBlurKeepAlpha(pix, s.Radius)
		} else {
			boxBlur(pix, s.Radius)
		}
	case BlurGaussian:
		if s.KeepAlpha {
			boxGaussianKeepAlpha(pix, s.Radius)
		} else {
			boxGaussian(pix, s.Radius)
		}

	}
}

func boxBlur(src *image.RGBA, radius float64) {
	var length = int(math.Ceil(radius)) + 1
	var rad = length / 2
	var mat = NewMatrix(length, length)
	mat.Clear(1)
	mat = mat.Normal()
	var pad = paddingExtend(src, rad, rad)
	var wg = new(sync.WaitGroup)
	var srcw, srcH = sizeEidt(src)

	for y := 0; y < srcH; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < srcw; x++ {
				var r, g, b, a float64
				for ky := 0; ky < length; ky++ {
					iy := y + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := pad.PixOffset(ix, iy)
						r += float64(pad.Pix[offset+R]) * kval
						g += float64(pad.Pix[offset+G]) * kval
						b += float64(pad.Pix[offset+B]) * kval
						a += float64(pad.Pix[offset+A]) * kval
					}
				}
				offset := src.PixOffset(x, y)
				src.Pix[offset+R] = uint8(clamp(r, 0, 255))
				src.Pix[offset+G] = uint8(clamp(g, 0, 255))
				src.Pix[offset+B] = uint8(clamp(b, 0, 255))
				src.Pix[offset+A] = uint8(clamp(a, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func boxBlurKeepAlpha(src *image.RGBA, radius float64) {
	var length = int(math.Ceil(radius)) + 1
	var rad = length / 2
	var mat = NewMatrix(length, length)
	mat.Clear(1)
	mat = mat.Normal()
	var pad = paddingExtend(src, rad, rad)
	var wg = new(sync.WaitGroup)
	var srcw, srch = sizeEidt(src)

	for y := 0; y < srch; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < srcw; x++ {
				var r, g, b float64
				for ky := 0; ky < length; ky++ {
					iy := y + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := pad.PixOffset(ix, iy)
						r += float64(pad.Pix[offset+R]) * kval
						g += float64(pad.Pix[offset+G]) * kval
						b += float64(pad.Pix[offset+B]) * kval
					}
				}
				offset := src.PixOffset(x, y)
				src.Pix[offset+R] = uint8(clamp(r, 0, 255))
				src.Pix[offset+G] = uint8(clamp(g, 0, 255))
				src.Pix[offset+B] = uint8(clamp(b, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func gaussian(x, y, sigma float64) float64 {
	return math.Exp(-(x*x/sigma + y*y/sigma))
}
func boxGaussian(src *image.RGBA, radius float64) {
	var length = int(math.Ceil(radius)) + 1
	var rad = length / 2
	var mat = NewMatrix(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			mat[x][y] = gaussian(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}
	var pad = paddingExtend(src, rad, rad)
	var wg = new(sync.WaitGroup)
	var srcw, srch = sizeEidt(src)

	for y := 0; y < srch; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < srcw; x++ {
				var r, g, b, a float64
				for ky := 0; ky < length; ky++ {
					iy := y + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := pad.PixOffset(ix, iy)
						r += float64(pad.Pix[offset+R]) * kval
						g += float64(pad.Pix[offset+G]) * kval
						b += float64(pad.Pix[offset+B]) * kval
						a += float64(pad.Pix[offset+A]) * kval
					}
				}
				offset := src.PixOffset(x, y)
				src.Pix[offset+R] = uint8(clamp(r, 0, 255))
				src.Pix[offset+G] = uint8(clamp(g, 0, 255))
				src.Pix[offset+B] = uint8(clamp(b, 0, 255))
				src.Pix[offset+A] = uint8(clamp(a, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func boxGaussianKeepAlpha(src *image.RGBA, radius float64) {
	var length = int(math.Ceil(radius)) + 1
	var rad = length / 2
	var mat = NewMatrix(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			mat[x][y] = gaussian(float64(x)-radius, float64(y)-radius, 4*radius)
		}
	}
	var pad = paddingExtend(src, rad, rad)
	var wg = new(sync.WaitGroup)
	var srcw, srch = sizeEidt(src)

	for y := 0; y < srch; y++ {
		wg.Add(1)
		go func(y int) {
			for x := 0; x < srcw; x++ {
				var r, g, b float64
				for ky := 0; ky < length; ky++ {
					iy := y + ky
					for kx := 0; kx < length; kx++ {
						ix := x + kx
						kval := mat[kx][ky]
						offset := pad.PixOffset(ix, iy)
						r += float64(pad.Pix[offset+R]) * kval
						g += float64(pad.Pix[offset+G]) * kval
						b += float64(pad.Pix[offset+B]) * kval
					}
				}
				offset := src.PixOffset(x, y)
				src.Pix[offset+R] = uint8(clamp(r, 0, 255))
				src.Pix[offset+G] = uint8(clamp(g, 0, 255))
				src.Pix[offset+B] = uint8(clamp(b, 0, 255))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
