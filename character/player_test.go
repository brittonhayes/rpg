package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	type args struct {
		options []PlayerOption
	}
	tests := []struct {
		name string
		args args
		want *Player
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPlayer(tt.args.options...))
		})
	}
}

func TestPlayer_HealthTo(t *testing.T) {
	type fields struct {
		Name      string
		Rank      int
		Health    Stat
		Stamina   Stat
		Abilities []Ability
		Inventory []Item
	}
	type args struct {
		health Stat
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name:      tt.fields.Name,
				Rank:      tt.fields.Rank,
				Health:    tt.fields.Health,
				Stamina:   tt.fields.Stamina,
				Abilities: tt.fields.Abilities,
				Inventory: tt.fields.Inventory,
			}
			p.HealthTo(tt.args.health)
		})
	}
}

func TestPlayer_RankTo(t *testing.T) {
	tests := []struct {
		name   string
		player Player
		rank   int
	}{
		{name: "rank down player", player: Player{Rank: 100}, rank: 1},
		{name: "rank up player from 0", player: Player{Rank: 0}, rank: 8},
		{name: "rank up player from nil", player: Player{}, rank: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name:      tt.player.Name,
				Rank:      tt.player.Rank,
				Health:    tt.player.Health,
				Stamina:   tt.player.Stamina,
				Abilities: tt.player.Abilities,
				Inventory: tt.player.Inventory,
			}
			p.RankTo(tt.rank)
			assert.Equal(t, tt.rank, p.Rank)
		})
	}
}

func TestPlayer_RankUp(t *testing.T) {
	tests := []struct {
		name     string
		player   Player
		increase int
		expected int
	}{
		{name: "rank up player", player: Player{Rank: 1}, increase: 1, expected: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				Name:      tt.player.Name,
				Rank:      tt.player.Rank,
				Health:    tt.player.Health,
				Stamina:   tt.player.Stamina,
				Abilities: tt.player.Abilities,
				Inventory: tt.player.Inventory,
			}
			p.RankUp(tt.increase)
			assert.Equal(t, tt.expected, p.Rank)
		})
	}
}
