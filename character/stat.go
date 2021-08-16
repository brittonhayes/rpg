package character

import (
	"fmt"
)

type Stat struct {
	value float64
}

func NewStat(value float64) *Stat {
	return &Stat{value: value}
}

type Modifier interface {
	To(value float64)
	Up(value float64)
	Down(value float64)
}

func (s *Stat) String() string {
	return fmt.Sprintf("%v", s.value)
}

func (s *Stat) Up(value float64) {
	if s.value+value <= 100.00 {
		s.value += value
		return
	}

	s.value = 100.00
}

func (s *Stat) To(value float64) {
	if value <= 100.00 || value >= 0.00 {
		s.value = value
	}
}

func (s *Stat) Down(value float64) {
	if s.value-value >= 0.00 {
		s.value -= value
		return
	}

	s.value = 0.00
}
