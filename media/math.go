package media

func clamp(val, min, max float64) uint8 {
	if val > max {
		return uint8(max)
	}
	if val < min {
		return uint8(min)
	}
	return uint8(val)
}
