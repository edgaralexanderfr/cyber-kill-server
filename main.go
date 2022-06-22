package main

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
)

func main() {
	//gameFactory := game.GameFactory{}
	//gameFactory.NewGame().Run()
	configFactory := config.ConfigFactory{}
	config := configFactory.NewConfig().GetConfig(true)
	println(config.Port, config.Server.TickRate)
}
