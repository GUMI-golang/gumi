package gumi

import (
	"github.com/GUMI-golang/giame"
)

type Pipeline struct {
	//Eventline  *etc.Manager
	Screen *Screen
	Root   *Pipe

	d giame.Driver
	renderer giame.Result
	todoR chan *giame.Contour
	postRender giame.Result
	todoPR chan *giame.Contour
	//
	//doneImage *image.RGBA
	//postRect  image.Rectangle
	//postImage *image.RGBA
}

func NewPipeline(Screen *Screen) *Pipeline {
	return &Pipeline{
		Screen: Screen,
		todoR: make(chan *giame.Contour, 16),
		todoPR: make(chan *giame.Contour, 16),
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
func (s *Pipeline) Rasterizer(dr giame.Driver, rasterzier, postRasterizer giame.Result)  {
	if dr == nil || rasterzier == nil || postRasterizer == nil{
		panic("Rasterizer error")
	}
	s.d = dr
	s.renderer = rasterzier
	s.postRender = postRasterizer
}
func (s *Pipeline) Rendering() {
	s.postRender.Clear()
	go func() {
		WorkingPipe(s.Root, postRenderWork)
		s.todoPR <- nil
	}()
	haveTo := WorkingPipe(s.Root, doRenderValid)
	if haveTo == needResizeAndRender{
		s.renderer.Clear()
		WorkingPipe(s.Root, doRenderResize)
		go func() {
			WorkingPipe(s.Root, doRenderWork)
			s.todoR <- nil
		}()
	}else if haveTo == needRender{
		s.renderer.Clear()
		go func() {
			WorkingPipe(s.Root, doRenderWork)
			s.todoR <- nil
		}()
	}else {
		s.todoR <- nil
	}
	for pr := range s.todoPR{
		if pr == nil{
			break
		}
		s.postRender.Request(s.d, pr)

	}
	for r := range s.todoR{
		if r == nil{
			break
		}
		s.renderer.Request(s.d, r)
	}
}

