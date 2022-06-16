package game

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/config"

func NewGame() GameInterface {
	return &Game{
		GameLoop:      NewGameLoop(),
		EntityManager: NewEntityManager(),
		Config:        config.NewConfig(),
	}
}
