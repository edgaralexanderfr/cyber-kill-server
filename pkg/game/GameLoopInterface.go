package game

type GameLoopInterface interface {
	Start(tickRate byte, update func())
}
