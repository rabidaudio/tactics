package game

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

type World struct {
	gameMap *tiled.Map
	baseimg *ebiten.Image
	overimg *ebiten.Image
	// origin point (top left) of camera
	camera     image.Point
	cameraAnim units.Anim2D
	StartPoint units.TPoint
	canvas     *ebiten.Image
}

func NewWorld(path string) (World, error) {
	gameMap, err := tiled.LoadFromFile(path)
	if err != nil {
		return World{}, err
	}
	overLayer := -1
	for i, l := range gameMap.Layers {
		if l.Name == "top" {
			l.Visible = false
			overLayer = i
			break
		}
	}
	// base layer
	r, err := render.NewRenderer(gameMap)
	if err != nil {
		return World{}, err
	}
	if err = r.RenderVisibleLayers(); err != nil {
		return World{}, err
	}
	baseimg, err := ebiten.NewImageFromImage(r.Result, ebiten.FilterDefault)
	if err != nil {
		return World{}, err
	}

	var overimg *ebiten.Image
	if overLayer > -1 {
		// overlayer
		gameMap.Layers[overLayer].Visible = true
		r.Clear()
		if err = r.RenderLayer(overLayer); err != nil {
			return World{}, err
		}
		overimg, err = ebiten.NewImageFromImage(r.Result, ebiten.FilterDefault)
		if err != nil {
			return World{}, err
		}
	}

	canvas, err := ebiten.NewImage(baseimg.Bounds().Dx(), baseimg.Bounds().Dy(), ebiten.FilterDefault)
	if err != nil {
		return World{}, err
	}
	start, err := startPoint(gameMap)
	if err != nil {
		return World{}, err
	}
	return World{
		gameMap:    gameMap,
		baseimg:    baseimg,
		overimg:    overimg,
		canvas:     canvas,
		StartPoint: start,
	}, nil
}

func MustNewWorld(path string) World {
	w, err := NewWorld(path)
	if err != nil {
		panic(err)
	}
	return w
}

func (w *World) IsBoundary(pt units.TPoint) bool {
	if !pt.IP().In(w.rect()) {
		return false
	}
	return !w.tileAt("boundaries", pt).Nil
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

func (w *World) Size() units.TPoint {
	return units.TPoint{X: w.gameMap.Width, Y: w.gameMap.Height}
}

func (w *World) Draw(drawEntities func(screen *ebiten.Image)) *ebiten.Image {
	w.canvas.Clear()
	w.canvas.DrawImage(w.baseimg, nil)
	drawEntities(w.canvas)
	if w.overimg != nil {
		w.canvas.DrawImage(w.overimg, nil)
	}
	return w.canvas
}
