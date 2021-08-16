package character

import (
	"fmt"
	"strings"

	"github.com/brittonhayes/rpg/stat"
)

type Rank int

type Character struct {
	Name      string
	Rank      Rank
	Stats     map[string]*Stat
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
	)

	p := &Character{
		Name:      defaultName,
		Rank:      defaultRank,
		Stats:     defaultStats,
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

func WithStamina(stamina float32) Option {
	return func(character *Character) {
		character.Stats[stat.Stamina] = &Stat{value: stamina}
	}
}

func WithHealth(health float32) Option {
	return func(character *Character) {
		character.Stats[stat.Health] = &Stat{value: health}
	}
}

func WithArmor(armor float32) Option {
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

func (c *Character) Attack(target *Character, damage float32) {
	target.hit(damage)
}

func (c *Character) IsDead() bool {
	return c.Stats[stat.Health].value <= 0.00
}

func (c *Character) hit(damage float32) {
	if c.Stats[stat.Health] == nil {
		return
	}

	dealt := damage
	for ok := true; ok; ok = dealt >= 0.00 {
		if c.Stats[stat.Armor] != nil {
			c.Stats[stat.Armor].Down(damage)
			dealt -= damage
			if dealt <= 0.00 {
				break
			}
		}

		c.Stats[stat.Health].Down(damage)
		dealt -= damage
		if dealt <= 0.00 {
			break
		}
	}
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
