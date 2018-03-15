package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

type NStyle struct {
	SingleNode
	s *Style
}

func (s *NStyle) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NStyle) GUMIStyle(style *Style) {
	s.child.GUMIStyle(s.s)
}
func (s *NStyle) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

func (s *NStyle) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.child.GUMIRenderSetup(man, parent)
}
func (s *NStyle) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func (s *NStyle) String() string {
	return fmt.Sprintf("%s", "NStyle")
}

func NStyle0(s *Style) *NStyle {
	if s == nil {
		s = DefaultStyle()
	}
	return &NStyle{
		s: s,
	}
}
func (s *NStyle) Set(st *Style) {
	s.s = st
}
func (s *NStyle) Get() *Style {
	return s.s
}
