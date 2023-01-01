package unit

import (
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/team"
	"github.com/rabidaudio/tactics/game/weapon"
)

func NewSpearman(loc units.TPoint, team team.Team, lvl int) *Unit {
	return new(UnitOptions{
		Location:   loc,
		Team:       team,
		Level:      lvl,
		Animations: assets.Spearman,
		Weapon:     weapon.Spear,
	})
}
