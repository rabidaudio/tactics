package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

type tile struct {
	prev *tile
	dir  units.Direction
}

type pathfinder struct {
	start, end units.TPoint
	queue      []units.TPoint
	tiles      *map[units.TPoint]*tile
	blocked    bool
	canMove    func(units.TPoint) bool
}

func (pf *pathfinder) cycle() ([]units.Direction, bool) {
	if len(pf.queue) == 0 {
		return nil, false
	}
	current := pf.queue[0]
	directions := _directions
	if !pf.blocked {
		// as an optimization, we try the most direct route
		// until we hit a barrier. From there we breadth-first search
		directions = []units.Direction{direction(current, pf.end)}
	}
	for _, dir := range directions {
		target := current.Add(dir.TP())
		if !pf.canMove(target) {
			if !pf.blocked {
				pf.blocked = true
				// don't pop the element off the queue, we want to search it again, but
				// from all directions
				return nil, false
			}
			continue
		}
		if _, ok := (*pf.tiles)[target]; ok {
			// already a shorter path there
			continue
		}
		(*pf.tiles)[target] = &tile{prev: (*pf.tiles)[current], dir: dir}
		if target == pf.end {
			results := make([]units.Direction, 0)
			t := (*pf.tiles)[pf.end]
			for t.prev != nil {
				results = append(results, t.dir)
				t = t.prev
			}
			return results, true
		}
		pf.queue = append(pf.queue, target)
	}
	if len(*pf.tiles) >= _maxsearch {
		panic("search exceeded the max number of steps")
	}
	pf.queue = pf.queue[1:]
	return nil, false
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
	pf := pathfinder{
		end:     start,
		start:   end,
		queue:   []units.TPoint{end},
		tiles:   &tiles,
		blocked: false,
		canMove: canMove,
	}
	for {
		if results, ok := pf.cycle(); ok {
			for i, d := range results {
				// because we went from end to start, we need to reverse
				// the directions to do the opposite
				results[i] = d.Opposite()
			}
			return results, true
		}
		if len(pf.queue) == 0 {
			// exausted all available paths
			return nil, false
		}
	}
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
