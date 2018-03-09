package temp

import (
	"fmt"
	"image"
	"github.com/iamGreedy/gumi/gumre"
	"github.com/iamGreedy/gumi/drawer"
)

// MTRadio Default Values
const (
	mtRadioMinWidth                  = 20
	mtRadioMinHeight                 = 20
	mtRadioAnimationOnOffDeltaMillis = 200
	mtRadioInnerRadiusDifference = 3
)

// MTRadio Animations
const (
	mtRadioAnimationOnOff  = iota
	//
	mtRadioAnimationLength = iota
)

// Material::Radio
//
// Material theme radio button(kind of toggle button)
type MTRadio struct {
	//
	VoidNode
	boundStore
	styleStore
	rendererStore
	//
	mtColorFromTo
	studio *gcore.Studio
	onoff *gcore.Percenting
	//
	cursorEnter, active bool
	onActive            MTRadioActive
}

// Material::Radio<Callback> -> Focus
//
// Click this occur it
type MTRadioActive func(self *MTRadio, active bool)

// GUMIFunction / GUMIInit 					-> Define
func (s *MTRadio) GUMIInit() {
	s.studio = gcore.Animation.Studio(mtRadioAnimationLength)
	s.onoff = s.studio.Set(mtRadioAnimationOnOff, &gcore.Percenting{
		Delta: gcore.Animation.PercentingByMillis(mtRadioAnimationOnOffDeltaMillis),
		Fn:    Material.DefaultAnimation.Radio,
	}).(*gcore.Percenting)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTRadio) GUMIInfomation(info Information) {
	if s.active {
		s.onoff.Request(1)
	} else {
		s.onoff.Request(0)
	}
	s.studio.Animate(float64(info.Dt))
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTRadio) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMIClip 					-> Define
func (s *MTRadio) GUMIClip(r image.Rectangle) {
	s.bound = r
}

// GUMIFunction / GUMIRender 				-> Define
func (s *MTRadio) GUMIRender(frame *image.RGBA) {
	var ctx = createContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor0 = s.GetFromMaterialColor().BaseColor()
	var mainColor1 = s.GetToMaterialColor().MainColor()
	var radius = h / 2
	var innerRadius = radius - mtRadioInnerRadiusDifference
	//
	ctx.SetColor(baseColor0)
	ctx.DrawCircle(w/2, h/2, radius)
	ctx.Fill()
	//
	ctx.SetColor(Scale.Color(baseColor0, mainColor1, s.onoff.Value()))
	ctx.DrawCircle(w/2, h/2, innerRadius)
	ctx.Fill()
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTRadio) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   gcore.FixLength(mtRadioMinHeight),
		Horizontal: gcore.FixLength(mtRadioMinWidth),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTRadio) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
	s.frame = frame
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *MTRadio) GUMIUpdate() {
	panic("implement me")
}

// GUMIRenderer / GUMIDraw 					-> Define
func (s *MTRadio) GUMIDraw() {
	s.GUMIRender(frame)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTRadio) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
	case EventKeyRelease:
		if ev.Key == KEY_MOUSE1 {
			if s.cursorEnter {
				s.active = !s.active
				if s.onActive != nil {
					s.onActive(s, s.active)
				}
			}
		}
	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		if (s.bound.Min.X <= x && x < s.bound.Max.X) && (s.bound.Min.Y <= y && y < s.bound.Max.Y) {
			s.cursorEnter = true
		} else {
			s.cursorEnter = false
		}
	}
}

// fmt.Stringer / String					-> Define
func (s *MTRadio) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTRadio", s.active)
}

// Constructor 0
func MTRadio0(active MTRadioActive) *MTRadio {
	temp := &MTRadio{
		onActive: active,
	}
	temp.SetFromMaterialColor(Material.Pallette.White)
	temp.SetToMaterialColor(Material.Pallette.White)
	return temp
}

// Constructor 1
func MTRadio1(from, to *MaterialColor, active MTRadioActive) *MTRadio {
	temp := &MTRadio{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}

// Method / Get -> GetActive()
func (s *MTRadio) Get() bool {
	return s.GetActive()
}

// Method / Set -> SetActive()
func (s *MTRadio) Set(active bool) {
	s.SetActive(active)
}

// Method / Get
func (s *MTRadio) GetActive() bool {
	return s.active
}

// Method / Set
func (s *MTRadio) SetActive(active bool) {
	s.active = active
}

// Method / Set Callback
func (s *MTRadio) OnActive(callback MTRadioActive) {
	s.onActive = callback
}

// Method / Get Callback
func (s *MTRadio) ReferActive() MTRadioActive {
	return s.onActive
}
