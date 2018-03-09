package gcore

type AnimationFunction func(t float64) float64


func (s _AnimationFunctions) Default(t float64) float64 {
	return s.Quad.Easing(t)
}
func (_AnimationFunctions) Linear(t float64) float64 {
	return t
}

type AnimFuncsQuad struct {
}

func (AnimFuncsQuad) EasingIn(t float64) float64 {
	return t * t
}
func (AnimFuncsQuad) EasingOut(t float64) float64 {
	return t * (2 - t)
}
func (AnimFuncsQuad) Easing(t float64) float64 {
	if t < .5 {
		return 2 * t * t
	}
	return -1 + (4-2*t)*t
}

type AnimFuncsCubic struct {
}

func (AnimFuncsCubic) EasingIn(t float64) float64 {
	return t * t * t
}
func (AnimFuncsCubic) EasingOut(t float64) float64 {
	return (t-1)*t*t + 1
}
func (AnimFuncsCubic) Easing(t float64) float64 {
	if t < .5 {
		return 4 * t * t * t
	}
	return (t-1)*(2*t-2)*(2*t-2) + 1
}

type AnimFuncsQuart struct {
}

func (AnimFuncsQuart) EasingIn(t float64) float64 {
	return t * t * t * t
}
func (AnimFuncsQuart) EasingOut(t float64) float64 {
	return 1 - (t-1)*t*t*t
}
func (AnimFuncsQuart) Easing(t float64) float64 {
	if t < .5 {
		return 8 * t * t * t * t
	}
	return 1 - 8*(t-1)*t*t*t
}

type AnimFuncsQuint struct {
}

func (AnimFuncsQuint) EasingIn(t float64) float64 {
	return t * t * t * t * t
}
func (AnimFuncsQuint) EasingOut(t float64) float64 {
	return 1 + (t-1)*t*t*t*t
}
func (AnimFuncsQuint) Easing(t float64) float64 {
	if t < .5 {
		return 16 * t * t * t * t * t
	}
	return 1 + 16*(t-1)*t*t*t*t
}

