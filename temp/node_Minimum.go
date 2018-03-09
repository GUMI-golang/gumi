package temp

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)
type NMinimum struct {
	SingleNode
	axis gcore.Axis
}

func (s *NMinimum) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NMinimum) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NMinimum) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
}
func (s *NMinimum) GUMIRender(frame *image.RGBA) {
}
func (s *NMinimum) GUMISize() gcore.Size {
	sz := s.child.GUMISize()
	if gcore.AxisVertical == gcore.AxisVertical & s.axis{
		sz.Vertical.Max = sz.Vertical.Min
	}
	if gcore.AxisHorizontal == gcore.AxisHorizontal & s.axis{
		sz.Horizontal.Max = sz.Horizontal.Min
	}
	return sz
}


func (s *NMinimum) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
}
func (s *NMinimum) GUMIDraw() {
	s.child.GUMIDraw()
}
func (s *NMinimum) GUMIUpdate() {
	panic("implement me")
}

func (s *NMinimum) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func (s *NMinimum) String() string {
	return fmt.Sprintf("%s", "NMinimum")
}

func NMinimum0(axis gcore.Axis, elem GUMI) *NMinimum {
	temp := &NMinimum{
		axis:axis,
	}
	elem.born(temp)
	temp.breed([]GUMI{elem})
	return temp
}

func (s *NMinimum) Get() gcore.Axis {
	return s.GetAxis()
}
func (s *NMinimum) Set(axis gcore.Axis) {
	s.Set(axis)
}
func (s *NMinimum) GetAxis() gcore.Axis {
	return s.axis
}
func (s *NMinimum) SetAxis(axis gcore.Axis) {
	s.axis = axis
}