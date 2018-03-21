package pipelines

import "github.com/GUMI-golang/gumi"

type Manager struct {
	Root *Node
}

func (s *Manager) New(parent *Node, elem gumi.GUMI) *Node {
	temp := &Node{
		Parent:parent,
		Elem:elem,
	}
	if parent == nil{
		s.Root = temp
	}
	return temp
}