package gcore

import (
	"fmt"
	"strings"
	"regexp"
	"github.com/pkg/errors"
)


type Size struct {
	Vertical   Length
	Horizontal Length
}

func (s Size) String() string {
	return MarshalSize(s)
}

var (
	ZEROSIZE = Size{MINLENGTH,MINLENGTH,}
	AUTOSIZE = Size{AUTOLENGTH,AUTOLENGTH,}
)
//([:], [:])
var re_size = regexp.MustCompile(`^\((?P<hori>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)]),(?P<vert>\[([+]?[1-9]\d*|0|):([+]?[1-9]\d*|0|)])\)$`)
func MarshalSize(s Size) string {
	return fmt.Sprintf(
		"(%s, %s)", MarshalLength(s.Horizontal), MarshalLength(s.Vertical),
	)
}
func UnmarshalSize(s string) (Size, error) {
	s = strings.Replace(s, " ", "", -1)
	m := re_size.FindStringSubmatch(s)
	if len(m) > 0{
		var vert, hori Length
		var err error
		for i, v := range re_size.SubexpNames(){
			if v == "hori"{
				hori, err = UnmarshalLength(m[i])
				if err != nil {
					return Size{}, err
				}
			}
			if v == "vert"{
				vert, err = UnmarshalLength(m[i])
				if err != nil {
					return Size{}, err
				}
			}
		}
		return Size{
			Horizontal:hori,
			Vertical:vert,
		}, nil
	}
	return Size{}, errors.New("Invalid size")
}
