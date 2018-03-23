package material

import (
	"github.com/GUMI-golang/gumi/media"
	"image"
	"image/color"
)

type _MaterialPallette struct {
	background color.Color
	silluet    color.Color

	White  *MaterialColor
	Red    *MaterialColor
	Green  *MaterialColor
	Blue   *MaterialColor
	Yellow *MaterialColor
}

func (s _MaterialPallette) BackgroundColor() color.Color {
	return s.background
}
func (s _MaterialPallette) BackgroundImage() image.Image {
	return image.NewUniform(s.background)
}
func (s _MaterialPallette) BackgroundDrawer() media.Drawer {
	return media.NewUniform(s.background)
}
func (s _MaterialPallette) SilluetColor() color.Color {
	return s.silluet
}
func (s _MaterialPallette) SilluetImage() image.Image {
	return image.NewUniform(s.silluet)
}
func (s _MaterialPallette) SilluetDrawer() media.Drawer {
	return media.NewUniform(s.silluet)
}

type MaterialColor struct {
	name     string
	colorset [2]color.Color
}

func (s MaterialColor) String() string {
	return s.name
}
func (s MaterialColor) Color() (base, main color.Color) {
	return s.BaseColor(), s.MainColor()
}
func (s MaterialColor) Image() (base, main image.Image) {
	return s.BaseImage(), s.MainImage()
}
func (s MaterialColor) Drawer() (base, main media.Drawer) {
	return s.BaseDrawer(), s.MainDrawer()
}
func (s MaterialColor) BaseColor() (base color.Color) {
	return s.colorset[0]
}
func (s MaterialColor) BaseImage() (base image.Image) {
	return image.NewUniform(s.colorset[0])
}
func (s MaterialColor) BaseDrawer() (base media.Drawer) {
	return media.NewUniform(s.colorset[0])
}
func (s MaterialColor) MainColor() (main color.Color) {
	return s.colorset[1]
}
func (s MaterialColor) MainImage() (main image.Image) {
	return image.NewUniform(s.colorset[1])
}
func (s MaterialColor) MainDrawer() (main media.Drawer) {
	return media.NewUniform(s.colorset[1])
}
