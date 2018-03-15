package gcore

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Sum(dts []float64) (res float64) {
	for _, v := range dts {
		res += v
	}
	return res
}
func Average(dts []float64) float64 {
	return Sum(dts) / float64(len(dts))
}
