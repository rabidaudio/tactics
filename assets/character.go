package assets

import "github.com/rabidaudio/tactics/sprite"

type AnimationState int

const (
	Idle AnimationState = iota
	Walk
	Attack
	Hit
	Death
)

type CharacterAnimation struct {
	sprites map[AnimationState]*sprite.Sprite
	State   AnimationState
}

func (ca *CharacterAnimation) Sprite() *sprite.Sprite {
	return ca.sprites[ca.State]
}

var castle = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Castle/CastleSpriteSheet.png", 16, 16)

func Spearman() CharacterAnimation {
	return CharacterAnimation{
		sprites: map[AnimationState]*sprite.Sprite{
			Idle:   castle.SpriteFromRow(0, 1, 4).WithRate(15).PlayState(true, true),
			Walk:   castle.SpriteFromRow(4, 1, 4).WithRate(15).PlayState(true, true),
			Attack: castle.SpriteFromRow(8, 1, 4).WithRate(15).PlayState(true, false),
			Hit:    castle.SpriteFromRow(12, 1, 4).WithRate(15).PlayState(true, false),
			Death:  castle.SpriteFromRow(16, 1, 4).WithRate(5).PlayState(true, false),
		},
		State: Idle,
	}
}
