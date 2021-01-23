package unit

import "github.com/rabidaudio/tactics/core/units"

type MoveCommand struct {
	Unit  *Unit
	Steps []units.Direction
}

type AttackCommand struct {
	Unit   *Unit
	Target *Unit
}

func (cmd AttackCommand) Dmg() int {
	// TODO [mechanics]
	return 0
}
