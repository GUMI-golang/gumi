package gcore

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"strings"
)

func StringBackSpace(str string, count int) string {
	temp := []rune(str)
	templen := len(temp)
	if count > templen {
		count = templen
	}
	return string(temp[:templen-count])
}
func StringControlBackSpace(str string, count int) (res string) {
	temp := []rune(str)
	to := len(temp) - 1
	if to < 0 {
		return ""
	}
	//
	for i := 0; i < count || to != 0; i++ {
		if temp[to] == ' ' {
			for i := to; i >= 0; i-- {
				to = i
				if temp[i] != ' ' {
					break
				}
			}
		}
		for i := to; i >= 0; i-- {
			to = i
			if temp[i] == ' ' {
				to += 1
				break
			}
		}
	}
	return string(temp[:to])
}

func HexToColor(h string) color.Color {
	if !strings.HasPrefix(h, "#") {
		return color.Transparent
	}
	h = strings.TrimPrefix(h, "#")

	if len(h) == 4 {
		h = fmt.Sprintf("%c%c%c%c%c%c%c%c", h[0], h[0], h[1], h[1], h[2], h[2], h[3], h[3])
	}
	d, _ := hex.DecodeString(h)
	if len(d) < 4 {
		return color.Transparent
	}
	return color.RGBA{
		uint8(d[0]), uint8(d[1]), uint8(d[2]), uint8(d[3]),
	}
}
func HexFromColor(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x%02x",
		uint8(r>>8),
		uint8(g>>8),
		uint8(b>>8),
		uint8(a>>8),
	)
}
