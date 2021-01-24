package unit

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/weapon"
)

type Team int

func (t Team) Color(cm *ebiten.ColorM) {
	// TODO [graphics] this is good enough for testing
	// but probably looks bad on different kinds of units
	h := 2 * math.Pi * (float64(t) / 8)
	cm.RotateHue(h)
}

type Unit struct {
	UnitOptions
	core.Drawable
	state  UnitState
	Stats  Stats
	status Status
}

type UnitOptions struct {
	Team       Team
	Level      int
	Weapon     weapon.Weapon
	Animations UnitAnimations
	Location   units.TPoint
}

type UnitAnimations struct {
	// from core/assets/generate
	Attack sprite.Template
	Death  sprite.Template
	Hit    sprite.Template
	Idle   sprite.Template
	Walk   sprite.Template
}

func offset(opts *ebiten.DrawImageOptions) {
	// offset by a partial-tile so feet are on the ground
	// TODO [graphics] this looks good on roads but
	// weird against other objects sometimes
	opts.GeoM.Translate(0, -6.0)
}

func new(opts UnitOptions) *Unit {
	stats := BaseStats(opts.Level).Bias(opts.Weapon.WeaponType)
	dopts := ebiten.DrawImageOptions{}
	opts.Team.Color(&dopts.ColorM)
	u := Unit{
		UnitOptions: opts,
		Drawable: core.Drawable{
			Coordinate:   opts.Location.IP(),
			DrawCallback: offset,
			Opts:         dopts,
		},
		Stats:  stats,
		status: Status(stats),
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

// animation speed of walking in units of pixels/second
// scales with the units speed stat (but independant of steps/turn)
func (u *Unit) moveSpeed() float64 {
	return (1 + (0.1 * float64(u.Stats.Spd))) * units.TileSize / float64(units.TickRate)
}

func (u *Unit) face(loc units.TPoint) {
	if loc.X > u.Location.X {
		u.ReverseFacing = false
	} else if loc.X < u.Location.X {
		u.ReverseFacing = true
	}
	// otherwise continue to face the same direction you are currently
}
