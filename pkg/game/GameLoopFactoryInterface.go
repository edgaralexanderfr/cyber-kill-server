package game

type GameLoopFactoryInterface interface {
	NewGameLoop() GameLoopInterface
}
