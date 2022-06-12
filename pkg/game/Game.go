package game

import (
	"fmt"
	"log"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
)

type Game struct {
	Config config.ConfigInterface
}

func (game Game) Run() {
	if game.Config == nil {
		log.Fatal("Missing required dependencies. Exiting.")
	}

	fmt.Println("The game is running!")
}
