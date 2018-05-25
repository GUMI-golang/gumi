package gcore

import (
	"image"
	"image/png"
	"os"
)

func Capture(name string, img image.Image) error {
	f, err := os.OpenFile(name+".png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}
