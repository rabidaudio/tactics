package unit

import (
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
)

type UnitState interface {
	Tick()
	Handle(c core.Command)
}

type idleState struct{}

func (u *Unit) idle() {
	u.Sprite = u.Animations.Idle().Loop(true)
	u.state = idleState{}
}

func (s idleState) Tick() {
}

func (s idleState) Handle(c core.Command) {
	// all commands accepted in idle mode
	switch cmd := c.(type) {
	case MoveCommand:
		cmd.Unit.walk(cmd)
	case AttackCommand:
		cmd.Unit.attack(cmd)
		cmd.Target.defend(cmd)
	default:
		core.Unexpected(c)
	}
}

type walkingState struct {
	unit      *Unit
	steps     []units.Direction
	animation units.Anim2D
}

func (u *Unit) walk(move MoveCommand) {
	s := &walkingState{steps: move.Steps, unit: u}
	s.nextStep()
	u.Sprite = u.Animations.Walk().Loop(true)
	u.state = s
}

func (s *walkingState) nextStep() {
	next := s.steps[0]
	s.steps = s.steps[1:]
	dest := s.unit.Location.Add(next.TP())
	s.unit.face(dest)
	speed := s.unit.moveSpeed()
	s.animation = units.Animate2D(s.unit.Location.IP(), dest.IP(), speed, nil)
}

func (s *walkingState) Tick() {
	s.unit.Drawable.Coordinate = s.animation.Tick()
	s.unit.Location = units.TPFromPoint(s.unit.Coordinate)
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

type attackingState struct{}

func (u *Unit) attack(cmd AttackCommand) {
	u.state = attackingState{}
	u.face(cmd.Target.Location)
	u.Sprite = u.Animations.Attack().OnComplete(func() {
		u.idle()
	})
}

func (s attackingState) Tick() {
}

func (s attackingState) Handle(c core.Command) {
	// commands are ignored while attacking
}

type defendingState struct{}

func (u *Unit) defend(cmd AttackCommand) {
	u.state = defendingState{}
	u.face(cmd.Unit.Location)
	u.Sprite = u.Animations.Hit().OnComplete(func() {
		u.status.HP -= cmd.Dmg()
		if u.status.HP <= 0 {
			u.die()
		} else {
			u.idle()
		}
	})
}

func (s defendingState) Tick() {
}

func (s defendingState) Handle(c core.Command) {
	// commands are ignored while defending
}

type deadState struct{}

func (u *Unit) die() {
	u.state = deadState{}
	u.Sprite = u.Animations.Death()
}

func (s deadState) Tick() {
}

func (s deadState) Handle(c core.Command) {
	// commands are ignored while dead
}
