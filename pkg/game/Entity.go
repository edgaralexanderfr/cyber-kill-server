// Package game contains the logic of the game itself.
package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

// Entity struct that implements (or must implement) EntityInterface.
type Entity struct {
	Id            uint32
	EntityManager EntityManagerInterface
	Position      physics.Vector2Interface
	Direction     physics.Vector2Interface
}
