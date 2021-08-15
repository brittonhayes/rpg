package character

type NPC struct {
	Name      string
	Health    float32
	Stamina   float32
	Abilities []Ability
	Inventory []Item
}
