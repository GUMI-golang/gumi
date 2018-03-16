package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/GUMI-golang/gumi/pipelines/renderline"
	"github.com/fogleman/gg"
	"image"
)

// MTEdit Default Values
const (
	mtEditMinWidth                    = 80
	mtEditMinHeight                   = 20
	mtEditAnimationTextCursorInterval = 400
)

// MTEdit Animations
const (
	mtEditAnimationTextCursor = iota
	mtEditAnimationDelPercent = iota
	//
	mtEditAnimationLength = iota
)

// Material::Edit
//
// Material theme editable text
type (
	MTEdit struct {
		//
		VoidNode
		styleStore
		rendererStore
		//
		mtColorSingle
		studio     *gcore.Studio
		textCursor *gcore.Switching
		delPercent *mtEditDeleteActor
		//
		ctrl bool
		//
		align       gcore.Align
		text        string
		editingRune rune
		inactive    bool
		//
		onChange            MTEditChange
		cursorEnter, active bool
	}
	// If text changed it occur
	MTEditChange func(self *MTEdit, text string)
	MTEditFocus  func(self *MTEdit, focus bool)
)

func (s *MTEdit) BaseRender(subimg *image.RGBA) {
	var baseColor, mainColor = s.GetMaterialColor().Color()
	var ctx = createContext(subimg)
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

func (s *MTEdit) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

// GUMIFunction / GUMIInit 					-> Define
func (s *MTEdit) GUMIInit() {
	s.studio = gcore.Animation.Studio(mtEditAnimationLength)
	s.textCursor = s.studio.Set(mtEditAnimationTextCursor, &gcore.Switching{
		Interval: mtEditAnimationTextCursorInterval,
	}).(*gcore.Switching)
	s.delPercent = s.studio.Set(mtEditAnimationDelPercent, newMTEditDeleteActor(100, 50)).(*mtEditDeleteActor)
}

// GUMIFunction / GUMIInfomation 			-> Define
func (s *MTEdit) GUMIInfomation(info Information) {
	if s.studio.Animate(float64(info.Dt)) {
		s.rnode.ThrowCache()
	}
	//
	if s.editingRune == 0 {
		var dels = s.delPercent.Pop()
		if dels > 0 {
			if s.ctrl {
				// ctrl + backspace
				temp := StringControlBackSpace(s.text, dels)
				if s.text != temp {
					s.emitChange()
				}
				s.text = temp
			} else {
				// backspace
				temp := StringBackSpace(s.text, dels)
				if s.text != temp {
					s.emitChange()
				}
				s.text = temp
			}
		}

	}
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *MTEdit) GUMIStyle(style *Style) {
	s.style = style
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
func (s *MTEdit) GUMIRenderSetup(man *renderline.Manager, parent renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent, nil)
	s.rnode.SetJob(s)
}

// GUMIEventer / GUMIHappen					-> Define
func (s *MTEdit) GUMIHappen(event Event) {
	switch ev := event.(type) {
	case EventKeyPress:
		switch ev.Key {
		case KEY_CONTROL:
			if !s.active {
				return
			}
			s.ctrl = true
		case KEY_BACKSPACE:
			if !s.active {
				return
			}
			if s.editingRune != 0 {
				return
			}
			s.delPercent.Start()
		}
	case EventKeyRelease:
		switch ev.Key {
		case KEY_CONTROL:
			if !s.active {
				return
			}
			s.ctrl = false
		case KEY_BACKSPACE:
			if !s.active {
				return
			}
			s.delPercent.Reset()
		case KEY_MOUSE1:
			if s.cursorEnter {
				if !s.inactive {
					s.active = true
					s.rnode.ThrowCache()
				}
			} else {
				s.active = false
				s.rnode.ThrowCache()
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
	case EventRuneEdit:
		if !s.active {
			return
		}
		s.editingRune = ev.Rune
		s.rnode.ThrowCache()
		s.emitChange()
	case EventRuneComplete:
		if !s.active {
			return
		}
		s.editingRune = 0
		s.text += string(ev.Rune)
		s.rnode.ThrowCache()
		s.emitChange()
	}

}

// fmt.Stringer / String					-> Define
func (s *MTEdit) String() string {
	return fmt.Sprintf("%s(text:%s)", "MTEdit", s.text)
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
	s.rnode.ThrowCache()
	s.emitChange()
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

func (s *MTEdit) emitChange() {
	if s.onChange != nil {
		if s.editingRune != 0 {
			s.onChange(s, s.text+string(s.editingRune))
		} else {
			s.onChange(s, s.text)
		}
	}
}
