package gcore

import "math"

func ToPolar(x, y float64) (r, radian float64) {
	if x == 0 {
		radian = math.Pi / 2
		if y < 0 {
			radian = math.Pi + math.Pi/2
		}
		return math.Abs(y), radian
	}
	if x > 0 {
		radian = math.Atan(y / x)
	} else {
		radian = math.Atan(y/x) - math.Pi
	}
	if radian < 0 {
		radian = math.Pi*2 + radian
	}
	return math.Sqrt(x*x + y*y), radian
}
func ToDegree(rad float64) (deg float64) {
	return rad * 180 / math.Pi
}
func ToRadian(deg float64) (rad float64) {
	return deg * math.Pi / 180
}
func ToZeroOne(rad float64) (zo float64) {
	for ; rad < 0; rad += 2 * math.Pi {
	}
	for ; rad > 2*math.Pi; rad -= 2 * math.Pi {
	}
	return rad / (2 * math.Pi)
}
