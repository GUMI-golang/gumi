package actor

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/media"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"image"
)

// Actor::Image
//
// AImage is an element for outputting images.
// The image uses iamGreedy / drawer.Drawer rather than image.Image
type AImage struct {
	// gumi.GUMI
	//renderline.RenderSetuper
	drawer media.Drawer
}

func (s *AImage) BaseRender(subimg *image.RGBA) {
	s.drawer.Draw(subimg)
}

func (s *AImage) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}


// GUMIFunction / GUMISize 		-> Define
func (s AImage) GUMISize() gcore.Size {
	bd := s.drawer.Bound()
	return gcore.Size{
		Horizontal: gcore.MinLength(uint16(bd.Dx())),
		Vertical:   gcore.MinLength(uint16(bd.Dy())),
	}
}


func (s *AImage) RenderSetup(man *renderline.Manager, parent renderline.Node) {
	man.New(parent, nil).SetJob(s)
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
