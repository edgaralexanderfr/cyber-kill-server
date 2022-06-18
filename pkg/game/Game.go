package game

import (
	"fmt"
	"log"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
)

type Game struct {
	GameLoop      GameLoopInterface
	EntityManager EntityManagerInterface
	EntityFactory EntityFactoryInterface
	Config        config.ConfigInterface
}

func (game *Game) Run() {
	game.validateDependencies()

	gameConfig := game.Config.GetConfig(true)

	game.GameLoop.Start(gameConfig.Server.TickRate, game.update)
}

func (game *Game) validateDependencies() {
	if game.GameLoop == nil ||
		game.EntityManager == nil ||
		game.Config == nil {
		log.Fatal("Missing required dependencies. Exiting.")
	}
}

func (game *Game) update() {
	if game.EntityManager != nil {
		game.EntityManager.UpdateAll()
	}

	fmt.Println("Updating game state...")
}
