package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
	"image"
	"image/color"
)

const (
	mtColorAnimFullExtendMillis = 300
	mtColorWidth                = 20
	mtColorHeight               = 20
	mtColorRadiusDifference     = 3
	mtColorExtendProportion     = 4
	mtColorVStep                = 32
)

const (
	mtColorAnimationExtend = iota
	//
	mtColorAnimationLength = iota
)

type (
	MTColor struct {
		VoidNode
		rendererStore
		mtColorSingle
		//
		studio *gcore.Studio
		extend *gcore.Percenting
		//
		c color.Color
		//
		hsv_H, hsv_S, hsv_V float64
		ignoreflag          bool
		cursorEnter, active bool
		//
		onChange MTColorChange
	}
	MTColorChange func(self *MTColor, c color.Color)
)

func getWHRadius(nd renderline.Node) (w, h, rad float64) {
	bd := nd.GetAllocation()
	w, h = float64(bd.Dx()), float64(bd.Dy())
	if w > h {
		rad = h / 2
	} else {
		rad = w / 2
	}
	return w, h, rad
}
func (s *MTColor) BaseRender(subimg *image.RGBA) {
	var ctx = createContext(subimg)
	var w, h, radius = getWHRadius(s.rnode)
	var innerRadius = radius - mtColorRadiusDifference
	var hw, hh = w / 2, h / 2
	//
	ctx.SetColor(s.mcl1.BaseColor())
	ctx.DrawCircle(hw, hh, radius)
	ctx.Fill()
	//
	ctx.SetColor(s.c)
	ctx.DrawCircle(hw, hh, innerRadius)
	ctx.Fill()

}

func (s *MTColor) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	var bd = s.rnode.GetAllocation()
	var center = bd.Min.Add(bd.Max).Div(2)
	var _, _, radius = getWHRadius(s.rnode)
	var realR = radius + radius*mtColorExtendProportion*s.extend.Value()
	var innerR = realR - mtColorRadiusDifference
	var decalbd = image.Rect(
		center.X-int(realR),
		center.Y-int(realR),
		center.X+int(realR),
		center.Y+int(realR),
	)
	//
	if realR > radius {
		var ctx = createContext(fullimg.SubImage(decalbd).(*image.RGBA))
		var hw, hh = float64(ctx.Width()) / 2, float64(ctx.Height()) / 2
		//
		ctx.SetColor(s.mcl1.BaseColor())
		ctx.DrawCircle(hw, hh, realR)
		ctx.Fill()
		//
		quadInnerR := int(innerR * innerR)
		for x := -int(innerR); x < int(innerR); x++ {
			for y := -int(innerR); y < int(innerR); y++ {
				if x*x+y*y > quadInnerR {
					continue
				}
				//
				r, radian := gcore.ToPolar(float64(x), float64(y))
				radianZO := gcore.ToZeroOne(radian)
				rZO := r / innerR
				//
				hsvtoRGB := gcore.HSVToRGB(radianZO, rZO, s.hsv_V)
				offset := fullimg.PixOffset(center.X+x, center.Y+y)
				fullimg.Pix[offset+0] = hsvtoRGB.R
				fullimg.Pix[offset+1] = hsvtoRGB.G
				fullimg.Pix[offset+2] = hsvtoRGB.B
				fullimg.Pix[offset+3] = 0xFF
			}
		}

		return decalbd
	}
	return image.ZR
}

func (s *MTColor) GUMIInit() {
	s.hsv_V = 1
	s.studio = gcore.Animation.Studio(mtColorAnimationLength)
	s.extend = s.studio.Set(mtColorAnimationExtend, &gcore.Percenting{
		Delta: gcore.Animation.PercentingByMillis(mtColorAnimFullExtendMillis),
		Fn:    Material.DefaultAnimation.ColorExtend,
	}).(*gcore.Percenting)
}
func (s *MTColor) GUMIInfomation(info Information) {
	if s.studio.Animate(float64(info.Dt)) {
		s.rnode.ThrowCache()
	}
	if s.extend.Value() == 0 {
		s.ignoreflag = false
	}
}

func (s *MTColor) GUMIStyle(style *Style) {
}

func (s *MTColor) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   gcore.FixLength(mtColorHeight),
		Horizontal: gcore.FixLength(mtColorWidth),
	}
}

func (s *MTColor) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.rnode.SetJob(s)
}

func (s *MTColor) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = true
				s.rnode.ThrowCache()
			}
		}
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.active {
				s.active = false
				s.c = gcore.HSVToRGB(s.hsv_H, s.hsv_S, s.hsv_V)
				s.ignoreflag = true
				s.cursorEnter = false
				if s.extend.To != 0 {
					s.extend.Request(0)
				}
				s.rnode.ThrowCache()
			}
		}
	case EventCursor:
		bd := s.rnode.GetAllocation()
		center := bd.Min.Add(bd.Max).Div(2)
		//
		_, _, r := getWHRadius(s.rnode)
		realR := r + r*mtColorExtendProportion*s.extend.Value()
		deltaR, deltaTheta := gcore.ToPolar(
			float64(ev.X)-float64(center.X),
			float64(ev.Y)-float64(center.Y),
		)
		if deltaR < realR && !s.ignoreflag{
			s.cursorEnter = true
			//
			s.hsv_H = gcore.ToZeroOne(deltaTheta)
			s.hsv_S = deltaR / realR
			if s.extend.To != 1 {
				s.extend.Request(1)
			}
		} else {
			s.cursorEnter = false
			if s.extend.To != 0 {
				s.extend.Request(0)
			}
		}
	case EventScroll:
		if s.cursorEnter {
			s.hsv_V = Clamp(s.hsv_V-float64(ev.Y+ev.X)/mtColorVStep, 0, 1)
		}
	}
}

func (s *MTColor) String() string {
	return fmt.Sprintf("%s(color:%v)", "MTColor", HexFromColor(s.c))
}

func MTColor0() *MTColor {
	temp := &MTColor{
		c: color.White,
	}
	temp.mcl1 = Material.Pallette.White
	return temp
}
