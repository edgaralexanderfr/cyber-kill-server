package game

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/config"

type GameFactory struct{}

func (gameFactory *GameFactory) NewGame() GameInterface {
	gameLoopFactory := GameLoopFactory{}
	entityManagerFactory := EntityManagerFactory{}
	configFactory := config.ConfigFactory{}

	return &Game{
		GameLoop:      gameLoopFactory.NewGameLoop(),
		EntityManager: entityManagerFactory.NewEntityManager(),
		Config:        configFactory.NewConfig(),
	}
}
