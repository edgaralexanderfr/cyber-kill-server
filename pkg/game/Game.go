package game

import (
	"fmt"
	"log"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/net"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

type Game struct {
	GameLoop      GameLoopInterface
	EntityManager EntityManagerInterface
	EntityFactory EntityFactoryInterface
	Config        config.ConfigInterface
	Matchmaking   net.MatchmakingInterface
}

func (game *Game) Run() {
	game.validateDependencies()

	gameConfig := game.Config.GetConfig(true)

	game.Matchmaking.SetInputEvents(InputEvents)
	game.Matchmaking.GetPlayerEntity(game.getPlayerEntity)
	game.Matchmaking.Disconnect(game.disconnect)
	game.Matchmaking.Listen(gameConfig.Port)
	game.GameLoop.Start(gameConfig.Server.TickRate, game.update)
}

func (game *Game) validateDependencies() {
	if game.GameLoop == nil ||
		game.EntityManager == nil ||
		game.EntityFactory == nil ||
		game.Config == nil ||
		game.Matchmaking == nil {
		log.Fatal("Missing required dependencies. Exiting.")
	}
}

func (game *Game) getPlayerEntity() net.EntityInterface {
	player := game.EntityFactory.NewSoldier(1, "soldier", game.EntityManager, physics.Vector2{}, physics.Vector2{})
	game.EntityManager.AddEntity(player)

	return player
}

func (game *Game) disconnect(entity net.EntityInterface) {
	player := entity.(EntityInterface)

	fmt.Println("Player", player.GetId(), "disconnected.")
}

func (game *Game) update() {
	if game.EntityManager != nil {
		game.EntityManager.UpdateAll()
	}

	fmt.Println("Updating game state...")
}
