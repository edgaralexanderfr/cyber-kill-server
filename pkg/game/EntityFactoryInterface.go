package game

import (
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/input"
	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

type EntityFactoryInterface interface {
	// NewSoldier instantiates a new Soldier to be used in the game.
	NewSoldier(
		id uint32,
		t string,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
		entityFactory EntityFactoryInterface,
		entityManager EntityManagerInterface,
		input input.InputInterface,
	) EntityInterface

	// NewBullet instantiates a new Bullet to be used in the game.
	NewBullet(
		id uint32,
		t string,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
		entityFactory EntityFactoryInterface,
		entityManager EntityManagerInterface,
		input input.InputInterface,
	) EntityInterface
}
