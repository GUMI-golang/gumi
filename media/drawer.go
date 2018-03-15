package media

import (
	"image"
	"image/draw"
)

type Drawer interface {
	Bound() image.Rectangle
	Effector
}
type Effector interface {
	Draw(dst draw.Image)
}
