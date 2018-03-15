// Existing structs for ease of development
//
package gumi

import (
	"github.com/GUMI-golang/gumi/renderline"
)

type VoidNode struct {
	parent GUMI
}

func (s *VoidNode) GUMIInit() {
}
func (s *VoidNode) born(gumi GUMI) {
	s.parent = gumi
}
func (s *VoidNode) breed(gumi []GUMI) {
}
func (s *VoidNode) Parent() GUMI {
	return s.parent
}
func (s *VoidNode) Childrun() []GUMI {
	return nil
}

type SingleNode struct {
	parent GUMI
	child  GUMI
}

func (s *SingleNode) GUMIInit() {
	s.child.GUMIInit()
}
func (s *SingleNode) born(gumi GUMI) {
	s.parent = gumi
}
func (s *SingleNode) breed(gumi []GUMI) {
	if len(gumi) > 0 {
		s.child = gumi[0]
	}
}
func (s *SingleNode) Parent() GUMI {
	return s.parent
}
func (s *SingleNode) Childrun() []GUMI {
	return []GUMI{s.child}
}

type MultipleNode struct {
	parent GUMI
	child  []GUMI
}

func (s *MultipleNode) GUMIInit() {
	for _, v := range s.child {
		v.GUMIInit()
	}
}
func (s *MultipleNode) born(gumi GUMI) {
	s.parent = gumi
}
func (s *MultipleNode) breed(gumi []GUMI) {
	s.child = gumi
}
func (s *MultipleNode) Parent() GUMI {
	return s.parent
}
func (s *MultipleNode) Childrun() []GUMI {
	res := make([]GUMI, len(s.child))
	for i, v := range s.child {
		res[i] = v
	}
	return res
}

type styleStore struct {
	style *Style
}
type rendererStore struct {
	rmana *renderline.Manager
	rnode renderline.Node
}
