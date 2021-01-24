package sprite

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
)

type TileAsset struct {
	img   image.Image
	stepX int
	stepY int
}

func OpenTileAsset(path string, stepX, stepY int) TileAsset {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open asset: %v", err)
	}
	i, err := png.Decode(f)
	if err != nil {
		log.Fatalf("png decode: %v", err)
	}
	return TileAsset{
		img:   i,
		stepX: stepX,
		stepY: stepY,
	}
}

type subImage interface {
	SubImage(r image.Rectangle) image.Image
}

func (ta *TileAsset) Get(x, y int) *ebiten.Image {
	si := ta.img.(subImage).SubImage(image.Rect(x*ta.stepX, y*ta.stepY, (x+1)*ta.stepX, (y+1)*ta.stepY))
	i, err := ebiten.NewImageFromImage(si, ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("load asset: %v", err)
	}
	return i
}

// func (ta *TileAsset) SpriteFromRow(x, y, frames int) *Sprite {
// 	ff := make([]*ebiten.Image, frames)
// 	for i := 0; i < frames; i++ {
// 		ff[i] = ta.Get(x+i, y)
// 	}
// 	return New(ff...)
// }

// func (ta *TileAsset) SpriteFromColumn(x, y, frames int) *Sprite {
// 	ff := make([]*ebiten.Image, frames)
// 	for i := 0; i < frames; i++ {
// 		ff[i] = ta.Get(x, y+i)
// 	}
// 	return New(ff...)
// }

func (ta *TileAsset) SpriteTemplate(frames [][]int) Template {
	ff := make([]*ebiten.Image, len(frames))
	for i := 0; i < len(frames); i++ {
		ff[i] = ta.Get(frames[i][0], frames[i][1])
	}
	return NewTemplate(ff...)
}
