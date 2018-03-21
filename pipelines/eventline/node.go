package eventline

import (
	"github.com/GUMI-golang/gumi"
)

type Node interface {
	PreHandle(event gumi.Event) gumi.Event
	Handle(event gumi.Event)
	Occur(event gumi.Event)
	//
	SetHandle(itf interface{})
	GetHandle() interface{}
	//
	Set(man *Manager, parent Node)
	GetManager() *Manager
	GetParent() Node
	GetChildrun() []Node
}

type BaseEventNode struct {
	man      *Manager
	parent   Node
	childrun []Node
	//
	Handler interface{}
}

func (s *BaseEventNode) Set(man *Manager, parent Node) {
	s.man = man
	s.parent = parent
}
func (s *BaseEventNode) GetManager() *Manager {
	return s.man
}
func (s *BaseEventNode) GetParent() Node {
	return s.parent
}
func (s *BaseEventNode) GetChildrun() []Node {
	return s.childrun
}

func (s *BaseEventNode) SetHandle(itf interface{}) {
	s.Handler = itf
}
func (s *BaseEventNode) GetHandle() interface{} {
	return s.Handler
}
