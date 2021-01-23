package core_test

import (
	"testing"

	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"gotest.tools/v3/assert"
)

func steps(s string) []units.Direction {
	d := make([]units.Direction, len(s))
	for i, r := range s {
		switch r {
		case 'N':
			d[i] = units.North
		case 'S':
			d[i] = units.South
		case 'E':
			d[i] = units.East
		case 'W':
			d[i] = units.West
		}
	}
	return d
}

func bounds(m []rune, width int) func(units.TPoint) bool {
	return func(t units.TPoint) bool {
		i := t.Y*width + t.X
		if t.Y < 0 || t.X < 0 || i < 0 || i >= len(m) {
			return false
		}
		return m[t.Y*width+t.X] == ' '
	}
}

func TestPathfinder(t *testing.T) {
	cases := []struct {
		name      string
		Start     units.TPoint
		End       units.TPoint
		CanMove   func(pt units.TPoint) bool
		expected  []units.Direction
		available bool
	}{
		{
			name:      "right-down",
			Start:     units.TP(0, 0),
			End:       units.TP(1, 1),
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("ES"),
			available: true,
		},
		{
			name:      "up-left",
			Start:     units.TP(0, 0),
			End:       units.TP(-1, -1),
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("WN"),
			available: true,
		},
		{
			name:      "same place",
			Start:     units.TP(0, 0),
			End:       units.TP(0, 0),
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  []units.Direction{},
			available: true,
		},
		{
			name:      "straight line",
			Start:     units.TP(0, 0),
			End:       units.TP(0, -3),
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("NNN"),
			available: true,
		},
		{
			name:      "must go south first",
			Start:     units.TP(-2, -1),
			End:       units.TP(0, 0),
			CanMove:   func(pt units.TPoint) bool { return pt != units.TP(-1, -1) },
			expected:  steps("SEE"),
			available: true,
		},
		{
			name:      "must go east first",
			Start:     units.TP(0, 0),
			End:       units.TP(2, 1),
			CanMove:   func(pt units.TPoint) bool { return pt != units.TP(0, 1) },
			expected:  steps("ESE"),
			available: true,
		},
		{
			name:      "no path",
			Start:     units.TP(0, 0),
			End:       units.TP(2, 2),
			CanMove:   func(pt units.TPoint) bool { return false },
			expected:  []units.Direction{},
			available: false,
		},
		{
			name:      "one step",
			Start:     units.TP(0, 0),
			End:       units.TP(1, 0),
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("E"),
			available: true,
		},
		{
			name:  "from the start",
			Start: units.TP(0, 0),
			End:   units.TP(2, 0),
			// ' ', ' ', ' ',
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("EE"),
			available: true,
		},
		{
			name:  "from the end",
			Start: units.TP(0, 0),
			End:   units.TP(3, 0),
			// ' ', ' ', ' ', ' ',
			CanMove:   func(pt units.TPoint) bool { return true },
			expected:  steps("EEE"),
			available: true,
		},
		{
			name:  "backtracking",
			Start: units.TP(0, 1),
			End:   units.TP(4, 1),
			CanMove: bounds([]rune{
				' ', ' ', ' ', ' ', ' ',
				' ', ' ', 'X', ' ', ' ',
				'X', 'X', 'X', 'X', 'X',
			}, 5),
			expected:  steps("ENEESE"),
			available: true,
		},
		{
			name:  "complex",
			Start: units.TP(4, 1),
			End:   units.TP(0, 3),
			CanMove: bounds([]rune{
				'X', 'X', 'X', 'X', 'X',
				' ', ' ', ' ', 'X', ' ',
				' ', 'X', ' ', 'X', ' ',
				' ', 'X', ' ', ' ', ' ',
				'X', 'X', 'X', 'X', 'X',
			}, 5),
			expected:  steps("SSWWNNWWSS"),
			available: true,
		},
		{
			name:  "short path one",
			Start: units.TP(2, 4),
			End:   units.TP(0, 2),
			CanMove: bounds([]rune{
				' ', ' ', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', ' ', ' ',
			}, 3),
			expected:  steps("SWWNNN"),
			available: true,
		},
		{
			name:  "short path two",
			Start: units.TP(2, 3),
			End:   units.TP(0, 1),

			CanMove: bounds([]rune{
				' ', ' ', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', 'X', ' ',
				' ', ' ', ' ',
			}, 3),
			expected:  steps("NNNWWS"),
			available: true,
		},
		{
			name:  "blocked",
			Start: units.TP(6, 4),
			End:   units.TP(3, 4),
			CanMove: bounds([]rune{
				'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
				'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X',
				'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X',
				'X', ' ', 'X', 'X', 'X', ' ', ' ', ' ', 'X',
				'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X',
				'X', ' ', 'X', 'X', 'X', ' ', ' ', ' ', 'X',
				'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X',
				'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X',
				'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
			}, 9),
			available: false,
		},
		{
			name:  "blocked unbounded",
			Start: units.TP(3, 4),
			End:   units.TP(6, 4),
			CanMove: func(pt units.TPoint) bool {
				switch pt {
				case units.TP(2, 3),
					units.TP(3, 3),
					units.TP(4, 3),
					units.TP(4, 4),
					units.TP(4, 5),
					units.TP(3, 5),
					units.TP(2, 5),
					units.TP(2, 4):
					return false
				default:
					return true
				}
			},
			available: false,
		},
		{
			name:  "maze",
			Start: units.TP(1, 1),
			End:   units.TP(3, 1),
			CanMove: bounds([]rune{
				'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
				'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X',
				'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X',
				'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X',
				'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X',
				'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X',
				'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X',
				'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X',
				'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
			}, 9),
			expected:  steps("SSSSSSEENEESEENNNNNNWWSWWN"),
			available: true,
		},
		{
			name:  "diagonal",
			Start: units.TP(6, 6),
			End:   units.TP(1, 0),
			CanMove: bounds([]rune{
				'X', ' ', 'X', 'X', 'X', 'X', 'X',
				'X', ' ', 'X', 'X', 'X', 'X', 'X',
				'X', ' ', 'X', 'X', 'X', 'X', 'X',
				'X', ' ', 'X', 'X', 'X', 'X', ' ',
				'X', ' ', 'X', 'X', 'X', ' ', ' ',
				'X', ' ', 'X', 'X', ' ', ' ', ' ',
				'X', ' ', 'X', ' ', ' ', ' ', ' ',
				'X', ' ', ' ', ' ', ' ', ' ', ' ',
			}, 7),
			expected:  steps("WSWWWWNNNNNNN"),
			available: true,
		},
	}
	t.Parallel()
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			r, ok := core.FindPath(test.Start, test.End, test.CanMove)
			if test.available {
				if !ok {
					t.Errorf("expected path")
				} else {
					assert.DeepEqual(t, test.expected, r)
				}
			} else {
				if ok {
					t.Errorf("expected no path but found %v", r)
				}
			}
		})
	}
}

//  1495 ns/op	    2103 B/op	      61 allocs/op
func BenchmarkFindPathOpen(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		canMove := bounds([]rune{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		}, 9)
		start, end := units.TP(0, 0), units.TP(8, 8)
		for pb.Next() {
			core.FindPath(start, end, canMove)
		}
	})
}

// 4247 ns/op	    5366 B/op	     100 allocs/op
func BenchmarkFindPathMaze(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		canMove := bounds([]rune{
			'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
			'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X',
			'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X',
			'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X',
			'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X',
			'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X',
			'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X',
			'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X',
			'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X',
		}, 9)
		start, end := units.TP(1, 1), units.TP(3, 1)
		for pb.Next() {
			core.FindPath(start, end, canMove)
		}
	})
}

// 62752 ns/op	   95450 B/op	    1569 allocs/op
func BenchmarkFindPathFar(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		canMove := func(_ units.TPoint) bool { return true }
		start, end := units.TP(0, 0), units.TP(256, 256)
		for pb.Next() {
			core.FindPath(start, end, canMove)
		}
	})
}

// 77175 ns/op	  102949 B/op	    1735 allocs/op
func BenchmarkFindPathFarBlocked(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		canMove := func(p units.TPoint) bool {
			return !((p.X == 127 || p.X == 128) && (p.Y == 127 || p.Y == 128))
		}
		start, end := units.TP(0, 0), units.TP(255, 255)
		for pb.Next() {
			core.FindPath(start, end, canMove)
		}
	})
}

// 558517218 ns/op	48775340 B/op	  535247 allocs/op
func BenchmarkFindPathFarWorstCase(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		canMove := func(p units.TPoint) bool {
			return p != units.TP(0, 1) && p != units.TP(1, 0) &&
				p != units.TP(254, 255) && p != units.TP(255, 254)
		}
		start, end := units.TP(0, 0), units.TP(255, 255)
		for pb.Next() {
			core.FindPath(start, end, canMove)
		}
	})
}
