package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"time"
)

type (
	NEventer struct {
		SingleNode
		//
		rnode       renderline.Node
		onClick     NEventerClick
		onFocus     NEventerFocus
		onInterval  NEventerInterval
		cursorEnter bool
		interval    time.Duration
		t           int64
	}
	NEventerOption struct {
		OnClick    NEventerClick
		OnFocus    NEventerFocus
		Interval   time.Duration
		OnInterval NEventerInterval
	}
	NEventerClick    func(self *NEventer)
	NEventerFocus    func(self *NEventer, focus bool)
	NEventerInterval func(self *NEventer, t time.Time)
)

func (s *NEventer) GUMIInfomation(info Information) {
	if s.interval != 0 && s.onInterval != nil {
		s.t += info.Dt
		temp := int64(s.interval.Seconds() * 1000)
		if s.t > temp {
			n := s.t / temp
			s.t = s.t % temp
			for i := int64(0); i < n; i++ {

				s.onInterval(s, time.Now())
			}
		}
	}
	s.child.GUMIInfomation(info)
}
func (s *NEventer) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s NEventer) GUMISize() gcore.Size {
	return s.child.GUMISize()
}
func (s *NEventer) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.child.GUMIRenderSetup(man, parent)
	s.rnode = parent
}
func (s *NEventer) GUMIHappen(event Event) {
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
			if s.onFocus != nil && s.cursorEnter == false {
				s.onFocus(s, true)
			}
			s.cursorEnter = true
		} else {
			if s.onFocus != nil && s.cursorEnter == true {
				s.onFocus(s, false)
			}
			s.cursorEnter = false
		}
	}
	s.child.GUMIHappen(event)
}
func (s *NEventer) String() string {
	return fmt.Sprintf("%s", "NClicker")
}

func NEventer0() *NEventer {
	return &NEventer{}
}
func NEventer1(onClick NEventerClick) *NEventer {
	return &NEventer{
		onClick: onClick,
	}
}
func NEventer2(onClick NEventerClick, onFocus NEventerFocus) *NEventer {
	return &NEventer{
		onClick: onClick,
		onFocus: onFocus,
	}
}
func NEventer3(onClick NEventerClick, onFocus NEventerFocus, interval time.Duration, onInterval NEventerInterval) *NEventer {
	return &NEventer{
		onClick:    onClick,
		onFocus:    onFocus,
		interval:   interval,
		onInterval: onInterval,
	}
}
func NEventer4(opt NEventerOption) *NEventer {
	return &NEventer{
		onClick:    opt.OnClick,
		onFocus:    opt.OnFocus,
		interval:   opt.Interval,
		onInterval: opt.OnInterval,
	}
}

func (s *NEventer) GetInterval() time.Duration {
	return s.interval
}
func (s *NEventer) SetInterval(interval time.Duration) {
	s.interval = interval
}

//
func (s *NEventer) OnClick(onClick NEventerClick) {
	s.onClick = onClick
}
func (s *NEventer) ReferClick() NEventerClick {
	return s.onClick
}

func (s *NEventer) OnFocus(onFocus NEventerFocus) {
	s.onFocus = onFocus
}
func (s *NEventer) ReferFocus() NEventerFocus {
	return s.onFocus
}

func (s *NEventer) OnInterval(onInterval NEventerInterval) {
	s.onInterval = onInterval
}
func (s *NEventer) ReferInterval() NEventerInterval {
	return s.onInterval
}

func (s *NEventer) OnSetInterval(interval time.Duration, onInterval NEventerInterval) {
	s.SetInterval(interval)
	s.OnInterval(onInterval)
}
