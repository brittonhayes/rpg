package character

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/brittonhayes/rpg/stat"
)

var (
	ErrInvalidAttack = errors.New("attack kind does not match")
)

type Rank int

type Character struct {
	Name      string
	Rank      Rank
	Stats     map[string]*Stat
	Attacks   Attacks
	Abilities []Ability
	Inventory []Item
	Playable  bool
}

type Option func(*Character)

func New(options ...Option) *Character {
	var (
		defaultName  = "Character 1"
		defaultRank  = Rank(0)
		defaultStats = map[string]*Stat{
			stat.Health:  NewStat(stat.Full),
			stat.Stamina: NewStat(stat.Full),
		}
		defaultAttacks = Attacks{
			Light: NewAttack("Basic Light", LightAttack, 5.00),
			Heavy: NewAttack("Basic Heavy", HeavyAttack, 8.00),
		}
	)

	p := &Character{
		Name:      defaultName,
		Rank:      defaultRank,
		Stats:     defaultStats,
		Attacks:   defaultAttacks,
		Abilities: nil,
		Inventory: nil,
		Playable:  false,
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

func IsPlayer() Option {
	return func(character *Character) {
		character.Playable = true
	}
}

func WithAttacks(light, heavy *Attack) Option {
	return func(character *Character) {
		if light.Kind != LightAttack || heavy.Kind != HeavyAttack {
			log.Fatalln(ErrInvalidAttack)
		}

		character.Attacks.Light = light
		character.Attacks.Heavy = heavy
	}
}

func WithName(name string) Option {
	return func(character *Character) {
		character.Name = name
	}
}

func WithStamina(stamina float64) Option {
	return func(character *Character) {
		character.Stats[stat.Stamina] = &Stat{value: stamina}
	}
}

func WithHealth(health float64) Option {
	return func(character *Character) {
		character.Stats[stat.Health] = &Stat{value: health}
	}
}

func WithArmor(armor float64) Option {
	return func(character *Character) {
		character.Stats[stat.Armor] = &Stat{value: armor}
	}
}

func WithRank(rank int) Option {
	return func(character *Character) {
		character.Rank = Rank(rank)
	}
}

func WithCustomStats(stats map[string]*Stat) Option {
	return func(character *Character) {
		for key, value := range stats {
			character.Stats[key] = value
		}
	}
}

func (c *Character) Stat(name string) *Stat {
	for key, value := range c.Stats {
		if key == strings.ToTitle(name) {
			return value
		}
	}

	return &Stat{value: 0.00}
}

func (c *Character) Health() *Stat {
	for key, value := range c.Stats {
		if key == stat.Health {
			return value
		}
	}

	return &Stat{value: 0.00}
}

func (r Rank) String() string {
	return fmt.Sprintf("Rank: %d", r)
}

func (c *Character) RankTo(rank int) {
	c.Rank = Rank(rank)
}

func (c *Character) RankUp(rank int) {
	c.Rank += Rank(rank)
}
