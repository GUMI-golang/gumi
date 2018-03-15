package media

import (
	"bytes"
	"github.com/GUMI-golang/gumi/gcore"
	"image"
)

func MustImageDecode(data []byte) image.Image {
	return gcore.MustValue(ImageDecode(data)).(image.Image)
}
func ImageDecode(data []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewBuffer(data))
	return img, err
}
