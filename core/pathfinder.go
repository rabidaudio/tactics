package core

import (
	"github.com/rabidaudio/tactics/core/units"
)

type tile struct {
	prev       *tile
	point      units.TPoint
	dir        units.Direction
	accessible bool
}

const maxsteps = 1024
const maxsearch = (((2*(maxsteps) - 1) * (2*(maxsteps) - 1)) - 1) / 2

// type pathfinder struct {
// 	start, end units.TPoint
// 	tiles      map[units.TPoint]*tile
// 	canMove    func(units.TPoint) bool
// }

// func (pf *pathfinder) cycle() ([]units.Direction, bool) {

// }

// func (pf *pathfinder) found() []units.Direction {
// 	results := make([]units.Direction, 0)
// 	t := pf.tiles[pf.start]
// 	for t.prev != nil {
// 		results = append(results, t.dir)
// 	}
// 	return results
// }

// func print(tiles map[units.TPoint]*tile) {
// 	// return
// 	// minx, miny := -12, -12
// 	// maxx, maxy := 12, 12
// 	// minx, miny, maxx, maxy := 9, 9, 9, 9
// 	// minx, miny, maxx, maxy := 0, 0, 0, 0
// 	// for t := range tiles {
// 	// 	if t.X < minx {
// 	// 		minx = t.X
// 	// 	}
// 	// 	if t.X > maxx {
// 	// 		maxx = t.X
// 	// 	}
// 	// 	if t.Y < miny {
// 	// 		miny = t.Y
// 	// 	}
// 	// 	if t.Y > maxy {
// 	// 		maxy = t.Y
// 	// 	}
// 	// }
// 	r := [25][25]rune{}
// 	// r := make([][]rune, maxx-minx+1)
// 	for i := 0; i < 25; i++ {
// 		// y := make([]rune, maxy-miny+1)
// 		for j := 0; j < 25; j++ {
// 			// y[j] = ' '
// 			r[i][j] = ' '
// 		}
// 		// r[i] = y
// 	}
// 	for _, t := range tiles {
// 		x := t.point.X + 12
// 		y := 12 - t.point.Y
// 		if !t.accessible {
// 			r[y][x] = 'X'
// 		} else if t.prev != nil {
// 			r[y][x] = '*'
// 		} else {
// 			r[y][x] = '?'
// 		}
// 	}
// 	for i := 24; i >= 0; i-- {
// 		log.Println(string(r[i][:]))
// 	}
// }

var directions = []units.Direction{units.North, units.East, units.South, units.West}

// func check(tiles map[units.TPoint]*tile, tp units.TPoint, start, end units.TPoint, canMove func(units.TPoint) bool) bool {
// 	current := tiles[tp]
// 	if current == nil {
// 		// put a tile there, but no path yet to get there. we may come back to it later
// 		tiles[tp] = &tile{prev: nil, point: tp, accessible: canMove(tp)}
// 		log.Printf("%v not currently accessible", tp)
// 		print(tiles)
// 		return false
// 	}
// 	if !canMove(tp) {
// 		return false
// 	}
// 	for _, d := range directions {
// 		target := current.point.Add(d.TP())
// 		if t, ok := tiles[target]; ok && t.prev != nil && target != end {
// 			// log.Printf("%v already visited", target)
// 			continue // already visited
// 		}
// 		if !canMove(target) {
// 			tiles[target] = &tile{prev: nil, point: target, accessible: false}
// 			log.Printf("%v blocked", target)
// 			print(tiles)
// 			continue // can't get there
// 		}
// 		if _, ok := tiles[target]; ok {
// 			log.Printf("previously visited but wasn't accessible at the time")
// 			tiles[target] = &tile{prev: current, point: target, dir: d.Opposite(), accessible: true}
// 			if check(tiles, target, start, end, canMove) {
// 				return true
// 			}
// 		}
// 		log.Printf("found path from %v to %v", current.point, target)
// 		tiles[target] = &tile{prev: current, point: target, dir: d.Opposite(), accessible: true}
// 		print(tiles)
// 		if target == start {
// 			return true
// 		}
// 	}
// 	return false
// }

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
			pf.tiles[target] = &tile{point: target, accessible: false} // TODO not required
			continue
		}
		if _, ok := pf.tiles[target]; ok {
			// already a shorter path there
			continue
		}
		pf.tiles[target] = &tile{prev: current, point: target, accessible: true, dir: d.Opposite()}
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
		// print(pf.tiles)
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
		canMove = canAlwaysMove
	}
	if !canMove(end) || !canMove(start) {
		return nil, false
	}

	pf := pathfinder{
		start: start,
		end:   end,
		tiles: map[units.TPoint]*tile{
			end: {prev: nil, point: end, accessible: true},
		},
		canMove: canMove,
		queue:   []units.TPoint{end},
	}
	return pf.find()

	// 	// pf := pathfinder{start: start, end: end, tiles: make(map[units.TPoint]*tile, 0), canMove: canMove}
	// 	tiles := make(map[units.TPoint]*tile, 0)
	// 	// start by checking around the start
	// 	tiles[end] = &tile{prev: nil, point: end, accessible: true}
	// 	t := end
	// FOUND:
	// 	for cycle := 0; true; cycle++ {
	// 		log.Printf("== cycle %v ==", cycle)
	// 		if cycle == 0 {
	// 			if check(tiles, end, start, end, canMove) {
	// 				break FOUND
	// 			}
	// 			continue
	// 		}
	// 		added := len(tiles)
	// 		// north
	// 		t = t.Add(units.TP(0, -1))
	// 		if check(tiles, t, start, end, canMove) {
	// 			break FOUND
	// 		}
	// 		// north ---> east
	// 		for i := 0; i < cycle; i++ {
	// 			t = t.Add(units.TP(1, 1))
	// 			if check(tiles, t, start, end, canMove) {
	// 				break FOUND
	// 			}
	// 		}
	// 		// east --> south
	// 		for i := 0; i < cycle; i++ {
	// 			t = t.Add(units.TP(-1, 1))
	// 			if check(tiles, t, start, end, canMove) {
	// 				break FOUND
	// 			}
	// 		}
	// 		// south --> west
	// 		for i := 0; i < cycle; i++ {
	// 			t = t.Add(units.TP(-1, -1))
	// 			if check(tiles, t, start, end, canMove) {
	// 				break FOUND
	// 			}
	// 		}
	// 		// west --> north
	// 		for i := 0; i < cycle; i++ {
	// 			t = t.Add(units.TP(1, -1))
	// 			if check(tiles, t, start, end, canMove) {
	// 				break FOUND
	// 			}
	// 		}

	// 		// print(tiles)
	// 		if len(tiles)-added == 0 {
	// 			// for k, v := range tiles {
	// 			// 	if !v.revisted {

	// 			// 	}
	// 			// }
	// 			// we went around the entire cycle and there was nowhere to go
	// 			return nil, false
	// 		}
	// 		if len(tiles) >= maxsearch {
	// 			panic("search exceeded the max number of steps")
	// 		}
	// 		// walkable := 0
	// 		// for _, tile := range tiles {
	// 		// 	if tile.accessible && tile.prev == nil && tile.point != end {
	// 		// 		// revisit that one
	// 		// 		t = tile.point
	// 		// 	}
	// 		// }
	// 	}
	// 	results := make([]units.Direction, 0)
	// 	tt := tiles[start]
	// 	for tt.prev != nil {
	// 		results = append(results, tt.dir)
	// 		tt = tt.prev
	// 	}
	// 	return results, true

	// TODO: might be able to resolve the infinite problem by alternating starting from the start and from the end

	// 	// vset := mapset.NewThreadUnsafeSet()
	// 	tiles := make(map[units.TPoint]*tile, 1)
	// 	tiles[end] = &tile{point: end}
	// FOUND:
	// 	for {
	// 		added := 0
	// 		// TODO can optimize by keeping track of where we can skip to
	// 		// TODO could potentially remove the set by knowing that we expand
	// 		// outwards in a diamond, so only check outwards
	// 		for point, current := range tiles {
	// 			if current == nil {
	// 				continue // can't go there
	// 			}
	// 			for _, d := range directions {
	// 				target := point.Add(d.TP())
	// 				// if vset.Contains(target) {
	// 				// 	continue
	// 				// }
	// 				if t, ok := tiles[target]; ok && t == nil {
	// 					continue // can't go there
	// 				}
	// 				if !canMove(target) {
	// 					tiles[target] = nil
	// 					continue
	// 				}
	// 				added++
	// 				// vset.Add(target)
	// 				tiles[target] = &tile{prev: current, point: target, dir: d.Opposite()}
	// 				if target == start {
	// 					break FOUND
	// 				}
	// 			}
	// 		}
	// 		if added == 0 {
	// 			// exausted all possible paths
	// 			return nil, false
	// 		}
	// 	}
	// 	results := make([]units.Direction, 0)
	// 	s := tiles[start]
	// 	for s.prev != nil {
	// 		results = append(results, s.dir)
	// 		s = s.prev
	// 	}
	// 	return results, true
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
