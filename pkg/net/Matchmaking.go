package net

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/types"

type Matchmaking struct {
	InputEvents             []InputEvent
	GetPlayerEntityCallback func() EntityInterface
	DisconnectCallback      func(entity EntityInterface)
}

func (matchmaking *Matchmaking) SetInputEvents(inputEvents []InputEvent) {
	matchmaking.InputEvents = inputEvents
}

func (matchmaking *Matchmaking) GetPlayerEntity(callback func() EntityInterface) {
	matchmaking.GetPlayerEntityCallback = callback
}

func (matchmaking *Matchmaking) Disconnect(callback func(entity EntityInterface)) {
	matchmaking.DisconnectCallback = callback
}

func (matchmaking *Matchmaking) Listen(port types.Port) {

}

func (matchmaking *Matchmaking) Sync(gameStateStr string) {

}
