package gbasic

import "github.com/GUMI-golang/gumi/gcore"

func (_Tool) MarginMinRegular(min uint16, gumi GUMI) GUMI {
	temp := NMargin0(gcore.RegularBlank(gcore.MinLength(min)))
	temp2 := LinkingFrom(temp, gumi)
	return temp2
}
func (_Tool) MarginMinSymmetry(minHorizontal, minVertical uint16, gumi GUMI) GUMI {
	temp := NMargin0(gcore.SymmetryBlank(gcore.MinLength(minHorizontal), gcore.MinLength(minVertical)))
	return LinkingFrom(temp, gumi)
}
