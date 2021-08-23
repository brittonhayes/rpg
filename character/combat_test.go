package character

import (
	"testing"

	"github.com/brittonhayes/rpg/stat"
	"github.com/stretchr/testify/assert"
)

func TestNewAttack(t *testing.T) {
	type args struct {
		name   string
		kind   AttackType
		damage float64
	}
	tests := []struct {
		name string
		args args
		want *Attack
	}{
		{name: "Test new attack", args: args{
			name:   "Test attack",
			kind:   LightAttack,
			damage: 20.00,
		},
			want: &Attack{Name: "Test attack", Kind: LightAttack, Damage: 20.00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewAttack(tt.args.name, tt.args.kind, tt.args.damage))
		})
	}
}

func TestCharacter_IsDead(t *testing.T) {
	type fields struct {
		Name      string
		Rank      Rank
		Stats     map[string]*Stat
		Attacks   Attacks
		Abilities []Ability
		Inventory []Item
		Playable  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test character is dead",
			fields: fields{
				Stats: map[string]*Stat{
					"HEALTH": {
						value: 0.00,
					},
				},
			},
			want: true,
		},
		{
			name: "Test character is not dead",
			fields: fields{
				Stats: map[string]*Stat{
					"HEALTH": {
						value: 50.00,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{
				Name:      tt.fields.Name,
				Rank:      tt.fields.Rank,
				Stats:     tt.fields.Stats,
				Attacks:   tt.fields.Attacks,
				Abilities: tt.fields.Abilities,
				Inventory: tt.fields.Inventory,
				Playable:  tt.fields.Playable,
			}
			assert.Equal(t, tt.want, c.IsDead())
		})
	}
}

func TestCharacter_Attack(t *testing.T) {
	type fields struct {
		Name      string
		Rank      Rank
		Stats     map[string]*Stat
		Attacks   Attacks
		Abilities []Ability
		Inventory []Item
		Playable  bool
	}
	type args struct {
		target *Character
		attack *Attack
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Test character attack",
			fields: fields{
				Stats: defaultStats,
			},
			args: args{
				target: &Character{
					Stats: map[string]*Stat{
						stat.Health: {
							value: 100.00,
						},
					},
				},
				attack: &Attack{
					Name:   "Test attack",
					Kind:   LightAttack,
					Damage: 10.00,
				},
			},
			want: 90.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{
				Name:      tt.fields.Name,
				Rank:      tt.fields.Rank,
				Stats:     tt.fields.Stats,
				Attacks:   tt.fields.Attacks,
				Abilities: tt.fields.Abilities,
				Inventory: tt.fields.Inventory,
				Playable:  tt.fields.Playable,
			}
			c.Attack(tt.args.target, tt.args.attack)
			assert.Equal(t, tt.want, tt.args.target.Stats["HEALTH"].value)
		})
	}
}
