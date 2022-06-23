package input

type InputInterface interface {
	// GetAction indicates whether the specified action is being performed or not.
	GetAction(action string) bool

	// SetAction sets the value for the specified performed action.
	SetAction(action string, value bool)

	// GetValue returns the analog value for the specified update/event.
	GetValue(value string) (float32, float32)

	// SetValue sets the analog value for the specified update/event.
	SetValue(action string, x float32, y float32)
}
