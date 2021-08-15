package game

import (
	"github.com/brittonhayes/rpg/character"
)

type Scene struct {
	ID          int
	Player      character.Player
	Environment Environment
	Events      []Event
}

type Event struct {
	Name        string
	Description string
}

type Environment struct {
	Name        string
	Description string
	Modifiers   []character.Modifier
}
