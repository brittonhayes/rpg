package main

import (
	"fmt"
	"os"

	"github.com/brittonhayes/rpg/character"
	"github.com/brittonhayes/rpg/dialogue"
	"github.com/brittonhayes/rpg/logger"
	"github.com/brittonhayes/rpg/stat"
)

func main() {
	player := character.New(
		character.WithName("Mario"),
		character.IsPlayer(),
		character.WithRank(1),
		character.WithHealth(stat.Full),
		character.WithArmor(stat.Full),
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
		character.WithHealth(stat.Half),
	)
	log := logger.NewLogger(os.Stderr)

	log.Status(player.Name, player.Health, player.Armor)
	log.Status(enemy.Name, enemy.Health, enemy.Armor)

	dialogue.Say(enemy, "It's me Bowser!")
	dialogue.Say(player, "Yeah, I'm aware of that.")

	for ok := true; ok; ok = !enemy.IsDead() {
		attack := player.Attacks.Heavy

		log.Attack(player.Name, enemy.Name, attack.Name, attack.Damage)
		player.Attack(enemy, player.Attacks.Heavy)
	}

	dialogue.Say(player, "Bowser has been killed.")
	dialogue.Say(enemy, "Yeah, I'm aware of that.")

	dialogue.Ask(player, "What should we do now?", func(answer string) error {
		log.Log(logger.LevelChat, fmt.Sprintf("Are you sure you want to %q?", answer))
		return nil
	})

}
