package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/chars/spearman"
)

type Game struct {
	spearman spearman.Spearman
	tick     uint64
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	g.spearman.Step()
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.spearman.Go(spearman.West)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.spearman.Go(spearman.South)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.spearman.Go(spearman.East)
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.spearman.Go(spearman.North)
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	g.spearman.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		spearman: spearman.New(),
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tactics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
