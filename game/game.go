package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/pkg/errors"
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/unit"
)

var ErrQuit = errors.Errorf("Quit due to user input")

type Game struct {
	Window *core.Window
	World  World
	Units  []*unit.Unit
	Tick   units.Tick
	Turn   unit.Team
}

const (
	PlayerTeam unit.Team = iota
	EnemyTeam
)

func New() *Game {
	// "raw/maps/map1.tmx"
	w := MustNewWorld("raw/maps/square.tmx")
	game := &Game{
		Window: &core.Window{
			// Size: image.Point{X: 230, Y: 240},
			Size: image.Point{X: 320, Y: 288},
			// Size: image.Point{X: 160, Y: 144},
		},
		World: w,
		Units: []*unit.Unit{
			unit.NewSpearman(w.StartPoint, PlayerTeam, 1),
			unit.NewSpearman(w.StartPoint.Add(units.TP(1, 0)), EnemyTeam, 1),
		},
		Turn: PlayerTeam,
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

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return ErrQuit
	}

	for _, u := range g.Units {
		u.Tick()
	}
	// u := g.Units[0]
	// TODO [arch] who's responsibility is it to verify
	// actions are legal? is it the game's? the unit's? the world's?
	// core.ActionHandler().
	// 	OnLeftMouseClick(func(_ image.Point) core.Action {
	// 		p := g.CursorPosition()
	// 		if target := g.UnitAt(p); target != nil {
	// 			if g.canAttack(u, target) {
	// 				return unit.AttackCommand{Unit: u, Target: target}
	// 			}
	// 			return nil
	// 		}
	// 		if !g.canMoveTo(p, u) {
	// 			return nil
	// 		}
	// 		canMove := func(pt units.TPoint) bool {
	// 			return g.canMoveThrough(pt, u)
	// 		}
	// 		if d, ok := core.FindPath(u.Location, p, canMove); ok {
	// 			return unit.MoveCommand{Unit: u, Steps: d}
	// 		}
	// 		return nil
	// 	}).
	// 	Execute(func(action core.Action) {
	// 		// u.Handle(action)
	// 	})
	return nil
}

func (g *Game) canAttack(attacker, target *unit.Unit) bool {
	if attacker.Team == target.Team {
		return false // TODO [mechanics] healing
	}
	if !attacker.IsReady() || !target.IsReady() {
		return false
	}
	if !attacker.CanReach(target) {
		return false
	}
	return true
}

func (g *Game) UnitAt(pt units.TPoint) *unit.Unit {
	for _, u := range g.Units {
		if u.Location == pt {
			return u
		}
	}
	return nil
}

func (g *Game) canMoveTo(dest units.TPoint, unit *unit.Unit) bool {
	// TODO [bug] doesn't check that all these steps are accessible
	if unit.Location.StepsTo(dest) > unit.StepsPerTurn() {
		return false
	}
	if g.World.IsBoundary(dest) {
		return false
	}
	if u := g.UnitAt(dest); u != nil {
		return false
	}
	return true
}

func (g *Game) canMoveThrough(dest units.TPoint, unit *unit.Unit) bool {
	// TODO [style] share logic better with canMoveTo
	if g.World.IsBoundary(dest) {
		return false
	}
	// can move through friendly units but not enemy units
	// TODO [mechanics] desired behavior?
	if u := g.UnitAt(dest); u != nil && u.Team != unit.Team {
		return false
	}
	return true
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	canvas := g.World.Draw(func(canvas *ebiten.Image) {
		g.drawCursor(canvas)
		for _, u := range g.Units {
			u.Draw(canvas)
		}
	})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.World.StartPoint.IP().X), float64(g.World.StartPoint.IP().Y))
	screen.DrawImage(canvas.SubImage(g.Window.Rect()).(*ebiten.Image), nil)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Size.X, g.Window.Size.Y
}

func (g *Game) CursorPosition() units.TPoint {
	return units.TPFromPoint(image.Pt(ebiten.CursorPosition()).Add(g.Window.CameraOrigin()))
}

func (g *Game) drawCursor(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	tp := g.CursorPosition()
	tile := assets.TileSelectable
	if !g.canMoveTo(tp, g.Units[0]) {
		tile = assets.TileNotSelectable
	}
	p := tp.IP()
	opts.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(tile, &opts)
}
