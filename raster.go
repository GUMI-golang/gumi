package gumi

import (
	"image"
	"github.com/fogleman/gg"
)

func createContext(rgba *image.RGBA) (*gg.Context) {
	temp := &image.RGBA{
		Rect:image.Rect(0,0, rgba.Rect.Dx(), rgba.Rect.Dy()),
		Pix:rgba.Pix,
		Stride:rgba.Stride,
	}
	ctx := gg.NewContextForRGBA(temp)
	return ctx
}