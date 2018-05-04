package gumi

type styleHandler struct {
	pipe *Pipe
	shape Shaper
	//
	focus          bool
	imeActive      bool
	imeCTRL        bool
	imeComplete    string
	imeEditing     rune
	imeDeleteActor *TextDeleteActor
	//
}

func (s *styleHandler) PreHandle(event Event) Event {
	return event
}
func (s *styleHandler) Handle(event Event) {
	switch e := event.(type) {
	case EventStyle:
		if v, ok := s.pipe.Elem.(Styler); ok {
			v.Style(e.StyleData)
		}
	}
}
