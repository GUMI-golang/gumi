package renderline

import (
	"github.com/GUMI-golang/gumi/gcore"
	"image"
	"image/draw"
	"sync"
)

type SilluetNode struct {
	baseNode Node
	silluetNode Node
	silluetEnable bool
}


func NewSilluetNode() Node {
	return &SilluetNode{}
}

func (s *SilluetNode) Setup() {
	//캐시 영역을 만들어 둠
	if s.do != nil {
		s.cache = image.NewRGBA(s.allocation)
	}
	s.cacheValid = false
	for _, child := range s.childrun {
		child.Setup()
	}
}
func (s *SilluetNode) BaseRender() {
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
func (s *SilluetNode) childrunBaseRender() {
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
func (s *SilluetNode) DecalRender(updated *image.Rectangle) {
	if s.do != nil {
		*updated = updated.Union(s.do.DecalRender(s.manager.decalImage))
	}
	s.childrunDecalRender(updated)
}
func (s *SilluetNode) childrunDecalRender(updated *image.Rectangle) {
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
func (s *SilluetNode) ThrowCache() {
	s.cacheValid = false
	for _, child := range s.childrun {
		child.ThrowCache()
	}
}

func (s *SilluetNode) GetAllocation() image.Rectangle {
	return s.allocation
}
func (s *SilluetNode) SetAllocation(alloc image.Rectangle) {
	s.allocation = alloc
}
func (s *SilluetNode) GetJob() Job {
	return s.do
}
func (s *SilluetNode) SetJob(j Job) {
	s.do = j
}

func (s *SilluetNode) setManager(man *Manager) {
	s.baseNode.setManager(man)
	s.silluetNode.setManager(man)
}
func (s *SilluetNode) Manager() *Manager {
	return s.baseNode.Manager()
}
func (s *SilluetNode) Parent() Node {
	return s.baseNode.Parent()
}
func (s *SilluetNode) Childrun() []Node {
	return []Node{s.baseNode, s.silluetNode}
}
func (s *SilluetNode) setParent(n Node) {
	s.baseNode.setParent(n)
	s.silluetNode.setParent(n)
}
func (s *SilluetNode) appendChildrun(c ...Node) {
	s.childrun = append(s.childrun, c...)
}
func (s *SilluetNode) clearChildrun(c ...Node) {
	s.childrun = nil
}

func (s *SilluetNode) valid() bool {
	return s.cacheValid
}
