package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/rabidaudio/tactics/assets"
	"github.com/rabidaudio/tactics/chars/spearman"
	"github.com/rabidaudio/tactics/core"
	"github.com/rabidaudio/tactics/core/sprite"
	"github.com/rabidaudio/tactics/core/units"
	"github.com/rabidaudio/tactics/core/window"
	"github.com/rabidaudio/tactics/world"
)

type Game struct {
	window   *window.Window
	world    world.World
	spearman spearman.Spearman
	i        int
	sprite   func(ca *CharAnimation) sprite.Sprite
	loop     bool
	player   *sprite.Player
	tick     units.Tick
	ready    bool
}

type CharacterMoveAction struct {
	Direction units.Direction
}

type ChangeSpriteSet struct {
	Step int
}

// type Animation int

// const (
// 	Idle Animation = iota
// 	Walk
// 	Attack
// 	Hit
// 	Death
// )

type PlayAnimation struct {
	Sprite func(ca *CharAnimation) sprite.Sprite
	Loop   bool
}

type CharAnimation struct {
	Attack, Death, Hit, Idle, Walk func() sprite.Sprite
}

var animations = []CharAnimation{
	assets.Spearman,
	assets.Halberd,
	assets.Swordsman,
	assets.Swordsman2,
	assets.Axeman,
	assets.Hammerman,
	assets.Crossbowman,
	assets.Crossbowman2,
	assets.Hunter,
	assets.Hunter2,
	assets.Monk,
	assets.Bishop,
	assets.Rider,
	assets.Rider2,
	assets.Knight,
	assets.Knight2,
	assets.Mercenary,
	assets.Brute,
	assets.Brute2,
	assets.GoblinRider,
	assets.GoblinRider2,
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	g.tick++
	g.window.Tick()
	// g.character.Tick()

	g.player.Tick()

	core.ActionHandler().OnKey(map[ebiten.Key]core.Action{
		ebiten.Key1:   PlayAnimation{func(ca *CharAnimation) sprite.Sprite { return ca.Idle() }, true},
		ebiten.Key2:   PlayAnimation{func(ca *CharAnimation) sprite.Sprite { return ca.Walk() }, true},
		ebiten.Key3:   PlayAnimation{func(ca *CharAnimation) sprite.Sprite { return ca.Attack() }, false},
		ebiten.Key4:   PlayAnimation{func(ca *CharAnimation) sprite.Sprite { return ca.Hit() }, false},
		ebiten.Key5:   PlayAnimation{func(ca *CharAnimation) sprite.Sprite { return ca.Death() }, false},
		ebiten.KeyTab: ChangeSpriteSet{Step: 1},
		ebiten.KeyZ:   ChangeSpriteSet{Step: -1},
	}).Execute(func(a core.Action) {
		if aa, ok := a.(PlayAnimation); ok {
			g.loop = aa.Loop
			g.sprite = aa.Sprite
		} else if aa, ok := a.(ChangeSpriteSet); ok {
			g.i = (g.i + aa.Step) % len(animations)
			if g.i < 0 {
				g.i = len(animations) - 1
			}
		}
		if g.loop {
			g.player.ReplaceLoop(g.sprite(&animations[g.i]))
		} else {
			g.player.ReplaceOnce(g.sprite(&animations[g.i]))
		}
	})
	// g.spearman.Tick()
	// g.ready = !(g.window.IsCameraMoving() || g.spearman.IsMoving())
	// if !g.ready {
	// 	return nil
	// }
	// core.ActionHandler().
	// 	OnKey(map[ebiten.Key]core.Action{
	// 		ebiten.KeyA: CharacterMoveAction{Direction: units.West},
	// 		ebiten.KeyS: CharacterMoveAction{Direction: units.South},
	// 		ebiten.KeyD: CharacterMoveAction{Direction: units.East},
	// 		ebiten.KeyW: CharacterMoveAction{Direction: units.North},
	// 	}).
	// 	OnLeftMouseClick(func(screenPoint image.Point) core.Action {
	// 		p := units.TPFromPoint(screenPoint.Add(g.window.CameraOrigin()))
	// 		if d, ok := units.TPFromPoint(g.spearman.Location).IsAdjacent(p); ok {
	// 			return CharacterMoveAction{Direction: d}
	// 		}
	// 		return nil
	// 	}).
	// 	Execute(func(action core.Action) {
	// 		dir := action.(CharacterMoveAction).Direction
	// 		t := units.TPFromPoint(g.spearman.Location).Add(dir.TP())
	// 		if !g.world.IsBoundary(t) {
	// 			g.spearman.Go(dir)
	// 			g.window.AnimateCamera(t)
	// 		}
	// 	})
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(g.world.Canvas)
	// g.spearman.Draw(g.world.Canvas)
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.world.StartPoint.IP().X), float64(g.world.StartPoint.IP().Y))
	g.world.Canvas.DrawImage(g.player.Frame(), &opts)
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

		sprite: func(ca *CharAnimation) sprite.Sprite { return ca.Idle() },
		loop:   true,
		player: sprite.NewPlayer().AppendLoop(animations[0].Idle()),
	}
	game.window.WorldSize(game.world.Size())
	game.window.JumpCamera(game.world.StartPoint)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Untitled")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
