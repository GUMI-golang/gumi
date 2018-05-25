package gumi

import (
	"image/color"
	"github.com/GUMI-golang/giame"
	"github.com/GUMI-golang/gumi/gcore"
)

type Style map[StyleData]interface{}

func DefaultStyle() Style {
	temp := Style(map[StyleData]interface{}{
		STYLE_Background: giame.NewUniformFiller(color.White),
		STYLE_LineWidth:  10,


		STYLE_Font:       giame.DefaultVFont,
		STYLE_TextHeight: 12,
		STYLE_TextColor:  color.White,
		STYLE_TextAlign: gcore.AlignLeft | gcore.AlignTop,
		//

	})
	return temp
}
