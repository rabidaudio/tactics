package core

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
)

type Drawable struct {
	Sprite        *sprite.Sprite
	ReverseFacing bool
	opts          ebiten.DrawImageOptions
	Coordinate    image.Point
	DrawCallback  func(*ebiten.DrawImageOptions)
}

func (d *Drawable) Tick() {
	d.Sprite.Tick()
}

func (d *Drawable) Draw(screen *ebiten.Image) {
	d.opts.GeoM.Reset()
	if d.ReverseFacing {
		d.opts.GeoM.Scale(-1.0, 1.0)
		d.opts.GeoM.Translate(float64(units.TileSize), 0)
	}
	d.opts.GeoM.Translate(float64(d.Coordinate.X), float64(d.Coordinate.Y))
	if d.DrawCallback != nil {
		d.DrawCallback(&d.opts)
	}
	screen.DrawImage(d.Sprite.Frame(), &d.opts)
}
