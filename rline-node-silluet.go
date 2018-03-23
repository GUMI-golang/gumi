package gumi

import (
	"image"
)

type SilluetNode struct {
	baseRenderNode

	allocation    image.Rectangle
	silluetEnable bool
}

func (s *SilluetNode) DoRender() {
	s.childrun[0].DoRender()
	if s.silluetEnable {
		s.childrun[1].DoRender()
	}
}
func (s *SilluetNode) PostRender(updated *image.Rectangle) {
	s.childrun[0].PostRender(updated)
	if s.silluetEnable {
		s.childrun[1].PostRender(updated)
	}
}

func NewSilluetNode(alloc image.Rectangle) Render {
	return &SilluetNode{
		allocation: alloc,
	}
}

// 상위 요소에서 캐시를 버리면 하위 요소들도 자동으로 캐시를 버려야 한다.
func (s *SilluetNode) ThrowCache() {
	for _, c := range s.childrun {
		c.ThrowCache()
	}
}

//
func (s *SilluetNode) GetAllocation() image.Rectangle {
	return s.allocation
}
func (s *SilluetNode) SetAllocation(alloc image.Rectangle) {
	s.allocation = alloc
	s.childrun[0].SetAllocation(alloc)
	s.childrun[1].SetAllocation(alloc)
}
func (s *SilluetNode) Enable() {
	s.silluetEnable = true
}
func (s *SilluetNode) Disable() {
	s.silluetEnable = false
}
func (s *SilluetNode) IsEnable() bool {
	return s.silluetEnable
}
