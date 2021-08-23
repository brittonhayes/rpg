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

func TestCharacter_IsDead(t *testing.T) {
	tests := []struct {
		name   string
		health float64
		want   bool
	}{
		{
			name:   "Test character is dead",
			health: 0.00,
			want:   true,
		},
		{
			name:   "Test character is not dead",
			health: 100.00,
			want:   false,
		},
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
	type args struct {
		target *Character
		attack *Attack
	}
	tests := []struct {
		name   string
		args   args
		want   float64
	}{
		{
			name: "Test character attack",
			args: args{
				target: &Character{
					Health: 100.00,
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
			c := &Character{}
			c.Attack(tt.args.target, tt.args.attack)
			assert.Equal(t, tt.want, tt.args.target.Health)
		})
	}
}
