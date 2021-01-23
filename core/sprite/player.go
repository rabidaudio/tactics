package sprite

import (
	"github.com/hajimehoshi/ebiten"
)

// A Player keeps a queue of animations and maintains it's own playback state
type Player struct {
	queue []queuedAnimation
}

func NewPlayer() *Player {
	return &Player{
		queue: []queuedAnimation{},
	}
}

type queuedAnimation struct {
	sprite Sprite
	loop   bool
}

// Clear stops all current animations and forget all queued animations
func (p *Player) clear() *Player {
	p.queue = []queuedAnimation{}
	return p
}

// AppendOnce adds this annimation to the end of the queue. If a loop is
// running, it will complete the current cycle before ending the loop
func (p *Player) AppendOnce(sprite Sprite) *Player {
	p.queue = append(p.queue, queuedAnimation{sprite: sprite, loop: false})
	return p
}

// AppendLoop will cause this annimation to play over and over until some
// other animation is queued
func (p *Player) AppendLoop(sprite Sprite) *Player {
	p.queue = append(p.queue, queuedAnimation{sprite: sprite, loop: true})
	return p
}

// ReplaceOnce stops all current and pending animations and
// plays the given animation once
func (p *Player) ReplaceOnce(sprite Sprite) *Player {
	return p.clear().AppendOnce(sprite)
}

// ReplaceLoop stops all current and pending animations and
// plays the given animation over and over until some
// other animation is queued
func (p *Player) ReplaceLoop(sprite Sprite) *Player {
	return p.clear().AppendLoop(sprite)
}

func (p *Player) Tick() {
	if !p.queue[0].sprite.IsPlaying() {
		if len(p.queue) > 1 {
			p.queue = p.queue[1:]
		} else if p.queue[0].loop {
			p.queue[0].sprite.Reset()
			p.queue[0].sprite.Play()
		}
	}
	p.queue[0].sprite.Tick()
}

func (p *Player) Frame() *ebiten.Image {
	return p.queue[0].sprite.Frame()
}
