package unit

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/sprite"
)

type Unit struct {
	UnitOptions
	core.Drawable
	state UnitState
}

type UnitOptions struct {
	Animations      UnitAnimations
	InitialLocation image.Point
	// Walk speed in units of pixels/second
	MoveSpeed   float64
	MaxHP       int
	AttackSpeed int
}

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
	if opts.MoveSpeed == 0 {
		opts.MoveSpeed = 1.0
	}
	u := Unit{
		UnitOptions: opts,
		Drawable: core.Drawable{
			Location:     opts.InitialLocation,
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
