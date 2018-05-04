package gumi

import (
	"image"
	"github.com/GUMI-golang/gorat"
)

type (
	Bounder interface {
		RelayBound() image.Rectangle
		setBound(rectangle image.Rectangle)
		GetBound() image.Rectangle
		RequestResize()
		resizeDone()
		needResize() bool
	}
	DoRender interface {
		DoRender(rst gorat.SubRasterizer)
		// Cache related
		Throw()
		Done()
		Valid() bool
		Bounder
	}
	PostRender interface {
		PostRender(rasterzier gorat.SubRasterizer)
	}
)

type ParentDoRender struct {
	ParentBounder
	caching     bool

}

// do not override
func (s *ParentDoRender) Done() {
	s.caching = true
}
// do not override
func (s *ParentDoRender) Throw() {
	s.caching = false
}
// do not override
func (s *ParentDoRender) Valid() bool {
	return s.caching
}

type ParentBounder struct {
	reverseresz bool
	rect        image.Rectangle
}
// if need, override
func (s *ParentBounder) RelayBound() image.Rectangle {
	return s.rect
}
// do not override
func (s *ParentBounder) GetBound() image.Rectangle {
	return s.rect
}
// do not override
func (s *ParentBounder) setBound(rectangle image.Rectangle) {
	s.rect = rectangle
}
// do not override
func (s *ParentBounder) RequestResize() {
	s.reverseresz = false
}
// do not override
func (s *ParentBounder) resizeDone() {
	s.reverseresz = true
}
// do not override
func (s *ParentBounder) needResize() bool {
	return !s.reverseresz
}
