
// Code generated by core/assets/generator; DO NOT EDIT.

package assets

import "github.com/rabidaudio/tactics/core/sprite"

var Castle = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Castle/CastleSpriteSheet.png", 16, 16)

var Bishop = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 13, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 13, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 13, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 13, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 13, 4)
	},
}

var Crossbowman = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 2, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 2, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 2, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 2, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 2, 4)
	},
}

var Crossbowman2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 10, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 10, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 10, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 10, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 10, 4)
	},
}

var Halberd = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 9, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 9, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 9, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 9, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 9, 4)
	},
}

var Knight = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 7, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 7, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 7, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 7, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 7, 4)
	},
}

var Knight2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 15, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 15, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 15, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 15, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 15, 4)
	},
}

var Monk = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 5, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 5, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 5, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 5, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 5, 4)
	},
}

var Rider = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 6, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 6, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 6, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 6, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 6, 4)
	},
}

var Rider2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 14, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 14, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 14, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 14, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 14, 4)
	},
}

var Spearman = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 1, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 1, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 1, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 1, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 1, 4)
	},
}

var Swordsman = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 4, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 4, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 4, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 4, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 4, 4)
	},
}

var Swordsman2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Castle.SpriteFromRow(8, 12, 4)
	},
	Death: func() sprite.Sprite {
		return Castle.SpriteFromRow(16, 12, 4)
	},
	Hit: func() sprite.Sprite {
		return Castle.SpriteFromRow(12, 12, 4)
	},
	Idle: func() sprite.Sprite {
		return Castle.SpriteFromRow(0, 12, 4)
	},
	Walk: func() sprite.Sprite {
		return Castle.SpriteFromRow(4, 12, 4)
	},
}


var Rampart = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Rampart/RampartSpriteSheet.png", 16, 16)

var Axeman = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Rampart.SpriteFromRow(8, 2, 4)
	},
	Death: func() sprite.Sprite {
		return Rampart.SpriteFromRow(16, 2, 4)
	},
	Hit: func() sprite.Sprite {
		return Rampart.SpriteFromRow(12, 2, 4)
	},
	Idle: func() sprite.Sprite {
		return Rampart.SpriteFromRow(0, 2, 4)
	},
	Walk: func() sprite.Sprite {
		return Rampart.SpriteFromRow(4, 2, 4)
	},
}

var Hammerman = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Rampart.SpriteFromRow(8, 10, 4)
	},
	Death: func() sprite.Sprite {
		return Rampart.SpriteFromRow(16, 10, 4)
	},
	Hit: func() sprite.Sprite {
		return Rampart.SpriteFromRow(12, 10, 4)
	},
	Idle: func() sprite.Sprite {
		return Rampart.SpriteFromRow(0, 10, 4)
	},
	Walk: func() sprite.Sprite {
		return Rampart.SpriteFromRow(4, 10, 4)
	},
}

var Hunter = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Rampart.SpriteFromRow(8, 4, 4)
	},
	Death: func() sprite.Sprite {
		return Rampart.SpriteFromRow(16, 4, 4)
	},
	Hit: func() sprite.Sprite {
		return Rampart.SpriteFromRow(12, 4, 4)
	},
	Idle: func() sprite.Sprite {
		return Rampart.SpriteFromRow(0, 4, 4)
	},
	Walk: func() sprite.Sprite {
		return Rampart.SpriteFromRow(4, 4, 4)
	},
}

var Hunter2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Rampart.SpriteFromRow(8, 12, 4)
	},
	Death: func() sprite.Sprite {
		return Rampart.SpriteFromRow(16, 12, 4)
	},
	Hit: func() sprite.Sprite {
		return Rampart.SpriteFromRow(12, 12, 4)
	},
	Idle: func() sprite.Sprite {
		return Rampart.SpriteFromRow(0, 12, 4)
	},
	Walk: func() sprite.Sprite {
		return Rampart.SpriteFromRow(4, 12, 4)
	},
}


var Stronghold = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Stronghold/StrongholdSpriteSheet.png", 16, 16)

var Brute = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(8, 7, 4)
	},
	Death: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(16, 7, 4)
	},
	Hit: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(12, 7, 4)
	},
	Idle: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(0, 7, 4)
	},
	Walk: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(4, 7, 4)
	},
}

var Brute2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(8, 15, 4)
	},
	Death: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(16, 15, 4)
	},
	Hit: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(12, 15, 4)
	},
	Idle: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(0, 15, 4)
	},
	Walk: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(4, 15, 4)
	},
}

var GoblinRider = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(8, 3, 4)
	},
	Death: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(16, 3, 4)
	},
	Hit: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(12, 3, 4)
	},
	Idle: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(0, 3, 4)
	},
	Walk: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(4, 3, 4)
	},
}

var GoblinRider2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(8, 11, 4)
	},
	Death: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(16, 11, 4)
	},
	Hit: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(12, 11, 4)
	},
	Idle: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(0, 11, 4)
	},
	Walk: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(4, 11, 4)
	},
}

var Mercenary = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(8, 1, 4)
	},
	Death: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(16, 1, 4)
	},
	Hit: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(12, 1, 4)
	},
	Idle: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(0, 1, 4)
	},
	Walk: func() sprite.Sprite {
		return Stronghold.SpriteFromRow(4, 1, 4)
	},
}


var Tower = sprite.OpenTileAsset("raw/HAS CreaturePack/HAS Creature Pack 1.2/Tower/TowerSpriteSheet.png", 16, 16)

var Shotgun = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Tower.SpriteFromRow(8, 1, 4)
	},
	Death: func() sprite.Sprite {
		return Tower.SpriteFromRow(16, 1, 4)
	},
	Hit: func() sprite.Sprite {
		return Tower.SpriteFromRow(12, 1, 4)
	},
	Idle: func() sprite.Sprite {
		return Tower.SpriteFromRow(0, 1, 4)
	},
	Walk: func() sprite.Sprite {
		return Tower.SpriteFromRow(4, 1, 4)
	},
}

var Shotgun2 = struct {
	Attack func() sprite.Sprite
	Death func() sprite.Sprite
	Hit func() sprite.Sprite
	Idle func() sprite.Sprite
	Walk func() sprite.Sprite
}{
	Attack: func() sprite.Sprite {
		return Tower.SpriteFromRow(8, 9, 4)
	},
	Death: func() sprite.Sprite {
		return Tower.SpriteFromRow(16, 9, 4)
	},
	Hit: func() sprite.Sprite {
		return Tower.SpriteFromRow(12, 9, 4)
	},
	Idle: func() sprite.Sprite {
		return Tower.SpriteFromRow(0, 9, 4)
	},
	Walk: func() sprite.Sprite {
		return Tower.SpriteFromRow(4, 9, 4)
	},
}

