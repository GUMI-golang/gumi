package gumi

import (
	"github.com/fogleman/gg"
	"github.com/iamGreedy/freetype"
	"golang.org/x/image/font/gofont/goregular"
	"sync"
	"github.com/iamGreedy/freetype/truetype"
	"github.com/GUMI-golang/gumi/gcore"
)

type Style struct {
	Default  StyleDefault
	Map      map[string]interface{}
}

func (s *Style) Equare(e *Style) bool {
	if s == e {
		return true
	}
	if s.Default.Font != e.Default.Font{
		return false
	}
	if s.Default.LineWidth != e.Default.LineWidth{
		return false
	}
	for k, v := range s.Map{
		if ev, eok := e.Map[k]; !eok || ev != v{
			return false
		}
	}
	return true
}

func (s *Style) useContext(ctx *gg.Context) {
	s.Default.Font.Use()
	ctx.SetLineWidth(s.Default.LineWidth)
	ctx.SetFontFace(s.Default.Font.Face())
}
func (s *Style) releaseContext(ctx *gg.Context) {
	s.Default.Font.Release()
}

type StyleDefault struct {
	Font *gcore.Font
	//
	LineWidth float64
}

var (
	defaultStyleSingleton      *Style = nil
	defaultStyleSingletonMutex        = new(sync.RWMutex)
)

func DefaultStyle() *Style {
	if defaultStyleSingleton == nil {
		defaultStyleSingletonMutex.Lock()
		defer defaultStyleSingletonMutex.Unlock()
		f, _ := freetype.ParseFont(goregular.TTF)
		temp := &Style{
			Default: StyleDefault{
				Font:      gcore.NewFont(f, 12),
				LineWidth: 1,
			},
			Map: map[string]interface{}{
				"dummy": nil,
			},
		}
		defaultStyleSingleton = temp
	} else {
		defaultStyleSingletonMutex.RLock()
		defer defaultStyleSingletonMutex.RUnlock()
	}
	return defaultStyleSingleton
}
func ModifyDefaultStyle(font *truetype.Font, size float64) {
	defaultStyleSingletonMutex.Lock()
	defer defaultStyleSingletonMutex.Unlock()
	if defaultStyleSingleton == nil {
		temp := &Style{
			Default: StyleDefault{
				Font:      gcore.NewFont(font, size),
				LineWidth: 1,
			},
			Map: map[string]interface{}{
				"dummy": nil,
			},
		}
		defaultStyleSingleton = temp
	} else {
		defaultStyleSingleton.Default.Font = gcore.NewFont(font, size)
	}
}
