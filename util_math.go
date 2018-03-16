package gumi

func Clamp(i float64, min, max float64) float64 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

