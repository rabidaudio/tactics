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

	tiles := map[units.TPoint]*tile{
		end: {prev: nil},
	}
	queue := []units.TPoint{end}
FOUND:
	for {
		if len(queue) == 0 {
			return nil, false
		}
		t := queue[0]
		current := tiles[t]
		for _, d := range _directions {
			target := t.Add(d.TP())
			if !canMove(target) {
				continue
			}
			if _, ok := tiles[target]; ok {
				// already a shorter path there
				continue
			}
			tiles[target] = &tile{prev: current, dir: d.Opposite()}
			if target == start {
				break FOUND
			}
			queue = append(queue, target)
		}
		queue = queue[1:]
		if len(tiles) >= _maxsearch {
			panic("search exceeded the max number of steps")
		}
	}
	results := make([]units.Direction, 0)
	tt := tiles[start]
	for tt.prev != nil {
		results = append(results, tt.dir)
		tt = tt.prev
	}
	return results, true
}

func canAlwaysMove(pt units.TPoint) bool {
	return true
}
