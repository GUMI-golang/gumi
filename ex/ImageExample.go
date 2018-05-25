package main

import (
	"fmt"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi-basic/actor"
	"github.com/GUMI-golang/gumi-basic/node"
	"github.com/GUMI-golang/gumi/gcore"
	"strings"
	"github.com/GUMI-golang/gumi-basic/layout"
	"runtime"
	"os"
	"github.com/GUMI-golang/gumi-basic/drawing"
	"github.com/GUMI-golang/giame/giamegl/giamegl43"
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.3-core/gl"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(1, 1, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		panic(err)
	}
	f := gcore.MustValue(os.Open("./ex/example-00.gumi.xml")).(*os.File)
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
	dr := giamegl43.NewDriver()
	gcore.Must(dr.Init())
	defer dr.Close()
	result := dr.MakeResult(800, 600).(giamegl43.GLResult)
	r0 := dr.MakeResult(800, 600).(giamegl43.GLResult)
	r1 := dr.MakeResult(800, 600).(giamegl43.GLResult)
	//==========================================================================================
	t := gcore.MustValue(xq.Build()).(*gumi.Screen)
	recurpipe(t.Pipeline.Root, 0)
	fmt.Println("==============================================")
	t.Pipeline.Rasterizer(dr, r0, r1)
	t.Pipeline.Rendering()
	//==========================================================================================
	giamegl43.Mixing(dr, result, r0, r1)
	//
	gcore.Capture("aout", result.Image())
	fmt.Println()
}
func recurpipe(pipe *gumi.Pipe, depth int) {
	if pipe == nil {
		return
	}
	fmt.Print(strings.Repeat("    ", depth))
	if bdr, ok := pipe.Elem.(gumi.Bounder);ok{
		fmt.Printf("%v, %v", bdr, bdr.GetBound())
	}else {
		fmt.Printf("%v", pipe.Elem)
	}
	fmt.Println()
	for _, c := range pipe.Childrun {
		recurpipe(c, depth+1)
	}
}
