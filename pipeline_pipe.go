package gumi

type Pipe struct {
	Pipeline *Pipeline
	Parent   *Pipe
	Childrun []*Pipe
	//
	Elem GUMI
	// nilable
	//Event  eventline.Render
	render Render
}

func (s *Pipe) NewRender(parent Render, tp Render) Render {
	tp.setup(baseRenderNode{
		manager:s.Pipeline,
		pipe:s,
		parent:parent,
		childrun:nil,
	})
	s.render = tp
	return tp
}
func (s *Pipe) GetRender() Render {
	return s.render
}

