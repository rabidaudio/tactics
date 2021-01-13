package chars

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/sprite"
)

type Character struct {
	// Anim is a sprite player for animating the character
	Anim          *sprite.Player
	Location      image.Point
	MoveSpeed     float64 // 1.0 * float64(units.TickRate)/ units.TileSize
	walk          units.Anim2D
	reverseFacing bool
}

func (c *Character) IsMoving() bool {
	return !c.walk.Value().Eq(c.Location)
}

// type walkAnim struct {
// 	m   float64
// 	b   float64
// 	end int
// 	t   int
// }

// pos = 5 tile
// step = 2 tiles
// pixels = 32
// walk speed = 12 ticks / tile

// t=0 ... pos = 5
// t=1 ... pos = 5 + (t/12) = 5.08333 tiles = 81.333 pixels = 81 pixels
// t=2 ... pos = 5 + (t/12) = 82.666
// t=3 ... 84
// t=11 ... 94.666
// t=12 ... 96

// func (wa *walkAnim) Tick() int {
// 	if wa.t < wa.end {
// 		wa.t++
// 	}
// 	return int(wa.b + (float64(wa.t) / wa.m))
// }

// type animStep image.Point
// type animComplete struct{}

// func (animStep) Action() string {
// 	return "animation/character/moving" // TODO this method is useless
// }

func (c *Character) Go(dir units.Direction) {
	if dir == units.East {
		c.reverseFacing = false
	} else if dir == units.West {
		c.reverseFacing = true
	}
	c.walk = units.Animate2D(c.Location, c.Location.Add(dir.TP().IP()), c.MoveSpeed*units.TileSize/float64(units.TickRate))

	// TODO how to chain? callbacks?

	// TODO: how to customize animation?

	// if c.listenerKey == "" {
	// 	c.listenerKey = core.Events.AddListener(func(event core.Event, tick units.Tick) {
	// 		if e, ok := event.(animStep); ok {
	// 			c.Location = image.Point(e)
	// 		} else if _, ok := event.(animComplete); ok {
	// 			c.moving = false
	// 		}
	// 	})
	// }
	// dir.TP()
}

// func lerp(start, end, t, rate int) int {
// 	s := float64(start)
// 	e := float64(end)
// 	tt := float64(t) / float64(rate)
// 	return int(s*(1-tt) + e*tt)
// }

func (c *Character) Tick() {
	c.Location = c.walk.Tick()
}

func (s *Character) Draw(screen *ebiten.Image, tick units.Tick) {
	opts := ebiten.DrawImageOptions{}
	if s.reverseFacing {
		opts.GeoM.Scale(-1.0, 1.0)
		opts.GeoM.Translate(16.0, 0)
	}
	opts.GeoM.Translate(float64(s.Location.X), float64(s.Location.Y))
	s.Anim.Tick()
	screen.DrawImage(s.Anim.Frame(), &opts)
}

// func (s *Spearman) Go(direction units.Direction, tiles int) {
// 	if s.IsMoving() {
// 		return
// 	}
// 	s.remainingSteps = int(float32(tiles*StepSize) / WalkSpeed)
// 	s.direction = direction
// 	if s.direction == units.East {
// 		s.reverseFacing = false
// 	} else if s.direction == units.West {
// 		s.reverseFacing = true
// 	}
// 	s.anim.
// 		ReplaceOnce(assets.BarbarianSholder()).
// 		AppendLoop(assets.BarbarianWalk())
// }

// func (c *Character) Draw(screen *ebiten.Image) {

// }

// func (c *Character) Move(steps []units.Direction) {

// }
