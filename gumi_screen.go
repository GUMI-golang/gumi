package gumi

import (
	"image"
	"math/rand"
	"github.com/iamGreedy/gumi/renderline"
)

type Screen struct {
	rline *renderline.Manager
	rstyle *Style
	//
	_hook  map[uint64]func(event Event) Event
	_defer map[uint64]func(rgba *image.RGBA)
	//

	//
	root GUMIRoot
}

func NewScreen(w, h int) *Screen {
	res := &Screen{
		rline: renderline.NewManager(w, h),
		rstyle: DefaultStyle(),
		_hook:  make(map[uint64]func(event Event) Event),
		_defer: make(map[uint64]func(rgba *image.RGBA)),
	}
	return res
}

func (s *Screen) Width() int {
	return s.rline.Width()
}
func (s *Screen) Height() int {
	return s.rline.Height()
}
func (s *Screen) Size() (width, height int) {
	return s.rline.Size()
}
func (s *Screen) Root(root GUMI) {
	s.root = newGUMIRoot(s, root)
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
	s.root.GUMIHappen(event)
}
//
func (s *Screen) Init() {
	s.root.GUMIInit()
	s.root.GUMIStyle(s.rstyle)

	// renderline은 렌더 트리를 완성시킨 이후 셋업해야 함 따라서 GUMIRenderSetup 이후 Setup을 함
	s.root.GUMIRenderSetup(s.rline, s.rline.New(nil))
	s.rline.Setup()
}
func (s *Screen) Update(info Information) {
	s.root.GUMIInfomation(info)
}
func (s *Screen) Draw() {
	s.rline.Render()
}
func (s *Screen) Frame() image.Image {
	return s.rline.Image()
}
func (s *Screen) RGBA() *image.RGBA {
	return s.rline.Image()
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
func (s *Screen) deferReserve() (id uint64) {
	defer func() {
		s._defer[id] = nil
	}()
	for {
		id = rand.Uint64()
		if id == 0 {
			continue
		}
		if _, ok := s._defer[id]; !ok {
			return
		}
	}
}
func (s *Screen) deferRequest(id uint64, d func(rgba *image.RGBA)) {
	s._defer[id] = d
}
