package unit

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/weapon"
)

type Unit struct {
	UnitOptions
	Drawable core.Drawable
	state    UnitState
	Stats    Stats
	HP       int
}

type UnitOptions struct {
	Name       string
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
	// opts.GeoM.Translate(0, -6.0)
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
		Stats: stats,
		HP:    stats.BaseHP,
	}
	u.Idle()
	return &u
}

func (u *Unit) Tick() {
	u.Drawable.Tick()
	u.state.Tick()
}

func (u *Unit) CanReach(other *Unit) bool {
	return u.Weapon.CanHit(u.Location, other.Location)
}

func (u *Unit) StepsPerTurn() int {
	return 2 + (u.Stats.Spd - BaseStat)
}

func (u *Unit) String() string {
	// TODO [archetecture] placeholder
	return fmt.Sprintf("%v{team: %v lvl: %v hp: %v/%v wpn: %v}", u.Name, u.Team, u.Level, u.HP, u.Stats.BaseHP, u.Weapon.Name)
}

// animation speed of walking in units of pixels/second
// scales with the units speed stat (but independant of steps/turn)
func (u *Unit) moveSpeed() float64 {
	return (1 + (0.1 * float64(u.Stats.Spd))) * units.TileSize / float64(units.TickRate)
}

func (u *Unit) Draw(screen *ebiten.Image) {
	u.Drawable.Draw(screen)
}

func (u *Unit) FaceTowards(loc units.TPoint) {
	if loc.X > u.Location.X {
		u.Drawable.ReverseFacing = false
	} else if loc.X < u.Location.X {
		u.Drawable.ReverseFacing = true
	}
	// otherwise continue to face the same direction you are currently
}
