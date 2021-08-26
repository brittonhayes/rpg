package character

import (
	"testing"

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

func Test(t *testing.T) {
	tests := []struct {
		name   string
		health float64
		want   bool
	}{
		{name: "Test character is dead", health: 0.00, want: true},
		{name: "Test character is not dead", health: 100.00, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{
				Health: tt.health,
			}
			assert.Equal(t, tt.want, c.IsDead())
		})
	}
}

func TestCharacter_Attack(t *testing.T) {
	tests := []struct {
		name   string
		target *Character
		attack *Attack
		want   float64
	}{
		{
			name:   "Test character light attack",
			target: &Character{Health: 100.00},
			attack: &Attack{
				Name:   "Light attack",
				Kind:   LightAttack,
				Damage: 10.00,
			},
			want: 90.00,
		},
		{
			name:   "Test character heavy attack",
			target: &Character{Health: 100.00},
			attack: &Attack{
				Name:   "Heavy attack",
				Kind:   HeavyAttack,
				Damage: 40.00,
			},
			want: 60.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{}
			c.Attack(tt.target, tt.attack)
			assert.Equal(t, tt.want, tt.target.Health)
		})
	}
}
