package gcore

import (
	"github.com/pkg/errors"
	"strings"
)

type Axis uint8

func (s Axis) String() string {
	switch s {
	default:
		return "Unknown"
	case AxisHorizontal:
		return "Horizontal"
	case AxisVertical:
		return "Vertical"
	case AxisBoth:
		return "Both"
	}
}

const (
	AxisVertical   Axis = 1 << iota
	AxisHorizontal Axis = 1 << iota
	AxisBoth            = AxisVertical | AxisHorizontal
)

func MarshalAxis(a Axis) string {
	switch a {
	case AxisHorizontal:
		return "Horizontal"
	case AxisVertical:
		return "Vertical"
	case AxisBoth:
		return "Both"
	default:
		panic(errors.New("Unknown axis"))
	}
}
func UnmarshalAxis(s string) (Axis, error) {
	s = strings.Replace(strings.ToLower(s), " ", "", -1)
	switch s {
	case "h":
		fallthrough
	case "horizontal":
		return AxisHorizontal, nil
	case "v":
		fallthrough
	case "vertical":
		return AxisVertical, nil
	case "b":
		fallthrough
	case "both":
		return AxisBoth, nil
	default:
		return 0, errors.New("Unknown axis")
	}
}