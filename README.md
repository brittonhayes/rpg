# RPG

> Create text-based role-playing games with Go

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
