package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/pkg/errors"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
)

var ErrQuit = errors.Errorf("Quit due to user input")

type Game struct {
	Window *core.Window
	Scene  core.Scene
	Tick   units.Tick
}

func New() *Game {
	game := &Game{
		Window: &core.Window{
			// Size: image.Point{X: 230, Y: 240},
			Size: image.Point{X: 320, Y: 288},
			// Size: image.Point{X: 160, Y: 144},
		},
	}
	return game
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ErrQuit
	}

	g.Tick++
	g.Scene.Tick()
	g.Window.Tick()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Size.X, g.Window.Size.Y
}
