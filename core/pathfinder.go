package core

import (
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
		// TODO can we detect when a solution is impossible, even if there
		// are no outer bounds?
		canMove = canAlwaysMove
	}
	if !canMove(end) {
		return nil, false
	}
	vset := mapset.NewThreadUnsafeSet()
	steps := make([]step, 1, guessPossibleSteps(start, end))
	steps[0] = step{point: end}
FOUND:
	for {
		added := 0
		// TODO can optimize by keeping track of where we can skip to
		// TODO could potentially remove the set by knowing that we expand
		// outwards in a diamond, so only check outwards
		for i, current := range steps {
			for _, d := range testDirections(current.point, start) {
				target := current.point.Add(d.TP())
				if vset.Contains(target) {
					continue
				}
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
	results := make([]units.Direction, 0, guessSteps(len(steps)))
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

func guessPossibleSteps(start, end units.TPoint) int {
	// search is in a diamond pattern from the start.
	// each round we add another layer to the diamond.
	// each layer ads 4*n steps.
	d := delta(start, end)
	s := 1
	for i := 0; i < d; i++ {
		s += i * 4
	}
	return s
}

func guessSteps(possibleSteps int) int {
	// this is essentially a revverse of guessPossibleSteps.
	// we take the number of possible steps we visited, and
	// determine how far away the point must be
	if possibleSteps == 0 {
		return 0
	}
	s := possibleSteps - 1
	i := 0
	for ; s >= 0; i++ {
		s -= i * 4
	}
	return i - 1
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
	x, y := start.Sub(end).XY()
	return abs(x) + abs(y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
