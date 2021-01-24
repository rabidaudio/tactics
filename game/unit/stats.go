package unit

import "github.com/rabidaudio/tactics/game/weapon"

type Stats struct {
	Atk    int
	Def    int
	Spd    int
	BaseHP int
}

func BaseStats(lvl int) Stats {
	// TODO [balance]
	// TODO [mechanics] leveling
	return Stats{
		Atk:    BaseStat,
		Def:    BaseStat,
		Spd:    BaseStat,
		BaseHP: BaseHP,
	}
}

func (s Stats) Bias(wtype weapon.WeaponType) Stats {
	switch wtype {
	case weapon.SwordType:
		s.Def += StatBias
		s.Spd -= StatBias
	case weapon.AxeType:
		s.Atk += StatBias
		s.Def -= StatBias
	case weapon.LanceType:
		s.Spd += StatBias
		s.Atk -= StatBias
	}
	return s
}
