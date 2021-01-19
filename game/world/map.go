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

type World struct {
	gameMap *tiled.Map
	img     *ebiten.Image
	// origin point (top left) of camera
	camera     image.Point
	cameraAnim units.Anim2D
	StartPoint units.TPoint
	Canvas     *ebiten.Image
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
	canvas, err := ebiten.NewImage(img.Bounds().Dx(), img.Bounds().Dy(), ebiten.FilterDefault)
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
		Canvas:     canvas,
		StartPoint: start,
	}, nil
}

func MustNew() World {
	w, err := New()
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

// func (w *World) FindShortestPath(from, to image.Point) (steps []units.Direction, ok bool) {

// }

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

func (w *World) Draw(screen *ebiten.Image) {
	screen.DrawImage(w.img, nil)
}
