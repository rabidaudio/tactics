package spearman

import (
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/chars"
	"github.com/rabidaudio/tactics/core/units"
)

type Spearman struct {
	chars.Character
}

func New(loc units.TPoint) Spearman {
	return Spearman{
		Character: chars.New(chars.CharacterOptions{
			Location:  loc.IP(),
			MoveSpeed: 2.0,
			AnimationHooks: chars.AnimationHooks{
				Idle:    chars.LoopOf(assets.SpearmanIdle().Rate(15)),
				Walking: chars.LoopOf(assets.SpearmanWalk().Rate(5)),
			},
		}),
	}
}
