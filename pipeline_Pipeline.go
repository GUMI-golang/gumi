package gumi

import (
	"github.com/GUMI-golang/gorat"
)

type Pipeline struct {
	//Eventline  *etc.Manager
	Screen *Screen
	Root   *Pipe

	renderer gorat.Rasterizer
	postRender gorat.Rasterizer
	//
	//doneImage *image.RGBA
	//postRect  image.Rectangle
	//postImage *image.RGBA
}

func NewPipeline(Screen *Screen) *Pipeline {
	return &Pipeline{
		Screen: Screen,
	}
}
func (s *Pipeline) New(parent *Pipe, elem GUMI) *Pipe {
	temp := &Pipe{
		Pipeline: s,
		Parent:   parent,
		Elem:     elem,
	}
	elem.setPipe(temp)

	if parent == nil {
		s.Root = temp
		s.Root.style = DefaultStyle()
	} else {
		parent.Childrun = append(parent.Childrun, temp)
	}
	return temp
}
func (s *Pipeline) Rasterizer(rasterzier, postRasterizer gorat.Rasterizer)  {
	if rasterzier == nil || postRasterizer == nil{
		panic("Rasterizer error")
	}
	s.renderer = rasterzier
	s.postRender = postRasterizer
}
func (s *Pipeline) Rendering() {
	s.postRender.Clear()
	WorkingPipe(s.Root, postRenderWork)
	haveTo := WorkingPipe(s.Root, doRenderValid)
	if haveTo == needResizeAndRender{
		s.renderer.Clear()
		WorkingPipe(s.Root, doRenderResize)
		WorkingPipe(s.Root, doRenderWork)
	}else if haveTo == needRender{
		s.renderer.Clear()
		WorkingPipe(s.Root, doRenderWork)
	}
}

