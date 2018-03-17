package eventline

import "github.com/GUMI-golang/gumi"

type Handler interface {
	Handle(event gumi.Event)
}
type Filter interface {
	Filt(event gumi.Event) gumi.Event
}