package units

// TileSize is the number of pixels to one tile
const TileSize = 16

// TilesPerSecond is a helper to conver to
// the standard unit of velocity is pixels/second.
const TilesPerSecond = float64(TileSize) / 1.0

// PixelsPerTick is a helper to conver to
// the standard unit of velocity is pixels/second.
const PixelsPerTick = 1.0 / float64(TickRate)

// TickRate is the ticks per second
const TickRate = 60
