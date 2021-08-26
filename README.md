# RPG ✨🕹️

[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/rpg)](https://goreportcard.com/report/github.com/brittonhayes/rpg)
[![Go Reference](https://pkg.go.dev/badge/github.com/brittonhayes/rpg.svg)](https://pkg.go.dev/github.com/brittonhayes/rpg)

> Create text-based role-playing games with Go

<image src="./logo.png" width=500/>

## Installation 📥

```shell
go get github.com/brittonhayes/rpg
```

## Examples

Check the [_examples](./_examples) directory for examples of how to use the library.

## Characters 👥

### Create a player

```go
player := character.New(
    character.WithName("Player 1"),
    character.WithRank(1),
)
```

### Customize your character

```go
player := character.New(
    character.WithName("Player 1"),
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
```

## Combat ⚔️

```go
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
player.Attack(enemy, player.Attacks.Heavy)

// Log the attack
log.Attack(player.Name, enemy.Name, attack.Name, attack.Damage)
```

## Dialogue 💬

```go
// Speak as a player
dialogue.Say(player, "Hey there! I'm the main character.")

// Speak as an NPC
dialogue.Say(enemy, "I'm an enemy. You can tell by my different output color.")

// Speak as an NPC that died in combat. This will be prepended with "[DEAD]".
enemy.Health = 0.00
dialogue.Say(enemy, "I've been defeated")

// Ask a question to the player and handle their answer
// with a custom function
dialogue.Ask(player, "What should we do next boss?", func(answer string) error {
    dialogue.Say(enemy, fmt.Sprintf("Do you you really want to %q?", answer))
    return nil
  }
)
```
