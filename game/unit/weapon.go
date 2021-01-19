package unit

type WeaponType int

const (
	Sword WeaponType = iota
	Spear
	Axe
	Bow
)

type Weapon struct {
	Name string
}
