package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/chars/spearman"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/world"
)

type Game struct {
	world    world.World
	spearman spearman.Spearman
	tick     units.Tick
	pressed  bool
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	g.spearman.Tick()
	g.world.Tick()
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	if !g.pressed {
	// 		p := image.Pt(ebiten.CursorPosition()).Add(g.world.CameraOrigin())
	// 		g.world.Go(units.TPFromPoint(p))
	// 		g.pressed = true
	// 	}
	// } else if ebiten.IsKeyPressed(ebiten.KeyA) {
	// 	if !g.pressed {
	// 		g.spearman.Go(units.West)
	// 		g.pressed = true
	// 	}
	// } else if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	if !g.pressed {
	// 		g.spearman.Go(units.South)
	// 		g.pressed = true
	// 	}
	// } else if ebiten.IsKeyPressed(ebiten.KeyD) {
	// 	if !g.pressed {
	// 		g.spearman.Go(units.East)
	// 		g.pressed = true
	// 	}
	// } else if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	if !g.pressed {
	// 		g.spearman.Go(units.North)
	// 		g.pressed = true
	// 	}
	// } else {
	// 	g.pressed = false
	// }
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)
	g.spearman.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return world.WindowSize.X, world.WindowSize.Y
}

func main() {
	world, err := world.New()
	if err != nil {
		log.Fatal(err)
	}
	game := &Game{
		spearman: spearman.New(world.StartPoint),
		world:    world,
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tactics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
