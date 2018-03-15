package gcore

var Animation _Animation

type _Animation struct {
	Functions _AnimationFunctions
}
type _AnimationFunctions struct {
	//
	Quad  AnimFuncsQuad
	Cubic AnimFuncsCubic
	Quart AnimFuncsQuart
	Quint AnimFuncsQuint
}

func (_Animation) PercentingByMillis(millis float64) (delta float64) {
	return 1000 / millis
}
func (_Animation) ReachingBySpeed(width, speedPerSecond float64) (delta float64) {
	return 1 / (width / speedPerSecond)
}
