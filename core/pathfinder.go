package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

type tile struct {
	prev *tile
	dir  units.Direction
}

//https://en.wikipedia.org/wiki/Centered_square_number
const _maxsteps = 1024
const _maxsearch = (((2*(_maxsteps) - 1) * (2*(_maxsteps) - 1)) - 1) / 2

var _directions = []units.Direction{units.North, units.East, units.South, units.West}

// FindPath returns a set of steps to get from start to end,
// if such a path is possible.
func FindPath(start, end units.TPoint, canMove func(pt units.TPoint) bool) ([]units.Direction, bool) {
	if start == end {
		return []units.Direction{}, true
	}

	if canMove == nil {
		// TODO can we detect when a solution is impossible, even if there
		// are no outer bounds?
		// might be able to resolve the problem by alternating starting from the start and from the end
		canMove = canAlwaysMove
	}
	if !canMove(end) || !canMove(start) {
		return nil, false
	}

	// note: we search backwards from end to start here, so
	// that we can walk the `prev` pointers backwards to find
	// the path
	tiles := map[units.TPoint]*tile{
		end: {prev: nil},
	}
	queue := []units.TPoint{end}
	steps := 0
	blocked := false

FOUND:
	for {
		if len(queue) == 0 {
			return nil, false
		}
		current := queue[0]
		directions := _directions
		if !blocked {
			// as an optimization, we try the most direct route
			// until we hit a barrier. From there we breadth-first search
			directions = []units.Direction{direction(current, start)}
		}
		for _, dir := range directions {
			target := current.Add(dir.TP())
			if !canMove(target) {
				if !blocked {
					blocked = true
					continue FOUND
				}
				continue
			}
			if _, ok := tiles[target]; ok {
				// already a shorter path there
				continue
			}
			tiles[target] = &tile{prev: tiles[current], dir: dir}
			if target == start {
				break FOUND
			}
			queue = append(queue, target)
		}
		if len(tiles) >= _maxsearch {
			panic("search exceeded the max number of steps")
		}
		queue = queue[1:]
		steps++
	}
	results := make([]units.Direction, 0, steps)
	t := tiles[start]
	for t.prev != nil {
		// because we went from end to start, we need to reverse
		// the directions to do the opposite
		results = append(results, t.dir.Opposite())
		t = t.prev
	}
	return results, true
}

func canAlwaysMove(_ units.TPoint) bool {
	return true
}

func direction(from, to units.TPoint) units.Direction {
	dx := abs(to.X - from.X)
	dy := abs(to.Y - from.Y)
	if dx >= dy {
		if to.X > from.X {
			return units.East
		}
		return units.West
	}
	if to.Y > from.Y {
		return units.South
	}
	return units.North
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
