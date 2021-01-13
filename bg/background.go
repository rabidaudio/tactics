package bg

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/sprite"
)

const CameraSpeed = 1.0 * float64(units.TileSize) / float64(units.TickRate)

type Background struct {
	img    *ebiten.Image
	camera image.Point
	pan    units.Anim2D
}

func New() Background {
	s, _ := sprite.Load("raw/HAS Overworld 2.1/Universal/Universal-Buildings-and-walls.png")
	img, _ := ebiten.NewImageFromImage(s, ebiten.FilterDefault)
	return Background{
		img:    img,
		camera: image.Pt(0, 0),
	}
}

func (b *Background) StepCamera(dir units.Direction) {
	b.pan = units.Animate2D(b.camera, b.camera.Add(dir.TP().IP()), CameraSpeed)
	// b.camera.AnimatePlus(dir.TP().IP(), units.Second(0.2).Ticks())
	// b.Camera = b.Camera.Add(dir.TP())
}

func (b *Background) Go(point units.TPoint) {
	b.pan = units.Animate2D(b.camera, point.IP(), CameraSpeed)
	// TODO: speed as property of animation
	// b.camera.AnimateTo(point.IP(), 12)

}

func (b *Background) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	t := b.camera.Mul(-1)
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(b.img, &opts)
}

func (b *Background) Tick() {
	b.pan.Tick()
}
