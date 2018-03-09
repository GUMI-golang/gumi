package gumi

import (
	"fmt"
	"image"
	"image/color"
	"github.com/iamGreedy/gumi/renderline"
	"github.com/iamGreedy/gumi/gcore"
)

// Actor::Text
//
// AText use for render text
type AText struct {
	VoidNode
	styleStore
	rendererStore
	//
	align     gcore.Align
	text      string
	textColor color.Color
	//
}

// renderline.Job / BaseRender
func (s *AText) BaseRender(subimg *image.RGBA) {
	ctx := createContext(subimg)
	s.style.useContext(ctx)
	defer s.style.releaseContext(ctx)
	ctx.SetColor(s.textColor)
	expectw, expecth := ctx.MeasureString(s.text)
	v, h := gcore.ParseAlign(s.align)
	var drawX, drawY float64
	switch v {
	case gcore.AlignBottom:
		drawY = float64(ctx.Height())
	case gcore.AlignVertical:
		drawY = float64(ctx.Height())/2 + expecth/2
	case gcore.AlignTop:
		drawY = expecth
	}
	switch h {
	case gcore.AlignRight:
		drawX = float64(ctx.Width()) - expectw
	case gcore.AlignHorizontal:
		drawX = float64(ctx.Width())/2 - expectw/2
	case gcore.AlignLeft:
		drawX = 0
	}
	ctx.DrawString(s.text, drawX, drawY - 1)
}

// renderline.Job / DecalRender
func (s *AText) DecalRender(fullimg *image.RGBA) (updated image.Rectangle) {
	return image.ZR
}

// GUMIFunction / GUMIInfomation 			-> Define::Empty
func (s *AText) GUMIInfomation(info Information) {
}

// GUMIFunction / GUMIStyle 				-> Define
func (s *AText) GUMIStyle(style *Style) {
	s.style = style
}

// GUMIFunction / GUMISize 					-> Define
func (s *AText) GUMISize() gcore.Size {
	s.style.Default.Font.Use()
	defer s.style.Default.Font.Release()

	h, v := s.style.Default.Font.CalculateSize(s.text)

	temp := gcore.Size{
		Horizontal: gcore.MinLength(uint16(h)),
		Vertical:   gcore.MinLength(uint16(v)),
	}

	return temp
}

// GUMITree / born 							-> VoidNode::Default

// GUMITree / breed 						-> VoidNode::Default

// GUMITree / parent()						-> VoidNode::Default

// GUMITree / childrun()					-> VoidNode::Default

// GUMIRenderer / GUMIRenderSetup			-> Define
func (s *AText) GUMIRenderSetup(man *renderline.Manager, parent *renderline.Node) {
	s.rmana = man
	s.rnode = man.New(parent)
	s.rnode.Do = s
}


// GUMIEventer / GUMIHappen					-> Define
func (s *AText) GUMIHappen(event Event) {
}

// fmt.Stringer / String					-> Define
func (s *AText) String() string {
	return fmt.Sprintf("%s(text:%s)", "AText", s.text)
}

// Constructor 0
func AText0(str string) *AText {
	temp := &AText{
		text:      str,
		align:     gcore.AlignCenter,
		textColor: color.White,
	}
	return temp
}

// Constructor 1
func AText1(str string, align gcore.Align) *AText {
	temp := &AText{
		text:str,
		align:align,
		textColor:color.White,
	}
	return temp
}

// Constructor 2
func AText2(str string, align gcore.Align, textColor color.Color) *AText {
	temp := &AText{
		text:str,
		align:align,
		textColor:textColor,
	}
	return temp
}

// Method / Set -> SetText(...)
func (s *AText) Set(text string) {
	s.SetText(text)
}

// Method / Get -> GetText()
func (s *AText) Get() string {
	return s.GetText()
}

// Method / SetText
func (s *AText) SetText(text string) {
	if s.text != text{
		s.text = text
		s.rnode.ThrowCache()
	}
}

// Method / GetText
func (s *AText) GetText() string {
	return s.text
}

// Method / SetAlign
func (s *AText) SetAlign(align gcore.Align) {
	if s.align != align{
		s.align = align
		s.rnode.ThrowCache()
	}
}

// Method / GetAlign
func (s *AText) GetAlign() gcore.Align {
	return s.align
}

// Method / SetColor
func (s *AText) SetColor(textColor color.Color) {
	r1,g1,b1,a1 := s.textColor.RGBA()
	r2,g2,b2,a2 := textColor.RGBA()
	if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2{
		s.textColor = textColor
		s.rnode.ThrowCache()
	}
}

// Method / GetColor
func (s *AText) GetColor() color.Color {
	return s.textColor
}
