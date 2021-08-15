package character

type Item struct {
	Name      string
	MinRank   int
	Modifiers []Modifier
}

type Health float32

type Modifier interface {
	Apply(stat float32) error
}
