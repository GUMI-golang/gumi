package main

import (
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/GUMI-golang/gumi"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi-basic/actor"
	"github.com/GUMI-golang/gumi-basic/node"
	"github.com/GUMI-golang/gumi-basic/layout"
	"fmt"
	"os"
	"github.com/go-gl/gl/v4.3-core/gl"
	"github.com/GUMI-golang/gumi-basic/drawing"
	"github.com/GUMI-golang/giame/giamegl/giamegl43"
	"runtime"
)

const (
	w, h = 800, 600
)
func main() {
	runtime.LockOSThread()
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
	//glfw.SwapInterval(10)
	//
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
	dr := giamegl43.NewDriver()
	gcore.Must(dr.Init())
	defer dr.Close()
	result := dr.MakeResult(800, 600).(giamegl43.GLResult)
	r0 := dr.MakeResult(800, 600).(giamegl43.GLResult)
	r1 := dr.MakeResult(800, 600).(giamegl43.GLResult)
	//==========================================================================================
	t := gcore.MustValue(xq.Build()).(*gumi.Screen)
	fmt.Println("==============================================")
	t.Pipeline.Rasterizer(dr, r0, r1)
	//
	var fbo uint32
	//
	gl.Enable(gl.BLEND)
	gl.ClearColor(0,0,0,1)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.GenFramebuffers(1, &fbo)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo)
	gl.FramebufferTexture2D(gl.READ_FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, result.GLTex(), 0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		t.Pipeline.Rendering()
		giamegl43.Mixing(dr, result, r0, r1)
		// Upload result to FrameBuffer using FrameBufferObject(FBO)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo)
		gl.BlitFramebuffer(0,0,w,h,0,0,w,h, gl.COLOR_BUFFER_BIT, gl.LINEAR)
		//
		window.SwapBuffers()
		glfw.PollEvents()
	}

	//gcore.Capture("aout0", res.Get())
}
