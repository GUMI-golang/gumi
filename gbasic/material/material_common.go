package material

import (
	"image/color"
)

var Material = _Material{
	DefaultAnimation: _MaterialAnimation{},
	Pallette: _MaterialPallette{
		background: color.RGBA{64, 74, 83, 255},
		silluet:    color.RGBA{54, 64, 73, 200},
		White: &MaterialColor{
			name: "White",
			colorset: [2]color.Color{
				color.RGBA{94, 97, 97, 255},
				color.RGBA{213, 217, 218, 255},
			},
		},
		Red: &MaterialColor{
			name: "Red",
			colorset: [2]color.Color{
				color.RGBA{127, 51, 49, 255},
				color.RGBA{255, 84, 74, 255},
			},
		},
		Green: &MaterialColor{
			name: "Green",
			colorset: [2]color.Color{
				color.RGBA{61, 127, 59, 255},
				color.RGBA{110, 204, 102, 255},
			},
		},
		Blue: &MaterialColor{
			name: "Blue",
			colorset: [2]color.Color{
				color.RGBA{59, 70, 127, 255},
				color.RGBA{118, 139, 255, 255},
			},
		},
		Yellow: &MaterialColor{
			name: "Yellow",
			colorset: [2]color.Color{
				color.RGBA{122, 127, 67, 255},
				color.RGBA{231, 235, 118, 255},
			},
		},
	},
}

type _Material struct {
	DefaultAnimation _MaterialAnimation
	Pallette         _MaterialPallette
}
