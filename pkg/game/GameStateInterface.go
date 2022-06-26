package game

type GameStateInterface[T any] interface {
	// Get returns the requested generic value through its key.
	// It retrieves the data from the internal map[string]T
	Get(key string) T

	// Set sets the provided generic value with its key.
	// It stores the data into the internal map[string]T
	Set(key string, value T)

	// Serialize returns a final string with all the complete game state data.
	// The string is ready to be broadcasted through the network.
	// It contains the encoded map[string]T in a section.
	// In the other section it contains the serialized entities.
	// The string is encoded in a JSON format.
	SerializeToJSON(gameStateEntities []GameStateEntityInterface) string
}
