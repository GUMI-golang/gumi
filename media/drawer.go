package media

import (
	"image/draw"
	"image"
)

type Drawer interface{
	Bound() image.Rectangle
	Effector
}
type Effector interface {
	Draw(dst draw.Image)
}