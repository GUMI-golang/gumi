// gumi.go define GUMI interface
// GUMI is elements of GUI
// Every elements can render(graphic), affect(event), update(information) must implements this interface
package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/gime"
)

// GUMI is a collection of basic elements
type GUMI interface {
	MaxChildrun() int
	Pipe() *Pipe
	setPipe(v *Pipe)
	fmt.Stringer
}
type Sizer interface {
	Size() gcore.Size
}

type ParentGUMI struct {
	pipe *Pipe
}
func (s *ParentGUMI) Pipe() *Pipe {
	return s.pipe
}
func (s *ParentGUMI) Pipeline() *Pipeline {
	return s.pipe.Pipeline
}
func (s *ParentGUMI) Screen() *Screen {
	return s.pipe.Pipeline.Screen
}
func (s *ParentGUMI) setPipe(v *Pipe) {
	 s.pipe = v
}

type ValueManager interface {
	ListValue() []string
	GetValue(t string) (gime.Value)
	SetValue(t string, v gime.Value) error
}
type DefaultManager interface {
	GetDefault() (gime.Value)
	SetDefault(v gime.Value) error
}
