package window

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core/units"
)

const CameraSpeed = 0.25 * units.TilesPerSecond

type Window struct {
	Size         image.Point
	worldBounds  image.Rectangle
	camera       image.Point
	cameraAnim   units.Anim2D
	cameraBounds image.Rectangle
}

func (w *Window) WorldSize(s units.TPoint) *Window {
	w.cameraBounds = image.Rectangle{Max: s.IP().Sub(w.Size)}
	w.camera = w.boundedCamera(w.camera)
	return w
}

func (w *Window) JumpCamera(point units.TPoint) *Window {
	w.camera = w.boundedCamera(w.cameraCenterToOrigin(point.IP()))
	return w
}

func (w *Window) StepCamera(dir units.Direction) {
	w.AnimateCamera(units.TPFromPoint(w.camera).Add(dir.TP()))
}

func (w *Window) AnimateCamera(point units.TPoint) {
	if w.cameraAnim.IsMoving() {
		return
	}
	dest := w.boundedCamera(w.cameraCenterToOrigin(point.IP()))
	w.cameraAnim = units.Animate2D(w.camera, dest, CameraSpeed, nil)
}

func (w *Window) cameraCenterToOrigin(center image.Point) image.Point {
	return center.Sub(w.Size.Div(2))
}

// boundedCamera doesn't let the camera show content beyond the
// edge of the world
func (w *Window) boundedCamera(origin image.Point) image.Point {
	return units.Bound(origin, w.cameraBounds)
}

func (w *Window) CameraOrigin() image.Point {
	return w.camera
}

func (w *Window) CameraCenter() image.Point {
	return w.camera.Add(w.Size.Div(2))
}

func (w *Window) Tick() {
	if w.cameraAnim.IsMoving() {
		w.camera = w.cameraAnim.Tick()
	}
}

func (w *Window) DrawOpts() *ebiten.DrawImageOptions {
	opts := ebiten.DrawImageOptions{}
	t := w.camera.Mul(-1)
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	return &opts
}

func (w *Window) WindowPoint(worldPoint image.Point) image.Point {
	return worldPoint.Sub(w.camera)
}

func (w *Window) Rect() image.Rectangle {
	return image.Rectangle{Min: w.camera, Max: w.camera.Add(w.Size)}
}
