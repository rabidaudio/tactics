package game

func (cmd AttackCommand) Dmg() int {
	// TODO [mechanics]
	dmg := cmd.atk() - cmd.def()
	if dmg < 0 {
		return 0
	}
	return dmg
}

func (cmd AttackCommand) Count() int {
	if (cmd.Attacker.Stats.Spd - cmd.Target.Stats.Spd) >= SecondAttackSpeedDelta {
		return 2
	}
	return 1
}

func (cmd AttackCommand) atk() int {
	a := cmd.Attacker.Stats.Atk + cmd.Attacker.Weapon.DamageLevel
	at := cmd.Attacker.Weapon.WeaponType
	dt := cmd.Target.Weapon.WeaponType
	if at.HasTypeAdvantage(dt) {
		a = a + (a / 2)
	}
	return a
}

func (cmd AttackCommand) def() int {
	return cmd.Target.Stats.Def
}
