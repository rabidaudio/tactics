package spearman

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/assets"
)

const StepSize = 16

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Spearman struct {
	Location       image.Point
	remainingSteps int
	direction      Direction
	reverseFacing  bool
	anim           assets.CharacterAnimation
}

func New() Spearman {
	return Spearman{
		Location: image.Point{X: 10, Y: 10},
		// anim:     assets.Spearman(),
		anim: assets.Barbarian(),
	}
}

func (s *Spearman) IsMoving() bool {
	return s.remainingSteps > 0
}

func (s *Spearman) Step() {
	if s.IsMoving() {
		switch s.direction {
		case North:
			s.Location.Y -= 1
		case South:
			s.Location.Y += 1
		case East:
			s.Location.X += 1
		case West:
			s.Location.X -= 1
		}
		s.remainingSteps--
		if s.remainingSteps == 0 {
			s.anim.SetState(assets.StateIdle)
		}
	}
	s.anim.Sprite().Tick()
}

func (s *Spearman) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	if s.reverseFacing {
		opts.GeoM.Scale(-1.0, 1.0)
		opts.GeoM.Translate(16.0, 0)
	}
	opts.GeoM.Translate(float64(s.Location.X), float64(s.Location.Y))
	screen.DrawImage(s.anim.Sprite().Frame(), &opts)
}

func (s *Spearman) Go(direction Direction) {
	if s.IsMoving() {
		return
	}
	s.remainingSteps = StepSize
	s.direction = direction
	if s.direction == East {
		s.reverseFacing = false
	} else if s.direction == West {
		s.reverseFacing = true
	}
	s.anim.SetState(assets.StateWalk)
}
