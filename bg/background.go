package bg

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/sprite"
	"github.com/rabidaudio/tactics/units"
)

type Background struct {
	img    *ebiten.Image
	camera units.APoint
}

func New() Background {
	s, _ := sprite.Load("raw/HAS Overworld 2.1/Universal/Universal-Buildings-and-walls.png")
	img, _ := ebiten.NewImageFromImage(s, ebiten.FilterDefault)
	return Background{
		img:    img,
		camera: units.AP(0, 0),
	}
}

func (b *Background) StepCamera(dir units.Direction) {
	b.camera.AnimatePlus(dir.TP().IP(), units.Second(0.2).Ticks())
	// b.Camera = b.Camera.Add(dir.TP())
}

func (b *Background) Go(point units.TPoint) {
	// TODO: speed as property of animation
	b.camera.AnimateTo(point.IP(), 12)
}

func (b *Background) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	t := b.camera.Point.Mul(-1)
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(b.img, &opts)
}

func (b *Background) Tick(tick units.Tick) {
	b.camera.Tick()
}
