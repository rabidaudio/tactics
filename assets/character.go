package assets

import "github.com/rabidaudio/tactics/sprite"

type AnimationState int

const (
	StateIdle AnimationState = iota
	StateWalk
	StateAttack
	StateHit
	StateDeath
)

type CharacterAnimation struct {
	sprites map[AnimationState]sprite.Sprite
	state   AnimationState
}

func (ca *CharacterAnimation) Sprite() sprite.Sprite {
	return ca.sprites[ca.state]
}

func (ca *CharacterAnimation) SetState(state AnimationState) {
	ca.sprites[ca.state].Reset()
	ca.state = state
	ca.sprites[ca.state].Reset()
}

var castle = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Castle/CastleSpriteSheet.png", 16, 16)
var tower = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Tower/TowerSpriteSheet.png", 16, 16)

func Spearman() CharacterAnimation {
	return CharacterAnimation{
		sprites: map[AnimationState]sprite.Sprite{
			StateIdle:   castle.SpriteFromRow(0, 1, 4).Loop().Rate(15).Sprite(),
			StateWalk:   castle.SpriteFromRow(4, 1, 4).Loop().Rate(5).Sprite(),
			StateAttack: castle.SpriteFromRow(8, 1, 4).Rate(15).Sprite(),
			StateHit:    castle.SpriteFromRow(12, 1, 4).Rate(15).Sprite(),
			StateDeath:  castle.SpriteFromRow(16, 1, 4).Rate(10).Sprite(),
		},
		state: StateIdle,
	}
}

func Barbarian() CharacterAnimation {
	return CharacterAnimation{
		sprites: map[AnimationState]sprite.Sprite{
			StateIdle: tower.SpriteFromRow(0, 1, 4).Loop().Rate(15).Sprite(),
			StateWalk: tower.SpriteFromRow(20, 1, 4).Rate(5).Concat(
				tower.SpriteFromRow(4, 1, 4).Loop().Rate(15).Sprite()).Sprite(),

			// StateWalk:   castle.SpriteFromRow(4, 1, 4).WithRate(5).PlayState(true, true),
			// StateAttack: castle.SpriteFromRow(8, 1, 4).WithRate(15).PlayState(true, false),
			// StateHit:    castle.SpriteFromRow(12, 1, 4).WithRate(15).PlayState(true, false),
			// StateDeath:  castle.SpriteFromRow(16, 1, 4).WithRate(10).PlayState(true, false),
		},
		state: StateIdle,
	}
}
