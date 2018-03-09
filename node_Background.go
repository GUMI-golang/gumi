package gumi

import (
	"fmt"
	"image"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
	"github.com/GUMI-golang/gumi/media"
)

type NBackground struct {
	SingleNode
	rendererStore
	//
	drawer media.Drawer
}

// renderline.Job / BaseRender
func (s *NBackground) BaseRender(subimg *image.RGBA) {
	s.drawer.Draw(subimg)
}

// renderline.Job / DecalRender
func (s *NBackground) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

func (s *NBackground) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NBackground) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}

func (s *NBackground) GUMIRender(frame *image.RGBA) {
	s.drawer.Draw(frame)
}
func (s NBackground) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

func (s *NBackground) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent)
	s.rnode.Do = s
	s.child.GUMIRenderSetup(s.rmana, s.rnode)
}
func (s *NBackground) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func (s *NBackground) String() string {
	return fmt.Sprintf("%s", "NBackground")
}

func NBackground0(draw media.Drawer) *NBackground {
	return &NBackground{
		drawer:draw,
	}
}

func (s *NBackground) Set(dw media.Drawer) {
	s.SetDrawer(dw)
}
func (s *NBackground) Get() media.Drawer {
	return s.GetDrawer()
}

func (s *NBackground) SetDrawer(dw media.Drawer) {
	s.drawer = dw
	s.rnode.ThrowCache()
}

func (s *NBackground) GetDrawer() media.Drawer {
	return s.drawer
}