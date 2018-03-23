package gumi

import "image"

type Render interface {
	setup(arg baseRenderNode)
	Manager() *RenderLine
	Parent() Render
	Pipe() *Pipe
	Childrun() []Render
	//
	DoRender()
	PostRender(updated *image.Rectangle)
	//
	ThrowCache()
	GetAllocation() image.Rectangle
	SetAllocation(alloc image.Rectangle)
}

type baseRenderNode struct {
	// 구조적인 요소의 성분들, 렌더링 파이프라인 트리를 이루는 정보들을 포함한다.,
	manager  *Pipeline
	parent   Render
	pipe     *Pipe
	childrun []Render
}

func (s *baseRenderNode) setup(arg baseRenderNode){
	s.manager = arg.manager
	s.parent = arg.parent
	s.pipe = arg.pipe
	s.childrun = s.childrun
}
func (s *baseRenderNode) Manager() *RenderLine {
	return s.manager.Renderline
}
func (s *baseRenderNode) Parent() Render {
	return s.parent
}
func (s *baseRenderNode) Pipe() *Pipe{
	return s.pipe
}
func (s *baseRenderNode) Childrun() []Render {
	return s.childrun
}
