package eventline

import (
	"image"
	"time"
	"github.com/GUMI-golang/gumi"
)



// PreHandling
type (
	PreFocuser interface {
		PreFocus(event gumi.Event, focus bool) gumi.Event
	}
)


// Handling
type (
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
		Style(style *gumi.Style)
	}
)

