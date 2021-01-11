package char

import "github.com/hajimehoshi/ebiten"

type Character interface {
	Step()
	Draw(screen *ebiten.Image)
}
