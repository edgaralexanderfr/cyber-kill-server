package game

import (
	"fmt"
	"log"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/config"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/input"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/net"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/time"
)

type Game struct {
	GameLoop      GameLoopInterface
	EntityManager EntityManagerInterface
	EntityFactory EntityFactoryInterface
	Config        config.ConfigInterface
	Matchmaking   net.MatchmakingInterface
	InputFactory  input.InputFactoryInterface
	Map           physics.MapInterface
	Time          time.TimeInterface
	GameState     GameStateInterface[any]
}

func (game *Game) Run() {
	game.validateDependencies()

	gameConfig := game.Config.GetConfig(true)

	game.Map.SetTileSize(gameConfig.Map.Size)
	game.Map.Generate(gameConfig.Map.Width, gameConfig.Map.Height)

	game.Time.SetTickRate(gameConfig.Server.TickRate)

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
		game.Matchmaking == nil ||
		game.InputFactory == nil ||
		game.Map == nil ||
		game.Time == nil ||
		game.GameState == nil {
		log.Fatal("Missing required dependencies. Exiting.")
	}
}

func (game *Game) getPlayerEntity() net.EntityInterface {
	player := game.EntityFactory.NewSoldier(
		game.EntityManager.GetNewId(),
		"soldier",
		nil,
		physics.Vector2{},
		physics.Vector2{},
		game.EntityFactory,
		game.EntityManager,
		game.InputFactory.NewInput(),
		game.Map,
		game.Time,
		game.GameState,
	)

	game.EntityManager.AddEntity(player)

	return player
}

func (game *Game) disconnect(entity net.EntityInterface) {
	player := entity.(EntityInterface)

	fmt.Println("Player", player.GetId(), "disconnected.")
}

func (game *Game) update() {
	game.Time.Update()
	game.EntityManager.UpdateAll()

	game.Matchmaking.Sync(
		game.GameState.SerializeToJSON(
			game.EntityManager.Serialize(),
		),
	)

	fmt.Println("Updating game state...")
}
