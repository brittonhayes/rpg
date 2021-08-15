package game

import (
	"fmt"
	"time"
)

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

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("done")
	case <-g.ctx.Done():
		fmt.Println("Done!")
	}
}

func (g *Game) Stop() {
	g.State = Stopped
	defer g.cancel()
}
