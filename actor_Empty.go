package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
)

// Actor::Empty
//
// AEmpty exists only for the GUMI Tree as an element that does nothing
type AEmpty struct {
	VoidNode
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s AEmpty) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s AEmpty) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMISize 			-> Define
func (AEmpty) GUMISize() gcore.Size {
	return gcore.Size{
		gcore.AUTOLENGTH,
		gcore.AUTOLENGTH,
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define::Empty
func (s AEmpty) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
}
// GUMIEventer / GUMIHappen					-> Define
func (AEmpty) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s AEmpty) String() string {
	return fmt.Sprintf("%s", "AEmpty")
}

// Constructor
func AEmpty0() *AEmpty {
	return &AEmpty{}
}
