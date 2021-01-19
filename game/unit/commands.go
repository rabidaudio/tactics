package unit

import "github.com/rabidaudio/tactics/core/units"

type MoveCommand struct {
	unit  *Unit
	steps []units.Direction
}

func Move(unit *Unit, steps ...units.Direction) MoveCommand {
	return MoveCommand{unit: unit, steps: steps}
}
