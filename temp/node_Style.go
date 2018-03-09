package temp

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

type NStyle struct {
	SingleNode
	s *Style
}

func (s *NStyle) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NStyle) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NStyle) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NStyle) GUMIRender(frame *image.RGBA) {
}
func (s *NStyle) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

func (s *NStyle) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
}
func (s *NStyle) GUMIUpdate() {
	panic("implement me")
}
func (s *NStyle) GUMIDraw(frame *image.RGBA) {
	s.GUMIDraw(frame)
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
