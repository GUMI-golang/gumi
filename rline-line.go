package gumi

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"sync"
	"github.com/pkg/errors"
)

type RenderLine struct {
	pl *Pipeline
	// 루트 노드를 저장함
	Root Render
	// 이미지를 참조하기 위해 정해짐, 기본적으로 Render()메소드의 결과는 여기에 저장
	completeImage *image.RGBA
	decalRect     image.Rectangle
	baseImage     *image.RGBA
	decalImage    *image.RGBA
	// __
	wgpool *sync.Pool
}

// Node의 Setup과정에 도움을 주는 메서드
// 편의성 이외의 의미는 없음
func (s *RenderLine) New(parent, Value Render) Render {
	if Value == nil {
		panic(errors.WithMessage(ErrorNotNil, "RenderLine.New not allow Value nil"))
	}
	Value.setManager(s)
	Value.setParent(parent)
	if parent == nil {
		Value.SetAllocation(s.completeImage.Rect)
		s.Root = Value
	} else {
		parent.appendChildrun(Value)
	}
	return Value
}

func (s *RenderLine) Render() {
	s.decalRect = image.ZR
	wg := s.wgpool.Get().(*sync.WaitGroup)
	defer s.wgpool.Put(wg)

	wg.Add(2)
	go func() {
		// 모든 요소들이 캐싱되 있고(즉 변경된 내용이 하나도 없는 경우) 별도의 렌더링 작업이 필요치 않다고 판단되는 경우
		if !AllCached(s.Root) {
			s.Root.BaseRender()
		}
		wg.Done()
	}()
	go func() {
		s.Root.DecalRender(&s.decalRect)
		wg.Done()
	}()
	wg.Wait()

}
func (s *RenderLine) Width() int {
	return s.completeImage.Rect.Dx()
}
func (s *RenderLine) Height() int {
	return s.completeImage.Rect.Dy()
}
func (s *RenderLine) Size() (w, h int) {
	return s.completeImage.Rect.Dx(), s.completeImage.Rect.Dy()
}
func (s *RenderLine) Rect() image.Rectangle {
	return s.completeImage.Rect
}
func (s *RenderLine) Image() *image.RGBA {
	if s.decalRect == image.ZR {
		return s.baseImage
	}
	draw.Draw(s.completeImage, s.completeImage.Rect, s.baseImage, s.baseImage.Rect.Min, draw.Src)
	draw.Draw(s.completeImage, s.decalRect, s.decalImage, s.decalRect.Min, draw.Over)
	s.DecalClear()
	return s.completeImage
}

//
func (s *RenderLine) BaseImage() *image.RGBA {
	return s.baseImage
}
func (s *RenderLine) DecalImage() *image.RGBA {
	return s.decalImage.SubImage(s.decalRect).(*image.RGBA)
}
func (s *RenderLine) DecalClear() {
	draw.Draw(s.decalImage, s.decalRect, clearIMG, image.ZP, draw.Src)
}

var clearIMG = image.NewUniform(color.RGBA{0, 0, 0, 0})

func NewManager(w, h int) *RenderLine {
	sz := image.Rect(0, 0, w, h)
	return &RenderLine{
		completeImage: image.NewRGBA(sz),
		baseImage:     image.NewRGBA(sz),
		decalImage:    image.NewRGBA(sz),
		wgpool: &sync.Pool{
			New: func() interface{} {
				return new(sync.WaitGroup)
			},
		},
	}
}
func AllCached(nd Render) bool {
	if !nd.valid() {
		return false
	}
	for _, child := range nd.Childrun() {
		if !AllCached(child) {
			return false
		}
	}

	return true
}
