package gcore

func Clamp(i float64, min, max float64) float64 {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

func GCD(a, b int64) (n int64){
	if a < b{
		a, b = b, a
	}
	for b != 0{
		n = a % b
		a = b
		b = n
	}
	return
}