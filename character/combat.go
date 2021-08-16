package character

import (
	"github.com/brittonhayes/rpg/stat"
)

type AttackType int

type Attacks struct {
	Light *Attack
	Heavy *Attack
}

type Attack struct {
	Name   string
	Kind   AttackType
	Damage float64
}

const (
	LightAttack AttackType = iota + 1
	HeavyAttack
)

func NewAttack(name string, kind AttackType, damage float64) *Attack {
	return &Attack{Name: name, Kind: kind, Damage: damage}
}

func (c *Character) IsDead() bool {
	return c.Stats[stat.Health].value <= 0.00
}

func (c *Character) Attack(target *Character, attack *Attack) {
	if attack == nil {
		attack = &Attack{Name: "Basic", Damage: 5.00}
	}

	if target.Stats[stat.Health] == nil {
		return
	}

	dealt := attack.Damage
	for ok := true; ok; ok = dealt >= 0.00 {
		if target.Stats[stat.Armor] != nil {
			target.Stats[stat.Armor].Down(attack.Damage)
			dealt -= attack.Damage
			if dealt <= 0.00 {
				break
			}
		}

		target.Stats[stat.Health].Down(attack.Damage)
		dealt -= attack.Damage
		if dealt <= 0.00 {
			break
		}
	}
}
