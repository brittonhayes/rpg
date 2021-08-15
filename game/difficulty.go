package game

type difficulty int

const (
	Easy difficulty = iota
	Medium
	Hard
)

func (d difficulty) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}
