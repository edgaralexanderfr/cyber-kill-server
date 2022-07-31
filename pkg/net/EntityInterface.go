package net

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/input"

type EntityInterface interface {
	// GetInput returns a reference to the entity's input interface.
	GetInput() input.InputInterface

	// Destroy calls game.EntityManagerInterface.Destroy(entity EntityInterface)
	Destroy()
}
