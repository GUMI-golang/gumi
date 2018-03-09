package gumi

import (
	"image"
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

// Layout::Horizontal
//
// Horizontal align
type LVertical struct {
	MultipleNode
	rendererStore
	rule gcore.Distribute
}

// GUMIFunction / GUMIInit 					-> SingleNode::Default

// GUMIFunction / GUMIInfomation 			-> Define
func (s *LVertical) GUMIInfomation(info Information) {
	for _, v := range s.child{
		v.GUMIInfomation(info)
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *LVertical) GUMIStyle(style *Style) {
	for _, v := range s.child{
		v.GUMIStyle(style)
	}
}


// GUMIFunction / GUMISize 					-> Define
func (s *LVertical) GUMISize() gcore.Size {
	var minMax, sum uint16 = 0, 0
	for _, v := range s.child{
		sz := v.GUMISize()
		if sz.Horizontal.Min > minMax{
			minMax = sz.Horizontal.Min
		}
		sum += sz.Vertical.Min
	}
	return gcore.Size{
		gcore.MinLength(sum),
		gcore.MinLength(minMax),
	}
}

// GUMITree / born 							-> MultipleNode::Default

// GUMITree / breed 						-> MultipleNode::Default

// GUMITree / parent()						-> MultipleNode::Default

// GUMITree / childrun()					-> MultipleNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define::Empty
func (s *LVertical) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent)
	// 렌더링 영역 할당
	var tempVert = make([]gcore.Length, len(s.child))
	var tempHori = make([]gcore.Length, len(s.child))
	for i, v := range s.child{
		tempVert[i] = v.GUMISize().Vertical
		tempHori[i] = v.GUMISize().Horizontal
	}
	dis := s.rule(s.rnode.Allocation.Dy(), tempVert)
	//
	var startat = s.rnode.Allocation.Min.Y
	for i, v := range s.child{
		inrect := image.Rect(
			s.rnode.Allocation.Min.X,
			startat,
			s.rnode.Allocation.Max.X,
			startat + dis[i],
		)
		temp := s.rmana.New(s.rnode)
		temp.Allocation = inrect
		v.GUMIRenderSetup(s.rmana, temp)
		startat += dis[i]
	}
}

// GUMIEventer / GUMIHappen					-> Define
func (s *LVertical) GUMIHappen(event Event) {
	for _, v := range s.child{
		go v.GUMIHappen(event)
	}
}

// fmt.Stringer / String					-> Define
func (s *LVertical) String() string{
	return fmt.Sprintf(
		"%s(childrun:%d)", "LVertical", len(s.Childrun()),
	)
}

// Constructor 0
func LVertical0(rule gcore.Distribute, childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule:rule,
	}
	for _, v := range childrun{
		v.born(s)
	}
	s.breed(childrun)
	return s
}

// Constructor 1
func LVertical1(childrun ...GUMI) *LVertical {
	s := &LVertical{
		rule: gcore.Distribution.Minimalize,
	}
	for _, v := range childrun{
		v.born(s)
	}
	s.breed(childrun)
	return s
}

// Get Elems
func (s *LVertical) LoadElements(index gcore.Index, count int) []GUMI {
	return loadGUMIChildrun(s.child, index, count)
}

// Count Elems
func (s *LVertical) SizeElements() int {
	return len(s.child)
}

// Set Elems
func (s *LVertical) SaveElements(mode gcore.Mode, index gcore.Index, elem ...GUMI) (input int) {
	return saveGUMIChildrun(&s.child, mode, index, elem...)
}