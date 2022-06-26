// Package game contains the logic of the game itself.
package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/input"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

// EntityInterface is used to represent a game entity (GameObject)
type EntityInterface interface {
	// NewEntity instantiates a new entity providing the required parameters.
	NewEntity(
		id uint32,
		t string,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
		entityFactory EntityFactoryInterface,
		entityManager EntityManagerInterface,
		input input.InputInterface,
	)

	// GetId returns the entity's ID provided to NewEntity.
	GetId() uint32

	// GetType returns the entity's Type (t) provided to NewEntity.
	GetType() string

	// GetPosition returns the current X/Y position of the entity.
	GetPosition() (x float32, y float32)

	// GetInput returns a reference to the entity's input interface.
	GetInput() input.InputInterface

	// Update updates the entity's state.
	Update()

	// Serialize returns a JSON with extra data for the current state of the entity.
	// It a returns a struct{} singleton to be broadcasted through the network.
	Serialize() GameStateEntityInterface

	// Destroy calls EntityManagerInterface.Destroy(entity EntityInterface)
	Destroy()
}
