package main

import (
	"context"

	"github.com/brittonhayes/rpg/game"
)

func main() {
	g := game.New(
		game.WithContext(context.Background(), "", ""),
		game.WithDifficulty(game.Hard),
		game.WithScenes(game.Scene{}),
	)

	go func() {
		g.Stop()
	}()

	g.Start()
}
