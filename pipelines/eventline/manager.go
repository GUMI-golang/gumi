package eventline

import (
	"github.com/GUMI-golang/gumi"
)

type Manager struct {
	Root Node
}

func NewManager() *Manager {
	return new(Manager)
}
func (s *Manager) New(parent *Node, tp Node) Node {
	if tp == nil{
		tp = new(AllNode)
	}
	tp.Set(s, parent)
	if parent == nil{
		s.Root = tp
	}
	return tp
}
func (s *Manager) Occur(event gumi.Event) {
	s.Root.Occur(event)
}