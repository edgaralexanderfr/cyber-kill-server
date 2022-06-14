package main

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/game"
)

func main() {
	var game game.GameInterface = game.Game{
		GameLoop: game.GameLoop{},
	}

	game.Run()
}
