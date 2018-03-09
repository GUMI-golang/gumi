package renderline

import (
	"image"
	"sync"
	"golang.org/x/image/draw"
)

type (
	Node struct {
		// 구조적인 요소의 성분들, 렌더링 파이프라인 트리를 이루는 정보들을 포함한다.,
		_NodeStructure
		// 정보 요소의 성분들, 렌더링 작업에서 무었을 해야 하는지를 정의한다.
		_NodeData
		// 렌더링 작업의 결과를 저장하는 부분
		// 렌더링 이후 이 노드가 활성화 되었었는지 등을 저장한다.
		_NodeRenderInfomation
	}
	_NodeStructure struct {
		// R_
		Manager  *Manager
		// R_
		Parent   *Node
		// R_
		Childrun []*Node
	}
	_NodeData struct {
		// RW
		Allocation image.Rectangle
		// RW
		Do         Job
		// R_
		Cache      *image.RGBA
	}
	_NodeRenderInfomation struct {
		// R_
		// 캐시 정보가 유효한지를 점검,
		CacheValid bool
	}
)

func (s *Node ) Setup()  {
	//캐시 영역을 만들어 둠
	if s.Do != nil{
		s.Cache = image.NewRGBA(s.Allocation)
	}
	s.CacheValid = false
	for _, child := range s.Childrun{
		child.Setup()
	}
}
func (s *Node ) BaseRender()  {
	if s.Do != nil{
		sub := s.Manager.baseImage.SubImage(s.Allocation).(*image.RGBA)
		if s.CacheValid {
			// 캐싱된 자료를 이용해도 되는 경우
			ParallelDraw(sub, sub.Rect, s.Cache, s.Cache.Rect.Min, draw.Src)
		}else {
			// 캐싱된 자료를 새로 만들어야 하는 경우
			s.Do.BaseRender(sub)
			s.CacheValid = true
			ParallelDraw(s.Cache, s.Cache.Rect, sub, s.Allocation.Min, draw.Src)
		}
	}else {
		s.CacheValid = true
	}
	s.childrunBaseRender()
}
func (s *Node ) childrunBaseRender()  {
	wg := s.Manager.wgpool.Get().(*sync.WaitGroup)
	defer s.Manager.wgpool.Put(wg)
	//
	wg.Add(len(s.Childrun))
	for _, child := range s.Childrun{
		go func(ch *Node) {

			ch.BaseRender()
			wg.Done()
		}(child)
	}
	wg.Wait()
}
func (s *Node ) DecalRender(updated *image.Rectangle)  {
	if s.Do != nil{
		*updated = updated.Union(s.Do.DecalRender(s.Manager.decalImage))
	}
	s.childrunDecalRender(updated)
}
func (s *Node ) childrunDecalRender(updated *image.Rectangle)  {
	wg := s.Manager.wgpool.Get().(*sync.WaitGroup)
	defer s.Manager.wgpool.Put(wg)
	//
	wg.Add(len(s.Childrun))
	for _, child := range s.Childrun{
		go func() {
			child.DecalRender(updated)
			wg.Done()
		}()
	}
	wg.Wait()
}

// 상위 요소에서 캐시를 버리면 하위 요소들도 자동으로 캐시를 버려야 한다.
func (s *Node) ThrowCache()  {
	s.CacheValid = false
	for _, child := range s.Childrun{
		child.ThrowCache()
	}
}
