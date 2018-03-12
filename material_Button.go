package gumi

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/renderline"
)

// MTButton Default Values
const (
	mtButtonMinPadding = 5
	mtButtonAnimDeltaMillis = 250
)

// MTButton Animations
const (
	mtButtonAnimationHover = iota
	//
	mtButtonAnimationLength = iota
)

// Material::Button
//
// Material theme button
type (
	MTButton struct {
		//
		VoidNode
		styleStore
		rendererStore
		//
		mtColorSingle
		studio *gcore.Studio
		hover *gcore.Percenting
		//
		text string
		//
		cursorEnter, active bool
		onClick             MTButtonClick
		onFocus             MTButtonFocus
	}

	// When Cursor enter, leave this Elem call this function
	MTButtonFocus func(self *MTButton, focus bool)
	// When Cursor click occur this
	MTButtonClick func(self *MTButton)
)



func (s *MTButton) BaseRender(subimg *image.RGBA) {
	var ctx = createContext(subimg)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor, mainColor = s.GetMaterialColor().Color()
	radius := h / 2
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	//

	if s.active {
		ctx.SetColor(baseColor)
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(mainColor)
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	} else {

		ctx.SetColor(Scale.Color(baseColor, mainColor, s.hover.Value()))
		// background
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		// text
		ctx.SetColor(Scale.Color(mainColor, baseColor, s.hover.Value()))
		txtw, txth := ctx.MeasureString(s.text)
		ctx.DrawString(s.text, (w-txtw)/2, (h+txth)/2-1)
		ctx.Stroke()
	}
}

func (s *MTButton) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

// GUMIFunction / GUMIInit 					-> Define
func (s *MTButton) GUMIInit() {
	s.studio = gcore.Animation.Studio(mtButtonAnimationLength)
	s.hover = s.studio.Set(mtButtonAnimationHover, &gcore.Percenting{
		Delta: gcore.Animation.PercentingByMillis(mtButtonAnimDeltaMillis),
		Fn:    Material.DefaultAnimation.Button,
	}).(*gcore.Percenting)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTButton) GUMIInfomation(info Information) {
	if s.studio.Animate(float64(info.Dt)){
		s.rnode.ThrowCache()
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTButton) GUMIStyle(style *Style) {
	s.style = style
}



// GUMIFunction / GUMISize 					-> Define
func (s *MTButton) GUMISize() gcore.Size {

	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()
	var hori, vert = s.style.Default.Font.CalculateSize(s.text)
	//
	return gcore.Size{
		Vertical:   gcore.MinLength(uint16(vert + mtButtonMinPadding*2)),
		Horizontal: gcore.MinLength(uint16(hori + mtButtonMinPadding*2)),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTButton) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.rnode.SetJob(s)
}


// GUMIEventer / GUMIHappen					-> Define
func (s *MTButton) GUMIHappen(event Event) {
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
				if s.onClick != nil {
					s.onClick(s)
				}
				s.active = false
				s.rnode.ThrowCache()
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		bd := s.rnode.GetAllocation()
		if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
			if s.onFocus != nil {
				s.onFocus(s, true)
			}
			//
			s.hover.Delta = gcore.Animation.PercentingByMillis(mtButtonAnimDeltaMillis)
			s.hover.Request(1)
			s.cursorEnter = true
		} else {
			if s.onFocus != nil {
				s.onFocus(s, false)
			}
			s.hover.Delta = gcore.Animation.PercentingByMillis(mtButtonAnimDeltaMillis)
			s.hover.Request(0)
			s.cursorEnter = false
		}
	}
}

// fmt.Stringer / String					-> Define
func (s *MTButton) String() string {
	return fmt.Sprintf("%s(text:%s)", "MTButton", s.GetText())
}

// Constructor 0
func MTButton0(text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructor 1
func MTButton1(mcl *MaterialColor, text string, onclick MTButtonClick) *MTButton {
	temp := &MTButton{
		text:    text,
		onClick: onclick,
	}
	temp.SetMaterialColor(mcl)
	return temp
}

// Method / Set -> SetText
func (s *MTButton) Set(txt string) {
	s.SetText(txt)
}

// Method / Get -> GetText
func (s *MTButton) Get() string {
	return s.GetText()
}

// Method / Set
func (s *MTButton) SetText(txt string) {
	s.text = txt
}

// Method / Get
func (s *MTButton) GetText() string {
	return s.text
}

// Method / Set Callback
func (s *MTButton) OnClick(callback MTButtonClick) {
	s.onClick = callback
}

// Method / Get Callback
func (s *MTButton) ReferClick() MTButtonClick {
	return s.onClick
}

// Method / Set Callback
func (s *MTButton) OnFocus(callback MTButtonClick) {
	s.onClick = callback
}

// Method / Get Callback
func (s *MTButton) ReferFocus() MTButtonClick {
	return s.onClick
}
