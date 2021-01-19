package unit

import (
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/core/units"
)

func NewSpearman(loc units.TPoint) *Unit {
	return new(UnitOptions{
		InitialLocation: loc.IP(),
		MoveSpeed:       2.0,
		Animations:      assets.Spearman,
	})
}
