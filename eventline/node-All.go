package eventline

import (
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"time"
)

type AllNode struct {
	BaseEventNode
	//
	shape   Shaper

	//
	focus          bool
	imeActive      bool
	imeCTRL        bool
	imeComplete    string
	imeEditing     rune
	imeDeleteActor *TextDeleteActor
	//
}
//
func (s *AllNode) Occur(event gumi.Event) {
	recurviePreOccurEvent(s, &event)
	if event == nil{
		return
	}
	recurvieOccurEvent(s, event)
}
func (s *AllNode) PreHandle(event gumi.Event) gumi.Event {
	switch e := event.(type) {
	case gumi.EventCursor:
		if v, ok := s.Handler.(PreFocuser); ok {
			return v.PreFocus(event, s.shape.In(e.ToPoint()))
		}
	}
	return event
}
func (s *AllNode) Handle(event gumi.Event) {
	switch e := event.(type) {
	case gumi.EventKeyPress:
		if s.focus {
			switch e.Key {
			case gumi.KEY_MOUSE1:
				if v, ok := s.Handler.(Clicker); ok {
					v.Click(true)
				}
			case gumi.KEY_CONTROL:
				if s.imeActive {
					s.imeCTRL = true
				}
			case gumi.KEY_BACKSPACE:
				if s.imeEditing == 0{
					s.imeDeleteActor.Start()
				}
			}
		}
	case gumi.EventKeyRelease:
		if s.focus {
			switch e.Key {
			case gumi.KEY_MOUSE1:
				if v, ok := s.Handler.(Clicker); ok {
					v.Click(false)
				}
			case gumi.KEY_CONTROL:
				if s.imeActive {
					s.imeCTRL = false
				}
			case gumi.KEY_BACKSPACE:
				s.imeDeleteActor.Reset()
			}
		}
	case gumi.EventScroll:
		if s.focus {
			if v, ok := s.Handler.(Scroller); ok {
				v.Scroll(e.X, e.Y)
			}
		}
	case gumi.EventCursor:
		curfocus := s.shape.In(e.ToPoint())
		if s.focus != curfocus {
			s.focus = curfocus
			if v, ok := s.Handler.(Focuser); ok {
				v.Focus(s.focus)
			}
		}
	case gumi.EventResize:
		if v, ok := s.Handler.(Resizer); ok {
			v.Resize(e.Bound)
		}
	case gumi.EventStyle:
		if v, ok := s.Handler.(Styler); ok {
			v.Style(e.Style)
		}
	case gumi.EventTick:
		if s.imeDeleteActor.Animate(float64(e.DeltaT / time.Millisecond)){
			var deleteCount = s.imeDeleteActor.Pop()
			if deleteCount > 0 {
				if s.imeCTRL {
					// ctrl + backspace
					temp := gcore.StringForwardControlBackSpace(s.imeComplete, deleteCount)
					if s.imeComplete != temp {
						s.imeComplete = temp
						if v, ok := s.Handler.(Editer); ok {
							v.Edit(textComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
						}
					}
				} else {
					// backspace
					temp := gcore.StringForwardBackSpace(s.imeComplete, deleteCount)
					if s.imeComplete != temp {
						s.imeComplete = temp
						if v, ok := s.Handler.(Editer); ok {
							v.Edit(textComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
						}
					}
				}
			}
		}
		if v, ok := s.Handler.(Ticker); ok {
			v.Tick(e.DeltaT)
		}
	case gumi.EventRuneEdit:
		if s.imeActive {
			s.imeEditing = e.Rune
			if v, ok := s.Handler.(Editer); ok {
				v.Edit(textComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
			}
		}
	case gumi.EventRuneComplete:
		if s.imeActive {
			s.imeComplete += string(e.Rune)
			s.imeEditing = 0
			if v, ok := s.Handler.(Editer); ok {
				v.Edit(textComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
			}
		}
	}
}
//
func (s *AllNode) TextActive(onoff bool) {
	if !onoff {
		s.imeCTRL = false
	}
	s.imeActive = onoff
}
func (s *AllNode) IsTextActive() bool {
	return s.imeActive
}
