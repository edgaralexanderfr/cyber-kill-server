package physics

type Vector2Interface interface {
	// GetDirectionTo calculates the direction vector to the target vector.
	// Current implementation serves as the origin vector.
	// It uses the speed parameter to calculate the step of that direction.
	GetDirectionTo(target Vector2Interface, speed float32) Vector2Interface
}
