package game

type GameLoopInterface interface {
	Start(tickRate int, update func())
}
