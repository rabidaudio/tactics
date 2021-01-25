package unit

import (
	"log"

	"github.com/rabidaudio/tactics/core/units"
)

type Attack struct {
	Attacker *Unit
	Target   *Unit
	Dmg      int
	Count    int
}

type UnitState interface {
	Tick()
}

type idleState struct{}

func (u *Unit) Idle() {
	u.Drawable.Sprite = u.Animations.Idle.Sprite()
	u.state = idleState{}
}

func (s idleState) Tick() {
}

type walkingState struct {
	unit      *Unit
	steps     []units.Direction
	animation units.Anim2D
}

func (u *Unit) IsReady() bool {
	_, ok := u.state.(idleState)
	return ok
}

func (u *Unit) Walk(steps ...units.Direction) {
	if !u.IsReady() {
		return // commands only accepted in idle state
	}
	s := &walkingState{steps: steps, unit: u}
	s.nextStep()
	u.Drawable.Sprite = u.Animations.Walk.Sprite()
	u.state = s
}

func (s *walkingState) nextStep() {
	next := s.steps[0]
	s.steps = s.steps[1:]
	dest := s.unit.Location.Add(next.TP())
	s.unit.FaceTowards(dest)
	speed := s.unit.moveSpeed()
	s.animation = units.Animate2D(s.unit.Location.IP(), dest.IP(), speed, nil)
}

func (s *walkingState) Tick() {
	c := s.animation.Tick()
	s.unit.Drawable.Coordinate = c
	s.unit.Location = units.TPFromPoint(c)
	if s.animation.IsMoving() {
		return // keep on going
	}
	if len(s.steps) == 0 {
		s.unit.Idle()
		return
	}
	s.nextStep()
}

type attackingState struct {
	attacksRemaining int
}

func (u *Unit) Attack(atk Attack) {
	if !u.IsReady() {
		return // commands only accepted in idle state
	}
	u.state = attackingState{attacksRemaining: atk.Count - 1}
	u.FaceTowards(atk.Target.Location)
	u.Drawable.Sprite = u.Animations.Attack.Sprite().OnComplete(func() {
		// TODO [bug] attack again if remaining
		u.Idle()
	})
}

func (a attackingState) Tick() {}

type defendingState struct {
	unit          *Unit
	dmg           int
	hitsRemaining int
}

func (a defendingState) Tick() {}

func (u *Unit) DefendAgainst(atk Attack) {
	if !u.IsReady() {
		return // commands only accepted in idle state
	}
	log.Printf("%v hit %v %v times for %v dmg each", atk.Attacker, u, atk.Count, atk.Dmg)
	state := defendingState{unit: u, dmg: atk.Dmg, hitsRemaining: atk.Count}
	u.state = state
	u.FaceTowards(atk.Attacker.Location)
	u.Drawable.Sprite = u.Animations.Hit.Sprite().OnComplete(state.animEnd)
}

func (ds *defendingState) animEnd() {
	ds.unit.HP -= ds.dmg
	ds.hitsRemaining--
	if ds.unit.HP <= 0 {
		ds.unit.Die()
	} else if ds.hitsRemaining > 0 {
		ds.unit.Drawable.Sprite = ds.unit.Animations.Hit.Sprite().OnComplete(ds.animEnd)
	} else {
		ds.unit.Idle()
	}
}

type deadState struct{}

func (a deadState) Tick() {}

func (u *Unit) Die() {
	u.state = deadState{}
	u.Drawable.Sprite = u.Animations.Death.Sprite()
}
