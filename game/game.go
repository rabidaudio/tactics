package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/core/window"
	"github.com/rabidaudio/tactics/game/unit"
	"github.com/rabidaudio/tactics/game/world"
)

type Game struct {
	Window *window.Window
	World  world.World
	Units  []*unit.Unit
	Tick   units.Tick
}

func New() *Game {
	w := world.MustNew()
	game := &Game{
		Window: &window.Window{Size: image.Point{X: 230, Y: 240}},
		World:  w,
		Units: []*unit.Unit{
			unit.NewSpearman(w.StartPoint),
		},
	}
	game.Window.WorldSize(w.Size())
	game.Window.JumpCamera(w.StartPoint)
	return game
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.Tick++
	g.Window.Tick()

	for _, u := range g.Units {
		u.Tick()
	}
	// g.ready = !(g.window.IsCameraMoving() || g.spearman.IsMoving())
	// if !g.ready {
	// 	return nil
	// }
	u := g.Units[0]
	core.ActionHandler().
		OnKey(map[ebiten.Key]core.Action{
			ebiten.KeyA: unit.Move(u, units.West),
			ebiten.KeyS: unit.Move(u, units.South),
			ebiten.KeyD: unit.Move(u, units.East),
			ebiten.KeyW: unit.Move(u, units.North),
		}).
		OnLeftMouseClick(func(screenPoint image.Point) core.Action {
			p := units.TPFromPoint(screenPoint.Add(g.Window.CameraOrigin()))
			if d, ok := units.TPFromPoint(u.Location).IsAdjacent(p); ok {
				return unit.Move(u, d)
			}
			return nil
		}).
		Execute(func(action core.Action) {
			u.Handle(action)
			// dir := action.(CharacterMoveAction).Direction
			// t := units.TPFromPoint(g.spearman.Location).Add(dir.TP())
			// if !g.world.IsBoundary(t) {
			// 	g.spearman.Go(dir)
			// 	g.window.AnimateCamera(t)
			// }
		})
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.World.Draw(g.World.Canvas)
	for _, u := range g.Units {
		u.Draw(g.World.Canvas)
	}
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.World.StartPoint.IP().X), float64(g.World.StartPoint.IP().Y))
	screen.DrawImage(g.World.Canvas.SubImage(g.Window.Rect()).(*ebiten.Image), nil)
	g.World.Canvas.Clear()
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Size.X, g.Window.Size.Y
}
