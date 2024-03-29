package weapon

import (
	"github.com/rabidaudio/tactics/core/units"
)

// Reach is the range of a weapon
type Reach int

const (
	Melee  Reach = 1
	Ranged Reach = 2
)

var meleeThreat = []units.TPoint{
	units.TP(0, -1),
	units.TP(1, 0),
	units.TP(0, 1),
	units.TP(-1, 0),
}

var rangedThreat = []units.TPoint{
	units.TP(0, -2),
	units.TP(1, -1),
	units.TP(2, 0),
	units.TP(1, 1),
	units.TP(0, 2),
	units.TP(-1, 1),
	units.TP(-2, 0),
	units.TP(-1, -1),
}

// Threatens returns the points that can be hit from the current
// point with a weapon with this Reach
func (r Reach) Threatens(from units.TPoint) []units.TPoint {
	// TODO [performance]
	var threat []units.TPoint
	switch r {
	case Melee:
		threat = meleeThreat
	case Ranged:
		threat = rangedThreat
	}
	res := make([]units.TPoint, len(threat))
	for i := range threat {
		res[i] = from.Add(threat[i])
	}
	return res
}

// CanHit determines if a weapon with this reach can
// hit the given point
func (r Reach) CanHit(from, to units.TPoint) bool {
	// TODO [performance]
	for _, t := range r.Threatens(from) {
		if t == to {
			return true
		}
	}
	return false
}

type WeaponType int

const (
	SwordType WeaponType = iota
	LanceType
	AxeType
	BowType
	MagicType
)

func (wt WeaponType) HasTypeAdvantage(other WeaponType) bool {
	switch wt {
	case LanceType:
		return other == SwordType
	case SwordType:
		return other == AxeType
	case AxeType:
		return other == LanceType
	default:
		return false
	}
}

type Weapon struct {
	Reach
	WeaponType
	Name        string
	DamageLevel int
}

// TODO [archetecture] generate/parse from config file

var Spear = Weapon{
	Name:        "spear",
	Reach:       Melee,
	WeaponType:  LanceType,
	DamageLevel: 3,
}

var Halberd = Weapon{
	Name:        "halberd",
	Reach:       Melee,
	WeaponType:  LanceType,
	DamageLevel: 5,
}

var Sword1 = Weapon{
	Name:        "iron sword",
	Reach:       Melee,
	WeaponType:  SwordType,
	DamageLevel: 3,
}

var Sword2 = Weapon{
	Name:        "steel sword",
	Reach:       Melee,
	WeaponType:  SwordType,
	DamageLevel: 5,
}

var Axe = Weapon{
	Name:        "axe",
	Reach:       Melee,
	WeaponType:  AxeType,
	DamageLevel: 3,
}

var Hammer = Weapon{
	Name:        "hammer",
	Reach:       Melee,
	WeaponType:  AxeType,
	DamageLevel: 5,
}

var Bow1 = Weapon{
	Name:        "bow",
	Reach:       Ranged,
	WeaponType:  BowType,
	DamageLevel: 3,
}

var Bow2 = Weapon{
	Name:        "compound bow",
	Reach:       Ranged,
	WeaponType:  BowType,
	DamageLevel: 5,
}

var Crossbow1 = Weapon{
	Name:        "light crossbow",
	Reach:       Ranged,
	WeaponType:  BowType,
	DamageLevel: 3,
}

var Crossbow2 = Weapon{
	Name:        "heavy crossbow",
	Reach:       Ranged,
	WeaponType:  BowType,
	DamageLevel: 5,
}

var Heal1 = Weapon{
	Name:        "mend",
	Reach:       Melee,
	WeaponType:  MagicType,
	DamageLevel: 3,
}

var Heal2 = Weapon{
	Name:        "heal",
	Reach:       Melee,
	WeaponType:  MagicType,
	DamageLevel: 5,
}

var Wound1 = Weapon{
	Name:        "harm",
	Reach:       Ranged,
	WeaponType:  MagicType,
	DamageLevel: 3,
}

var Wound2 = Weapon{
	Name:        "curse",
	Reach:       Ranged,
	WeaponType:  MagicType,
	DamageLevel: 5,
}
