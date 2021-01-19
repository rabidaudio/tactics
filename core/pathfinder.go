package core

import (
	"log"
	"sort"

	mapset "github.com/deckarep/golang-set"
	"github.com/rabidaudio/tactics/core/units"
)

type step struct {
	prev  *step
	point units.TPoint
	dir   units.Direction
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
	if !canMove(end) {
		return nil, false
	}
	// TODO heuristic for capacity
	vset := mapset.NewSet()
	steps := make([]step, 1, 25)
	steps[0] = step{point: end}
FOUND:
	for {
		added := 0
		for i, current := range steps {
			for _, d := range testDirections(current.point, start) {
				target := current.point.Add(d.TP())
				if vset.Contains(target) {
					continue
				}
				log.Printf("check %v -> %v go %v to %v", current.point, start, d, target)
				if !canMove(target) {
					continue
				}
				added++
				vset.Add(target)
				steps = append(steps, step{prev: &steps[i], point: target, dir: d.Opposite()})
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
	// TODO: heuristic for capacity
	results := make([]units.Direction, 0, 5)
	s := &steps[len(steps)-1]
	for s.prev != nil {
		results = append(results, s.dir)
		s = s.prev
	}
	return results, true
}

func canAlwaysMove(pt units.TPoint) bool {
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func testDirections(start, end units.TPoint) []units.Direction {
	// try all legal directions, but try the most direct routes first
	directions := []units.Direction{
		units.North, units.South, units.East, units.West,
	}
	sort.Slice(directions, func(i, j int) bool {
		return delta(start.Add(directions[i].TP()), end) < delta(start.Add(directions[j].TP()), end)
	})
	return directions
}

func delta(start, end units.TPoint) int {
	return int(end.Sub(start).Mag())
}
