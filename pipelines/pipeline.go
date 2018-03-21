package pipelines

import (
	"github.com/GUMI-golang/gumi/pipelines/eventline"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"github.com/GUMI-golang/gumi"
)

type Pipeline struct {
	Render *renderline.Manager
	Event  *eventline.Manager
	Tree   *Manager
}

func test()  {
	m := new(Manager)
	m.New(nil, gumi.NDrawing0())
}