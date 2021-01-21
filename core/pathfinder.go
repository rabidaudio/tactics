package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

type tile struct {
	prev *tile
	dir  units.Direction
}

const maxsteps = 1024
const maxsearch = (((2*(maxsteps) - 1) * (2*(maxsteps) - 1)) - 1) / 2

var directions = []units.Direction{units.North, units.East, units.South, units.West}

type pathfinder struct {
	start, end units.TPoint
	canMove    func(units.TPoint) bool
	queue      []units.TPoint
	tiles      map[units.TPoint]*tile
}

func (pf *pathfinder) check(t units.TPoint) bool {
	current := pf.tiles[t]
	for _, d := range directions {
		target := t.Add(d.TP())
		if !pf.canMove(target) {
			continue
		}
		if _, ok := pf.tiles[target]; ok {
			// already a shorter path there
			continue
		}
		pf.tiles[target] = &tile{prev: current, dir: d.Opposite()}
		if target == pf.start {
			return true
		}
		pf.queue = append(pf.queue, target)
	}
	pf.queue = pf.queue[1:]
	return false
}

func (pf *pathfinder) find() ([]units.Direction, bool) {
	for {
		if len(pf.queue) == 0 {
			return nil, false
		}
		if pf.check(pf.queue[0]) {
			results := make([]units.Direction, 0)
			tt := pf.tiles[pf.start]
			for tt.prev != nil {
				results = append(results, tt.dir)
				tt = tt.prev
			}
			return results, true
		}
		if len(pf.tiles) >= maxsearch {
			panic("search exceeded the max number of steps")
		}
	}
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
		// might be able to resolve the problem by alternating starting from the start and from the end
		canMove = canAlwaysMove
	}
	if !canMove(end) || !canMove(start) {
		return nil, false
	}

	pf := pathfinder{
		start: start,
		end:   end,
		tiles: map[units.TPoint]*tile{
			end: {prev: nil},
		},
		canMove: canMove,
		queue:   []units.TPoint{end},
	}
	return pf.find()
}

func canAlwaysMove(pt units.TPoint) bool {
	return true
}
