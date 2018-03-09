package gcore

import (
	"github.com/iamGreedy/freetype"
	"github.com/iamGreedy/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"sync"
)

type Font struct {
	font *truetype.Font
	ctx  *freetype.Context
	face font.Face
	sz   float64
	mtx  *sync.Mutex
}

func (s *Font) Use() {
	s.mtx.Lock()
}
func (s *Font) CalculateSize(text string) (h, v int) {
	temp := font.Drawer{Face: s.face}
	bd, _ := temp.BoundString(text)
	return (bd.Max.X - bd.Min.X).Ceil(), (bd.Max.Y - bd.Min.Y).Ceil()
}
func (s *Font) Draw(rect image.Rectangle, dst draw.Image, text string, point fixed.Point26_6) {
	s.ctx.SetDst(dst)
	s.ctx.SetClip(rect)
	s.ctx.DrawString(text, point)
}
func (s *Font) ChangeSize(size float64) {
	s.ctx.SetFontSize(size)
	s.face = truetype.NewFace(s.font, &truetype.Options{
		Size:    size,
		Hinting: font.HintingFull,
		DPI:     72,
	})
	s.sz = size
}
func (s *Font) FontHeight() fixed.Int26_6 {
	return s.ctx.PointToFixed(s.sz)
}
func (s *Font) FontSize() float64 {
	return s.sz
}
func (s *Font) ChangeSource(image2 image.Image) {
	s.ctx.SetSrc(image2)
}
func (s *Font) Face() font.Face {
	return s.face
}
func (s *Font) Release() {
	s.mtx.Unlock()
}
func NewFont(font *truetype.Font, size float64) *Font {
	temp := &Font{
		font: font,
		ctx:  freetype.NewContext(),
		mtx:  new(sync.Mutex),
	}
	temp.ctx.SetFont(temp.font)
	temp.ChangeSize(size)
	return temp
}
