package gumi

import (
	"image/color"
)

type Style map[StyleData]interface{}

func DefaultStyle() Style {
	temp := Style(map[StyleData]interface{}{
		STYLE_LineWidth:  10,
		STYLE_Font:       nil,
		STYLE_TextHeight: 12.,
		STYLE_TextColor:  color.Black,
	})
	return temp
}
