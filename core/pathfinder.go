package core

import (
	"sort"

	"github.com/rabidaudio/tactics/core/units"
)

type Pathfinder struct {
	Start   units.TPoint
	End     units.TPoint
	CanMove func(pt units.TPoint) bool
	visited []units.TPoint
}

// Find returns a set of steps to get from Start to End,
// if such a path is possible.
func (p *Pathfinder) Find() (steps []units.Direction, ok bool) {
	// TODO [performance] the recursion makes this really elegant but
	// requires a lot of allocations which can be optimized away
	if p.Start == p.End {
		return []units.Direction{}, true
	}
	steps = make([]units.Direction, 0, int(p.End.Sub(p.Start).Mag()))
	p.visited = append(p.visited, p.Start)
LOOP:
	for _, d := range p.testDirections() {
		target := p.Start.Add(d.TP())
		if !p.CanMove(target) {
			continue
		}
		for _, t := range p.visited {
			// TODO use set
			if t == target {
				continue LOOP
			}
		}
		rp := Pathfinder{
			Start:   target,
			End:     p.End,
			CanMove: p.CanMove,
			visited: p.visited,
		}
		if path, ok := rp.Find(); ok {
			steps = append(steps, d)
			steps = append(steps, path...)
			return steps, true
		}
	}
	return nil, false
}

func (p *Pathfinder) testDirections() []units.Direction {
	// try all legal directions, but try the direct route first

	directions := []units.Direction{
		units.East, units.West, units.North, units.South,
	}
	sort.Slice(directions, func(i, j int) bool {
		return p.End.Sub(p.Start.Add(directions[i].TP())).Mag() < p.End.Sub(p.Start.Add(directions[j].TP())).Mag()
	})
	return directions
}
