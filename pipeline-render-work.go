package gumi

import (
	"github.com/pkg/errors"
	"image"
)

// doRenderValid return this when need render
var needRender = errors.New("Need Render this pipe")
var needResizeAndRender = errors.New("Need Render, Resize this pipe")

// It recursively cycles through the pipeline
// Checks whether the DoRender interface needs rendering or resizing
func doRenderValid(pipe *Pipe) error {
	if r, ok := pipe.Elem.(DoRender); ok {
		// check whether the DoRender or not
		if r.needResize() {
			return needResizeAndRender
		}
		if !r.Valid() {
			// Request to render again
			return needRender
		}
	} else if b, ok := pipe.Elem.(Bounder); ok {
		// check whether the Bounder or not
		if b.needResize() {
			return needResizeAndRender
		}
	}

	return nil
}
func doRenderResize(pipe *Pipe) error {
	if b, ok := pipe.Elem.(Bounder); ok {
		if b.needResize() {
			// If the size changes, must render again.
			if proxPB := pipe.ProximateParentBound(); proxPB != image.ZR {
				b.setBound(proxPB)
				b.resizeDone()
			}
		}
	}
	return nil
}
// It recursively cycles through the pipeline
//
func doRenderWork(pipe *Pipe) error {
	if r, ok := pipe.Elem.(DoRender); ok {

		r.DoRender(pipe.Pipeline.renderer.SubRasterizer(r.GetBound()))
		r.Done()
	}
	return nil
}
func postRenderWork(pipe *Pipe) error {
	if r, ok := pipe.Elem.(PostRender); ok {
		r.PostRender(pipe.Pipeline.postRender)
	}
	return nil
}
