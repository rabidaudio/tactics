package core_test

import (
	"testing"

	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"gotest.tools/v3/assert"
)

func TestPathfinder(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name      string
		pf        core.Pathfinder
		expected  []units.Direction
		available bool
	}{
		{
			name: "right-down",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(1, 1),
				CanMove: func(pt units.TPoint) bool { return true },
			},
			expected:  []units.Direction{units.East, units.South},
			available: true,
		},
		{
			name: "up-left",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(-1, -1),
				CanMove: func(pt units.TPoint) bool { return true },
			},
			expected:  []units.Direction{units.West, units.North},
			available: true,
		},
		{
			name: "same place",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(0, 0),
				CanMove: func(pt units.TPoint) bool { return true },
			},
			expected:  []units.Direction{},
			available: true,
		},
		{
			name: "straight line",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(0, -3),
				CanMove: func(pt units.TPoint) bool { return true },
			},
			expected:  []units.Direction{units.North, units.North, units.North},
			available: true,
		},
		{
			name: "straight line",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(0, -3),
				CanMove: func(pt units.TPoint) bool { return true },
			},
			expected:  []units.Direction{units.North, units.North, units.North},
			available: true,
		},
		{
			name: "must go south first",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(2, 1),
				CanMove: func(pt units.TPoint) bool { return pt != units.TP(1, 0) },
			},
			expected:  []units.Direction{units.South, units.East, units.East},
			available: true,
		},
		{
			name: "must go east first",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(2, 1),
				CanMove: func(pt units.TPoint) bool { return pt != units.TP(0, 1) },
			},
			expected:  []units.Direction{units.East, units.East, units.South},
			available: true,
		},
		{
			name: "no path",
			pf: core.Pathfinder{
				Start:   units.TP(0, 0),
				End:     units.TP(2, 2),
				CanMove: func(pt units.TPoint) bool { return false },
			},
			expected:  []units.Direction{},
			available: false,
		},
		{
			name: "backtracking",
			pf: core.Pathfinder{
				Start: units.TP(0, 1),
				End:   units.TP(4, 1),
				CanMove: func(pt units.TPoint) bool {
					return [15]int{
						0, 0, 0, 0, 0,
						0, 0, 1, 0, 0,
						1, 1, 1, 1, 1,
					}[pt.Y*5+pt.X] == 0
				},
			},
			expected: []units.Direction{
				units.East, units.North, units.East, units.East,
				units.East, units.South,
			},
			available: true,
		},
		{
			name: "complex",
			pf: core.Pathfinder{
				Start: units.TP(4, 1),
				End:   units.TP(0, 3),
				CanMove: func(pt units.TPoint) bool {
					return [25]int{
						1, 1, 1, 1, 1,
						0, 0, 0, 1, 0,
						0, 1, 0, 1, 0,
						0, 1, 0, 0, 0,
						1, 1, 1, 1, 1,
					}[pt.Y*5+pt.X] == 0
				},
			},
			expected: []units.Direction{
				units.South, units.South,
				units.West, units.West,
				units.North, units.North,
				units.West, units.West,
				units.South, units.South,
			},
			available: true,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			r, ok := test.pf.Find()
			if test.available {
				if !ok {
					t.Errorf("expected path")
				} else {
					assert.DeepEqual(t, r, test.expected)
				}
			} else {
				if ok {
					t.Errorf("expected no path but found %v", r)
				}
			}
		})
	}
}
