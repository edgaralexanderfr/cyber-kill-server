package physics

type Vector2Interface interface {
	// GetDirectionTo calculates the direction vector to the target vector.
	// Current implementation serves as the origin vector.
	// It uses the speed parameter to calculate the step of that direction.
	// You can run some examples following the examples link.
	// https://edgaralexanderfr.github.io/cyber-kill-server/examples/get-direction-to/
	GetDirectionTo(target Vector2, speed float32) Vector2
}
