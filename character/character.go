package character

import (
	"errors"
	"log"
	"strings"
	"sync/atomic"

	"github.com/brittonhayes/rpg/stat"
)

var (
	ErrInvalidAttack = errors.New("attack kind does not match")
)

var (
	defaultID    uint64
	defaultName  = "Character 1"
	defaultRank  = uint64(1)
	defaultStats = map[string]*Stat{
		stat.Stamina: NewStat(stat.Full),
	}
	defaultAttacks = Attacks{
		Light: NewAttack("Basic Light", LightAttack, 5.00),
		Heavy: NewAttack("Basic Heavy", HeavyAttack, 8.00),
	}
)

type Character struct {
	ID        uint64
	Name      string
	Rank      uint64
	Health    float64
	Armor     float64
	Stats     map[string]*Stat
	Attacks   Attacks
	Abilities []Ability
	Inventory []Item
	Playable  bool
}

type Option func(*Character)

func New(options ...Option) *Character {

	p := &Character{
		ID:        atomic.AddUint64(&defaultID, 1),
		Name:      defaultName,
		Health:    100.00,
		Armor:     0.00,
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
		character.Health = health
	}
}

func WithArmor(armor float64) Option {
	return func(character *Character) {
		character.Armor = armor
	}
}

func WithRank(rank int) Option {
	return func(character *Character) {
		character.Rank = uint64(rank)
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

func (c *Character) IsCalled() string {
	return c.Name
}

func (c *Character) IsPlayable() bool {
	return c.Playable
}

func (c *Character) RankTo(rank uint64) {
	c.Rank = rank
}

func (c *Character) RankUp(rank uint64) {
	c.Rank += rank
}
