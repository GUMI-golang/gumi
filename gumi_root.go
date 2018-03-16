package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
)

// _::Root
//
// gumiRoot is Special case. private struct
type gumiRoot struct {
	SingleNode
	scr *Screen
}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *gumiRoot) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *gumiRoot) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}

// GUMIFunction / GUMISize 					-> Define
func (s *gumiRoot) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

// GUMITree / born 							-> SingleNode::Default

// GUMITree / breed 						-> SingleNode::Default

// GUMITree / parent()						-> SingleNode::Default

// GUMITree / childrun()					-> SingleNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define::Empty
func (s *gumiRoot) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.child.GUMIRenderSetup(man, parent)
}

// GUMIEventor
func (s *gumiRoot) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}

func (s *gumiRoot) String() string {
	return fmt.Sprintf("%s", "GUMI Root")
}

// GUMIRoot / Screen
func (s *gumiRoot) Screen() *Screen {
	return s.scr
}

// Constructor
func newGUMIRoot(scr *Screen, under GUMI) GUMIRoot {
	temp := &gumiRoot{
		scr: scr,
	}
	LinkingFrom(temp, under)
	return temp
}

// Utility Function
func Root(g GUMI) GUMIRoot {
	if g == nil {
		return nil
	}
	if v, ok := g.(*gumiRoot); ok {
		return v
	}
	return Root(g.Parent())
}
