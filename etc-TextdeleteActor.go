package gumi

import "github.com/GUMI-golang/gumi/gcore"

const (
	newMTEditDeleteActor_PMILLIS = 4096
	newMTEditDeleteActor_MAX     = 8
)

type TextDeleteActor struct {
	Thershold float64
	Interval  float64
	sum       float64
	prev      float64
	active    bool
	first     bool
	//
	needDel int32
	p       *gcore.Percenting
}

func NewTextDeleteActor(Thershold float64, Interval float64) *TextDeleteActor {
	return &TextDeleteActor{
		Thershold: Thershold,
		Interval:  Interval,
		p: &gcore.Percenting{
			Delta: gcore.Animation.PercentingByMillis(newMTEditDeleteActor_PMILLIS),
			Fn:    gcore.Animation.Functions.Quart.EasingOut,
		},
	}
}
func (s *TextDeleteActor) Start() {
	s.sum = 0
	s.prev = 0
	s.needDel = 1
	s.active = true
	s.first = true
	s.p.Request(1)
}
func (s *TextDeleteActor) Reset() {
	s.sum = 0
	s.prev = 0
	s.needDel = 0
	s.active = false
	s.first = false
	s.p.Reset()
}

func (s *TextDeleteActor) Animate(delta float64) bool {
	if !s.active {
		return false
	}
	if s.first {
		s.first = false
		return true
	}
	//
	s.sum += delta
	if s.sum > s.Thershold {
		s.p.Animate(delta)
		if s.sum-s.prev > s.Interval {
			s.prev = s.sum
			s.needDel += int32(s.p.Value() * newMTEditDeleteActor_MAX)
			return true
		}
	}
	return false
}
func (s *TextDeleteActor) Pop() int {
	temp := int(s.needDel)
	s.needDel = 0
	return temp
}

func TextComplete(s string, r rune) string {
	return s + string(r)
}