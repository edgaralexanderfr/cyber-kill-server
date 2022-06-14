package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

type Entity struct {
	Position  physics.Vector2Interface
	Direction physics.Vector2Interface
}
