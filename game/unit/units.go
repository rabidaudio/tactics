package unit

import (
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/weapon"
)

func NewSpearman(loc units.TPoint, team Team) *Unit {
	return new(UnitOptions{
		Location:   loc,
		Team:       team,
		Animations: assets.Spearman,
		Weapon:     weapon.Spear,
		Stats: Stats{
			Attack:  1,
			Defense: 2,
			Speed:   3,
		},
	})
}
