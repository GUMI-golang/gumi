package renderline

import (
	"github.com/GUMI-golang/gumi/gcore"
	"image"
	"image/draw"
	"sync"
)

type SimpleNode struct {
	// 구조적인 요소의 성분들, 렌더링 파이프라인 트리를 이루는 정보들을 포함한다.,
	manager  *Manager
	parent   Node
	childrun []Node
	// 정보 요소의 성분들, 렌더링 작업에서 무었을 해야 하는지를 정의한다.
	allocation image.Rectangle
	do         Job
	cache      *image.RGBA
	// 렌더링 작업의 결과를 저장하는 부분
	// 렌더링 이후 이 노드가 활성화 되었었는지 등을 저장한다.
	cacheValid bool
}

func (s *SimpleNode) setManager(man *Manager) {
	s.manager = man
}

func NewSimpleNode() Node {
	return &SimpleNode{}
}

func (s *SimpleNode) Setup() {
	//캐시 영역을 만들어 둠
	if s.do != nil {
		s.cache = image.NewRGBA(s.allocation)
	}
	s.cacheValid = false
	for _, child := range s.childrun {
		child.Setup()
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
		go func(ch Node) {

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
		go func() {
			child.DecalRender(updated)
			wg.Done()
		}()
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
}
func (s *SimpleNode) GetJob() Job {
	return s.do
}
func (s *SimpleNode) SetJob(j Job) {
	s.do = j
}

func (s *SimpleNode) Manager() *Manager {
	return s.manager
}
func (s *SimpleNode) Parent() Node {
	return s.parent
}
func (s *SimpleNode) Childrun() []Node {
	return s.childrun
}
func (s *SimpleNode) setParent(n Node) {
	s.parent = n
}
func (s *SimpleNode) appendChildrun(c ...Node) {
	s.childrun = append(s.childrun, c...)
}
func (s *SimpleNode) clearChildrun(c ...Node) {
	s.childrun = nil
}

func (s *SimpleNode) valid() bool {
	return s.cacheValid
}
