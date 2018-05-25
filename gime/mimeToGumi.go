package gime

import (
	"github.com/GUMI-golang/gorat/textrat"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"strings"
	"golang.org/x/image/font"
	"github.com/golang/freetype/truetype"
)

// can be
// - <nil>
// - error
// - string
//		- text/plain
// - giame.Filler
//		- image/x-uniform
//		- image/png
//		- image/bmp
//		- image/jpeg
//		- image/webp
// - textrat.Font
//		- application/x-font-opentype
//
type Value interface {
}

func IsValidValue(sample interface{}) bool {
	switch sample.(type) {
	case error:
		return true
	case string:
		// text
		return true
	case image.Image:
		// image
		return true
	case textrat.Font:
		// font, application/x-font
		return true
	}
	return false
}
func Parse(mimetext string, data io.Reader) Value {
	main, sub := splitMime(mimetext)
	switch main {
	case "text":
		return parseText(sub, data)
	case "image":
		return parseImage(sub, data)
	case "application":
		return parseApplication(sub, data)
	}
	return errors.WithMessage(UnknownMime, mimetext+" is not support mime type on GUMI")
}
func splitMime(mimetext string) (main, sub string) {
	temp := strings.SplitN(mimetext, "/", 2)
	if len(temp) < 2 {
		return temp[0], ""
	}
	return temp[0], temp[1]

}
func parseText(sub string, data io.Reader) Value {
	bts, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	return string(bts)
}
func parseImage(sub string, data io.Reader) Value {
	switch sub {
	case "png":
		res, err := png.Decode(data)
		if err != nil {
			return err
		}
		return res
	case "jpeg":
		res, err := jpeg.Decode(data)
		if err != nil {
			return err
		}
		return res
	case "bmp":
		res, err := bmp.Decode(data)
		if err != nil {
			return err
		}
		return res
	case "webp":
		res, err := webp.Decode(data)
		if err != nil {
			return err
		}
		return res
	}
	if sub == "jpeg" {
		return errors.WithMessage(UnknownMime, "mime 'image/"+sub+"' is unknown, instead image/jpg, use image/jpeg")
	}
	return errors.WithMessage(UnknownMime, "mime 'image/"+sub+"' is unknown")
}
func parseApplication(sub string, data io.Reader) Value {
	switch sub {
	case "x-font-opentype":
		bts, err := ioutil.ReadAll(data)
		if err != nil {
			return err
		}
		f, err := truetype.Parse(bts)
		if err != nil {
			return err

		}
		return textrat.NewVectorFont(f, 12, font.HintingFull)
	}
	return errors.WithMessage(UnknownMime, "mime 'application/"+sub+"' is unknown")
}