package main

import (
	"github.com/brittonhayes/rpg/character"
	"github.com/brittonhayes/rpg/color"
	"github.com/brittonhayes/rpg/dialogue"
	"github.com/brittonhayes/rpg/stat"
)

func main() {

	player := character.New(
		character.WithName("Player"),
		character.IsPlayer(),
		character.WithColor(color.Cyan),
		character.WithRank(5),
		character.WithHealth(stat.Full),
	)

	enemy := character.New(
		character.WithName("Enemy"),
		character.WithRank(1),
		character.WithHealth(stat.Half),
	)

	npc := character.New(
		character.WithName("Steve"),
	)

	// Speak as a player
	dialogue.Say(player, "...oh boy. uh. I should probably get out of here.")

	// Speak as an NPC
	dialogue.Say(enemy, "Yeah this is for sure a felony. I think that guy over there definitely saw you do it too.")

	// Speak as an NPC that died in combat. This will be prepended with "[DEAD]".
	enemy.Health = 0.00
	dialogue.Say(enemy, "Welp. Now look what you did. I'm dead. You're absolutely gonna need a lawyer.")

	dialogue.Say(player, "You... uh.. You know a good lawyer?")

	// Ask a question to the player and handle their answer
	// with a custom function
	dialogue.Ask(npc, "Oh. Are you talking to me? (yes/no)", func(answer string) error {
		switch answer {
		case "yes":
			dialogue.Say(npc, "I feel like you probably shouldn't be asking me for legal advice.")
		case "no":
			dialogue.Say(npc, "Okay cool, I'm gonna back away into the forrest and let you sort all this out.")
		default:
			dialogue.Say(npc, "I have no idea what you want me to say right now. Can I leave?")
			return nil
		}

		return nil
	})
}
