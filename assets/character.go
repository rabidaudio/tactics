package assets

import "github.com/rabidaudio/tactics/sprite"

// TODO: use go generate for this

var castle = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Castle/CastleSpriteSheet.png", 16, 16)
var tower = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Tower/TowerSpriteSheet.png", 16, 16)

func BarbarianIdle() sprite.Sprite {
	return tower.SpriteFromRow(0, 1, 4).Rate(15)
}

func BarbarianWalk() sprite.Sprite {
	return tower.SpriteFromRow(4, 1, 4).Rate(15)
}

func BarbarianSholder() sprite.Sprite {
	return tower.SpriteFromRow(20, 1, 4).Rate(15)
}
