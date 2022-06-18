package game

type GameLoopFactory struct{}

func (gameLoopFactory *GameLoopFactory) NewGameLoop() GameLoopInterface {
	return &GameLoop{}
}
