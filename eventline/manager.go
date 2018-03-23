package eventline

import (
	"github.com/GUMI-golang/gumi"
)

type Manager struct {
	Root     Node
}
func (s *Manager) New(parent Node, tp Node) Node {
	if tp == nil{
		tp = new(AllNode)
	}
	tp.setManager(s)
	tp.setParent(parent)
	if parent == nil{
		s.Root = tp
	}else {
		parent.appendChildrun(tp)
	}
	return tp
}
func (s *Manager) Occur(event gumi.Event) {
	s.Root.Occur(event)
}