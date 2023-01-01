package team

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Team int

const (
	Player Team = iota
	Enemy
)

func (t Team) Color(cm *ebiten.ColorM) {
	// TODO [graphics] this is good enough for testing
	// but probably looks bad on different kinds of units
	h := 2 * math.Pi * (float64(t) / 8)
	cm.RotateHue(h)
}
