package character

import (
	"testing"

	"github.com/brittonhayes/rpg/color"
	"github.com/brittonhayes/rpg/stat"
	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	t.Run("Test new character", func(t *testing.T) {
		expect := &Character{
			ID:        1,
			Name:      defaultName,
			Health:    100.00,
			Armor:     0.00,
			Rank:      defaultRank,
			Stats:     defaultStats,
			Attacks:   defaultAttacks,
			Abilities: nil,
			Inventory: nil,
			Playable:  false,
			Color:     color.Gray,
		}

		assert.Equal(t, expect, New())
	})

	t.Run("Test character stat", func(t *testing.T) {
		arg := "CHARM"
		field := map[string]*Stat{"CHARM": {value: 99.00}}
		expect := &Stat{
			value: 99.00,
		}

		c := &Character{Stats: field}
		assert.Equal(t, expect, c.Stat(arg))
	})

	t.Run("Test character is player", func(t *testing.T) {
		arg := false
		expect := false

		c := &Character{Playable: arg}
		assert.Equal(t, expect, c.IsPlayable())
	})
}

func TestWithOption(t *testing.T) {
	t.Run("Test character with name", func(t *testing.T) {
		arg := "Mario"
		expect := "Mario"

		c := New(WithName(arg))
		assert.Equal(t, expect, c.Name)
	})

	t.Run("Test with rank", func(t *testing.T) {
		arg := 2
		expect := uint64(2)

		c := New(WithRank(arg))
		assert.Equal(t, expect, c.Rank)
	})

	t.Run("Test with health", func(t *testing.T) {
		arg := 20.00
		expect := 20.00

		c := New(WithHealth(arg))
		assert.Equal(t, expect, c.Health)
	})

	t.Run("Test with armor", func(t *testing.T) {
		arg := 20.00
		expect := 20.00

		c := New(WithArmor(arg))
		assert.Equal(t, expect, c.Armor)
	})

	t.Run("Test with stamina", func(t *testing.T) {
		arg := 20.00
		expect := 20.00

		c := New(WithStamina(arg))
		assert.Equal(t, expect, c.Stats[stat.Stamina].value)
	})

	t.Run("Test with attacks", func(t *testing.T) {
		args := map[string]*Attack{
			"light": defaultAttacks.Light,
			"heavy": defaultAttacks.Heavy,
		}

		expect := map[string]*Attack{
			"light": defaultAttacks.Light,
			"heavy": defaultAttacks.Heavy,
		}

		c := New(WithAttacks(args["light"], args["heavy"]))
		assert.Equal(t, expect["light"], c.Attacks.Light)
		assert.Equal(t, expect["heavy"], c.Attacks.Heavy)
	})

	t.Run("Test with custom stats", func(t *testing.T) {
		arg := defaultStats
		expect := defaultStats

		c := New(WithCustomStats(arg))
		assert.Equal(t, expect, c.Stats)
	})
}

func TestRank(t *testing.T) {
	t.Run("Test rank to", func(t *testing.T) {
		arg := uint64(10)
		expect := uint64(10)

		c := &Character{
			Rank: uint64(25),
		}

		c.RankTo(arg)
		assert.Equal(t, expect, c.Rank)
	})

	t.Run("Test rank up", func(t *testing.T) {
		base := uint64(1)
		arg := uint64(10)
		expect := uint64(11)

		c := &Character{
			Rank: base,
		}

		c.RankUp(arg)
		assert.Equal(t, expect, c.Rank)
	})
}
