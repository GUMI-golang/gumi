package gumi

import "github.com/GUMI-golang/gumi/gcore"

type Meta struct {
	scr         *Screen
	Title       string
	Description string
	size        gcore.FixedSize
	Framerate   gcore.Framerate
	Locale      string
}
//
func (s *Meta ) Size() gcore.FixedSize {
	return s.size
}
func (s *Meta ) SetSize(size gcore.FixedSize)  {
	s.size = size
	s.commit()
}
func (s *Meta) commit() {
	if s.scr != nil{
		// Todo
		s.scr.Pipeline.renderer.Setup(s.size.Width, s.size.Height)
	}
}