package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/chars/spearman"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/window"
	"github.com/rabidaudio/tactics/world"
)

type Game struct {
	window   *window.Window
	world    world.World
	spearman spearman.Spearman
	tick     units.Tick
	ready    bool
}

type CharacterMoveAction struct {
	Direction units.Direction
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	g.window.Tick()
	g.spearman.Tick()
	g.ready = !(g.window.IsCameraMoving() || g.spearman.IsMoving())
	if !g.ready {
		return nil
	}
	core.ActionHandler().
		OnKey(map[ebiten.Key]core.Action{
			ebiten.KeyA: CharacterMoveAction{Direction: units.West},
			ebiten.KeyS: CharacterMoveAction{Direction: units.South},
			ebiten.KeyD: CharacterMoveAction{Direction: units.East},
			ebiten.KeyW: CharacterMoveAction{Direction: units.North},
		}).
		OnLeftMouseClick(func(screenPoint image.Point) core.Action {
			p := units.TPFromPoint(screenPoint.Add(g.window.CameraOrigin()))
			if d, ok := units.TPFromPoint(g.spearman.Location).IsAdjacent(p); ok {
				return CharacterMoveAction{Direction: d}
			}
			return nil
		}).
		Execute(func(action core.Action) {
			dir := action.(CharacterMoveAction).Direction
			t := units.TPFromPoint(g.spearman.Location).Add(dir.TP())
			if !g.world.IsBoundary(t) {
				g.spearman.Go(dir)
				g.window.AnimateCamera(t)
			}
		})
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(g.world.Canvas)
	g.spearman.Draw(g.world.Canvas)
	screen.DrawImage(g.world.Canvas.SubImage(g.window.Rect()).(*ebiten.Image), nil)
	g.world.Canvas.Clear()
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.window.Size.X, g.window.Size.Y
}

func main() {
	world, err := world.New()
	if err != nil {
		log.Fatal(err)
	}
	game := &Game{
		window:   &window.Window{Size: image.Point{X: 230, Y: 240}},
		world:    world,
		spearman: spearman.New(world.StartPoint),
	}
	game.window.WorldSize(game.world.Size())
	game.window.JumpCamera(game.world.StartPoint)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tactics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
