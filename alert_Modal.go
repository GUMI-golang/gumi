package gumi

import (
	"github.com/GUMI-golang/gumi/renderline"
	"github.com/GUMI-golang/gumi/gcore"
	"fmt"
)

// ALert::Modal
//
//
type ALModal struct {
	SingleNode
	//
	rmana *renderline.Manager
	rnode *renderline.SelectNode
	//
	lastCursorEvent EventCursor
	//
	modal GUMI
}

// GUMIFunction / GUMIInit 					-> Define
func (s *ALModal) GUMIInit() {
	s.modal.GUMIInit()
	s.child.GUMIInit()
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *ALModal) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
	s.modal.GUMIInfomation(info)
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *ALModal) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
	s.modal.GUMIStyle(style)
}

// GUMIFunction / GUMISize 					-> Define
func (s *ALModal) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

// GUMITree / born 							-> SingleNode::Default

// GUMITree / breed 						-> SingleNode::Default

// GUMITree / parent()						-> SingleNode::Default

// GUMITree / childrun()					-> SingleNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ALModal) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, renderline.NewSelectNode(2)).(*renderline.SelectNode)
	s.rnode.Select(renderline.NewSelects(0))
	s.child.GUMIRenderSetup(man, s.rnode)
	s.rnode.Select(renderline.NewSelects(1))
	s.modal.GUMIRenderSetup(man, s.rnode)
	s.rnode.Select(renderline.NewSelects(0))
}

// GUMIEventer / GUMIHappen					-> Define
func (s *ALModal) GUMIHappen(event Event) {
	if s.rnode.GetSelected().Check(1){
		s.modal.GUMIHappen(event)
	}else {
		if event.Kind() == EVENT_CURSOR {
			s.lastCursorEvent = event.(EventCursor)
		}
		s.child.GUMIHappen(event)
	}
}

// fmt.Stringer / String				-> Define
func (s *ALModal) String() string {
	return fmt.Sprintf("%s", "ALModal")
}

// Constructor 0
func ALModal0() *ALModal {
	temp := &ALModal{}
	return temp
}

// Constructor 1
func ALModal1(modal GUMI) *ALModal {
	temp := ALModal{
		modal:modal,
	}
	modal.born(&temp)
	return &temp
}

// Method / SetShow
func (s *ALModal ) Set(show bool)  {
	s.SetShow(show)
}

// Method / GetShow
func (s *ALModal ) Get() bool {
	return s.GetShow()
}

// Method / SetModal
func (s *ALModal ) SetModal(modal GUMI)  {
	s.modal = modal
	modal.born(s)
}

// Method / GetModal
func (s *ALModal ) GetModal() GUMI {
	return s.modal
}

// Method / SetShow
func (s *ALModal ) SetShow(show bool)  {
	s.modal.GUMIHappen(s.lastCursorEvent)
	if show{
		s.rnode.Select(renderline.NewSelects(0, 1))
	}else {
		s.rnode.Select(renderline.NewSelects(0))
	}
	s.rnode.ThrowCache()
}

// Method / GetShow
func (s *ALModal ) GetShow() bool {
	return s.rnode.GetSelected().Check(1)
}