package temp

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/iamGreedy/gumi/drawer"
	"github.com/iamGreedy/gumi/gumre"
	"image"
)

// MTEdit Default Values
const (
	mtEditMinWidth                    = 80
	mtEditMinHeight                   = 20
	mtEditAnimationTextCursorInterval = 400
	mtEditDeleteInterval              = 50
	mtEditDeleteThereshold            = 200
)


// MTEdit Animations
const (
	mtEditAnimationTextCursor = iota
	//
	mtEditAnimationLength = iota
)

// Material::Edit
//
// Material theme editable text
type MTEdit struct {
	//
	VoidNode
	boundStore
	styleStore
	rendererStore
	//
	mtColorSingle
	studio                *gcore.Studio
	textCursor            *gcore.Switching
	deleteSum             int64
	deleteOn              bool
	deleteCount           int
	deleteTheresholdStack int64
	deleteRequestCount    uint
	keyCTRL               bool
	editingRune           rune
	editingNow            bool
	//
	align    gcore.Align
	text     string
	inactive bool
	//
	onChange            MTEditChange
	cursorEnter, active bool
}

// Material::Edit<Callback> -> Change
//
// If text changed it occur
// TODO : CJK Editing change
type MTEditChange func(self *MTEdit, text string)


// GUMIFunction / GUMIInit 					-> Define
func (s *MTEdit) GUMIInit() {
	s.studio = gcore.Animation.Studio(mtEditAnimationLength)
	s.textCursor = s.studio.Set(mtEditAnimationTextCursor, &gcore.Switching{
		Interval: mtEditAnimationTextCursorInterval,
	}).(*gcore.Switching)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTEdit) GUMIInfomation(info Information) {
	s.studio.Animate(float64(info.Dt))
	if s.deleteRequestCount > 0 {
		if !s.editingNow {
			if s.keyCTRL {
				s.text = StringControlBackSpace(s.text)
			} else {
				s.text = StringBackSpace(s.text, 1)
			}
		} else {
			if s.editingRune == 0 {
				s.editingNow = false
			}
		}
		s.deleteRequestCount = 0
	}
	s.deleteSum += info.Dt
	s.deleteTheresholdStack += info.Dt
	temp := s.deleteSum / mtEditDeleteInterval
	if temp >= 1 {
		s.deleteSum -= temp * mtEditDeleteInterval
		if s.deleteOn && s.deleteTheresholdStack > mtEditDeleteThereshold {
			if !s.editingNow {
				for i := int64(0); i < temp; i++ {
					s.deleteCount += 1
					if s.keyCTRL {
						s.text = StringControlBackSpace(s.text)
					} else {
						s.text = StringBackSpace(s.text, 1)
					}
				}
			} else {
				if s.editingRune == 0 {
					s.editingNow = false
				}
			}
		}
	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTEdit) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMIClip 					-> Define
func (s *MTEdit) GUMIClip(r image.Rectangle) {
	s.bound = r
}

// GUMIFunction / GUMIRender 				-> Define
func (s *MTEdit) GUMIRender(frame *image.RGBA) {
	var baseColor, mainColor = s.GetMaterialColor().Color()
	var ctx = createContextRGBASub(frame, s.bound)
	var w, h = float64(ctx.Width()), float64(ctx.Height())
	var radius = h / 2
	//
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	// string position make
	var drawtext = s.text
	if s.editingRune != 0 {
		drawtext += string(s.editingRune)
	} else {
		if s.active && s.textCursor.Switch {
			drawtext += "_"
		}
	}

	var expectw, expecth = ctx.MeasureString(drawtext)
	var stringposX, stringposY float64
	var vert, hori = gcore.ParseAlign(s.align)

	switch vert {
	case gcore.AlignBottom:
		stringposY = h
	case gcore.AlignVertical:
		stringposY = h/2 + expecth/2
	case gcore.AlignTop:
		stringposY = 0 + expecth
	}
	switch hori {
	case gcore.AlignRight:
		stringposX = w - radius - expectw
	case gcore.AlignHorizontal:
		stringposX = w/2 - expectw/2
	case gcore.AlignLeft:
		stringposX = radius
	}
	//
	ctx.SetColor(baseColor)
	ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
	ctx.DrawRectangle(radius, 0, w-radius*2, h)
	ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
	ctx.Fill()
	//
	ctx.SetColor(mainColor)
	if s.active {
		ctx.DrawArc(radius, radius, radius, gg.Radians(90), gg.Radians(270))
		ctx.DrawLine(radius, 0, w-radius, 0)
		ctx.DrawLine(radius, h, w-radius, h)
		ctx.Stroke()
		ctx.DrawArc(w-radius, radius, radius, gg.Radians(-90), gg.Radians(90))
		ctx.Stroke()
	}
	ctx.DrawString(drawtext, stringposX, stringposY)
	ctx.Stroke()
}

// GUMIFunction / GUMISize 					-> Define
func (s *MTEdit) GUMISize() gcore.Size {
	return gcore.Size{
		Vertical:   gcore.MinLength(mtEditMinHeight),
		Horizontal: gcore.MinLength(mtEditMinWidth),
	}
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup 			-> Define
func (s *MTEdit) GUMIRenderSetup(frame *image.RGBA, tree *media.RenderTree, parentnode *media.RenderNode) {
	s.frame = frame
}

// GUMIRenderer / GUMIUpdate 				-> Define
func (s *MTEdit) GUMIUpdate() {
	// TODO
	panic("implement me")
}

// GUMIRenderer / GUMIDraw 					-> Define
func (s *MTEdit) GUMIDraw() {
	s.GUMIRender(s.frame)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTEdit) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		switch ev.Key {
		case KEY_CONTROL:
			s.keyCTRL = true
		case KEY_BACKSPACE:
			if s.active {
				s.deleteOn = true
				s.deleteCount = 0
				s.deleteTheresholdStack = 0
			}
		}
	case EventKeyRelease:
		switch ev.Key {
		case KEY_CONTROL:
			s.keyCTRL = false
		case KEY_BACKSPACE:
			if s.active {
				if s.deleteCount == 0 {
					s.deleteRequest(1)
				}
				s.deleteOn = false
				s.deleteCount = 0
				s.deleteTheresholdStack = 0
			}
		case KEY_MOUSE1:
			if s.cursorEnter {
				if !s.inactive {
					s.active = true
				}

			} else {
				s.active = false
				s.deleteOn = false
				s.deleteCount = 0
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
	case EventRuneEdit:
		if s.active {
			s.editingRune = ev.Rune
			s.editingNow = true
		}
	case EventRuneComplete:
		if s.active {
			s.editingRune = 0
			s.editingNow = false
			s.text += string(ev.Rune)
			if s.onChange != nil {
				s.onChange(s, s.text)
			}
		}
	}

}

// fmt.Stringer / String					-> Define
func (s *MTEdit) String() string {
	return fmt.Sprintf("%s(text:%s)", "MTEdit", s.text)
}

//
func (s *MTEdit) deleteRequest(count uint) {
	s.deleteRequestCount += count
}

// Constructors 0
func MTEdit0() *MTEdit {
	temp := &MTEdit{
		text:  "",
		align: gcore.AlignLeft | gcore.AlignVertical,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors 1
func MTEdit1(str string) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: gcore.AlignLeft | gcore.AlignVertical,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors 2
func MTEdit2(str string, align gcore.Align) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: align,
	}
	temp.SetMaterialColor(Material.Pallette.White)
	return temp
}

// Constructors 3
func MTEdit3(mcl *MaterialColor, str string, align gcore.Align) *MTEdit {
	temp := &MTEdit{
		text:  str,
		align: align,
	}
	temp.SetMaterialColor(mcl)
	return temp
}

// Method / Set -> SetText
func (s *MTEdit) Set(str string) {
	s.SetText(str)
}

// Method / Get -> GetText
func (s *MTEdit) Get() string {
	return s.GetText()
}

// Method / Set
func (s *MTEdit) SetText(str string) {
	s.text = str
	if s.onChange != nil {
		s.onChange(s, s.text)
	}
}

// Method / Get
func (s *MTEdit) GetText() string {
	return s.text
}

// Method / Set
func (s *MTEdit) SetAlign(align gcore.Align) {
	s.align = align
}

// Method / Get
func (s *MTEdit) GetAlign() gcore.Align {
	return s.align
}

// Method / Set
func (s *MTEdit) SetActive(active bool) {
	s.inactive = !active
}

// Method / Get
func (s *MTEdit) GetActive() bool {
	return !s.inactive
}

// Method / Get Callback
func (s *MTEdit) OnChange(callback MTEditChange) {
	s.onChange = callback
}

// Method / Get Callback
func (s *MTEdit) ReferChange() MTEditChange {
	return s.onChange
}
