package unit

import "github.com/rabidaudio/tactics/game/weapon"

type Stats struct {
	Atk   int
	Def   int
	Spd   int
	HP    int
	Steps int
}

func BaseStats(lvl int) Stats {
	// TODO [balance]
	return Stats{
		Atk:   lvl,
		Def:   lvl,
		Spd:   lvl,
		HP:    5 + 2*lvl,
		Steps: 2,
	}
}

func (s Stats) Bias(wtype weapon.WeaponType) Stats {
	scale := 2 // TODO [balance]
	switch wtype {
	case weapon.SwordType:
		s.Def += scale
		s.Spd -= scale
	case weapon.AxeType:
		s.Atk += scale
		s.Def -= scale
	case weapon.LanceType:
		s.Spd += scale
		s.Atk -= scale
	}
	return s
}

// Status is the current level of the unit's stats,
// as opposed to their max level
type Status Stats
