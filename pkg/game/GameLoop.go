package game

import (
	"time"
)

type GameLoop struct{}

func (gameLoop GameLoop) Start(tickRate int, update func()) {
	for range time.Tick(time.Second / time.Duration(tickRate)) {
		update()
	}
}
