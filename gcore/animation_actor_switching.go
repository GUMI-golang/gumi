package gcore

type Switching struct {
	Switch            bool
	Current, Interval float64
}

func (s *Switching) Reset() {
	s.Switch = false
	s.Current = 0
}
func (s *Switching) Animate(delta float64) bool {
	s.Current += delta
	temp := (int(s.Current)/int(s.Interval))%2 == 1
	if temp != s.Switch {
		s.Switch = temp
		return true
	}
	return false
}
