package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

// We gotta create an EntityFactory too...
type EntityInterface interface {
	NewEntity(position physics.Vector2Interface, direction physics.Vector2Interface)
	Update()  // Create separated function to handle this
	Destroy() // Create separated function to handle this
}
