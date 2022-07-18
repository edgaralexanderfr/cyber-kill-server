package main

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/game"

func main() {
	gameFactory := game.GameFactory{}
	gameFactory.NewGame().Run()
}
