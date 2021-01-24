package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/unit"
)

type Game struct {
	Window *core.Window
	World  World
	Units  []*unit.Unit
	Tick   units.Tick
}

const (
	PlayerTeam unit.Team = iota
	EnemyTeam
)

func New() *Game {
	// "raw/maps/map1.tmx"
	w := MustNewWorld("raw/maps/square.tmx")
	game := &Game{
		Window: &core.Window{Size: image.Point{X: 230, Y: 240}},
		World:  w,
		Units: []*unit.Unit{
			unit.NewSpearman(w.StartPoint, PlayerTeam, 1),
			unit.NewSpearman(w.StartPoint.Add(units.TP(1, 0)), EnemyTeam, 1),
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
	u := g.Units[0]
	// TODO [arch] who's responsibility is it to verify
	// actions are legal? is it the game's? the unit's? the world's?
	core.ActionHandler().
		OnLeftMouseClick(func(screenPoint image.Point) core.Action {
			p := units.TPFromPoint(screenPoint.Add(g.Window.CameraOrigin()))
			if target, ok := g.unitAt(p); ok {
				if target.Team == PlayerTeam {
					// TODO [gameplay] switch players
					return nil
				}
				// otherwise enemy player
				if !u.Weapon.CanHit(u.Location, target.Location) {
					return nil
				}
				return unit.AttackCommand{Unit: u, Target: target}
			}
			if !g.canMoveTo(p, u) {
				return nil
			}
			canMove := func(pt units.TPoint) bool {
				return g.canMoveThrough(pt, u)
			}
			if d, ok := core.FindPath(u.Location, p, canMove); ok {
				return unit.MoveCommand{Unit: u, Steps: d}
			}
			return nil
		}).
		Execute(func(action core.Action) {
			u.Handle(action)
		})
	return nil
}

func (g *Game) unitAt(pt units.TPoint) (*unit.Unit, bool) {
	for _, u := range g.Units {
		if u.Location == pt {
			return u, true
		}
	}
	return nil, false
}

func (g *Game) canMoveTo(dest units.TPoint, unit *unit.Unit) bool {
	if unit.Location.StepsTo(dest) > unit.Stats.Steps {
		return false
	}
	if g.World.IsBoundary(dest) {
		return false
	}
	if _, ok := g.unitAt(dest); ok {
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
	if u, ok := g.unitAt(dest); ok && u.Team != unit.Team {
		return false
	}
	return true
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	canvas := g.World.Draw(func(canvas *ebiten.Image) {
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
