// Package game contains the logic of the game itself.
package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

// EntityInterface is used to represent a game entity (GameObject)
type EntityInterface interface {
	// NewEntity instantiates a new entity providing the required parameters.
	NewEntity(
		id uint32,
		t string,
		entityManager EntityManagerInterface,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
	)

	// GetId returns the entity's ID provided to NewEntity.
	GetId() uint32

	// GetType returns the entity's Type (t) provided to NewEntity.
	GetType() string

	// Update updates the entity's state.
	Update()

	// Destroy calls EntityManagerInterface.Destroy(entity EntityInterface)
	Destroy()
}
