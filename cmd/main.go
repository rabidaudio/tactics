package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/game"
)

func main() {
	g := game.New()
	// TODO [graphics] window size
	ebiten.SetWindowSize(640, 480)
	// TODO [story] game title
	ebiten.SetWindowTitle("Untitled")
	if err := ebiten.RunGame(g); err != nil {
		if err == game.ErrQuit {
			return
		}
		log.Fatal(err)
	}
}
