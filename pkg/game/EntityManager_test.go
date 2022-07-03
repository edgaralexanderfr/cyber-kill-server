package game

import (
	"testing"

	"github.com/edgaralexanderfr/cyber-kill-server/pkg/physics"
)

func TestEntityManager(t *testing.T) {
	entityFactory := &EntityFactory{}
	entityManagerFactory := &EntityManagerFactory{}
	entityManager := entityManagerFactory.NewEntityManager()

	if entityManager == nil {
		t.Fatalf(`entityManagerFactory.NewEntityManager() = nil, want EntityManagerInterface`)
	}

	id := entityManager.GetNewId()

	if id < 1 {
		t.Fatalf(`entityManager.GetNewId() = %d, want > 0`, id)
	}

	newId := entityManager.GetNewId()

	if newId == id {
		t.Fatalf(`entityManager.GetNewId() = %d, want != %d`, newId, id)
	}

	entityManager.AddEntity(entityFactory.NewSoldier(
		entityManager.GetNewId(),
		"soldier",
		nil,
		physics.Vector2{X: 5.0, Y: 5.0},
		physics.Vector2{X: -2.0, Y: 7.5},
		entityFactory,
		entityManager,
		nil,
		nil,
		nil,
		nil,
	))

	soldiers := entityManager.GetEntitiesByType("soldier")

	if soldiers == nil {
		t.Fatalf(`entityManager.GetEntitiesByType("soldier") = nil, want []EntityInterface`)
	}

	total := len(soldiers)

	if total != 1 {
		t.Fatalf(`len(entityManager.GetEntitiesByType("soldier")) = %d, want 1`, total)
	}

	if soldiers[0].GetType() != "soldier" {
		t.Fatalf(`soldiers[0].GetType() = "%s", want "soldier"`, soldiers[0].GetType())
	}

	x, y := soldiers[0].GetPosition()

	if x != 5.0 || y != 5.0 {
		t.Fatalf(`soldiers[0].GetPosition() = {%f %f}, want {5.0 5.0}`, x, y)
	}

	entityManager.AddEntity(entityFactory.NewBullet(
		entityManager.GetNewId(),
		"bullet",
		nil,
		physics.Vector2{X: 10.0, Y: -11.0},
		physics.Vector2{X: 0.0, Y: 9.3},
		entityFactory,
		entityManager,
		nil,
		nil,
		nil,
		nil,
	))

	bullets := entityManager.GetEntitiesByType("bullet")

	if bullets == nil {
		t.Fatalf(`entityManager.GetEntitiesByType("bullet") = nil, want []EntityInterface`)
	}

	total = len(bullets)

	if total != 1 {
		t.Fatalf(`len(entityManager.GetEntitiesByType("bullet")) = %d, want 1`, total)
	}

	if bullets[0].GetType() != "bullet" {
		t.Fatalf(`bullets[0].GetType() = "%s", want "bullet"`, bullets[0].GetType())
	}

	x, y = bullets[0].GetPosition()

	if x != 10.0 || y != -11.0 {
		t.Fatalf(`bullets[0].GetPosition() = {%f %f}, want {10.0 -11.0}`, x, y)
	}

	if soldiers[0].GetId() == bullets[0].GetId() {
		t.Fatalf(`soldiers[0].GetId() = bullets[0].GetId(), want %d != %d`, soldiers[0].GetId(), bullets[0].GetId())
	}

	soldier := entityFactory.NewSoldier(
		entityManager.GetNewId(),
		"soldier",
		nil,
		physics.Vector2{X: 0.0, Y: 0.0},
		physics.Vector2{X: 0.0, Y: 0.0},
		entityFactory,
		entityManager,
		nil,
		nil,
		nil,
		nil,
	)

	entityManager.AddEntity(soldier)
	entityManager.UpdateAll()
	entityManager.Destroy(soldier)

	soldiers = entityManager.GetEntitiesByType("soldiers")
	total = len(soldiers)

	if total != 1 {
		t.Fatalf(`entityManager.Destroy(soldier) failed, len(entityManager.GetEntitiesByType("soldiers")) = %d, want 1`, total)
	}

	for _, s := range soldiers {
		if s.GetId() == soldier.GetId() {
			t.Fatalf(`entityManager.Destroy(soldier) failed, soldier.GetId() found in entityManager.GetEntitiesByType("soldiers")`)
		}
	}

	data := entityManager.Serialize()

	if data == nil {
		t.Fatalf(`entityManager.Serialize() = nil, want []GameStateEntityInterface`)
	}

	total = len(data)

	if total != 2 {
		t.Fatalf(`len(entityManager.Serialize()) = %d, want 2`, total)
	}

	if data[0] == nil {
		t.Fatalf(`entityManager.Serialize()[0] = nil, want GameStateEntityInterface`)
	}

	if data[1] == nil {
		t.Fatalf(`entityManager.Serialize()[1] = nil, want GameStateEntityInterface`)
	}

	entityManager.UpdateAll()
}
