package time

type TimeInterface interface {
	// SetTickRate sets internally how much time passes for every call to the Update() method.
	// e.g.
	// tickRate = 10
	// That means that the time between each update = 1000/10 = 100
	// The idea is to accumulate 100 in every update to control each timer.
	SetTickRate(tickRate byte)

	// GetDeltaTime returns the elapsed time in seconds since previous iteration.
	// It is used to compensate entities interpolation.
	// It is also used to keep a cycle/frame-independent game.
	// e.g.
	// The update time between the previous frame and current was 100ms.
	// That means that GetDeltaTime() = 100/1000 = 0.1
	// The Delta Time is a cached calculated value from the Update() method.
	// Keep in mind that GetDeltaTime() uses real system time for evaluation.
	GetDeltaTime() float32

	// SetTimeout sets a function to be executed after the specific given time.
	// Once the function is executed it's removed from the internal stack.
	// The given time is represented in milliseconds.
	// It returns the ID of the created/stacked timeout.
	SetTimeout(time uint32, callback func()) (id uint32)

	// ClearTimeout removes the function from the internal stack.
	ClearTimeout(id uint32)

	// SetInterval sets a function to be executed every specific given time.
	// The given time is represented in milliseconds.
	// It returns the ID of the created/stacked interval.
	SetInterval(time uint32, callback func()) (id uint32)

	// ClearInterval removes the function from the internal stack.
	ClearInterval(id uint32)

	// Update updates all the logic from within the TimeInterface implementation.
	// It calculates the current DeltaTime in comparison to the previous frame.
	// It accumulates the elapsed time established with the tickRate.
	// It checks through every timer, compares each independent elapsed time.
	// It executes and removes the callbacks and timers in the stack.
	Update()
}
