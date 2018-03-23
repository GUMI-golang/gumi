package eventline

import (
	"image"
)

type (
	Shaper interface {
		In(point image.Point) bool
	}
	RectShape struct {
		Bound image.Rectangle
	}
	CircleShape struct {
		Center     image.Point
		Radius     int
	}
	ComplexShape struct {
		shapers []Shaper
	}
)

func NewComplexShape(shapers ... Shaper) *ComplexShape {
	return &ComplexShape {
		shapers:shapers,
	}
}
func NewRectShape(rect image.Rectangle) *RectShape {
	return &RectShape{
		Bound: rect,
	}
}
func NewCircleShape(center image.Point, radius int) *CircleShape {
	return &CircleShape{
		Center:     center,
		Radius:radius,
	}
}
func (s *ComplexShape) In(point image.Point) bool {
	for _, v := range s.shapers{
		if v.In(point){
			return true
		}
	}
	return false
}
func (s *ComplexShape) Clear() {
	s.shapers = nil
}
func (s *ComplexShape) Append(shapers ... Shaper) {
	s.shapers = append(s.shapers, shapers...)
}

func (s *RectShape) In(point image.Point) bool {
	return (s.Bound.Min.X < point.X && point.X < s.Bound.Max.X) && (s.Bound.Min.Y < point.Y && point.Y < s.Bound.Max.Y)
}

func (s *CircleShape) In(point image.Point) bool {
	diff := point.Sub(s.Center)
	diffRadQuad := diff.X * diff.X + diff.Y * diff.Y
	radQuad := s.Radius * s.Radius
	return diffRadQuad < radQuad
}



