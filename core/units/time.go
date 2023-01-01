package units

// A Tick is the smallest unit of time of the game.
// It represents one update to the game, which
// is typically the frame rate, i.e. 1/60 seconds.
type Tick uint64

// Second is a unit of time which allows us to do math
// in seconds and convert it to ticks.
type Second float32

// Ticks converts seconds to ticks
func (s Second) Ticks() Tick {
	return Tick(s * Second(TickRate))
}

// Seconds converts ticks to seconds
func (t Tick) Seconds() Second {
	return Second(t) / Second(TickRate)
}
