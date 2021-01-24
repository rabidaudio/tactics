package unit

import (
	"log"

	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
)

type UnitState interface {
	Tick()
	Handle(c core.Command)
	AcceptingCommands() bool
}

// basicState is for states that are logic-free
type basicState struct{}

func (s basicState) Tick() {
}

func (s basicState) AcceptingCommands() bool {
	return false
}

func (s basicState) Handle(c core.Command) {
	core.Unexpected(c)
}

type idleState struct{}

func (u *Unit) idle() {
	u.Sprite = u.Animations.Idle.Sprite()
	u.state = idleState{}
}

func (s idleState) Tick() {
}

func (s idleState) AcceptingCommands() bool {
	return true
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
	u.Sprite = u.Animations.Walk.Sprite()
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

func (s *walkingState) AcceptingCommands() bool {
	return false
}

func (s *walkingState) Handle(c core.Command) {
	// commands are ignored while moving
}

type attackingState struct {
	basicState
}

func (u *Unit) attack(cmd AttackCommand) {
	u.state = attackingState{}
	u.face(cmd.Target.Location)
	u.Sprite = u.Animations.Attack.Sprite().OnComplete(func() {
		u.idle()
	})
}

type defendingState struct {
	basicState
	unit          *Unit
	dmg           int
	hitsRemaining int
}

func (u *Unit) defend(cmd AttackCommand) {
	log.Printf("%v hit %v for %v", cmd.Unit, cmd.Target, cmd.Dmg())
	state := defendingState{unit: u, dmg: cmd.Dmg(), hitsRemaining: cmd.Count()}
	u.state = state
	u.face(cmd.Unit.Location)
	u.Sprite = u.Animations.Hit.Sprite().OnComplete(state.animEnd)
}

func (ds *defendingState) animEnd() {
	ds.unit.HP -= ds.dmg
	ds.hitsRemaining--
	if ds.unit.HP <= 0 {
		ds.unit.die()
	} else if ds.hitsRemaining > 0 {
		ds.unit.Sprite = ds.unit.Animations.Hit.Sprite().OnComplete(ds.animEnd)
	} else {
		ds.unit.idle()
	}
}

type deadState struct {
	basicState
}

func (u *Unit) die() {
	u.state = deadState{}
	u.Sprite = u.Animations.Death.Sprite()
}
