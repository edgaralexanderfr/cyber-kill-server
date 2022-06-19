package net

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/types"

type MatchmakingInterface interface {
	// SetInputEvents establishes the input events to process from the players.
	SetInputEvents([]InputEvent)

	// GetPlayerEntity sets the function to call to retrieve the player's entity instance.
	GetPlayerEntity(callback func() EntityInterface)

	// Disconnect event to be called once a player disconnects from the server.
	Disconnect(callback func(entity EntityInterface))

	// Listen initiates the server to listen for the events from the players.
	Listen(port types.Port)
}
