package media

import (
	"image/draw"
	"math"
	"math/rand"
	"sync"
)

type Noise struct {
	Intense float64
}

func NewNoise(intense float64) *Noise {
	return &Noise{
		Intense: intense,
	}
}

func (s Noise) Draw(target draw.Image) {
	var dst = startEdit(target)
	defer endEdit(target, dst)
	var wg = new(sync.WaitGroup)
	//
	for y := dst.Rect.Min.Y; y < dst.Rect.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			for x := dst.Rect.Min.X; x < dst.Rect.Max.X; x++ {
				offset := dst.PixOffset(x, y)
				delta := noiseGausian(s.Intense)
				dst.Pix[offset+R] = uint8(clamp(delta+float64(dst.Pix[offset+R]), 0, math.MaxUint8))
				dst.Pix[offset+G] = uint8(clamp(delta+float64(dst.Pix[offset+G]), 0, math.MaxUint8))
				dst.Pix[offset+B] = uint8(clamp(delta+float64(dst.Pix[offset+B]), 0, math.MaxUint8))
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func noiseGausian(intense float64) float64 {
	return rand.NormFloat64() * intense
}
