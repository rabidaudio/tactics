package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

//https://en.wikipedia.org/wiki/Centered_square_number
const _maxsteps = 1024
const _maxsearch = (((2*(_maxsteps) - 1) * (2*(_maxsteps) - 1)) - 1) / 2

var _directions = []units.Direction{units.North, units.East, units.South, units.West}

type tile struct {
	prev    *tile
	dir     units.Direction
	fromEnd bool
}

type pathfinder struct {
	dest    units.TPoint
	queue   []units.TPoint
	blocked bool
	fromEnd bool
}

// FindPath returns a set of steps to get from start to end,
// if such a path is possible.
func FindPath(start, end units.TPoint, canMove func(pt units.TPoint) bool) ([]units.Direction, bool) {
	if start == end {
		return []units.Direction{}, true
	}
	if canMove == nil {
		canMove = canAlwaysMove
	}
	if !canMove(end) || !canMove(start) {
		return nil, false
	}

	// we make two finders, one that searches from the end
	// and one from the start. They alternate, each searching
	// one more tile.
	// searching backwards allows us to walk the `prev` pointers
	// backwards so we don't have to reverse the array
	tiles := map[units.TPoint]*tile{
		start: {prev: nil, fromEnd: false},
		end:   {prev: nil, fromEnd: true},
	}
	forward := pathfinder{
		dest:    end,
		queue:   []units.TPoint{start},
		blocked: false,
		fromEnd: false,
	}
	reverse := pathfinder{
		dest:    start,
		queue:   []units.TPoint{end},
		blocked: false,
		fromEnd: true,
	}

	i := 0
	for {
		// alternate reverse and forward paths
		pf := &reverse
		if i%2 != 0 {
			pf = &forward
		}
		if len(pf.queue) == 0 {
			// exausted all available paths
			return nil, false
		}
		if pf.cycle(tiles, canMove) {
			results := make([]units.Direction, 0, i)
			t := tiles[start]
			for t.prev != nil {
				results = append(results, t.dir)
				t = t.prev
			}
			return results, true
		}
		if len(tiles) >= _maxsearch {
			panic("search exceeded the max number of steps")
		}
		i++
	}
}

func (pf *pathfinder) cycle(tiles map[units.TPoint]*tile, canMove func(units.TPoint) bool) bool {
	current := pf.queue[0]
	for _, dir := range pf.directions(current) {
		target := current.Add(dir.TP())
		if !canMove(target) {
			if !pf.blocked {
				pf.blocked = true
				// search it again, but next time from all directions instead
				// of just the most direct one
				return pf.cycle(tiles, canMove)
			}
			continue
		}
		if t, ok := tiles[target]; ok {
			if t.fromEnd == pf.fromEnd {
				continue // already found a shorter path to this point
			}

			// we met up with the other direction!
			if !pf.fromEnd {
				// let the reverse path find it, so that we don't have
				// to reverse the result
				return false
			}
			pf.resolvePaths(tiles[current], t, dir.Opposite())
			return true
		}
		if pf.fromEnd {
			// because we search backwards from end to start, we need to reverse
			// the direction
			dir = dir.Opposite()
		}
		tiles[target] = &tile{prev: tiles[current], dir: dir, fromEnd: pf.fromEnd}
		pf.queue = append(pf.queue, target)
	}
	pf.queue = pf.queue[1:] // pop
	return false
}

func (pf *pathfinder) resolvePaths(current, target *tile, dir units.Direction) {
	for {
		next := target.prev
		target.prev = current
		dtemp := target.dir
		target.dir = dir
		if next == nil {
			return
		}
		dir = dtemp
		current = target
		target = next
	}
}

func (pf *pathfinder) directions(from units.TPoint) []units.Direction {
	if pf.blocked {
		// try all directions
		return _directions
	}
	// as an optimization, we try the most direct route
	// until we hit a barrier. From there we breadth-first search
	return []units.Direction{direction(from, pf.dest)}
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
