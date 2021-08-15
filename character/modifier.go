package character

type Modifiable interface {
	HealthTo(health float32)
	RankTo(rank int)
	RankUp(rank int)
}
