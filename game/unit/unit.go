package unit

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/weapon"
)

type Team int

type Unit struct {
	UnitOptions
	core.Drawable
	state  UnitState
	status Status
}

type UnitOptions struct {
	Team       Team
	Weapon     weapon.Weapon
	Animations UnitAnimations
	Location   units.TPoint
	// Walk speed in units of pixels/second
	Stats
}

type Stats struct {
	Attack    int
	Defense   int
	Speed     int
	HitPoints int
}

// Status is the current level of the unit's stats,
// as opposed to their max level
type Status Stats

type UnitAnimations struct {
	// from core/assets/generate
	Attack func() *sprite.Sprite
	Death  func() *sprite.Sprite
	Hit    func() *sprite.Sprite
	Idle   func() *sprite.Sprite
	Walk   func() *sprite.Sprite
}

func offset(opts *ebiten.DrawImageOptions) {
	// offset by a quarter-tile so feet are on the ground
	opts.GeoM.Translate(0, -6.0)
}

func new(opts UnitOptions) *Unit {
	u := Unit{
		UnitOptions: opts,
		Drawable: core.Drawable{
			Coordinate:   opts.Location.IP(),
			DrawCallback: offset,
		},
	}
	u.idle()
	return &u
}

func (u *Unit) Tick() {
	u.Drawable.Tick()
	u.state.Tick()
}

func (u *Unit) Handle(cmd core.Command) {
	u.state.Handle(cmd)
}

// animation speed of walking
func (u *Unit) moveSpeed() float64 {
	return (1 + (0.1 * float64(u.Stats.Speed))) * units.TileSize / float64(units.TickRate)
}
