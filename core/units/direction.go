package units

type Direction int

//go:generate stringer -type=Direction

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

func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("invalid direction")
	}
}
