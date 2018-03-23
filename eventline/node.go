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
	setManager(man *Manager)
	GetManager() *Manager
	setParent(parent Node)
	GetParent() Node
	appendChildrun(nds ... Node)
	GetChildrun() []Node

}

type BaseEventNode struct {
	man      *Manager
	parent   Node
	childrun []Node
	//
	Handler interface{}
}

func (s *BaseEventNode) setManager(man *Manager) {
	s.man = man
}
func (s *BaseEventNode) setParent(parent Node) {
	s.parent = parent
}

func (s *BaseEventNode) GetManager() *Manager {
	return s.man
}
func (s *BaseEventNode) GetParent() Node {
	return s.parent
}
func (s *BaseEventNode) appendChildrun(nds ... Node) {
	s.childrun = append(s.childrun, nds...)
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
