package game

import (
	"fmt"
	"log"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
)

type Game struct {
	GameLoop GameLoopInterface
	Config   config.ConfigInterface
}

func (game Game) Run() {
	if game.Config == nil {
		log.Fatal("Missing required dependencies. Exiting.")
	}

	gameConfig := game.Config.GetConfig(true)

	game.GameLoop.Start(gameConfig.Server.TickRate, update)
}

func update() {
	fmt.Println("Updating game state...")
}
