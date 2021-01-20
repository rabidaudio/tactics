package core

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/rabidaudio/tactics/core/units"
)

type tile struct {
	prev  *tile
	point units.TPoint
	dir   units.Direction
}

var directions = []units.Direction{units.North, units.South, units.East, units.West}

// FindPath returns a set of steps to get from start to end,
// if such a path is possible.
func FindPath(start, end units.TPoint, canMove func(pt units.TPoint) bool) ([]units.Direction, bool) {
	if start == end {
		return []units.Direction{}, true
	}
	if canMove == nil {
		// TODO can we detect when a solution is impossible, even if there
		// are no outer bounds?
		canMove = canAlwaysMove
	}
	if !canMove(end) {
		return nil, false
	}
	vset := mapset.NewThreadUnsafeSet()
	tiles := make(map[units.TPoint]*tile, 1)
	tiles[end] = &tile{point: end}
FOUND:
	for {
		added := 0
		// TODO can optimize by keeping track of where we can skip to
		// TODO could potentially remove the set by knowing that we expand
		// outwards in a diamond, so only check outwards
		for _, current := range tiles {
			for _, d := range directions {
				target := current.point.Add(d.TP())
				if vset.Contains(target) {
					continue
				}
				if !canMove(target) {
					continue
				}
				added++
				vset.Add(target)
				tiles[target] = &tile{prev: current, point: target, dir: d.Opposite()}
				if target == start {
					break FOUND
				}
			}
		}
		if added == 0 {
			// exausted all possible paths
			return nil, false
		}
	}
	results := make([]units.Direction, 0)
	s := tiles[start]
	for s.point != end {
		results = append(results, s.dir)
		s = s.prev
	}
	return results, true
}

func canAlwaysMove(pt units.TPoint) bool {
	return true
}

func delta(start, end units.TPoint) int {
	x, y := start.Sub(end).XY()
	return abs(x) + abs(y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
