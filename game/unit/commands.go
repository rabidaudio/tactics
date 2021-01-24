package unit

import (
	"github.com/rabidaudio/tactics/core/units"
)

type MoveCommand struct {
	Unit  *Unit
	Steps []units.Direction
}

const (
	BaseStat               = 5
	BaseHP                 = 10
	StatBias               = 1
	SecondAttackSpeedDelta = 3
)

type AttackCommand struct {
	Unit   *Unit
	Target *Unit
}

func (cmd AttackCommand) Dmg() int {
	// TODO [mechanics]
	dmg := cmd.atk() - cmd.def()
	if dmg < 0 {
		return 0
	}
	return dmg
}

func (cmd AttackCommand) Count() int {
	if (cmd.Unit.Stats.Spd - cmd.Target.Stats.Spd) >= SecondAttackSpeedDelta {
		return 2
	}
	return 1
}

func (cmd AttackCommand) atk() int {
	a := cmd.Unit.Stats.Atk + cmd.Unit.Weapon.DamageLevel
	at := cmd.Unit.Weapon.WeaponType
	dt := cmd.Target.Weapon.WeaponType
	if at.HasTypeAdvantage(dt) {
		a = a + (a / 2)
	}
	return a
}

func (cmd AttackCommand) def() int {
	return cmd.Target.Stats.Def
}
