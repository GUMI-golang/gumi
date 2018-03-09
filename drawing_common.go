package gumi

import (
	"github.com/fogleman/gg"
	"math"
	"fmt"
	"github.com/iamGreedy/gumi/gcore"
)


type Drawer interface {
	Draw(context *gg.Context, style *Style)
}
type DrawerWithInformation interface {
	Drawer
	Inform(info Information) (changed bool)
}
type FunctionDrawer struct {
	fn func(context *gg.Context, style *Style)
}
func (s FunctionDrawer)Draw(context *gg.Context, style *Style)  {
	s.fn(context, style)
}

var Drawing _Drawing
type _Drawing struct {
	Ruler _Ruler
}

const FPSPos = 10
func (_Drawing) FPS() Drawer {
	return &fpsDrawer{}
}

const fpsDrawerHistory = 32

type fpsDrawer struct {
	dts [fpsDrawerHistory]float64
	i int
}
func (s *fpsDrawer ) Draw(context *gg.Context, style *Style)  {
	style.useContext(context)
	defer style.releaseContext(context)
	//
	context.SetColor(rulerColor)

	avg := Clamp(gcore.Average(s.dts[:]), 0.001, math.MaxFloat64)
	txt := fmt.Sprintf("FPS : %.2f - AVG : %2.5f", 1000 / float64(avg), avg)
	w := float64(context.Width())
	mw, mh := context.MeasureString(txt)
	context.DrawString(txt, w - FPSPos - mw, FPSPos + mh)
	context.Stroke()
}
func (s *fpsDrawer) Inform(info Information) (changed bool) {
	s.dts[s.i] = float64(info.Dt)
	s.i = (s.i + 1) % fpsDrawerHistory
	if s.i == 0{
		return true
	}
	return false
}