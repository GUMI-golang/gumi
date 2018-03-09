package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

// Actor::Spacer
//
// Aspacer is responsible for creating the margins.
// It can also be created with a combination of NMargin and AEmpty,
// But it exists separately for convenience.
type ASpacer struct {
	VoidNode
	//
	verical    gcore.Length
	horizontal gcore.Length
}

// GUMIFunction / GUMIInit 					-> VoidNode::Default

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s ASpacer) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 			-> Define::Empty
func (s ASpacer) GUMIStyle(style *Style) {
}

// GUMIFunction / GUMISize 				-> Define
func (s *ASpacer) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   s.verical,
		Horizontal: s.horizontal,
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent						-> VoidNode::Default

// GUMITree / childrun						-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s ASpacer) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
}

// GUMIEventer / GUMIHappen					-> Define
func (ASpacer) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (ASpacer) String() string {
	return fmt.Sprintf("%s", "ASpacer")
}

// Constructor 0
func ASpacer0(horizontal, vertical gcore.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(horizontal, vertical)
	return temp
}

// Constructor 1
func ASpacer1(horizontal gcore.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(horizontal, gcore.AUTOLENGTH)
	return temp
}

// Constructor 2
func ASpacer2(vertical gcore.Length) *ASpacer {
	temp := &ASpacer{}
	temp.Set(gcore.AUTOLENGTH, vertical)
	return temp
}

// Method / Get -> GetHorizontal(), GetVertical()
func (s *ASpacer) Get() (horizontal, vertical gcore.Length) {
	return s.horizontal, s.verical
}

// Method / Set -> SetHorizontal(...), SetVertical(...)
func (s *ASpacer) Set(horizontal, vertical gcore.Length) {
	s.horizontal, s.verical = horizontal, vertical
}

// Method / GetHorizontal
func (s *ASpacer) GetHorizontal() gcore.Length {
	return s.horizontal
}

// Method / SetHorizontal
func (s *ASpacer) SetHorizontal(horizontal gcore.Length) {
	s.horizontal = horizontal
}

// Method / GetVertical
func (s *ASpacer) GetVertical() gcore.Length {
	return s.verical
}

// Method / SetVertical
func (s *ASpacer) SetVertical(vertical gcore.Length) {
	s.verical = vertical
}
