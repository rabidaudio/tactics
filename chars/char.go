package chars

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
)

type Character struct {
	CharacterOptions
	// Anim is a sprite player for animating the character
	player        *sprite.Player
	reverseFacing bool
	walk          units.Anim2D
	opts          ebiten.DrawImageOptions
}

type CharacterOptions struct {
	AnimationHooks AnimationHooks
	Location       image.Point
	// Walk speed in units of pixels/second
	MoveSpeed   float64
	MaxHP       int
	AttackSpeed int
}

type WeaponType int

const (
	Sword WeaponType = iota
	Spear
	Axe
	Bow
)

type Weapon struct {
	Name string
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

func LoopOf(s sprite.Sprite) AnimationHook {
	return func(p *sprite.Player) *sprite.Player {
		return p.ReplaceLoop(s)
	}
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

func (c *Character) IsMoving() bool {
	return c.walk.IsMoving()
}

func (c *Character) Draw(screen *ebiten.Image) {
	c.opts.GeoM.Reset()
	if c.reverseFacing {
		c.opts.GeoM.Scale(-1.0, 1.0)
		c.opts.GeoM.Translate(float64(units.TileSize), 0)
	}
	c.opts.GeoM.Translate(float64(c.Location.X), float64(c.Location.Y))
	c.opts.GeoM.Translate(0, -6.0) // offset by a quarter-tile so feet are on the ground
	screen.DrawImage(c.player.Frame(), &c.opts)
}
