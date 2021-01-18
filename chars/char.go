package chars

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/sprite"
)

type Character struct {
	CharacterOptions
	// Anim is a sprite player for animating the character
	player        *sprite.Player
	reverseFacing bool
	walk          units.Anim2D
}

type CharacterOptions struct {
	AnimationHooks AnimationHooks
	Location       image.Point
	// Walk speed in units of pixels/second
	MoveSpeed float64
}

func New(opts CharacterOptions) Character {
	p := sprite.NewPlayer()
	p = opts.AnimationHooks.Idle(p)
	if opts.MoveSpeed == 0 {
		opts.MoveSpeed = 1.0
	}
	return Character{
		CharacterOptions: opts,
		player:           p,
	}
}

type AnimationHook func(*sprite.Player) *sprite.Player

type AnimationHooks struct {
	Idle    AnimationHook
	Walking AnimationHook
}

func (c *Character) Go(dir units.Direction) {
	if c.walk.IsMoving() {
		return
	}
	if dir == units.East {
		c.reverseFacing = false
	} else if dir == units.West {
		c.reverseFacing = true
	}
	c.player = c.AnimationHooks.Walking(c.player)
	speed := c.MoveSpeed * units.TileSize / float64(units.TickRate)
	c.walk = units.Animate2D(c.Location, c.Location.Add(dir.TP().IP()), speed, func() {
		c.player = c.AnimationHooks.Idle(c.player)
	})
}

func (c *Character) Tick() {
	if c.walk.IsMoving() {
		c.Location = c.walk.Tick()
	}
	c.player.Tick()
}

func (c *Character) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	if c.reverseFacing {
		opts.GeoM.Scale(-1.0, 1.0)
		opts.GeoM.Translate(float64(units.TileSize), 0)
	}
	opts.GeoM.Translate(float64(c.Location.X), float64(c.Location.Y))
	screen.DrawImage(c.player.Frame(), opts)
}
