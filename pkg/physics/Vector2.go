package physics

type Vector2 struct {
	X float32
	Y float32
}

func (vector2 Vector2) GetDirectionTo(target Vector2Interface, speed float32) Vector2Interface {
	return Vector2{}
}
