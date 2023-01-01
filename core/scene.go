package core

import (
	"github.com/hajimehoshi/ebiten"
)

type Ticker interface {
	Tick()
}

type Drawer interface {
	Draw(screen *ebiten.Image)
}

type Scene interface {
	Ticker
	Drawer
}

type EmptyScene struct{}

func (es EmptyScene) Tick()                     {}
func (es EmptyScene) Draw(screen *ebiten.Image) {}
