package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/media"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"image"
)

type (
	NCursor struct {
		SingleNode
		rendererStore
		//
		cursor *media.Fixed
		//
		x, y  uint16
		align gcore.Align
	}
)

func (s *NCursor) BaseRender(subimg *image.RGBA) {

}

func (s *NCursor) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	if s.cursor == nil {
		return image.ZR
	}
	sz := s.cursor.Bound().Size()
	x, y := int(s.x), int(s.y)
	v, h := gcore.ParseAlign(s.align)
	switch v {
	case gcore.AlignBottom:
		y = y
	case gcore.AlignVertical:
		y = y - sz.Y/2
	case gcore.AlignTop:
		y = y - sz.Y
	}
	switch h {
	case gcore.AlignRight:
		x = x - sz.X
	case gcore.AlignHorizontal:
		x = x - sz.X/2
	case gcore.AlignLeft:
		x = 0
	}
	bd := image.Rect(x, y, x+sz.X, y+sz.Y).Add(s.rnode.GetAllocation().Min)
	//
	s.cursor.Draw(fullimg.SubImage(bd).(*image.RGBA))
	return bd
}

func (s *NCursor) GUMIInfomation(info Information) {

}

func (s *NCursor) GUMIStyle(style *Style) {

}

func (s *NCursor) GUMISize() gcore.Size {
	return s.child.GUMISize()
}

func (s *NCursor) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.child.GUMIRenderSetup(man, s.rnode)
}

func (s *NCursor) GUMIHappen(event Event) {

}

func (s *NCursor) String() string {
	return fmt.Sprintf("%s", "NCursor")
}

func (s *NCursor) Get() *media.Fixed {
	return s.GetCursor()
}
func (s *NCursor) Set(cursor *media.Fixed) {
	s.SetCursor(cursor)
}
func (s *NCursor) GetCursor() *media.Fixed {
	return s.cursor
}
func (s *NCursor) SetCursor(cursor *media.Fixed) {
	s.cursor = cursor
}
func (s *NCursor) Position() (x, y uint16) {
	return s.x, s.y
}
func (s *NCursor) GetAlign() gcore.Align {
	return s.align
}
func (s *NCursor) SetAlign(align gcore.Align) {
	s.align = align
}
