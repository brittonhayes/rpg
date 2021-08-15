package game

import (
	"context"
)

type Runner interface {
	Start(ctx context.Context) error
	Stop() error
	Pause() error
}

type Game struct {
	ctx    context.Context
	cancel context.CancelFunc

	Difficulty difficulty
	Scenes     []Scene
	State      State
}

type Option func(*Game)

func New(options ...Option) *Game {
	const (
		defaultDifficulty = Medium
		defaultState      = Paused
	)

	game := &Game{
		ctx:        context.Background(),
		cancel:     nil,
		Difficulty: defaultDifficulty,
		State:      defaultState,
		Scenes:     nil,
	}

	for _, opt := range options {
		opt(game)
	}
	return game
}

func WithContext(ctx context.Context, key, value interface{}) Option {
	return func(game *Game) {
		game.ctx = context.WithValue(ctx, key, value)
	}
}

func WithDifficulty(difficulty difficulty) Option {
	return func(game *Game) {
		game.Difficulty = difficulty
	}
}

func WithScenes(scenes ...Scene) Option {
	return func(game *Game) {
		game.Scenes = scenes
	}
}
