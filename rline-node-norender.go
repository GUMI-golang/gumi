package gumi

import (
	"image"
	"sync"
)

type NoRenderNode struct {
	baseRenderNode
	allocation image.Rectangle
	cacheValid bool
}

func (s *NoRenderNode) DoRender() {
}

func (s *NoRenderNode) PostRender(updated *image.Rectangle) {
}

func (s *NoRenderNode) ThrowCache() {
	s.cacheValid = false
	for _, child := range s.childrun {
		child.ThrowCache()
	}
}


func (s *NoRenderNode) valid() bool {
	return s.cacheValid
}

func (s *NoRenderNode) GetAllocation() image.Rectangle {
	return s.allocation
}

func (s *NoRenderNode) SetAllocation(alloc image.Rectangle) {
	s.allocation = alloc
}