package gcore

type Studio struct {
	ani []Actor
}

func (_Animation) Studio(anicount int)  *Studio {
	temp := &Studio{
		ani: make([]Actor, anicount),
	}
	return temp
}
func (s *Studio) Get(idx int) Actor {
	return s.ani[idx]
}
func (s *Studio) Set(idx int, actor Actor) Actor {
	s.ani[idx] = actor
	return actor
}
func (s *Studio) Reset() {
	for _, v := range s.ani {
		v.Reset()
	}
}
func (s *Studio) Animate(delta float64) bool {
	ch := false
	for _, v := range s.ani {
		ch = ch || v.Animate(delta)
	}
	return ch
}

type Actor interface {
	Reset()
	Animate(delta float64) bool
}
