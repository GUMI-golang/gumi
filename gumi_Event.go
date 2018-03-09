package gumi

const (
	EVENT_KEYPRESS     EventKind = iota
	EVENT_KEYRELEASE   EventKind = iota
	EVENT_CURSOR       EventKind = iota
	EVENT_SCROLL       EventKind = iota
	EVENT_RUNECOMPLETE EventKind = iota
	EVENT_RUNEEDIT     EventKind = iota
	// RESIZE     EventKind = iota
)

type EventKind uint8
type Event interface {
	Kind() EventKind
}

type EventCursor struct {
	X, Y uint16
}
func (EventCursor) Kind() EventKind {
	return EVENT_CURSOR
}

type EventScroll struct {
	X, Y int32
}
func (EventScroll) Kind() EventKind {
	return EVENT_SCROLL
}

type EventKeyPress struct {
	Key GUMIKey
}
func (EventKeyPress) Kind() EventKind {
	return EVENT_KEYPRESS
}

type EventKeyRelease struct {
	Key GUMIKey
}
func (EventKeyRelease) Kind() EventKind {
	return EVENT_KEYRELEASE
}

type EventRuneComplete struct {
	Rune rune
}
func (EventRuneComplete) Kind() EventKind {
	return EVENT_RUNECOMPLETE
}

type EventRuneEdit struct {
	Rune rune
}
func (EventRuneEdit) Kind() EventKind {
	return EVENT_RUNEEDIT
}
