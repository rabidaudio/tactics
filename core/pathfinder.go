package core

import (
	"sort"

	mapset "github.com/deckarep/golang-set"
	"github.com/rabidaudio/tactics/core/units"
)

// FindPath returns a set of steps to get from start to end,
// if such a path is possible.
func FindPath(start, end units.TPoint, canMove func(pt units.TPoint) bool) ([]units.Direction, bool) {
	p := pathfinder{
		start:   start,
		current: end,
		canMove: canMove,
		steps:   make([]units.Direction, 0, int(end.Sub(start).Mag())),
		visited: mapset.NewSet(),
	}
	if canMove == nil {
		p.canMove = canAlwaysMove
	}
	ok := p.find()
	return p.steps, ok
}

func canAlwaysMove(pt units.TPoint) bool {
	return true
}

type pathfinder struct {
	start   units.TPoint
	current units.TPoint
	canMove func(pt units.TPoint) bool
	visited mapset.Set
	steps   []units.Direction
}

func (p *pathfinder) find() bool {
	// so that we don't have to reverse the steps when we get a match,
	// we actually search from end to start
	if p.current == p.start {
		return true
	}
	p.visited.Add(p.current)
	for _, d := range p.testDirections() {
		target := p.current.Add(d.TP())
		if !p.canMove(target) {
			continue
		}
		if p.visited.Contains(target) {
			continue
		}
		p.current = target
		if p.find() {
			// since we're going backwards, we need to track the reverse
			p.steps = append(p.steps, d.Opposite())
			return true
		}
	}
	return false
}

func (p *pathfinder) testDirections() []units.Direction {
	// try all legal directions, but try the most direct routes first
	directions := []units.Direction{
		units.North, units.South, units.East, units.West,
	}
	sort.Slice(directions, func(i, j int) bool {
		return p.distWith(directions[i]) < p.distWith(directions[j])
	})
	return directions
}

func (p *pathfinder) distWith(d units.Direction) int {
	return int(p.start.Sub(p.current.Add(d.TP())).Mag())
	// q := p.current.Add(d.TP())
	// dx := p.start.X - q.X
	// dy := p.start.Y - q.Y
	// // log.Printf("dir: %v int: %v -- float: %v", d, abs(dx)+abs(dy), int(p.start.Sub(p.current.Add(d.TP())).Mag()))
	// return abs(dx) + abs(dy)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
