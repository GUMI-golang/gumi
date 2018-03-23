package gumi

import (
	"github.com/GUMI-golang/gumi/gcore"
	"image"
	"image/draw"
	"sync"
)

type SimpleNode struct {
	baseRenderNode
	// 정보 요소의 성분들, 렌더링 작업에서 무었을 해야 하는지를 정의한다.
	allocation image.Rectangle
	do         Rendering
	cache      *image.RGBA
	// 렌더링 작업의 결과를 저장하는 부분
	// 렌더링 이후 이 노드가 활성화 되었었는지 등을 저장한다.
	cacheValid bool
}


func NewSimpleNode(alloc image.Rectangle) Render {
	return &SimpleNode{
		allocation:alloc,
		cache:image.NewRGBA(alloc),
	}
}

func (s *SimpleNode) BaseRender() {
	if s.do != nil {
		sub := s.manager.baseImage.SubImage(s.allocation).(*image.RGBA)
		if s.cacheValid {
			// 캐싱된 자료를 이용해도 되는 경우
			gcore.ParallelDraw(sub, sub.Rect, s.cache, s.cache.Rect.Min, draw.Src)
		} else {
			// 캐싱된 자료를 새로 만들어야 하는 경우
			s.do.BaseRender(sub)
			s.cacheValid = true
			gcore.ParallelDraw(s.cache, s.cache.Rect, sub, s.allocation.Min, draw.Src)
		}
	} else {
		s.cacheValid = true
	}
	s.childrunBaseRender()
}
func (s *SimpleNode) childrunBaseRender() {
	wg := s.manager.wgpool.Get().(*sync.WaitGroup)
	defer s.manager.wgpool.Put(wg)
	//
	wg.Add(len(s.childrun))
	for _, child := range s.childrun {
		go func(ch Render) {

			ch.BaseRender()
			wg.Done()
		}(child)
	}
	wg.Wait()
}
func (s *SimpleNode) DecalRender(updated *image.Rectangle) {
	if s.do != nil {
		*updated = updated.Union(s.do.DecalRender(s.manager.decalImage))
	}
	s.childrunDecalRender(updated)
}
func (s *SimpleNode) childrunDecalRender(updated *image.Rectangle) {
	wg := s.manager.wgpool.Get().(*sync.WaitGroup)
	defer s.manager.wgpool.Put(wg)
	//
	wg.Add(len(s.childrun))
	for _, child := range s.childrun {
		go func(ch Render) {
			ch.DecalRender(updated)
			wg.Done()
		}(child)
	}
	wg.Wait()
}

// 상위 요소에서 캐시를 버리면 하위 요소들도 자동으로 캐시를 버려야 한다.
func (s *SimpleNode) ThrowCache() {
	s.cacheValid = false
	for _, child := range s.childrun {
		child.ThrowCache()
	}
}

func (s *SimpleNode) GetAllocation() image.Rectangle {
	return s.allocation
}
func (s *SimpleNode) SetAllocation(alloc image.Rectangle) {
	s.allocation = alloc
	s.cache = image.NewRGBA(alloc)
}
func (s *SimpleNode) GetJob() Rendering {
	return s.do
}
func (s *SimpleNode) SetJob(j Rendering) {
	s.do = j
}

func (s *SimpleNode) valid() bool {
	return s.cacheValid
}
