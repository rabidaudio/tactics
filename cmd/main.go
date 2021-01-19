package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/game"
)

func main() {
	g := game.New()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Untitled")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
