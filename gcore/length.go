package gcore

import (
	"fmt"
	"math"
)

type Blank struct {
	L, B, R, T Length
}

func (s *Blank) String() string {
	return fmt.Sprintf(
		"[L:%v, R:%v, T:%v, B:%v]", s.L, s.R, s.T, s.B,
	)
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

type Size struct {
	Vertical   Length
	Horizontal Length
}

func (s Size) String() string {
	return fmt.Sprintf(
		"[Horizontal:%v, Verical:%v]", s.Horizontal, s.Vertical,
	)
}

type Length struct {
	Min, Max uint16
}

func (s Length) String() string {
	return fmt.Sprintf("[Min:%d, Max:%d]", s.Min, s.Max)
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
