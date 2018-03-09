package temp

import (
	"fmt"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/media"
)

// Actor::Canvas
//
// ACanvas using for render Vector image
type ACanvas struct {
	VoidNode
	styleStore
	rendererStore
	//
	w, h uint16
	fn   Drawer
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation			-> Define
func (s *ACanvas) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle					-> Define
func (s *ACanvas) GUMIStyle(style *Style) {
	if s.style != style{
		s.style = style
		s.rnode.Require()
	}
}

// GUMIFunction / GUMIClip					-> Define
func (s *ACanvas) GUMIClip(rect image.Rectangle) {
	s.rnode.ChangeRect(rect)
}

// GUMIFunction / GUMIRender				-> Define
func (s *ACanvas) GUMIRender(frame *image.RGBA) {
	ctx := createContext(frame)
	s.fn.Draw(ctx, s.style)
}

// GUMIFunction / GUMISize					-> Define
func (s ACanvas) GUMISize() gcore.Size {
	return gcore.Size{
		Horizontal: gcore.FixLength(uint16(s.w)),
		Vertical:   gcore.FixLength(uint16(s.h)),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas) GUMIRenderSetup(tree *media.RenderTree, parentnode *media.RenderNode) {
	s.rtree = tree
	s.rnode = tree.New(parentnode)
}

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas) GUMIDraw() {
	sub := s.rnode.SubImage()
	s.GUMIRender(sub)
	s.rnode.Complete()
}

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *ACanvas ) GUMIUpdate()  {
	if s.rnode.Check(){
		s.GUMIDraw()
	}
}

// GUMIEventer / GUMIHappen					-> Define
func (s *ACanvas) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *ACanvas) String() string {
	return fmt.Sprintf("%s", "ACanvas")
}


// Constructor
func ACanvas0(w, h uint16, fn Drawer) *ACanvas {
	return &ACanvas{
		w:  w,
		h:  h,
		fn: fn,
	}
}
