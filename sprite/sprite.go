package sprite

import (
	"github.com/hajimehoshi/ebiten"
)

type Animation interface {
	Frame() *ebiten.Image
	Tick()
}

// A Sprite plays an animation
type Sprite struct {
	frames  []*ebiten.Image
	index   int
	playing bool
	rate    int
}

func New(frames ...*ebiten.Image) Sprite {
	return Sprite{frames: frames, playing: true, rate: 1}
}

func (s Sprite) Rate(rate int) Sprite {
	return Sprite{frames: s.frames, playing: true, rate: rate}
}

func (s Sprite) Reversed() Sprite {
	reversed := make([]*ebiten.Image, len(s.frames))
	for i, f := range s.frames {
		reversed[len(s.frames)-i-1] = f
	}
	return Sprite{frames: reversed, playing: true, rate: s.rate}
}

func (s *Sprite) Frame() *ebiten.Image {
	return s.frames[s.index/s.rate]
}

func (s *Sprite) Tick() {
	if !s.playing {
		return
	}
	if s.index == (s.rate*len(s.frames))-1 {
		s.playing = false
	} else {
		s.index++
	}
}

func (s *Sprite) IsPlaying() bool {
	return s.playing
}

func (s *Sprite) Play() {
	s.playing = true
}

func (s *Sprite) Stop() {
	s.playing = false
}

func (s *Sprite) Reset() {
	s.index = 0
}
