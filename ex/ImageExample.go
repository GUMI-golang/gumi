package main

import (
	"fmt"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi-basic/actor"
	"github.com/GUMI-golang/gumi-basic/node"
	"github.com/GUMI-golang/gumi/gcore"
	"strings"
	"github.com/GUMI-golang/gumi-basic/layout"
	"runtime"
	"github.com/GUMI-golang/gorat/oglSupport/v43"
	"github.com/GUMI-golang/gorat/fwrat"
	"os"
	"github.com/GUMI-golang/gumi-basic/drawing"
)

func main() {
	f := gcore.MustValue(os.Open("./ex/example-03.gumi.xml")).(*os.File)
	defer f.Close()
	//
	xq, err := gumi.NewMLGBuilder(
		f,
		gcore.MustValue(gumi.NewSpace(
			actor.ActorSpace,
			node.NodeSpace,
			layout.LayoutSpace,
			drawing.DrawingSpace,
		)).(*gumi.Space),
	)
	if err != nil {
		panic(err)
	}
	//==========================================================================================
	runtime.LockOSThread()
	octx := gcore.MustValue(fwrat.CreateOffscreenContext(v43.Driver())).(*fwrat.Offscreen)
	defer octx.Delete()
	gorat.Use(octx)

	result := octx.Driver().Result(800, 600)
	defer result.Delete()
	res0 := octx.Driver().Result(800, 600)
	defer res0.Delete()
	rast0 := gorat.NewHardware(res0)
	res1 := octx.Driver().Result(800, 600)
	defer res1.Delete()
	rast1 := gorat.NewHardware(res1)
	//==========================================================================================

	t := gcore.MustValue(xq.Build()).(*gumi.Screen)
	recurpipe(t.Pipeline.Root, 0)
	fmt.Println("==============================================")
	t.Pipeline.Rasterizer(rast0, rast1)
	t.Pipeline.Rendering()
	//==========================================================================================
	gorat.Mixing(result, res0, res1)
	//
	gcore.Capture("aout", result.Get())
	gcore.Capture("aout0", res0.Get())
	gcore.Capture("aout1", res1.Get())
	fmt.Println()
}
func recurpipe(pipe *gumi.Pipe, depth int) {
	if pipe == nil {
		return
	}
	fmt.Println(strings.Repeat("    ", depth), pipe.Elem)
	for _, c := range pipe.Childrun {
		recurpipe(c, depth+1)
	}
}
