package gumi

import (
	"image"
	"time"
)

const (
	EVENT_KEYPRESS     EventKind = iota
	EVENT_KEYRELEASE   EventKind = iota
	EVENT_CURSOR       EventKind = iota
	EVENT_SCROLL       EventKind = iota
	EVENT_RUNECOMPLETE EventKind = iota
	EVENT_RUNEEDIT     EventKind = iota
	EVENT_TICK         EventKind = iota
	EVENT_RESIZE       EventKind = iota
	EVENT_STYLE        EventKind = iota
)

type (
	EventKind uint8
	Event     interface {
		Kind() EventKind
	}
)

type (
	EventCursor struct {
		image.Point
	}
	EventScroll struct {
		image.Point
	}
	EventKeyPress struct {
		Key GUMIKey
	}
	EventKeyRelease struct {
		Key GUMIKey
	}
	EventRuneComplete struct {
		Rune rune
	}
	EventRuneEdit struct {
		Rune rune
	}
	EventTick struct {
		DeltaT time.Duration
	}
	EventResize struct {
		Bound image.Rectangle
	}
	EventStyle struct {
		Style *Style
	}
)

func (EventCursor) Kind() EventKind {
	return EVENT_CURSOR
}
func (EventScroll) Kind() EventKind {
	return EVENT_SCROLL
}
func (EventKeyPress) Kind() EventKind {
	return EVENT_KEYPRESS
}
func (EventKeyRelease) Kind() EventKind {
	return EVENT_KEYRELEASE
}
func (EventRuneComplete) Kind() EventKind {
	return EVENT_RUNECOMPLETE
}
func (EventRuneEdit) Kind() EventKind {
	return EVENT_RUNEEDIT
}
func (EventTick) Kind() EventKind {
	return EVENT_TICK
}
func (EventResize) Kind() EventKind {
	return EVENT_RESIZE
}
func (EventStyle) Kind() EventKind {
	return EVENT_STYLE
}

// Utils
func (s EventCursor) ToPoint() image.Point {
	return image.Pt(s.X, s.Y)
}