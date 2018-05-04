package gcore

import (
	"errors"
	"fmt"
	"image/color"
	"regexp"
	"strconv"
)

var (
	re_color_rgb      = regexp.MustCompile(`^#(?P<r>[0-9a-fA-F])(?P<g>[0-9a-fA-F])(?P<b>[0-9a-fA-F])$`)
	re_color_rgba     = regexp.MustCompile(`^#(?P<r>[0-9a-fA-F])(?P<g>[0-9a-fA-F])(?P<b>[0-9a-fA-F])(?P<a>[0-9a-fA-F])$`)
	re_color_rrggbb   = regexp.MustCompile(`^#(?P<r>[0-9a-fA-F]{2})(?P<g>[0-9a-fA-F]{2})(?P<b>[0-9a-fA-F]{2})$`)
	re_color_rrggbbaa = regexp.MustCompile(`^#(?P<r>[0-9a-fA-F]{2})(?P<g>[0-9a-fA-F]{2})(?P<b>[0-9a-fA-F]{2})(?P<a>[0-9a-fA-F]{2})$`)
)

const (
	shortToLong = 0x11
)

func sameUnder4Upper4(val uint8) bool {
	return val&0x0F == val>>4
}
func MarshalColor(c color.Color) string {
	r, g, b, a := c.RGBA()
	r8 := uint8(r)
	g8 := uint8(g)
	b8 := uint8(b)
	a8 := uint8(a)
	if a8 == 0xFF {
		if sameUnder4Upper4(r8) && sameUnder4Upper4(g8) && sameUnder4Upper4(b8) {
			return fmt.Sprintf("#%1x%1x%1x", r8&0x0F, g8&0x0F, b8&0x0F)
		} else {
			return fmt.Sprintf("#%02x%02x%02x", r8, g8, b8)
		}
	} else {
		if sameUnder4Upper4(r8) && sameUnder4Upper4(g8) && sameUnder4Upper4(b8) && sameUnder4Upper4(a8) {
			return fmt.Sprintf("#%1x%1x%1x%1x", r8&0x0F, g8&0x0F, b8&0x0F, a8&0x0F)
		} else {
			return fmt.Sprintf("#%02x%02x%02x%02x", r8, g8, b8, a8)
		}
	}
	return ""
}
func UnmarshalColor(s string) (color.Color, error) {
	if res := re_color_rgb.FindStringSubmatch(s); len(res) > 0 {
		return color.RGBA{
			R: uint8(MustValue(strconv.ParseUint(res[1], 16, 8)).(uint64)) * shortToLong,
			G: uint8(MustValue(strconv.ParseUint(res[2], 16, 8)).(uint64)) * shortToLong,
			B: uint8(MustValue(strconv.ParseUint(res[3], 16, 8)).(uint64)) * shortToLong,
			A: 0xFF,
		}, nil
	}
	if res := re_color_rgba.FindStringSubmatch(s); len(res) > 0 {
		return color.RGBA{
			R: uint8(MustValue(strconv.ParseUint(res[1], 16, 8)).(uint64)) * shortToLong,
			G: uint8(MustValue(strconv.ParseUint(res[2], 16, 8)).(uint64)) * shortToLong,
			B: uint8(MustValue(strconv.ParseUint(res[3], 16, 8)).(uint64)) * shortToLong,
			A: uint8(MustValue(strconv.ParseUint(res[4], 16, 8)).(uint64)) * shortToLong,
		}, nil
	}
	if res := re_color_rrggbb.FindStringSubmatch(s); len(res) > 0 {
		return color.RGBA{
			R: uint8(MustValue(strconv.ParseUint(res[1], 16, 8)).(uint64)),
			G: uint8(MustValue(strconv.ParseUint(res[2], 16, 8)).(uint64)),
			B: uint8(MustValue(strconv.ParseUint(res[3], 16, 8)).(uint64)),
			A: 0xFF,
		}, nil
	}
	if res := re_color_rrggbbaa.FindStringSubmatch(s); len(res) > 0 {
		return color.RGBA{
			R: uint8(MustValue(strconv.ParseUint(res[1], 16, 8)).(uint64)),
			G: uint8(MustValue(strconv.ParseUint(res[2], 16, 8)).(uint64)),
			B: uint8(MustValue(strconv.ParseUint(res[3], 16, 8)).(uint64)),
			A: uint8(MustValue(strconv.ParseUint(res[4], 16, 8)).(uint64)),
		}, nil
	}
	return color.Black, errors.New("Invalid hexcolor")
}
