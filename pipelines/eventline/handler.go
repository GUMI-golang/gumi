package eventline

import (
	"image"
	"time"
)

type Handler interface {
	Handle()
}

type (
	Focuser interface {
		Focus(focus bool)
	}
	Clicker interface {
		Focuser
		Click(ispress bool)
	}
	Scroll interface {
		Focuser
		Scroll(x, y int)
	}
	Editer interface {
		Focuser
		Edit(text string)
	}
)

//
type Resizer interface {
	Resize(r image.Rectangle)
}

//
type Ticker interface {
	Tick(t time.Time)
}

//
//type Resizer interface {
//	Resize(r image.Rectangle)
//}
