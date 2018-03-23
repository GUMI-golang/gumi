package gumi

import (
	"image"
	"math/rand"
)

type Screen struct {
	//RenderingPipeline *renderline.RenderLine
	//EventPipeline     *eventline.RenderLine
	rstyle            *Style
	//
	_hook map[uint64]func(event Event) Event
	//

	//
	root GUMIRoot
	//
	title string `xml:"name"`
	width int `xml:"width"`
	height int `xml:"height"`
}


func NewScreen() *Screen {
	res := &Screen{

		rstyle:            DefaultStyle(),
		_hook:             make(map[uint64]func(event Event) Event),
	}
	return res
}

func (s *Screen) Width() int {
	//return s.RenderingPipeline.Width()
	return 0
}
func (s *Screen) Height() int {
	//return s.RenderingPipeline.Height()
	return 0
}
func (s *Screen) Size() (width, height int) {
	//return s.RenderingPipeline.Size()
	return 0, 0
}
func (s *Screen) Root(root GUMI) {
	//s.root = newGUMIRoot(s, root)
}

//
func (s *Screen) Event(event Event) {
	for _, v := range s._hook {
		if v != nil {
			event = v(event)
		}
	}
	if event == nil {
		return
	}
	//s.root.GUMIHappen(event)
}
func (s *Screen) Init() {
	//s.root.GUMIInit()
	//s.root.GUMIStyle(s.rstyle)

	// renderline은 렌더 트리를 완성시킨 이후 셋업해야 함 따라서 GUMIRenderSetup 이후 Setup을 함
	//s.root.GUMIRenderSetup(s.RenderingPipeline, s.RenderingPipeline.New(nil, nil))
}
func (s *Screen) Update(info Information) {
	//s.root.GUMIInfomation(info)
}
func (s *Screen) Draw() {
	//s.RenderingPipeline.Render()
}

//
func (s *Screen) Frame() image.Image {
	//return s.RenderingPipeline.Image()
	return nil
}
func (s *Screen) RGBA() *image.RGBA {
	//return s.RenderingPipeline.Image()
	return nil
}

//
func (s *Screen) hookReserve() (id uint64) {
	defer func() {
		s._hook[id] = nil
	}()
	for {
		id = rand.Uint64()
		if id == 0 {
			continue
		}
		if _, ok := s._hook[id]; !ok {
			return id
		}
	}
}
func (s *Screen) hookRequest(id uint64, hooking func(event Event) Event) {
	s._hook[id] = hooking
}
