package gumi

import "image/color"

type (
	MTColor struct {

	}
	MTColorChange func(self *MTColor, c color.Color)
)