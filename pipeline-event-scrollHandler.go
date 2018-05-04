package gumi


type scrollHandler struct {
	pipe *Pipe
	shape Shaper
	//
	focus          bool
	//
}

func (s *scrollHandler) PreHandle(event Event) Event {
	if v, ok := s.pipe.Elem.(PreRawHandler); ok {
		return v.PreHandle(event)
	}
	switch e := event.(type) {
	case EventCursor:
		if v, ok := s.pipe.Elem.(PreFocuser); ok {
			return v.PreFocus(event, s.shape.In(e.ToPoint()))
		}
	}
	return event
}
func (s *scrollHandler) Handle(event Event) {
	switch e := event.(type) {
	case EventKeyPress:
		if s.focus {
			switch e.Key {
			case KEY_MOUSE1:
				if v, ok := s.pipe.Elem.(Clicker); ok {
					v.Click(true)
				}
			}
		}
	case EventKeyRelease:
		if s.focus {
			switch e.Key {
			case KEY_MOUSE1:
				if v, ok := s.pipe.Elem.(Clicker); ok {
					v.Click(false)
				}
			}
		}
	case EventScroll:
		if s.focus {
			if v, ok := s.pipe.Elem.(Scroller); ok {
				v.Scroll(e.X, e.Y)
			}
		}
	case EventCursor:
		curfocus := s.shape.In(e.ToPoint())
		if s.focus != curfocus {
			s.focus = curfocus
			if v, ok := s.pipe.Elem.(Focuser); ok {
				v.Focus(s.focus)
			}
		}
	}
	if v, ok := s.pipe.Elem.(RawHandler); ok {
		v.Handle(event)
	}
}