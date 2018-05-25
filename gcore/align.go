package gcore

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

const (
	AlignTop        Align = 0x00
	AlignVertical   Align = 0x01
	AlignBottom     Align = 0x02
	AlignLeft       Align = 0x00
	AlignHorizontal Align = 0x10
	AlignRight      Align = 0x20
	//

	AlignCenter = AlignHorizontal | AlignVertical
)

var re_align = regexp.MustCompile(`^\((?P<hori>l|L|left|Left|LEFT|h|H|horizontal|Horizontal|HORIZONTAL|r|R|right|Right|RIGHT|),(?P<vert>t|T|top|Top|TOP|v|V|vertical|Vertical|VERTICAL|b|B|bottom|Bottom|BOTTOM|)\)$`)

type Align uint8

func (s Align) String() string {
	return MarshalAlign(s)
}

func SplitAlign(a Align) (v Align, h Align) {
	if a&AlignBottom == AlignBottom {
		v = AlignBottom
	} else if a&AlignVertical == AlignVertical {
		v = AlignVertical
	} else {
		v = AlignTop
	}
	if a&AlignRight == AlignRight {
		h = AlignRight
	} else if a&AlignHorizontal == AlignHorizontal {
		h = AlignHorizontal
	} else {
		h = AlignLeft
	}
	return

}
func MarshalAlign(a Align) string {
	v, h := SplitAlign(a)
	var s = "("
	switch h {
	case AlignLeft:
		s += "left"
	case AlignHorizontal:
		s += "horizontal"
	case AlignRight:
		s += "right"
	}
	s += ", "
	switch v {
	case AlignTop:
		s += "top"
	case AlignVertical:
		s += "vertical"
	case AlignBottom:
		s += "bottom"
	}
	s += ")"

	return s
}
func UnmarshalAlign(s string) (Align, error) {
	s = strings.Replace(s, " ", "", -1)
	res := re_align.FindStringSubmatch(s)
	if len(res) == 0 {
		return 0, errors.New("Unmatched align")
	}
	//
	var a Align
	switch strings.ToLower(res[2]) {
	case "t":
		fallthrough
	case "top":
		a |= AlignTop
	case "v":
		fallthrough
	case "vertical":
		a |= AlignVertical
	case "b":
		fallthrough
	case "bottom":
		a |= AlignBottom
	default:

	}
	switch strings.ToLower(res[1]) {
	case "l":
		fallthrough
	case "left":
		a |= AlignLeft
	case "h":
		fallthrough
	case "horizontal":
		a |= AlignHorizontal
	case "r":
		fallthrough
	case "right":
		a |= AlignRight
	default:

	}
	return a, nil
}
