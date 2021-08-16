package game

type State int

const (
	Crashed State = iota + 1
	Stopped
	Paused
	Running
)

func (s State) String() string {
	return [...]string{"Crashed", "Stopped", "Paused", "Running"}[s]
}

func (g *Game) Start() {
	g.State = Running

}

func (g *Game) Stop() {
	g.State = Stopped
}
