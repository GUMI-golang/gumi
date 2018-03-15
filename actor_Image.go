package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/media"
	"github.com/GUMI-golang/gumi/renderline"
	"image"
)

// Actor::Image
//
// AImage is an element for outputting images.
// The image uses iamGreedy / drawer.Drawer rather than image.Image
type AImage struct {
	VoidNode
	rendererStore
	//
	drawer media.Drawer
}

func (s *AImage) BaseRender(subimg *image.RGBA) {
	s.drawer.Draw(subimg)
}

func (s *AImage) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s *AImage) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s *AImage) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMISize 		-> Define
func (s AImage) GUMISize() gcore.Size {
	bd := s.drawer.Bound()
	return gcore.Size{
		Horizontal: gcore.MinLength(uint16(bd.Dx())),
		Vertical:   gcore.MinLength(uint16(bd.Dy())),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AImage) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.rnode.SetJob(s)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *AImage) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *AImage) String() string {
	return fmt.Sprintf("%s", "AImage")
}

// Constructor
func AImage0(drawer media.Drawer) *AImage {
	return &AImage{
		drawer: drawer,
	}
}
