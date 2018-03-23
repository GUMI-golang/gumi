package material

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"github.com/fogleman/gg"
	"image"
)

// MTToggle Default Values
const (
	mtToggleMinWidth                  = 45
	mtToggleMinHeight                 = 20
	mtToggleAnimationOnOffDeltaMillis = 200
	mtToggleInnerRadiusDifference     = 3
)

// MTToggle Animations
const (
	mtToggleAnimationOnOff  = iota
	mtToggleAnimationLength = iota
)

// Material::Toggle
//
// Material theme toggle
type (
	MTToggle struct {
		//
		VoidNode
		styleStore
		rendererStore
		//
		mtColorFromTo
		studio *gcore.Studio
		onoff  *gcore.Percenting
		//
		cursorEnter, active bool
		onActive            MTToggleActive
	}
	// Click this occur it
	MTToggleActive func(self *MTToggle, active bool)
)

func (s *MTToggle) BaseRender(subimg *image.RGBA) {
	var ctx = createContext(subimg)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var baseColor0, mainColor0 = s.GetFromMaterialColor().Color()
	var mainColor1 = s.GetToMaterialColor().MainColor()
	var radius = h / 2
	var innerRadius = radius - mtToggleInnerRadiusDifference
	//
	ctx.SetColor(Scale.Color(baseColor0, mainColor1, s.onoff.Value()))
	//ctx.SetColor(color.RGBA{94, 97, 97, 255})
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	ctx.SetColor(mainColor0)
	//ctx.SetColor(color.RGBA{213, 217, 255, 255})
	ctx.DrawCircle(radius+Scale.Length(w-2*radius, s.onoff.Value()), radius, innerRadius)
	ctx.Fill()
}

func (s *MTToggle) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

// GUMIFunction / GUMIInit 					-> Define
func (s *MTToggle) GUMIInit() {
	s.studio = gcore.Animation.Studio(mtToggleAnimationLength)
	s.onoff = s.studio.Set(mtToggleAnimationOnOff, &gcore.Percenting{
		Delta: gcore.Animation.PercentingByMillis(mtToggleAnimationOnOffDeltaMillis),
		Fn:    Material.DefaultAnimation.Toggle,
	}).(*gcore.Percenting)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTToggle) GUMIInfomation(info Information) {
	if s.active {
		s.onoff.Request(1)
	} else {
		s.onoff.Request(0)
	}
	if s.studio.Animate(float64(info.Dt)) {
		s.rnode.ThrowCache()
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTToggle) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTToggle) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   gcore.MinLength(mtToggleMinHeight),
		Horizontal: gcore.MinLength(mtToggleMinWidth),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTToggle) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.rnode.SetJob(s)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTToggle) GUMIHappen(event Event) {
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
		bd := s.rnode.GetAllocation()
		if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
			s.cursorEnter = true
		} else {
			s.cursorEnter = false
		}
	}
}

// fmt.Stringer / String					-> Define
func (s *MTToggle) String() string {
	return fmt.Sprintf("%s(active:%v)", "MTToggle", s.active)
}

// Constructor 0
func MTToggle0(active MTToggleActive) *MTToggle {
	temp := &MTToggle{
		onActive: active,
	}
	temp.SetFromMaterialColor(Material.Pallette.White)
	temp.SetToMaterialColor(Material.Pallette.White)
	return temp
}

// Constructor 1
func MTToggle1(from, to *MaterialColor, active MTToggleActive) *MTToggle {
	temp := &MTToggle{
		onActive: active,
	}
	temp.SetFromMaterialColor(from)
	temp.SetToMaterialColor(to)
	return temp
}

// Method / Get -> GetActive
func (s *MTToggle) Get() bool {
	return s.GetActive()
}

// Method / Set -> SetActive
func (s *MTToggle) Set(active bool) {
	s.SetActive(active)
}

// Method / Get
func (s *MTToggle) GetActive() bool {
	return s.active
}

// Method / Set
func (s *MTToggle) SetActive(active bool) {
	s.active = active
}

// Method / Set Callback
func (s *MTToggle) OnActive(callback MTToggleActive) {
	s.onActive = callback
}

// Method / Get Callback
func (s *MTToggle) ReferActive() MTToggleActive {
	return s.onActive
}
