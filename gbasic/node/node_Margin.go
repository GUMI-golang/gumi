package node

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"image"
)

type NMargin struct {
	SingleNode
	rendererStore
	b gcore.Blank
}

func (s *NMargin) GUMIInfomation(info Information) {
	s.child.GUMIInfomation(info)
}
func (s *NMargin) GUMIStyle(style *Style) {
	s.child.GUMIStyle(style)
}
func (s *NMargin) GUMIClip(rect image.Rectangle) {

}
func (s *NMargin) GUMISize() gcore.Size {
	sz := s.child.GUMISize()

	hmin := sz.Horizontal.Min + s.b.L.Min + s.b.R.Min
	var hmax uint16
	if uint(sz.Horizontal.Max)+uint(s.b.L.Max)+uint(s.b.R.Max) > uint(gcore.AUTOLENGTH.Max) {
		hmax = gcore.AUTOLENGTH.Max
	} else {
		hmax = sz.Horizontal.Max + s.b.L.Max + s.b.R.Max
	}

	vmin := sz.Vertical.Min + s.b.B.Min + s.b.T.Min
	var vmax uint16
	if uint(sz.Vertical.Max)+uint(s.b.B.Max)+uint(s.b.T.Max) > uint(gcore.AUTOLENGTH.Max) {
		vmax = gcore.AUTOLENGTH.Max
	} else {
		vmax = sz.Vertical.Max + s.b.L.Max + s.b.R.Max
	}
	return gcore.Size{
		gcore.Length{vmin, vmax},
		gcore.Length{hmin, hmax},
	}
}

func (s *NMargin) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)

	var sz = s.child.GUMISize()
	var palloc = parent.GetAllocation()
	var w, l, _ = calcMargin(palloc.Dx(), sz.Horizontal, s.b.L, s.b.R)
	var h, _, t = calcMargin(palloc.Dy(), sz.Vertical, s.b.B, s.b.T)
	s.rnode.SetAllocation(image.Rect(
		palloc.Min.X+l,
		palloc.Min.Y+t,
		palloc.Min.X+l+w,
		palloc.Min.Y+t+h,
	))
	s.child.GUMIRenderSetup(s.rmana, s.rnode)
}

func (s *NMargin) GUMIHappen(event Event) {
	s.child.GUMIHappen(event)
}
func (s *NMargin) String() string {
	return fmt.Sprintf("%s(margin:%v)", "NMargin", s.b)
}

//
func calcMargin(have int, l, a, b gcore.Length) (resl, resa, resb int) {
	if int(l.Max)+int(a.Max)+int(b.Max) <= have {
		// 최대값도 만족 가능
		resl = int(l.Max)
		resa = int(a.Max)
		resb = int(b.Max)
	} else if int(l.Max)+int(a.Min)+int(b.Min) <= have {
		// 최대길이 만족, 최대여백 불가
		resl = int(l.Max)
		temp := have - resl
		resa = (temp) / (int(a.Min) + int(b.Min)) * int(a.Min)
		resb = temp - resa
	} else if int(l.Min)+int(a.Min)+int(b.Min) <= have {
		// 최저길이만 만족가능
		resa = int(a.Min)
		resb = int(b.Min)
		resl = have - resa - resb
	} else if int(l.Min) <= have {
		resl = int(l.Min)
		temp := have - resl
		resa = (temp) / (int(a.Min) + int(b.Min)) * int(a.Min)
		resb = temp - resa
	} else {
		if int(a.Max)+int(b.Max) <= have {
			resa = int(a.Max)
			resb = int(b.Max)
			resl = have - resa - resb
		} else if int(a.Min)+int(b.Min) <= have {
			resa = int(a.Min)
			resb = int(b.Min)
			resl = have - resa - resb
		} else {
			resa = 0
			resb = 0
			resl = 0
		}
	}
	return
}

//
func NMargin0(sz gcore.Blank) *NMargin {
	return &NMargin{
		b: sz,
	}
}

func (s *NMargin) Set(sz gcore.Blank) {
	s.SetMargin(sz)
}
func (s *NMargin) Get() gcore.Blank {
	return s.GetMargin()
}
func (s *NMargin) SetMargin(sz gcore.Blank) {
	s.b = sz
}
func (s *NMargin) GetMargin() gcore.Blank {
	return s.b
}
