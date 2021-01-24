package assets

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/units"
)

var TileSelectable *ebiten.Image
var TileNotSelectable *ebiten.Image

func init() {
	TileSelectable = createTile(color.RGBA{B: 255, A: 63})
	TileNotSelectable = createTile(color.RGBA{R: 255, A: 63})
}

func createTile(c color.Color) *ebiten.Image {
	img, err := ebiten.NewImage(units.TileSize, units.TileSize, ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("create img: %v", err)
	}
	img.Fill(c)
	return img
}
