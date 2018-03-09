package gumi

import (
	"fmt"
	"github.com/GUMI-golang/gumi/gcore"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"strconv"
)

var (
	rulerColor = color.White
	rulerDash2 = 4.
	rulerDash1 = 2.
	rulerWidth = 10.
)

type _Ruler struct {
	Graduation
	Grid
	Hint
	Dashgrid
}
type Graduation struct{}

func (Graduation) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, rulerWidth, f)
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (Graduation) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, rulerWidth)
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}

type Grid struct{}

func (Grid) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, float64(context.Height()))
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (Grid) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, float64(context.Width()), f)
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}

type Hint struct{}

func (Hint) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := float64(pivot); f <= float64(context.Width()); f += float64(pivot) {
			txt := strconv.FormatInt(int64(f), 10)
			w, _ := context.MeasureString(txt)
			context.DrawString(txt, f-w, float64(style.Default.Font.FontHeight().Round()))
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (Hint) Horizontal(pivot float64) Drawer {

	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for f := float64(pivot); f <= float64(context.Height()); f += float64(pivot) {
			context.DrawString(strconv.FormatInt(int64(f), 10), 0, f)
		}

		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}

type Dashgrid struct{}

func (Dashgrid) Vertical(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)

		context.SetDash(rulerDash1, rulerDash2)
		for f := 0.0; f <= float64(context.Width()); f += float64(pivot) {
			context.DrawLine(f, 0, f, float64(context.Height()))
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (Dashgrid) Horizontal(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		context.SetDash(rulerDash1, rulerDash2)
		for f := 0.0; f <= float64(context.Height()); f += float64(pivot) {
			context.DrawLine(0, f, float64(context.Width()), f)
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}

func (_Ruler) Size() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)

		context.DrawLine(0, rulerWidth, float64(context.Width()), rulerWidth)
		context.DrawLine(0, rulerWidth, rulerWidth/2, rulerWidth/2)
		context.DrawLine(0, rulerWidth, rulerWidth/2, rulerWidth/2*3)
		context.DrawLine(float64(context.Width()), rulerWidth, float64(context.Width())-rulerWidth/2, rulerWidth/2)
		context.DrawLine(float64(context.Width()), rulerWidth, float64(context.Width())-rulerWidth/2, rulerWidth/2*3)
		//
		context.DrawLine(rulerWidth, 0, rulerWidth, float64(context.Height()))
		context.DrawLine(rulerWidth, 0, rulerWidth/2, rulerWidth/2)
		context.DrawLine(rulerWidth, 0, rulerWidth/2*3, rulerWidth/2)
		context.DrawLine(rulerWidth, float64(context.Height()), rulerWidth/2, float64(context.Height())-rulerWidth/2)
		context.DrawLine(rulerWidth, float64(context.Height()), rulerWidth/2*3, float64(context.Height())-rulerWidth/2)
		//
		w, h := context.MeasureString(strconv.FormatInt(int64(context.Width()), 10))
		context.DrawString(strconv.FormatInt(int64(context.Width()), 10), float64(context.Width()/2)-w/2, rulerWidth+h/4)
		w, h = context.MeasureString(strconv.FormatInt(int64(context.Height()), 10))
		context.DrawString(strconv.FormatInt(int64(context.Height()), 10), rulerWidth, float64(context.Height()/2)+h/4)
		//
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (_Ruler) Proportion() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)

		w := context.Width()
		h := context.Height()
		//
		gcd := int(gcore.GCD(int64(w), int64(h)))
		txt := fmt.Sprintf("%d : %d", w/gcd, h/gcd)
		//
		context.DrawLine(0, 0, float64(w), float64(h))
		bdw, bdh := context.MeasureString(txt)
		context.DrawString(txt, float64(w)/2-bdw/2, float64(h)/2+bdh/4)
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (_Ruler) Screen() Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for _, v := range DefinedResolutions.Smaller(context.Width(), context.Height()) {
			context.DrawRectangle(0, 0, float64(v.Width), float64(v.Height))
			w, _ := context.MeasureString(v.Name[0])
			context.DrawString(v.Name[0], float64(v.Width)-w-5, float64(v.Height)-5)
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
func (_Ruler) Dots(pivot float64) Drawer {
	return FunctionDrawer{func(context *gg.Context, style *Style) image.Rectangle {
		style.useContext(context)
		defer style.releaseContext(context)
		context.SetColor(rulerColor)
		for x := 0.; x <= float64(context.Width()); x += float64(pivot) {
			for y := 0.; y <= float64(context.Height()); y += float64(pivot) {
				context.DrawPoint(x, y, style.Default.LineWidth)
			}
		}
		context.Stroke()
		return image.Rect(0, 0, context.Width(), context.Height())
	}}
}
