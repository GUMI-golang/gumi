package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

type (
	NClicker struct {
		SingleNode
		//
		rnode       renderline.Node
		onClick     NClickerClick
		cursorEnter bool
	}
	NClickerClick func(self *NClicker)
)

func (s *NClicker) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NClicker) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}

func (s NClicker) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

func (s *NClicker) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.child.GUMIRenderSetup(man, parent)
	s.rnode = parent
}
func (s *NClicker) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.onClick != nil {
				s.onClick(s)
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		bd := s.rnode.GetAllocation()
		if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
			s.cursorEnter = true
		} else {
			s.cursorEnter = false
		}
	}
	s.child.GUMIHappen(event)
}
func (s *NClicker) String() string {
	return fmt.Sprintf("%s", "NClicker")
}

func NClicker0(onClick NClickerClick) *NClicker {
	return &NClicker{
		onClick: onClick,
	}
}
func (s *NClicker) OnClick(onClick NClickerClick) {
	s.onClick = onClick
}

func (s *NClicker) ReferClick() NClickerClick {
	return s.onClick
}
