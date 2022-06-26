package game

type GameStateFactoryInterface interface {
	NewGameState() GameStateInterface[any]
}
