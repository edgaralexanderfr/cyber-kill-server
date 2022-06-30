package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/input"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/time"
)

type EntityFactoryInterface interface {
	// NewSoldier instantiates a new Soldier to be used in the game.
	NewSoldier(
		id uint32,
		t string,
		parent EntityInterface,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
		entityFactory EntityFactoryInterface,
		entityManager EntityManagerInterface,
		input input.InputInterface,
		physics physics.MapInterface,
		time time.TimeInterface,
		gameState GameStateInterface[any],
	) EntityInterface

	// NewBullet instantiates a new Bullet to be used in the game.
	NewBullet(
		id uint32,
		t string,
		parent EntityInterface,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
		entityFactory EntityFactoryInterface,
		entityManager EntityManagerInterface,
		input input.InputInterface,
		physics physics.MapInterface,
		time time.TimeInterface,
		gameState GameStateInterface[any],
	) EntityInterface
}
