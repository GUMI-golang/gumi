package gumi

type cfHandler struct {
	pipe *Pipe
	shape Shaper
	//
	focus          bool
	//
}

func (s *cfHandler) PreHandle(event Event) Event {
	switch e := event.(type) {
	case EventCursor:
		if v, ok := s.pipe.Elem.(PreFocuser); ok {
			return v.PreFocus(event, s.shape.In(e.ToPoint()))
		}
	}
	return event
}
func (s *cfHandler) Handle(event Event) {
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
	case EventCursor:
		curfocus := s.shape.In(e.ToPoint())
		if s.focus != curfocus {
			s.focus = curfocus
			if v, ok := s.pipe.Elem.(Focuser); ok {
				v.Focus(s.focus)
			}
		}
	}
}
