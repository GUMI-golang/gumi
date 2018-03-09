package temp

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
	"math"
)

// MTDropbox Default Values
const (
	mtDropboxMinWidth              = 80
	mtDropboxMinHeight             = 20
	mtDropboxScroolWidth           = 8
	mtDropboxElemMargin            = 8
	mtDropboxElemUnderline         = 2
	mtDropboxStretchSpeedPerSecond = 500
	mtDropboxScroolSpeedPerSecond  = 200
	mtDropboxScroolModify          = 16
)

// MTDropbox Animations
const (
	mtDropboxAnimationStreching = iota
	mtDropboxAnimationScroll
	mtDropboxAnimationLength
)

// Material::Dropbox
//
// Material theme Dropbox(Combo Box)
type MTDropbox struct {
	VoidNode
	boundStore
	styleStore
	rendererStore
	//
	scr     *Screen
	deferid uint64
	hookid  uint64
	//
	mtColorSingle
	studio  *gcore.Studio
	stretch *gcore.Reaching
	scroll  *gcore.Reaching
	//
	Elems    mtDropboxElemList
	selected int
	hover    int
	inactive bool
	//
	boxHeight  int
	boxCut     int
	boxMaximum int
	//
	onChange MTDropboxChange
	//
	cursorEnter, active bool
}

// Material::Dropbox<Callback> -> Change
//
// When changed selected, it occur
type MTDropboxChange func(self *MTDropbox, selected string)

// GUMIFunction / GUMIInit 					-> Define
func (s *MTDropbox) GUMIInit() {
	s.scr = Root(s).Screen()
	s.deferid = s.scr.deferReserve()
	s.hookid = s.scr.hookReserve()
	//
	s.studio = gcore.Animation.Studio(mtDropboxAnimationLength)
	s.stretch = s.studio.Set(mtDropboxAnimationStreching, &gcore.Reaching{
		Delta: mtDropboxStretchSpeedPerSecond,
		Fn:    Material.DefaultAnimation.DropboxStretch,
	}).(*gcore.Reaching)
	s.scroll = s.studio.Set(mtDropboxAnimationScroll, &gcore.Reaching{
		Delta: mtDropboxScroolSpeedPerSecond,
		Fn:    Material.DefaultAnimation.DropboxStretch,
	}).(*gcore.Reaching)

}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTDropbox) GUMIInfomation(info Information) {
	s.studio.Animate(float64(info.Dt))
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTDropbox) GUMIStyle(style *Style) {
	if s.style != style || s.Elems.needUpdate() {
		s.Elems.update(style)
		s.style = style
	}
}

// GUMIFunction / GUMIClip 					-> Define
func (s *MTDropbox) GUMIClip(r image.Rectangle) {
	s.bound = r
}

// GUMIFunction / GUMIRender 				-> Define
func (s *MTDropbox) GUMIRender(frame *image.RGBA) {
	var baseColor, mainColor = s.GetMaterialColor().Color()
	s.boxCut = 0
	s.boxHeight = mtDropboxElemMargin*(s.Elems.Length()+1) + s.Elems.heightSum()
	if s.boxMaximum < s.boxHeight {
		s.boxCut += s.boxHeight - s.boxMaximum
		s.boxHeight = s.boxMaximum
	}
	if s.bound.Max.Y+s.boxHeight > frame.Rect.Max.Y {
		s.boxCut += (s.bound.Max.Y + s.boxHeight) - frame.Rect.Max.Y
	}
	var val = s.stretch.Value()
	var per = s.stretch.Percent()
	var scr = s.scroll.Value()
	if val > 0 && s.Elems.Length() > 0 {
		// TODO : DEFERCALL
		s.scr.deferRequest(s.deferid, func(rgba *image.RGBA) {
			var radius = float64(s.bound.Dy()) / 2
			// selecte, background, scrool
			func() {
				bd := s.bound
				bd.Max.Y += int(val)
				var ctx = createContextRGBASub(rgba, bd)
				s.style.useContext(ctx)
				defer s.style.releaseContext(ctx)
				//
				var w, h = float64(ctx.Width()), float64(ctx.Height())
				// background
				ctx.SetColor(baseColor)
				ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(270))
				ctx.DrawArc(radius, h-radius, radius, gg.Radians(90), gg.Radians(180))
				ctx.DrawRectangle(radius-1, 0, w-radius*2+1, h)
				ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(0))
				ctx.DrawArc(w-radius, h-radius, radius, gg.Radians(0), gg.Radians(90))
				ctx.Fill()
				// outline
				ctx.SetColor(Scale.Color(baseColor, mainColor, per))
				ctx.DrawArc(radius, radius, radius, gg.Radians(180), gg.Radians(270))
				ctx.DrawLine(radius, 0, w-radius, 0)
				ctx.DrawArc(w-radius, radius, radius, gg.Radians(270), gg.Radians(360))
				ctx.DrawLine(w, radius, w, h-radius)
				ctx.DrawArc(w-radius, h-radius, radius, gg.Radians(0), gg.Radians(90))
				ctx.DrawLine(w-radius, h, radius, h)
				ctx.DrawArc(radius, h-radius, radius, gg.Radians(90), gg.Radians(180))
				ctx.DrawLine(0, h-radius, 0, radius)
				ctx.Stroke()
				// selected underline
				ctx.SetColor(mainColor)
				ctx.Push()
				ctx.SetLineWidth(.25)
				ctx.DrawLine(radius, float64(s.bound.Dy()), w-2*radius, float64(s.bound.Dy()))
				ctx.Stroke()
				//

				ctx.Pop()
				// selected
				selectedElem := s.Elems.getForDraw(s.selected)
				if len(selectedElem.content) > 0 {
					ctx.DrawString(selectedElem.content, radius, (float64(s.bound.Dy())-float64(selectedElem.h))/2+float64(selectedElem.h))
					ctx.Stroke()
				}
				// scroll

				percent := float64(s.boxHeight-s.boxCut) / float64(s.boxHeight)
				scroolPercent := scr / float64(s.boxHeight)
				if percent < 0 {
					percent = 0
				}
				if percent > 1 {
					percent = 1
				}

				ctx.DrawArc(w-radius, radius+(scroolPercent)*(h-radius*2), mtDropboxScroolWidth/2, gg.Radians(180), gg.Radians(360))
				ctx.DrawRectangle(w-radius-mtDropboxScroolWidth/2, radius+(scroolPercent)*(h-radius*2), mtDropboxScroolWidth, percent*(h-radius*2))
				ctx.DrawArc(w-radius, radius+(scroolPercent)*(h-radius*2)+percent*(h-radius*2), mtDropboxScroolWidth/2, gg.Radians(0), gg.Radians(180))
				ctx.Fill()
			}()
			// elems, hover
			func() {
				bd := s.bound
				bd.Min.Y = s.bound.Max.Y
				bd.Max.Y = s.bound.Max.Y + int(val)
				var ctx = createContextRGBASub(rgba, bd)
				s.style.useContext(ctx)
				defer s.style.releaseContext(ctx)
				sum := mtDropboxElemMargin
				ctx.SetColor(mainColor)
				for i, v := range s.Elems.refer() {
					drawY := float64(sum+v.h) - float64(s.scroll.Value())
					ctx.DrawString(v.content, radius, drawY)
					ctx.Stroke()
					if i == s.hover {
						ctx.DrawLine(radius, drawY+mtDropboxElemUnderline, radius+float64(v.w), drawY+mtDropboxElemUnderline)
						ctx.Stroke()
					}
					sum += v.h + mtDropboxElemMargin
				}
			}()
		})
	} else {
		s.scr.deferRequest(s.deferid, nil)
		var ctx = createContextRGBASub(frame, s.bound)
		s.style.useContext(ctx)
		defer s.style.releaseContext(ctx)
		//
		var w, h = float64(ctx.Width()), float64(ctx.Height())
		var radius = float64(s.bound.Dy()) / 2
		//
		ctx.SetColor(baseColor)
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawRectangle(radius, 0, w-radius*2, h)
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Fill()
		//
		ctx.SetColor(mainColor)
		elem := s.Elems.getForDraw(s.selected)
		if len(elem.content) > 0 {
			ctx.DrawString(elem.content, radius, (h-float64(elem.h))/2+float64(elem.h))
			ctx.Stroke()
		}
		ctx.DrawCircle(w-radius, radius, mtDropboxScroolWidth/2)
		ctx.Fill()
	}
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTDropbox) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   gcore.MinLength(mtDropboxMinHeight),
		Horizontal: gcore.MinLength(mtDropboxMinWidth),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTDropbox) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
	s.frame = frame
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *MTDropbox) GUMIUpdate() {
	// TODO
	panic("implement me")
}

// GUMIRenderer / GUMIDraw 					-> Define
func (s *MTDropbox) GUMIDraw() {
	s.GUMIRender(s.frame)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTDropbox) GUMIHappen(event Event) {
	if s.inactive {
		s.cursorEnter = false
		s.active = false
		return
	}
	switch ev := event.(type) {
	case EventKeyRelease:
		switch ev.Key {
		case KEY_MOUSE1:
			if s.cursorEnter {
				// 정지상태인 경우 클릭 무시
				if !s.active {
					// 선택상태가 아니였을 경우 선택상태로 전환, 이벤트 후킹 실시
					s.active = true
					s.scroll.Range = float64(s.boxCut)
					s.stretch.Range = float64(s.boxHeight - s.boxCut)
					s.stretch.To = float64(s.boxHeight - s.boxCut)
					s.scr.hookRequest(s.hookid, func(event Event) Event {
						if v, ok := event.(EventCursor); ok {
							bd := s.bound
							bd.Max.Y = bd.Max.Y + s.boxHeight - s.boxCut
							if bd.Min.X <= int(v.X) && int(v.X) < bd.Max.X && bd.Min.Y <= int(v.Y) && int(v.Y) < bd.Max.Y {
								s.GUMIHappen(event)
								return nil
							}
						}
						return event
					})
				} else {
					// 선택상태, 커서 진입인 경우 selected, 선택하기
					if s.hover >= 0 {
						s.selected = s.hover

					}
					s.active = false
					s.stretch.To = 0
					s.scr.hookRequest(s.hookid, nil)
					if s.onChange != nil {
						s.onChange(s, s.Elems.Get(s.selected))
					}
				}
			} else {
				// 다른 곳 클릭시 선택 취소
				s.active = false
				s.stretch.To = 0
				s.scr.hookRequest(s.hookid, nil)
			}
		}

	case EventCursor:
		x := int(ev.X)
		y := int(ev.Y)
		bd := s.bound
		if !s.active {
			if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
				s.cursorEnter = true
			} else {
				s.cursorEnter = false
			}
		} else {
			bd.Max.Y += s.boxHeight - s.boxCut
			if (bd.Min.X <= x && x < bd.Max.X) && (bd.Min.Y <= y && y < bd.Max.Y) {
				s.cursorEnter = true
				sum := s.bound.Max.Y + mtDropboxElemMargin
				if y >= s.bound.Max.Y {
					for i, elem := range s.Elems.refer() {
						if sum <= y+int(s.scroll.Value()) && y+int(s.scroll.Value()) < sum+elem.h+mtDropboxElemMargin {
							s.hover = i
							break
						}
						sum += elem.h + mtDropboxElemMargin
					}
				} else {
					s.hover = -1
				}
			} else {
				s.cursorEnter = false
				s.hover = -1
			}
		}
	case EventScroll:
		s.scroll.To += float64(mtDropboxScroolModify * ev.Y)
		if s.scroll.To < 0 {
			s.scroll.To = 0
		}
		if s.scroll.To > float64(s.boxCut) {
			s.scroll.To = float64(s.boxCut)
		}
	}
}

// fmt.Stringer / String					-> Define
func (s *MTDropbox) String() string {
	return fmt.Sprintf("%s(select:%s)", "MTDropbox", s.Elems[s.selected])
}

// Constructor 0
func MTDropbox0() *MTDropbox {
	res := &MTDropbox{
		Elems:      mtDropboxElemList{},
		selected:   0,
		hover:      -1,
		boxMaximum: math.MaxInt32,
	}
	res.SetMaterialColor(Material.Pallette.White)
	return res
}

// Constructor 1
func MTDropbox1(maxboxlen uint16) *MTDropbox {
	res := &MTDropbox{
		Elems:      mtDropboxElemList{},
		selected:   0,
		hover:      -1,
		boxMaximum: int(maxboxlen),
	}
	res.SetMaterialColor(Material.Pallette.White)
	return res
}

// Constructor 2
func MTDropbox2(change MTDropboxChange) *MTDropbox {
	res := &MTDropbox{
		Elems:      mtDropboxElemList{},
		selected:   0,
		boxMaximum: math.MaxInt32,
		hover:      -1,
	}
	res.SetMaterialColor(Material.Pallette.White)
	res.OnChange(change)
	return res
}

// Constructor 3
func MTDropbox3(change MTDropboxChange, elems ...string) *MTDropbox {
	res := &MTDropbox{
		Elems:      mtDropboxElemList{},
		selected:   0,
		boxMaximum: math.MaxInt32,
		hover:      -1,
	}
	res.SetMaterialColor(Material.Pallette.White)
	res.OnChange(change)
	for i, v := range elems {
		res.Elems.Set(i, v)
	}
	return res
}

// Elems -> Method

// Method / Get
func (s *MTDropbox) GetMaxboxLength() uint16 {
	return uint16(s.boxMaximum)
}

// Method / Set
func (s *MTDropbox) SetMaxboxLength(l uint16) {
	s.boxMaximum = int(l)
}

// Method / Get Callback
func (s *MTDropbox) OnChange(callback MTDropboxChange) {
	s.onChange = callback
}

// Method / Get Callback
func (s *MTDropbox) ReferChange() MTDropboxChange {
	return s.onChange
}
