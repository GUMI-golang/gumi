package gcore

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/pkg/errors"
)

type Blank struct {
	L, B, R, T Length
}

func (s Blank) String() string {
	return MarshalBlank(s)
}

func SymmetryBlank(horizontal, vertical Length) Blank {
	return Blank{
		L: horizontal,
		R: horizontal,
		B: vertical,
		T: vertical,
	}
}
func RegularBlank(regular Length) Blank {
	return Blank{
		L: regular,
		R: regular,
		B: regular,
		T: regular,
	}
}


var re_blank = regexp.MustCompile(`^\((?P<l>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)]),(?P<r>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)]),(?P<t>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)]),(?P<b>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)])\)$`)
func MarshalBlank(s Blank) string {
	return fmt.Sprintf(
		"(%s, %s, %s, %s)", MarshalLength(s.L), MarshalLength(s.R), MarshalLength(s.T), MarshalLength(s.B),
	)
}
func UnmarshalBlank(s string) (Blank, error) {
	s = strings.Replace(s, " ", "", -1)
	m := re_blank.FindStringSubmatch(s)
	if len(m) > 0{
		return Blank{
			L:MustValue(UnmarshalLength(m[1])).(Length),
			R:MustValue(UnmarshalLength(m[2])).(Length),
			T:MustValue(UnmarshalLength(m[3])).(Length),
			B:MustValue(UnmarshalLength(m[4])).(Length),
		}, nil
	}
	return Blank{}, errors.New("Invalid size")
}
