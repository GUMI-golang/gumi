package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
	"image"
)

type NDrawing struct {
	SingleNode
	styleStore
	rendererStore
	//
	drawfuncs []Drawer
}

func (s *NDrawing) BaseRender(subimg *image.RGBA) {
}

func (s *NDrawing) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	sub := fullimg.SubImage(s.rnode.Allocation).(*image.RGBA)
	ctx := createContext(sub)
	var res image.Rectangle
	for _, v := range s.drawfuncs {
		ctx.Push()
		res = res.Union(v.Draw(ctx, s.style).Add(s.rnode.Allocation.Min))
		ctx.Pop()
	}
	return res
}

func (s *NDrawing) GUMIInfomation(info Information) {
	var changed bool
	for _, v := range s.drawfuncs{
		if v2, ok := v.(DrawerWithInformation); ok{
			changed = changed || v2.Inform(info)
		}
	}
	s.child.GUMIInfomation(info)
}
func (s *NDrawing) GUMIStyle(style *Style) {
	s.style = style
	s.child.GUMIStyle(style)
}

func (s *NDrawing) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent)
	s.rnode.Do = s
	s.child.GUMIRenderSetup(man, s.rnode)
}

func (s *NDrawing) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NDrawing) GUMISize() gcore.Size {
	return s.child.GUMISize()
}
func (s *NDrawing) String() string {
	return fmt.Sprintf("%s(drawing:%d GUMIRender)", "NDrawing", len(s.drawfuncs))
}
//
func NDrawing0(drawFuncs ...Drawer) *NDrawing {
	return &NDrawing{
		drawfuncs: drawFuncs,
	}
}
