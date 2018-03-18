package eventline

import (
	"image"
	"golang.org/x/image/math/fixed"
)

type (
	Shaper interface {
		In(point image.Point) bool
		AABB() image.Rectangle
	}
	RectShape struct {
		bound fixed.Rectangle26_6
	}
	CircleShape struct {
		center fixed.Point26_6
		radius fixed.Int26_6
	}
)

func (s *RectShape) In(point image.Point) bool {
	pt := fixed.P(point.X, point.Y)
	return (s.bound.Min.X < pt.X && pt.X < s.bound.Max.X) && (s.bound.Min.Y < pt.Y && pt.Y < s.bound.Max.Y)
}

func (s *RectShape) AABB() image.Rectangle {
	return image.Rect(
		s.bound.Min.X.Round(),
		s.bound.Min.Y.Round(),
		s.bound.Max.X.Round(),
		s.bound.Max.Y.Round(),
	)
}

func (s *CircleShape) In(point image.Point) bool {
	panic("implement me")
}

func (s *CircleShape) AABB() image.Rectangle {
	r := s.radius.Round()
	return image.Rect(
		s.center.X.Round() - r,
		s.center.Y.Round() - r,
		s.center.X.Round() + r,
		s.center.Y.Round() + r,
	)
}



