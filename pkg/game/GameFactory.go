package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/input"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/net"
)

type GameFactory struct{}

func (gameFactory *GameFactory) NewGame() GameInterface {
	gameLoopFactory := &GameLoopFactory{}
	entityManagerFactory := &EntityManagerFactory{}
	entityFactory := &EntityFactory{}
	configFactory := &config.ConfigFactory{}
	matchmakingFactory := &net.MatchmakingFactory{}
	inputFactory := &input.InputFactory{}

	return &Game{
		GameLoop:      gameLoopFactory.NewGameLoop(),
		EntityManager: entityManagerFactory.NewEntityManager(),
		EntityFactory: entityFactory,
		Config:        configFactory.NewConfig(),
		Matchmaking:   matchmakingFactory.NewMatchmaking(),
		InputFactory:  inputFactory,
	}
}
