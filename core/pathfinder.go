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
	forward := pathfinder{
		start:   start,
		dest:    end,
		queue:   []units.TPoint{start},
		blocked: false,
		fromEnd: false,
		tiles: map[units.TPoint]*tile{
			start: {prev: nil},
		},
	}
	reverse := pathfinder{
		start:   end,
		dest:    start,
		queue:   []units.TPoint{end},
		blocked: false,
		fromEnd: true,
		tiles: map[units.TPoint]*tile{
			end: {prev: nil},
		},
	}

	i := 0
	for {
		if len(reverse.queue) == 0 || len(forward.queue) == 0 {
			// exausted all available paths
			return nil, false
		}
		if reverse.cycle(&forward, canMove) || forward.cycle(&reverse, canMove) {
			break
		}
		// log.Printf("blocked: %v queue: %v", forward.blocked, forward.queue)
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

func (pf *pathfinder) cycle(other *pathfinder, canMove func(units.TPoint) bool) bool {
	current := pf.queue[0]
	for _, dir := range pf.directions(current) {
		target := current.Add(dir.TP())
		if t, ok := pf.tiles[target]; ok {
			if t.steps <= pf.tiles[current].steps+1 {
				continue // already found a shorter path to this point
			}
			// log.Printf("shorter path")
			// otherwise replace it with this shorter path
		}
		if t, ok := other.tiles[target]; ok {
			// we met up with the other direction!
			if !pf.fromEnd {
				// let the reverse path find it, so that we don't have
				// to reverse the result
				return false
			}
			pf.resolvePaths(pf.tiles[current], t, dir.Opposite())
			return true
		}
		if !canMove(target) {
			if !pf.blocked {
				pf.blocked = true
				pf.queue = []units.TPoint{current, pf.start}
				// for k := range pf.tiles {
				// 	if k != pf.start {
				// 		delete(pf.tiles, k)
				// 	}
				// }
				// search it again, but next time from all directions instead
				// of just the most direct one
				// return pf.cycle(other, canMove)
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

// var _directions = []units.Direction{units.North, units.East, units.South, units.West}

func (pf *pathfinder) directions(from units.TPoint) []units.Direction {
	return directions(from, pf.dest, !pf.blocked)
	// optimal := direction(from, pf.dest)
	// if pf.blocked {
	// 	// try all directions, starting with the most optimal one
	// 	return []units.Direction{optimal, units.North, units.East, units.South, units.West}
	// }
	// // as an optimization, we try the most direct route
	// // until we hit a barrier. From there we breadth-first search
	// return []units.Direction{optimal}
}

func canAlwaysMove(_ units.TPoint) bool {
	return true
}

func directions(from, to units.TPoint, single bool) []units.Direction {
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

// func direction(from, to units.TPoint) units.Direction {
// 	dx := abs(to.X - from.X)
// 	dy := abs(to.Y - from.Y)
// 	if dx >= dy {
// 		if to.X > from.X {
// 			return units.East
// 		}
// 		return units.West
// 	}
// 	if to.Y > from.Y {
// 		return units.South
// 	}
// 	return units.North
// }

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
