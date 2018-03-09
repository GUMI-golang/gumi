package gumi

import (
	"fmt"

	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

type NSize struct {
	SingleNode
	sz gcore.Size
}

func (s *NSize) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NSize) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NSize) GUMISize() gcore.Size {
	temp := s.sz
	c := s.child.GUMISize()
	if temp.Vertical == gcore.AUTOLENGTH {
		temp.Vertical = c.Vertical
	} else if temp.Vertical == gcore.MINLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Max = c.Vertical.Min
	} else if temp.Vertical == gcore.MAXLENGTH {
		temp.Vertical = c.Vertical
		temp.Vertical.Min = c.Vertical.Max
	}
	if temp.Horizontal == gcore.AUTOLENGTH {
		temp.Horizontal = c.Horizontal
	} else if temp.Horizontal == gcore.MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Max = c.Horizontal.Min
	} else if temp.Horizontal == gcore.MINLENGTH {
		temp.Horizontal = c.Horizontal
		temp.Horizontal.Min = c.Horizontal.Max
	}
	return temp
}

func (s *NSize) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
	s.child.GUMIRenderSetup(man,parent)
}

func (s *NSize) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NSize) String() string {
	return fmt.Sprintf("%s(GUMISize:%v)", "NSize", s.sz)
}

func NSize0(sz gcore.Size) *NSize {
	return &NSize{
		sz: sz,
	}
}

func (s *NSize) Get() gcore.Size {
	return s.GetSize()
}
func (s *NSize) Set(sz gcore.Size) {
	s.SetSize(sz)
}
func (s *NSize) GetSize() gcore.Size {
	return s.sz
}
func (s *NSize) SetSize(sz gcore.Size) {
	s.sz = sz
}
