package main

import (
	"github.com/brittonhayes/rpg/character"
	"github.com/brittonhayes/rpg/stat"
)

func main() {
	// g := game.New(
	// 	game.WithContext(context.Background()),
	// 	game.WithDifficulty(game.Hard),
	// 	game.WithScenes(game.Scene{}),
	// )

	player := character.New(
		character.IsPlayer(),
		character.WithRank(1),
		character.WithHealth(stat.Full),
		character.WithArmor(stat.Full),
		character.WithCustomStats(
			map[string]*character.Stat{
				"Intelligence": character.NewStat(stat.Half),
				"Charm":        character.NewStat(stat.Low),
			},
		),
	)

	enemy := character.New(
		character.WithRank(1),
		character.WithHealth(stat.Low),
	)

	player.Attack(enemy, 20)
	player.RankUp(1)
}
