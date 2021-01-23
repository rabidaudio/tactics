package unit

import (
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
)

type UnitState interface {
	Tick()
	Handle(c core.Command)
}

type idleState struct {
	unit *Unit
}

func (u *Unit) idle() *Unit {
	u.Sprite = u.Animations.Idle().Loop(true)
	u.state = idleState{unit: u}
	return u
}

func (s idleState) Tick() {
}

func (s idleState) Handle(c core.Command) {
	switch cmd := c.(type) {
	case MoveCommand:
		if len(cmd.steps) > 0 {
			s.unit.walk(cmd)
		}
	default:
		core.Unexpected(c)
	}
}

type walkingState struct {
	unit      *Unit
	steps     []units.Direction
	animation units.Anim2D
}

func (u *Unit) walk(move MoveCommand) *Unit {
	s := &walkingState{steps: move.steps, unit: u}
	s.nextStep()
	u.Sprite = u.Animations.Walk().Loop(true)
	u.state = s
	return u
}

func (s *walkingState) nextStep() {
	next := s.steps[0]
	s.steps = s.steps[1:]
	speed := s.unit.MoveSpeed * units.TileSize / float64(units.TickRate)
	dist := s.unit.Location.Add(next.TP().IP())
	if next == units.East {
		s.unit.ReverseFacing = false
	} else if next == units.West {
		s.unit.ReverseFacing = true
	}
	s.animation = units.Animate2D(s.unit.Location, dist, speed, nil)
}

func (s *walkingState) Tick() {
	s.unit.Location = s.animation.Tick()
	if s.animation.IsMoving() {
		return // keep on going
	}
	if len(s.steps) == 0 {
		s.unit.idle()
		return
	}
	s.nextStep()
}

func (s *walkingState) Handle(c core.Command) {
	// commands are ignored while moving
}
