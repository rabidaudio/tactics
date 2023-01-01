package game

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/game/team"
	"github.com/rabidaudio/tactics/game/unit"
)

type BattleScene struct {
	core.WorldScene
	core.Cursor
	Units    map[team.Team][]*unit.Unit
	Selected *unit.Unit
	Turn     team.Team
}

func NewBattleScene(g *Game, scene core.Scene) BattleScene {
	// "raw/maps/map1.tmx"
	b := BattleScene{}
	// Size: image.Point{X: 230, Y: 240},
	// Size: image.Point{X: 320, Y: 288},
	// Size: image.Point{X: 160, Y: 144},
	mapname := "raw/maps/square.tmx"
	size := image.Point{X: 320, Y: 288}
	b.WorldScene = core.SceneWithWorld(&b.WorldScene, mapname, size)
	b.Units = map[team.Team][]*unit.Unit{
		team.Player: {
			unit.NewSpearman(b.World.StartPoint, team.Player, 1),
		},
		team.Enemy: {
			unit.NewSpearman(b.World.StartPoint.Add(units.TP(1, 0)), team.Enemy, 1),
		},
	}
	b.Cursor = core.Cursor{
		Window: b.WorldScene.Window,
		IsSelectable: func(t units.TPoint) bool {
			if b.Turn != team.Player {
				return false
			}
			cmd := b.NewMoveCommand(g.Selected, t)
			return cmd.IsLegal(b)
		},
	}
	b.Turn = team.Player
	b.Selected = b.Units[b.Turn][0]
	return b
}

func (b *BattleScene) Tick() {
	for _, uu := range b.Units {
		for _, u := range uu {
			u.Tick()
		}
	}
	// core.ActionHandler().
	// 	OnLeftMouseClick(func(_ image.Point) core.Action {
	// 		// Here we just need to generate the command, it
	// 		// doesn't actually need to be valid
	// 		p := b.CursorPosition()
	// 		if target := b.UnitAt(p); target != nil {
	// 			return AttackCommand{Attacker: b.Selected, Target: target}
	// 		}
	// 		return b.NewMoveCommand(w, b.Selected, p)
	// 	}).
	// 	Execute(func(action core.Action) {
	// 		cmd := action.(Command)
	// 		if cmd.IsLegal(g) {
	// 			cmd.Execute(g)
	// 		}
	// 	})
}

func (b *BattleScene) Draw(screen *ebiten.Image) {
	b.Cursor.Draw(screen)
	for _, uu := range b.Units {
		for _, u := range uu {
			u.Draw(screen)
		}
	}
}

func (b *BattleScene) UnitAt(pt units.TPoint) *unit.Unit {
	for _, uu := range b.Units {
		for _, u := range uu {
			if u.Location == pt {
				return u
			}
		}
	}
	return nil
}
