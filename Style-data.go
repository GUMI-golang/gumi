package gumi

import "github.com/pkg/errors"

// [0:16]	: Identifier
// [16:32] 	: Style Datatype
type StyleData uint32

func (s StyleData) Type() StyleDatatype {
	return StyleDatatype(s)
}
func (s StyleData) String() string {
	for k, v := range stylemap {
		if s == v {
			return k
		}
	}
	panic(errors.New("Unknown Style"))
}
