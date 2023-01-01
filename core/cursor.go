package core

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/core/units"
)

type Cursor struct {
	Window       *Window
	IsSelectable func(units.TPoint) bool
}

func (c Cursor) Tick() {}

func (c Cursor) Draw(screen *ebiten.Image) {
	if c.IsSelectable == nil {
		return
	}
	opts := ebiten.DrawImageOptions{}
	tp := c.Position()
	tile := assets.TileNotSelectable
	if c.IsSelectable(tp) {
		tile = assets.TileSelectable
	}
	p := tp.IP()
	opts.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(tile, &opts)
}

func (c Cursor) Position() units.TPoint {
	return units.TPFromPoint(image.Pt(ebiten.CursorPosition()).Add(c.Window.CameraOrigin()))
}
