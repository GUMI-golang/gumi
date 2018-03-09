package gcore



type Percenting struct {
	From, Current, Delta, To float64
	Fn                       AnimationFunction
}

func (s *Percenting) Reset() {
	s.From = 0
	s.Current = 0
	s.To = 0
}
func (s *Percenting) Function(Fn AnimationFunction) {
	s.Fn = Fn
}
func (s *Percenting) Request(to float64) {
	if s.To != to {
		s.From = s.Current
		s.To = to
	}
}
func (s *Percenting) Animate(delta float64) bool {
	if s.Current == s.To {
		return false
	}

	if s.To > s.From {
		s.Current += s.Delta * delta / 1000
		if s.Current > s.To {
			s.Current = s.To
		}
	} else if s.To < s.From {
		s.Current -= s.Delta * delta / 1000
		if s.Current < s.To {
			s.Current = s.To
		}
	}
	return true
}
func (s *Percenting) Value() float64 {
	if s.From == s.To {
		return s.To
	}
	return s.Fn(s.Current)
}