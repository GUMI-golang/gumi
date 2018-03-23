package gumi

import (
	"image"
	"sync"
	"image/draw"
)

type Pipeline struct {
	wgpool *sync.Pool
	//
	Renderline *RenderLine
	//Eventline  *eventline.Manager
	Root *Pipe

	// Rendering, 관련
	// 이미지를 참조하기 위해 정해짐, 기본적으로 Render()메소드의 결과는 여기에 저장

	doneImage *image.RGBA
	postRect  image.Rectangle
	postImage *image.RGBA
}

func (s *Pipeline) New(parent *Pipe, elem GUMI) *Pipe {
	temp := &Pipe{
		Parent: parent,
		Elem:   elem,
	}

	if parent == nil {
		s.Root = temp
	} else {
		parent.Childrun = append(parent.Childrun, temp)
	}
	return temp
}

func (s *Pipeline) Rendering() {
	s.postRect = image.ZR
	wg := s.wgpool.Get().(*sync.WaitGroup)
	defer s.wgpool.Put(wg)

	wg.Add(2)
	go func() {
		// 모든 요소들이 캐싱되 있고(즉 변경된 내용이 하나도 없는 경우) 별도의 렌더링 작업이 필요치 않다고 판단되는 경우
		if !AllCached(s.Root.render) {
			s.Root.render.DoRender()
		}
		wg.Done()
	}()
	go func() {
		s.Root.render.PostRender(&s.postRect)
		wg.Done()
	}()
	wg.Wait()
}
func (s *Pipeline) SetSize(w, h int) {
	sz := image.Rect(0, 0, w, h)

	s.doneImage = image.NewRGBA(sz)
	s.postImage = image.NewRGBA(sz)
}
func (s *Pipeline) DoneImage() *image.RGBA {
	return s.doneImage
}
func (s *Pipeline) PostImage() (*image.RGBA, image.Rectangle) {
	return s.postImage, s.postRect
}
func (s *Pipeline) MergedImage() *image.RGBA {
	res := image.NewRGBA(s.doneImage.Rect)
	draw.Draw(res, s.doneImage.Rect, s.doneImage, s.doneImage.Rect.Min, draw.Src)
	draw.Draw(res, s.postRect, s.postImage, s.postRect.Min, draw.Over)
	return res
}

func (s *Pipeline ) Prepare()  {

}