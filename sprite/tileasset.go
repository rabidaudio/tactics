package sprite

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
)

type TileAsset struct {
	img   image.Image
	stepX int
	stepY int
}

func OpenTileAsset(path string, stepX, stepY int) TileAsset {
	i, err := load(path)
	if err != nil {
		log.Fatalf("open asset: %v", err)
	}
	return TileAsset{
		img:   i,
		stepX: stepX,
		stepY: stepY,
	}
}

func (ta *TileAsset) Get(x, y int) *ebiten.Image {
	si := ta.img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(x*ta.stepX, y*ta.stepY, (x+1)*ta.stepX, (y+1)*ta.stepY))
	i, err := ebiten.NewImageFromImage(si, ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("load asset: %v", err)
	}
	return i
}

func (ta *TileAsset) SpriteFromRow(x, y, frames int) *Builder {
	ff := make([]*ebiten.Image, frames)
	for i := 0; i < frames; i++ {
		ff[i] = ta.Get(x+i, y)
	}
	return New(ff...)
}
