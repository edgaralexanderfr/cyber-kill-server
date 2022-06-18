package game

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"

type EntityFactoryInterface interface {
	// NewSoldier instantiates a new Soldier to be used in the game.
	NewSoldier(
		id uint32,
		t string,
		entityManager EntityManagerInterface,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
	) EntityInterface

	// NewBullet instantiates a new Bullet to be used in the game.
	NewBullet(
		id uint32,
		t string,
		entityManager EntityManagerInterface,
		position physics.Vector2Interface,
		direction physics.Vector2Interface,
	) EntityInterface
}
