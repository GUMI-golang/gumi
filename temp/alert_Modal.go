package temp

import (
	"image"
	"fmt"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// ALert::Modal
//
//
type ALModal struct {
	SingleNode
	//
	lastCursorEvent EventCursor
	//
	modal GUMI
	show bool
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

// GUMIFunction / GUMIClip 					-> Define
func (s *ALModal) GUMIClip(r image.Rectangle) {
	s.child.GUMIClip(r)
	s.modal.GUMIClip(r)
}

// GUMIFunction / GUMIRender 				-> Define::Empty
func (s *ALModal) GUMIRender(frame *image.RGBA) {
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
func (s *ALModal) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
	s.child.GUMIRenderSetup(tree, parentnode)
	s.modal.GUMIRenderSetup(tree, parentnode)
}

// GUMIRenderer / GUMIUpdate			-> Define
func (s *ALModal) GUMIUpdate() {
	if s.show{
		s.child.GUMIUpdate()
		s.modal.GUMIUpdate()
	}else {
		s.child.GUMIUpdate()
	}
}

// GUMIEventer / GUMIHappen					-> Define
func (s *ALModal) GUMIHappen(event Event) {
	if s.show{
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
	temp := &ALModal{
		modal:modal,
	}
	temp.modal.born(temp)
	return temp
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
	s.modal.born(s)
}

// Method / GetModal
func (s *ALModal ) GetModal() GUMI {
	return s.modal
}

// Method / SetShow
func (s *ALModal ) SetShow(show bool)  {
	s.show = show
	s.modal.GUMIHappen(s.lastCursorEvent)
}

// Method / GetShow
func (s *ALModal ) GetShow() bool {
	return s.show
}