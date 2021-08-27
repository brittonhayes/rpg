package main

import (
	"os"

	"github.com/brittonhayes/rpg/character"
	"github.com/brittonhayes/rpg/dialogue"
	"github.com/brittonhayes/rpg/logger"
	"github.com/brittonhayes/rpg/stat"
)

func main() {

	log := logger.NewLogger(os.Stderr)

	// Create a player
	player := character.New(
		character.WithName("Player 1"),
	)

	// Create an enemy
	enemy := character.New(
		character.WithName("Enemy 1"),
		character.WithHealth(stat.Low),
		character.WithArmor(stat.Half),
	)

	// Attack the enemy
	dialogue.Say(player, "Let's see how a heavy attack works")
	player.Attack(enemy, player.Attacks.Heavy)

	// Log the attack
	log.Attack(player.Name, enemy.Name, player.Attacks.Heavy)
}
