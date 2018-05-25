package gcore

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"github.com/pkg/errors"
	"strconv"
)

type Length struct {
	Min, Max uint16
}

func (s Length) String() string {
	return MarshalLength(s)
}

var (
	AUTOLENGTH = Length{
		Min: 0,
		Max: math.MaxUint16,
	}
	MINLENGTH = Length{
		Min: 0,
		Max: 0,
	}
	MAXLENGTH = Length{
		Min: math.MaxUint16,
		Max: math.MaxUint16,
	}
)

func MinLength(min uint16) Length {
	return Length{
		Min: min,
		Max: math.MaxUint16,
	}
}
func MaxLength(max uint16) Length {
	return Length{
		Min: 0,
		Max: max,
	}
}
func FixLength(fix uint16) Length {
	return Length{
		Min: fix,
		Max: fix,
	}
}

//
var re_length = regexp.MustCompile(`^\[(?P<min>[+]?[1-9]\d*|0|):(?P<max>[+]?[1-9]\d*|0|)]$`)
func MarshalLength(l Length) string {
	if l.Min == 0 && l.Max == math.MaxUint16{
		return "[:]"
	}else if l.Min == 0{
		return fmt.Sprintf("[:%d]", l.Max)
	}else if l.Max == math.MaxUint16{
		return fmt.Sprintf("[%d:]", l.Min)
	}
	return fmt.Sprintf("[%d:%d]", l.Min, l.Max)
}
func UnmarshalLength(s string) (Length, error) {
	s = strings.Replace(s, " ", "", -1)
	m := re_length.FindStringSubmatch(s)
	if len(m) > 0{
		var min, max uint16 = 0, math.MaxUint16
		if len(m[1]) != 0 {
			temp, err := strconv.ParseUint(m[1], 10, 16)
			if err != nil {
				return MINLENGTH, err
			}
			min = uint16(temp)
		}
		if len(m[2]) != 0 {
			temp, err := strconv.ParseUint(m[2], 10, 16)
			if err != nil {
				return MINLENGTH, err
			}
			max = uint16(temp)
		}
		return Length{Min:min, Max:max}, nil
	}
	return MINLENGTH, errors.New("Invalid length")
}