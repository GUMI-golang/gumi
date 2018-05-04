package gumi

import (
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
	"github.com/GUMI-golang/gumi/gcore"
	"io"
)



type MLGBuilder struct {
	// namespace
	space *Space
	// node
	nd tree.Node
	// setting data
	PrintComment  bool
	Verbose       bool
	IgnoreWarning bool
	// share data
	meta *Meta
	resource *Resource
	document *Document
	screen *Screen
}

func NewMLGBuilder(r io.Reader, space *Space) (*MLGBuilder, error) {
	n, err := xmltree.ParseXML(r)
	if err != nil {
		return nil, err
	}
	return &MLGBuilder{
		nd:    n,
		space: space,
	}, nil
}
func (s *MLGBuilder) Build() (*Screen, error) {
	// document
	gcore.Must(s.parseDocument())
	// resource
	gcore.Must(s.parseResource())
	// meta
	gcore.Must(s.parseMeta())
	// screen
	gcore.Must(s.parseScreen())

	return s.screen, nil
}
