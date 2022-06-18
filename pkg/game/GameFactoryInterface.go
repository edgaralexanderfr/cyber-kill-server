package game

type GameFactoryInterface interface {
	NewGame() GameInterface
}
