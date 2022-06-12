package game

import "fmt"

type Game struct{}

func (game Game) Run() {
	fmt.Println("Is running!")
}
