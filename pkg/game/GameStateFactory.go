package game

type GameStateFactory struct{}

func (gameStateFactory *GameStateFactory) NewGameState() GameStateInterface[any] {
	return nil
}
