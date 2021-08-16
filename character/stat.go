package character

import (
	"fmt"
)

type Stat struct {
	value float32
}

func NewStat(value float32) *Stat {
	return &Stat{value: value}
}

type Modifier interface {
	To(value float32)
	Up(value float32)
	Down(value float32)
}

func (s *Stat) String() string {
	return fmt.Sprintf("%v", s.value)
}

func (s *Stat) Up(value float32) {
	if s.value+value <= 100.00 {
		s.value += value
		return
	}

	s.value = 100.00
}

func (s *Stat) To(value float32) {
	if value <= 100.00 || value >= 0.00 {
		s.value = value
	}
}

func (s *Stat) Down(value float32) {
	if s.value-value >= 0.00 {
		s.value -= value
		return
	}

	s.value = 0.00
}
