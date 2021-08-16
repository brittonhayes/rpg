package main

import (
	"fmt"
	"os"

	"github.com/brittonhayes/rpg/character"
	"github.com/brittonhayes/rpg/logger"
	"github.com/brittonhayes/rpg/stat"
)

func main() {
	player := character.New(
		character.WithName("Mario"),
		character.IsPlayer(),
		character.WithRank(1),
		character.WithHealth(stat.Full),
		character.WithAttacks(
			character.NewAttack("Smack", character.LightAttack, 8.00),
			character.NewAttack("Stomp", character.HeavyAttack, 20.00),
		),
		character.WithCustomStats(
			map[string]*character.Stat{
				"Intelligence": character.NewStat(stat.Half),
				"Charm":        character.NewStat(stat.Low),
			},
		),
	)

	enemy := character.New(
		character.WithName("Bowser"),
		character.WithRank(1),
		character.WithHealth(stat.Full),
	)

	log := logger.NewLogger(os.Stderr)

	for ok := true; ok; ok = !enemy.IsDead() {
		attack := player.Attacks.Heavy

		log.Attack(player.Name, enemy.Name, attack.Name, attack.Damage)

		player.Attack(enemy, player.Attacks.Heavy)

		log.Status(player.Name, player.Stat(stat.Health).String(), player.Stat(stat.Armor).String())
		log.Status(enemy.Name, enemy.Stat(stat.Health).String(), enemy.Stat(stat.Armor).String())
	}

	fmt.Printf("\n%s won the fight\n", player.Name)
}
