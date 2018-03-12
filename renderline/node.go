package renderline

import "image"

type Node interface {
	Manager() *Manager
	Parent() Node
	setParent(n Node)
	Childrun() []Node
	appendChildrun(c ...Node)
	clearChildrun(c ...Node)
	Setup()
	BaseRender()
	DecalRender(updated *image.Rectangle)
	ThrowCache()
	GetAllocation() image.Rectangle
	SetAllocation(alloc image.Rectangle)
	GetJob() Job
	SetJob(j Job)
	setManager(man *Manager)
	valid() bool
}