package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/assets"
)

type Game struct {
	spear assets.CharacterAnimation
	tick  uint64
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	g.spear.Sprite().Step()
	if g.tick%100 == 0 {
		g.spear.State = assets.Death
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
	f := g.spear.Sprite().Frame()
	screen.DrawImage(f, &ebiten.DrawImageOptions{
		GeoM: ebiten.TranslateGeo(100, 100),
	})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	game := &Game{
		spear: assets.Spearman(),
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tactics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
