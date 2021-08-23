package character

import (
	"testing"

	"github.com/brittonhayes/rpg/stat"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Character
	}{
		{
			name: "test new default character",
			args: args{
				options: nil,
			},
			want: &Character{
				Name:      defaultName,
				Health:    100.00,
				Armor:     0.00,
				Rank:      defaultRank,
				Stats:     defaultStats,
				Attacks:   defaultAttacks,
				Abilities: nil,
				Inventory: nil,
				Playable:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.options...))
		})
	}
}

func TestIsPlayer(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		want      bool
	}{
		{
			name:      "Test is player",
			character: &Character{Playable: false},
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(IsPlayer())
			assert.Equal(t, tt.want, c.Playable)
		})
	}
}

func TestWithAttacks(t *testing.T) {
	type args struct {
		light *Attack
		heavy *Attack
	}

	type want struct {
		light *Attack
		heavy *Attack
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Test with default attacks",
			args: args{
				light: defaultAttacks.Light,
				heavy: defaultAttacks.Heavy,
			},
			want: want{
				light: defaultAttacks.Light,
				heavy: defaultAttacks.Heavy,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithAttacks(tt.args.light, tt.args.heavy))
			assert.Equal(t, tt.want.light, c.Attacks.Light)
			assert.Equal(t, tt.want.heavy, c.Attacks.Heavy)
		})
	}
}

func TestWithName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test character with name", args: args{name: "Mario"}, want: "Mario"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithName(tt.args.name))
			assert.Equal(t, tt.want, c.Name)
		})
	}
}

func TestWithStamina(t *testing.T) {
	type args struct {
		stamina float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test with stamina",
			args: args{
				stamina: float64(20.00),
			},
			want: float64(20.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithStamina(tt.args.stamina))
			assert.Equal(t, tt.args.stamina, c.Stats[stat.Stamina].value)
		})
	}
}

func TestWithHealth(t *testing.T) {
	tests := []struct {
		name   string
		health float64
		want   float64
	}{
		{name: "Test with health", health: 20.00, want: 20.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithHealth(tt.health))
			assert.Equal(t, tt.want, c.Health)
		})
	}
}

func TestWithArmor(t *testing.T) {
	tests := []struct {
		name  string
		armor float64
		want  float64
	}{
		{name: "Test with armor", armor: 20.00, want: 20.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithArmor(tt.armor))
			assert.Equal(t, tt.want, c.Armor)
		})
	}
}

func TestWithRank(t *testing.T) {
	type args struct {
		rank int
	}
	tests := []struct {
		name string
		args args
		want Rank
	}{
		{name: "Test with rank", args: args{rank: 2}, want: Rank(2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithRank(tt.args.rank))
			assert.Equal(t, tt.want, c.Rank)
		})
	}
}

func TestWithCustomStats(t *testing.T) {
	type args struct {
		stats map[string]*Stat
	}
	tests := []struct {
		name string
		args args
		want map[string]*Stat
	}{
		{name: "Test with custom stats", args: args{stats: defaultStats}, want: defaultStats},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(WithCustomStats(tt.args.stats))
			assert.Equal(t, tt.want, c.Stats)
		})
	}
}

func TestCharacter_Stat(t *testing.T) {
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
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stat
	}{
		{
			name: "Test character stat",
			fields: fields{
				Name: defaultName,
				Rank: defaultRank,
				Stats: map[string]*Stat{
					"CHARM": {
						value: 100.00,
					},
				},
				Attacks:   defaultAttacks,
				Abilities: nil,
				Inventory: nil,
				Playable:  true,
			},
			args: args{
				name: "CHARM",
			},
			want: &Stat{
				value: 100.00,
			},
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
			assert.Equal(t, tt.want, c.Stat(tt.args.name))
		})
	}
}

func TestCharacter_Health(t *testing.T) {
	tests := []struct {
		name   string
		health float64
		want   float64
	}{
		{
			name:   "Test character health",
			health: 100.00,
			want:   100.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{
				Health: tt.health,
			}
			assert.Equal(t, tt.want, c.Health)
		})
	}
}

func TestRank_String(t *testing.T) {
	tests := []struct {
		name string
		r    Rank
		want string
	}{
		{name: "Test rank string", r: Rank(1), want: "Rank: 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.r.String())
		})
	}
}

func TestCharacter_RankTo(t *testing.T) {
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
		rank int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Rank
	}{
		{
			name: "Test rank to",
			fields: fields{
				Rank: Rank(1),
			},
			args: args{
				rank: 20,
			},
			want: Rank(20),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Character{
				Rank: tt.fields.Rank,
			}
			c.RankTo(tt.args.rank)
			assert.Equal(t, tt.want, c.Rank)
		})
	}
}

func TestCharacter_RankUp(t *testing.T) {
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
		rank int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Rank
	}{
		{
			name: "Test rank up",
			fields: fields{
				Rank: Rank(1),
			},
			args: args{
				rank: 3,
			},
			want: Rank(4),
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
			c.RankUp(tt.args.rank)
			assert.Equal(t, tt.want, c.Rank)
		})
	}
}
