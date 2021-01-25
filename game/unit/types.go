package unit

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	BaseStat               = 5
	BaseHP                 = 10
	StatBias               = 1
	SecondAttackSpeedDelta = 3
)

type Team int

func (t Team) Color(cm *ebiten.ColorM) {
	// TODO [graphics] this is good enough for testing
	// but probably looks bad on different kinds of units
	h := 2 * math.Pi * (float64(t) / 8)
	cm.RotateHue(h)
}

type Attack struct {
	Attacker *Unit
	Target   *Unit
	Dmg      int
	Count    int
}
