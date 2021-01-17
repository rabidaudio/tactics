package world

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten"
	tiled "github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
	"github.com/rabidaudio/tactics/core/units"
)

const CameraSpeed = 0.25 * units.TilesPerSecond

var WindowSize = image.Point{X: 230, Y: 240}

type World struct {
	gameMap *tiled.Map
	img     *ebiten.Image
	// origin point (top left) of camera
	camera     image.Point
	cameraAnim units.Anim2D
	StartPoint units.TPoint
}

func New() (World, error) {
	gameMap, err := tiled.LoadFromFile("raw/maps/map1.tmx")
	if err != nil {
		return World{}, err
	}
	r, err := render.NewRenderer(gameMap)
	if err != nil {
		return World{}, err
	}
	if err = r.RenderVisibleLayers(); err != nil {
		return World{}, err
	}
	img, err := ebiten.NewImageFromImage(r.Result, ebiten.FilterDefault)
	if err != nil {
		return World{}, err
	}

	start, err := startPoint(gameMap)
	if err != nil {
		return World{}, err
	}
	return World{
		gameMap:    gameMap,
		img:        img,
		camera:     cameraCenterToOrigin(start.IP()),
		StartPoint: start,
	}, nil
}

func (w *World) IsBoundary(pt units.TPoint) bool {
	if !pt.IP().In(w.rect()) {
		return false
	}
	return w.tileAt("boundary", pt).Nil
}

func startPoint(m *tiled.Map) (tp units.TPoint, err error) {
	s := m.Properties.Get("start")
	if len(s) == 0 {
		return tp, fmt.Errorf("map property 'start' not found")
	}
	p := strings.Split(s[0], ",")
	if len(p) != 2 {
		return tp, fmt.Errorf("start property format: %v", s[0])
	}
	x, err := strconv.Atoi(p[0])
	if err != nil {
		return
	}
	y, err := strconv.Atoi(p[1])
	if err != nil {
		return
	}
	return units.TPoint{X: x, Y: y}, nil
}

func (w *World) tileAt(layer string, point units.TPoint) *tiled.LayerTile {
	return w.layerByName(layer).Tiles[point.Y*w.gameMap.Width+point.X]
}

func (w *World) rect() image.Rectangle {
	return image.Rect(0, 0, w.gameMap.Width*units.TileSize, w.gameMap.Height*units.TileSize)
}

func (w *World) layerByName(name string) *tiled.Layer {
	for _, l := range w.gameMap.Layers {
		if l.Name == name {
			return l
		}
	}
	return nil
}

func (w *World) StepCamera(dir units.Direction) {
	w.Go(units.TPFromPoint(w.camera).Add(dir.TP()))
}

func (w *World) Go(point units.TPoint) {
	if w.cameraAnim.IsMoving() {
		return
	}
	dest := w.boundedCamera(cameraCenterToOrigin(point.IP()))
	w.cameraAnim = units.Animate2D(w.camera, dest, CameraSpeed, nil)
}

func cameraCenterToOrigin(center image.Point) image.Point {
	return center.Sub(WindowSize.Div(2))
}

// boundedCamera doesn't let the camera show content beyond the
// edge of the world
func (w *World) boundedCamera(origin image.Point) image.Point {
	return units.Bound(origin, image.Rectangle{Max: w.size().IP().Sub(WindowSize)})
}

func (w *World) size() units.TPoint {
	return units.TPoint{X: w.gameMap.Width, Y: w.gameMap.Height}
}

func (w *World) CameraOrigin() image.Point {
	return w.camera
}

func (w *World) CameraCenter() image.Point {
	return w.camera.Add(WindowSize.Div(2))
}

func (w *World) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	t := w.camera.Mul(-1)
	opts.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(w.img, &opts)
}

func (w *World) Tick() {
	if w.cameraAnim.IsMoving() {
		w.camera = w.cameraAnim.Tick()
	}
}
