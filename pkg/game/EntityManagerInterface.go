// Package game contains the logic of the game itself.
package game

// EntityManagerInterface is used for handling all game entities.
// (Players, NPCs, projectiles, etc...)
type EntityManagerInterface interface {
	// AddEntity adds a new entity of any type to the slice.
	AddEntity(entity EntityInterface)

	// GetEntitiesByType returns a slice with all the entities with the given type.
	// It compares calling the EntityInterface.GetType() method.
	GetEntitiesByType(t string) []EntityInterface

	// UpdateAll iterates over all the stacked entities on the slice.
	// Calls the corresponding implemented EntityInterface.Update() method.
	UpdateAll()

	// Destroy removes the specified entity from the slice.
	// It uses EntityInterface.GetId() to look, compare and remove the entity.
	Destroy(entity EntityInterface)
}
