package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

//https://en.wikipedia.org/wiki/Centered_square_number
const _maxsteps = 1024
const _maxsearch = (((2*(_maxsteps) - 1) * (2*(_maxsteps) - 1)) - 1) / 2

type tile struct {
	prev  *tile
	dir   units.Direction
	steps int
}

type pathfinder struct {
	start   units.TPoint
	dest    units.TPoint
	queue   []units.TPoint
	tiles   map[units.TPoint]*tile
	canMove func(units.TPoint) bool
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
	// in addition to the performance optimization, searching
	// from both ends helps us avoid searching forever when
	// the destination is boxed in and can't be reached
	forward := pathfinder{
		start:   start,
		dest:    end,
		queue:   []units.TPoint{start},
		blocked: false,
		fromEnd: false,
		canMove: canMove,
		tiles: map[units.TPoint]*tile{
			start: {prev: nil, steps: 0},
		},
	}
	reverse := pathfinder{
		start:   end,
		dest:    start,
		queue:   []units.TPoint{end},
		blocked: false,
		fromEnd: true,
		canMove: canMove,
		tiles: map[units.TPoint]*tile{
			end: {prev: nil, steps: 0},
		},
	}

	i := 0
	for {
		if len(reverse.queue) == 0 || len(forward.queue) == 0 {
			// exausted all available paths
			return nil, false
		}
		if reverse.cycle(&forward) || forward.cycle(&reverse) {
			break
		}
		if len(forward.tiles)+len(reverse.tiles) >= _maxsearch {
			panic("search exceeded the max number of steps")
		}
		i++
	}
	results := make([]units.Direction, 0, i*2)
	t := forward.tiles[start]
	for t.prev != nil {
		results = append(results, t.dir)
		t = t.prev
	}
	return results, true
}

func (pf *pathfinder) cycle(other *pathfinder) bool {
	current := pf.queue[0]
	for _, dir := range optimalDirections(current, pf.dest, !pf.blocked) {
		target := current.Add(dir.TP())
		if t, ok := pf.tiles[target]; ok {
			if t.steps <= pf.tiles[current].steps+1 {
				continue // already found a shorter path to this point
			}
			// otherwise replace it with this shorter path
		}
		if t, ok := other.tiles[target]; ok {
			// we met up with the other direction!
			if !pf.fromEnd {
				// if we're on the forward direction, let the reverse direction
				// be the one to find it.
				// searching backwards allows us to walk the `prev` pointers
				// backwards so we don't have to reverse the array at the end
				return false
			}
			pf.resolvePaths(pf.tiles[current], t, dir.Opposite())
			return true
		}
		if !pf.canMove(target) {
			if !pf.blocked {
				pf.blocked = true
				pf.queue = []units.TPoint{current, pf.start}
				// otherwise we'll keep searching from here in all directions,
				// but also be on the lookout for a shorter path via the start
				return false
			}
			continue
		}

		if pf.fromEnd {
			// because we search backwards from end to start, we need to reverse
			// the direction
			dir = dir.Opposite()
		}
		pf.tiles[target] = &tile{prev: pf.tiles[current], dir: dir, steps: pf.tiles[current].steps + 1}
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

func canAlwaysMove(_ units.TPoint) bool {
	return true
}

func optimalDirections(from, to units.TPoint, single bool) []units.Direction {
	dx := abs(to.X - from.X)
	dy := abs(to.Y - from.Y)
	var opta, optb units.Direction
	if to.X > from.X {
		opta = units.East
	} else {
		opta = units.West
	}
	if to.Y > from.Y {
		optb = units.South
	} else {
		optb = units.North
	}
	if dy >= dx {
		opta, optb = optb, opta
	}
	if single {
		return []units.Direction{opta}
	}
	return []units.Direction{opta, optb, optb.Opposite(), opta.Opposite()}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
