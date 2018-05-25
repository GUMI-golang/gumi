package gumi

import (
	"image"
	"github.com/GUMI-golang/gumi/gcore"
)

type Pipe struct {
	Pipeline *Pipeline
	Parent   *Pipe
	Childrun []*Pipe
	//
	Elem  GUMI
	style Style
}

func (s *Pipe) ProximateParentBound() image.Rectangle {
	if s.Parent == nil {
		w, h := s.Pipeline.renderer.Size()
		return image.Rect(0,0,w,h)
	}
	if r, ok := s.Parent.Elem.(Bounder); ok {
		return r.RelayBound()
	}
	return s.Parent.ProximateParentBound()
}
func (s *Pipe) ProximateChildrunSize() []*gcore.Size {
	if s.Childrun == nil {
		return nil
	}
	var res = make([]*gcore.Size, len(s.Childrun))
	for i, v := range s.Childrun {
		if szr, ok := v.Elem.(Sizer); ok {
			temp := szr.Size()
			res[i] = &temp
		}
	}
	return res
}

func (s *Pipe) GetStyle(data StyleData) interface{} {
	if v, ok := s.style[data]; ok {
		return v
	}
	if s.Parent == nil {
		return nil
	}
	return s.Parent.GetStyle(data)
}
func (s *Pipe) SetStyle(data StyleData, value interface{}) bool {
	if s.style == nil{
		s.style = make(Style)
	}
	if data.Type().Valid(value) {
		s.style[data] = value
		return true
	}
	return false
}

func (s *Pipe) ListStyle(each func(sdata StyleData, value interface{})) {
	for k, v := range s.style {
		each(k, v)
	}
}
