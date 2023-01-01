package unit

const (
	BaseStat               = 5
	BaseHP                 = 10
	StatBias               = 1
	SecondAttackSpeedDelta = 3
)

type Attack struct {
	Attacker *Unit
	Target   *Unit
	Dmg      int
	Count    int
}
