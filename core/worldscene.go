package core

import (
	"image"

	"github.com/hajimehoshi/ebiten"
)

// A scene which uses a tile map and camera
type WorldScene struct {
	Scene
	World  *World
	Window *Window
}

func (ws *WorldScene) Draw(screen *ebiten.Image) {
	canvas := ws.World.Draw(func(canvas *ebiten.Image) {
		ws.Scene.Draw(canvas)
	})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(ws.World.StartPoint.IP().X), float64(ws.World.StartPoint.IP().Y))
	screen.DrawImage(canvas.SubImage(ws.Window.Rect()).(*ebiten.Image), nil)
}

func SceneWithWorld(scene Scene, mapname string, windowsize image.Point) WorldScene {
	window := Window{Size: windowsize}
	world := MustNewWorld("raw/maps/square.tmx")
	window.BoundCamera(image.Rectangle{Max: world.Size().IP().Sub(window.Size)})
	window.JumpCamera(world.StartPoint)
	ws := WorldScene{
		Scene:  scene,
		World:  &world,
		Window: &window,
	}
	return ws
}
