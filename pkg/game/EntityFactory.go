package game

import "github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"

type EntityFactory struct{}

func (entityFactory *EntityFactory) NewSoldier(
	id uint32,
	t string,
	entityManager EntityManagerInterface,
	position physics.Vector2Interface,
	direction physics.Vector2Interface,
) EntityInterface {
	return nil
}

func (entityFactory *EntityFactory) NewBullet(
	id uint32,
	t string,
	entityManager EntityManagerInterface,
	position physics.Vector2Interface,
	direction physics.Vector2Interface,
) EntityInterface {
	return nil
}
