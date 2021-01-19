package units

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) TP() TPoint {
	switch d {
	case North:
		return TP(0, -1)
	case South:
		return TP(0, 1)
	case East:
		return TP(1, 0)
	case West:
		return TP(-1, 0)
	default:
		panic("invalid direction")
	}
}

func (d Direction) IsOpposite(other Direction) bool {
	x, y := d.TP().Sub(other.TP()).XY()
	return x == 0 && y == 0
}
