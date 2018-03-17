package renderline

import "image"

type Node interface {
	Manager() *Manager
	Parent() Node
	setParent(n Node)
	Childrun() []Node
	appendChildrun(c ...Node)
	clearChildrun()
	Setup()
	BaseRender()
	DecalRender(updated *image.Rectangle)
	ThrowCache()
	GetAllocation() image.Rectangle
	SetAllocation(alloc image.Rectangle)
	GetJob() Worker
	SetJob(j Worker)
	setManager(man *Manager)
	valid() bool
}
