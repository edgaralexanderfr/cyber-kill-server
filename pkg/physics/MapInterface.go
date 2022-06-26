package physics

type MapInterface interface {
	// GetMap returns the generated/stored matrix (width*height) for the map.
	// If the indexed tile/block is true, it means there's a collision there.
	// The rows represent the X axis and the cols the Y axis.
	// e.g. map[x][y]
	GetMap() [][]bool

	// GetRandomValidSpawnPoint returns a random spawn point from the map.
	// This spawn point is typically used to spawn the players.
	// It checks that the spawn point is not true inside the matrix to be valid.
	// The X/Y coordinates for the spawn point are the center of the valid tile.
	GetRandomValidSpawnPoint() Vector2Interface

	// SetTileSize stores the size for each tile/block.
	// It is used by PredictCollision to detect if the resultant point collides.
	SetTileSize(size uint16)

	// Generate generates/stores the [][]bool matrix to be returned in GetMap()
	// It ensures that all the boundaries of the map are set to true.
	// That way players and entities can't get off the map.
	// It also generates random walls and paths inside the map.
	Generate(width uint16, height uint16)

	// PredictCollision checks collision detection for the resultant point.
	// Resultant point = position param + direction param.
	// Returns true if resultant point happens to be inside a tile with true
	// It multiplies the tile size for each block to do the detection check.
	PredictCollision(position Vector2Interface, direction Vector2Interface) bool
}
