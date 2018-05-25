package gumi

import (
	"time"
	"image"
)

type Handler interface {
	PreHandle(event Event) Event
	Handle(event Event)

}


// PreHandling
type (
	PreRawHandler interface {
		PreHandle(event Event) Event
	}
	PreFocuser interface {
		PreFocus(event Event, focus bool) Event
	}
)


// Handling
type (
	RawHandler interface {
		Handle(event Event)
	}
	Focuser interface {
		Focus(focus bool)
	}
	Clicker interface {
		Click(ispress bool)
	}
	Scroller interface {
		Scroll(x, y int)
	}
	Editer interface {
		Edit(current string, completed string, editing rune)
	}
	Ticker interface {
		Tick(t time.Duration)
	}
	Resizer interface {
		Resize(r image.Rectangle)
	}
	Styler interface {
		Style(data StyleData)
	}
)
