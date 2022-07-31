package net

type AnalogicInputEvent struct {
	Event string

	Value struct {
		X float32
		Y float32
	}
}
