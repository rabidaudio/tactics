package spearman

import (
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/chars"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/sprite"
)

type Spearman struct {
	chars.Character
}

func New() Spearman {
	return Spearman{
		Character: chars.New(chars.CharacterOptions{
			Location:  units.TP(5, 5).IP(),
			MoveSpeed: 0.1,
			AnimationHooks: chars.AnimationHooks{
				Idle: func(p *sprite.Player) *sprite.Player {
					return p.ReplaceOnce(assets.BarbarianSholder().Reversed()).
						AppendLoop(assets.BarbarianIdle())
				},
				Walking: func(p *sprite.Player) *sprite.Player {
					return p.ReplaceOnce(assets.BarbarianSholder()).
						AppendLoop(assets.BarbarianWalk())
				},
			},
		}),
	}
}
