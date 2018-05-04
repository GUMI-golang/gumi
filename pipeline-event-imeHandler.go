package gumi

type imeHandler struct {
	pipe *Pipe
	//focus          bool
	imeActive      bool
	imeCTRL        bool
	imeComplete    string
	imeEditing     rune
	imeDeleteActor *TextDeleteActor
	//
}

func (s *imeHandler) PreHandle(event Event) Event {
	//if v, ok := s.pipe.Elem.(PreRawHandler); ok {
	//	return v.PreHandle(event)
	//}
	//switch e := event.(type) {
	//case EventCursor:
	//	if v, ok := s.pipe.Elem.(PreFocuser); ok {
	//		return v.PreFocus(event, s.shape.In(e.ToPoint()))
	//	}
	//}
	return event
}
func (s *imeHandler) Handle(event Event) {
	switch e := event.(type) {
	//case EventKeyPress:
	//	if s.focus {
	//		switch e.Key {
	//		case KEY_MOUSE1:
	//			if v, ok := s.pipe.Elem.(Clicker); ok {
	//				v.Click(true)
	//			}
	//		case KEY_CONTROL:
	//			if s.imeActive {
	//				s.imeCTRL = true
	//			}
	//		case KEY_BACKSPACE:
	//			if s.imeEditing == 0 {
	//				s.imeDeleteActor.Start()
	//			}
	//		}
	//	}
	//case EventKeyRelease:
	//	if s.focus {
	//		switch e.Key {
	//		case KEY_MOUSE1:
	//			if v, ok := s.pipe.Elem.(Clicker); ok {
	//				v.Click(false)
	//			}
	//		case KEY_CONTROL:
	//			if s.imeActive {
	//				s.imeCTRL = false
	//			}
	//		case KEY_BACKSPACE:
	//			s.imeDeleteActor.Reset()
	//		}
	//	}
	//case EventScroll:
	//	if s.focus {
	//		if v, ok := s.pipe.Elem.(Scroller); ok {
	//			v.Scroll(e.X, e.Y)
	//		}
	//	}
	//case EventCursor:
	//	curfocus := s.shape.In(e.ToPoint())
	//	if s.focus != curfocus {
	//		s.focus = curfocus
	//		if v, ok := s.pipe.Elem.(Focuser); ok {
	//			v.Focus(s.focus)
	//		}
	//	}
	//case EventResize:
	//	if v, ok := s.pipe.Elem.(Resizer); ok {
	//		v.Resize(e.Bound)
	//	}
	//case EventStyle:
	//	if v, ok := s.pipe.Elem.(Styler); ok {
	//		v.Style(e.Style)
	//	}
	//case EventTick:
	//	if s.imeDeleteActor.Animate(float64(e.DeltaT / time.Millisecond)) {
	//		var deleteCount = s.imeDeleteActor.Pop()
	//		if deleteCount > 0 {
	//			if s.imeCTRL {
	//				// ctrl + backspace
	//				temp := gcore.StringForwardControlBackSpace(s.imeComplete, deleteCount)
	//				if s.imeComplete != temp {
	//					s.imeComplete = temp
	//					if v, ok := s.pipe.Elem.(Editer); ok {
	//						v.Edit(TextComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
	//					}
	//				}
	//			} else {
	//				// backspace
	//				temp := gcore.StringForwardBackSpace(s.imeComplete, deleteCount)
	//				if s.imeComplete != temp {
	//					s.imeComplete = temp
	//					if v, ok := s.pipe.Elem.(Editer); ok {
	//						v.Edit(TextComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
	//					}
	//				}
	//			}
	//		}
	//	}
	//	if v, ok := s.pipe.Elem.(Ticker); ok {
	//		v.Tick(e.DeltaT)
	//	}
	case EventRuneEdit:
		if s.imeActive {
			s.imeEditing = e.Rune
			if v, ok := s.pipe.Elem.(Editer); ok {
				v.Edit(TextComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
			}
		}
	case EventRuneComplete:
		if s.imeActive {
			s.imeComplete += string(e.Rune)
			s.imeEditing = 0
			if v, ok := s.pipe.Elem.(Editer); ok {
				v.Edit(TextComplete(s.imeComplete, s.imeEditing), s.imeComplete, s.imeEditing)
			}
		}
	}
	if v, ok := s.pipe.Elem.(RawHandler); ok {
		v.Handle(event)
	}
}

//
func (s *imeHandler) TextActive(onoff bool) {
	if !onoff {
		s.imeCTRL = false
	}
	s.imeActive = onoff
}
func (s *imeHandler) IsTextActive() bool {
	return s.imeActive
}
