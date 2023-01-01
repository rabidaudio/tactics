package game

import (
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/unit"
)

type Command interface {
	IsLegal(s core.Scene) bool
	Execute(s core.Scene)
}

type MoveCommand struct {
	Unit  *unit.Unit
	Dest  units.TPoint
	Steps []units.Direction
}

func (b *BattleScene) NewMoveCommand(w *core.World, unit *unit.Unit, dest units.TPoint) MoveCommand {
	if unit.Location.StepsTo(dest) > core.MaxSearchSteps {
		return MoveCommand{Unit: unit, Dest: dest, Steps: nil}
	}
	steps, _ := core.FindPath(unit.Location, dest, func(pt units.TPoint) bool {
		// TODO [style] share logic better with canMoveTo
		if w.IsBoundary(pt) {
			return false
		}
		// can move through friendly units but not enemy units
		// TODO [mechanics] desired behavior?
		if u := b.UnitAt(pt); u != nil && u.Team != unit.Team {
			return false
		}
		return true
	})
	return MoveCommand{Unit: unit, Dest: dest, Steps: steps}
}

var _ Command = (*MoveCommand)(nil)

func (cmd MoveCommand) IsLegal(s core.Scene) bool {
	b, ok := s.(*BattleScene)
	if !ok {
		return false
	}
	if cmd.Steps == nil || len(cmd.Steps) == 0 {
		return false
	}
	if len(cmd.Steps) > cmd.Unit.StepsPerTurn() {
		return false
	}
	if u := b.UnitAt(cmd.Dest); u != nil {
		return false
	}
	if !cmd.Unit.IsReady() {
		return false
	}
	return true
}

func (cmd MoveCommand) Execute(s core.Scene) {
	cmd.Unit.Walk(cmd.Steps...)
}

const (
	BaseStat               = 5
	BaseHP                 = 10
	StatBias               = 1
	SecondAttackSpeedDelta = 3
)

type AttackCommand struct {
	Attacker *unit.Unit
	Target   *unit.Unit
}

var _ Command = (*AttackCommand)(nil)

func (cmd AttackCommand) IsLegal(s core.Scene) bool {
	if _, ok := s.(*BattleScene); !ok {
		return false
	}
	if cmd.Attacker == nil || cmd.Target == nil {
		return false
	}
	if cmd.Attacker.Team == cmd.Target.Team {
		return false // TODO [mechanics] healing
	}
	if !cmd.Attacker.IsReady() || !cmd.Target.IsReady() {
		return false
	}
	if !cmd.Attacker.CanReach(cmd.Target) {
		return false
	}
	return true
}

func (cmd AttackCommand) Execute(s core.Scene) {
	// TODO [style] would an interface be better here?
	atk := unit.Attack{
		Attacker: cmd.Attacker,
		Target:   cmd.Target,
		Count:    cmd.Count(),
		Dmg:      cmd.Dmg(),
	}
	cmd.Attacker.Attack(atk)
	cmd.Target.DefendAgainst(atk)
}
