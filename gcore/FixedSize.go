package gcore

import (
	"fmt"
	"github.com/pkg/errors"
	"image"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var InvalidSize = FixedSize{math.MinInt32, math.MinInt32}
var re_FixedSize = regexp.MustCompile(`^\((?P<width>[+]?[1-9]\d*|0),(?P<height>[+]?[1-9]\d*|0)\)$`)

type FixedSize struct {
	Width  int
	Height int
}

func (s FixedSize) Rect() image.Rectangle {
	return image.Rect(0, 0, s.Width, s.Height)
}
func (s FixedSize) Size() (w, h int) {
	return s.Width, s.Height
}
func (s FixedSize) String() string {
	if s == InvalidSize {
		return "FixedSize(Invalid)"
	}
	return fmt.Sprintf("(%d, %d)", s.Width, s.Height)
}

func MarshalFixedSize(f FixedSize) string {
	return fmt.Sprintf("(%d, %d)", f.Width, f.Height)
}
func UnmarshalFixedSize(s string) (FixedSize, error) {
	size := FixedSize{}
	s = strings.Replace(s, " ", "", -1)
	res := re_FixedSize.FindStringSubmatch(s)
	if len(res) == 0 {
		if w, h := DefinedResolutions.Get(s); w == 0 && h == 0 {
			return size, errors.New("Unmatched size")
		} else {
			size.Width = w
			size.Height = h
		}
	} else {
		size.Width = MustValue(strconv.Atoi(res[1])).(int)
		size.Height = MustValue(strconv.Atoi(res[2])).(int)
	}
	return size, nil
}
