package renderline

import (
	"image"
)

type SilluetNode struct {
	manager       *Manager
	parent        Node
	allocation    image.Rectangle
	baseNode      Node
	silluetNode   Node
	silluetEnable bool
}

func NewSilluetNode() Node {
	return &SilluetNode{
		baseNode:    NewSimpleNode(),
		silluetNode: NewSimpleNode(),
	}
}

func (s *SilluetNode) Setup() {
	//캐시 영역을 만들어 둠
	s.baseNode.Setup()
	s.silluetNode.Setup()
	//
	s.baseNode.setParent(s)
	s.silluetNode.setParent(s)
	//
	s.baseNode.setManager(s.manager)
	s.silluetNode.setManager(s.manager)

}
func (s *SilluetNode) BaseRender() {
	s.baseNode.BaseRender()
	if s.silluetEnable {
		s.silluetNode.BaseRender()
	}
}
func (s *SilluetNode) DecalRender(updated *image.Rectangle) {
	s.baseNode.DecalRender(updated)
	if s.silluetEnable {
		s.silluetNode.DecalRender(updated)
	}
}

// 상위 요소에서 캐시를 버리면 하위 요소들도 자동으로 캐시를 버려야 한다.
func (s *SilluetNode) ThrowCache() {
	s.baseNode.ThrowCache()
	s.silluetNode.ThrowCache()
}

//
func (s *SilluetNode) GetAllocation() image.Rectangle {
	return s.allocation
}
func (s *SilluetNode) SetAllocation(alloc image.Rectangle) {
	s.allocation = alloc
	s.baseNode.SetAllocation(alloc)
	s.silluetNode.SetAllocation(alloc)
}
func (s *SilluetNode) GetJob() Worker {
	return nil
}
func (s *SilluetNode) SetJob(j Worker) {

}

func (s *SilluetNode) Manager() *Manager {
	return s.manager
}
func (s *SilluetNode) setManager(man *Manager) {
	s.manager = man
}
func (s *SilluetNode) Parent() Node {
	return s.parent
}
func (s *SilluetNode) setParent(n Node) {
	s.parent = n
}
func (s *SilluetNode) Childrun() []Node {
	return []Node{s.baseNode, s.silluetNode}
}
func (s *SilluetNode) appendChildrun(c ...Node) {
	cl := len(c)
	if cl >= 1 {
		s.baseNode = c[0]
	}
	if cl >= 2 {
		s.silluetNode = c[1]
	}
}
func (s *SilluetNode) clearChildrun() {
	s.baseNode = nil
	s.silluetNode = nil
}

func (s *SilluetNode) valid() bool {
	return s.baseNode.valid() && s.silluetNode.valid()
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
