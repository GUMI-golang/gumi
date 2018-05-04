package main

import (
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi-basic/actor"
	"github.com/GUMI-golang/gumi-basic/node"
	"github.com/GUMI-golang/gumi-basic/layout"
	"runtime"
	"github.com/GUMI-golang/gorat"
	"github.com/GUMI-golang/gorat/oglSupport/v43"
	"fmt"
	"os"
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/GUMI-golang/gumi-basic/drawing"
)

const (
	w, h = 800, 600
)
func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(w, h, "Cube", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
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
	drv := v43.Driver()
	gcore.Must(drv.Init())
	sctx := gorat.CreateSimpleContext(drv)
	gorat.Use(sctx)
	res0 := sctx.Driver().Result(800, 600)
	defer res0.Delete()
	rast0 := gorat.NewHardware(res0)
	res1 := sctx.Driver().Result(800, 600)
	defer res1.Delete()
	rast1 := gorat.NewHardware(res1)
	//==========================================================================================
	t := gcore.MustValue(xq.Build()).(*gumi.Screen)
	fmt.Println("==============================================")
	t.Pipeline.Rasterizer(rast0, rast1)
	//
	var fbo [2]uint32
	//
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.GenFramebuffers(2, &fbo[0])
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo[0])
	gl.FramebufferTexture2D(gl.READ_FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, uint32(res0.(v43.GLResult)), 0)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo[1])
	gl.FramebufferTexture2D(gl.READ_FRAMEBUFFER, gl.COLOR_ATTACHMENT1, gl.TEXTURE_2D, uint32(res1.(v43.GLResult)), 0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		//
		t.Pipeline.Rendering()
		//
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo[0])
		gl.BlitFramebuffer(0,0,w,h,0,0,w,h, gl.COLOR_BUFFER_BIT, gl.LINEAR)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo[1])
		gl.BlitFramebuffer(0,0,w,h,0,0,w,h, gl.COLOR_BUFFER_BIT, gl.LINEAR)
		window.SwapBuffers()
		glfw.PollEvents()
	}

	//gcore.Capture("aout0", res.Get())
}
