package character

type (
	Rank int
)

type Player struct {
	Name      string
	Rank      int
	Health    Stat
	Stamina   Stat
	Abilities []Ability
	Inventory []Item
}

type PlayerOption func(*Player)

func NewPlayer(options ...PlayerOption) *Player {
	const (
		defaultName    = "Player 1"
		defaultRank    = 0
		defaultHealth  = Stat(100.0)
		defaultStamina = Stat(100.0)
	)

	p := &Player{
		Name:      defaultName,
		Rank:      defaultRank,
		Health:    defaultHealth,
		Stamina:   defaultStamina,
		Abilities: nil,
		Inventory: nil,
	}

	for _, opt := range options {
		opt(p)
	}

	return p
}

func (p *Player) HealthTo(health Stat) {
	p.Health = health
}

func (p *Player) RankTo(rank int) {
	p.Rank = rank
}

func (p *Player) RankUp(rank int) {
	p.Rank += rank
}
