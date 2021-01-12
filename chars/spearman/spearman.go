package spearman

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/sprite"
	"github.com/rabidaudio/tactics/units"
)

const StepSize = 16

const WalkSpeed = 0.25

type Spearman struct {
	Location       Point
	remainingSteps int
	direction      units.Direction
	reverseFacing  bool
	anim           *sprite.Player
}

type Point struct {
	X float32
	Y float32
}

func New() Spearman {
	return Spearman{
		Location: Point{X: 50.0, Y: 50.0},
		anim:     sprite.NewPlayer().AppendLoop(assets.BarbarianIdle()),
	}
}

func (s *Spearman) IsMoving() bool {
	return s.remainingSteps > 0
}

func (s *Spearman) Step() {
	if s.IsMoving() {
		switch s.direction {
		case units.North:
			s.Location.Y -= WalkSpeed
		case units.South:
			s.Location.Y += WalkSpeed
		case units.East:
			s.Location.X += WalkSpeed
		case units.West:
			s.Location.X -= WalkSpeed
		}
		s.remainingSteps--
		if s.remainingSteps == 0 {
			s.anim.
				ReplaceOnce(assets.BarbarianSholder()).
				AppendLoop(assets.BarbarianIdle())
		}
	}
}

func (s *Spearman) Draw(screen *ebiten.Image, tick units.Tick) {
	opts := ebiten.DrawImageOptions{}
	if s.reverseFacing {
		opts.GeoM.Scale(-1.0, 1.0)
		opts.GeoM.Translate(16.0, 0)
	}
	opts.GeoM.Translate(float64(s.Location.X), float64(s.Location.Y))
	if tick%15 == 0 {
		s.anim.Tick()
	}
	screen.DrawImage(s.anim.Frame(), &opts)
}

func (s *Spearman) Go(direction units.Direction, tiles int) {
	if s.IsMoving() {
		return
	}
	s.remainingSteps = int(float32(tiles*StepSize) / WalkSpeed)
	s.direction = direction
	if s.direction == units.East {
		s.reverseFacing = false
	} else if s.direction == units.West {
		s.reverseFacing = true
	}
	s.anim.
		ReplaceOnce(assets.BarbarianSholder()).
		AppendLoop(assets.BarbarianWalk())
}
