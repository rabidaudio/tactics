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
	loop    bool
}

func New(frames ...*ebiten.Image) *Sprite {
	return &Sprite{frames: frames, playing: true, rate: 1}
}

func (s *Sprite) Rate(rate int) *Sprite {
	s.rate = rate
	return s
}

func (s *Sprite) Reversed() *Sprite {
	reversed := make([]*ebiten.Image, len(s.frames))
	for i, f := range s.frames {
		reversed[len(s.frames)-i-1] = f
	}
	s.frames = reversed
	return s
}

func (s *Sprite) Loop(loop bool) *Sprite {
	s.loop = loop
	return s
}

func (s *Sprite) Frame() *ebiten.Image {
	return s.frames[s.index/s.rate]
}

func (s *Sprite) Tick() {
	if !s.playing {
		return
	}
	if s.index == (s.rate*len(s.frames))-1 {
		if s.loop {
			s.Reset()
		} else {
			s.Stop()
		}
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
