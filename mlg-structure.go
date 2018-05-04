package gumi

import (
	"github.com/GUMI-golang/gumi/gime"
	"github.com/GUMI-golang/gumi/verGo"
	"github.com/pkg/errors"
)

type (
	Document struct {
		Version verGo.Version
	}
)


type MLGResourceStore interface {
	ID() string
	Mime() string
	Data() (gime.Value)
	//
}
type (
	MLGResourceMemory struct {
		id   string
		mime string
		data gime.Value
	}
	MLGResourceLocation struct {
		id   string
		mime string
		uri  string
	}
)

func (s *MLGResourceMemory) ID() string {
	return s.id
}
func (s *MLGResourceMemory) Mime() string {
	return s.mime
}
func (s *MLGResourceMemory) Data() (gime.Value) {
	return s.data
}

func (s *MLGResourceLocation) ID() string {
	return s.id
}
func (s *MLGResourceLocation) Mime() string {
	return s.mime
}
func (s *MLGResourceLocation) Data() (gime.Value) {
	u := gime.URI(s.uri)
	if u == nil{
		return errors.WithMessage(ErrorCantParse, s.uri + "can't parse")
	}
	return gime.Request(*u)
}
