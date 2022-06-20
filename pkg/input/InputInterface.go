package input

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"

// TODO: Create Factory...
type InputInterface interface {
	// GetAction indicates whether the specified action is being performed or not.
	GetAction(action string) bool

	// GetValue returns the analog value for the specified update/event.
	GetValue(value string) physics.Vector2Interface
}
