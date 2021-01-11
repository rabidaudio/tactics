package sprite

import (
	"github.com/hajimehoshi/ebiten"
)

// A Sprite plays an animation
type Sprite interface {
	Frame() *ebiten.Image
	Tick()
	IsPlaying() bool
	Play()
	Stop()
	Reset()
}

func New(frames ...*ebiten.Image) *Builder {
	return &Builder{&pauseSprite{frames: frames, playing: true}}
}

type Builder struct {
	s Sprite
}

func (b *Builder) Loop() *Builder {
	b.s = &loopingSprite{b.s}
	return b
}

func (b *Builder) Rate(rate int) *Builder {
	b.s = &rateSprite{Sprite: b.s, rate: rate}
	return b
}

func (b *Builder) Concat(sprites ...Sprite) *Builder {
	ss := append([]Sprite{b.s}, sprites...)
	b.s = &concatSprite{sprites: ss}
	return b
}

func (b *Builder) Sprite() Sprite {
	return b.s
}

type pauseSprite struct {
	frames  []*ebiten.Image
	index   int
	playing bool
}

var _ Sprite = (*pauseSprite)(nil)

func (s *pauseSprite) Frame() *ebiten.Image {
	return s.frames[s.index]
}

func (s *pauseSprite) Tick() {
	if !s.playing {
		return
	}
	if s.index == len(s.frames)-1 {
		s.playing = false
	} else {
		s.index++
	}
}

func (s *pauseSprite) IsPlaying() bool {
	return s.playing
}

func (s *pauseSprite) Play() {
	s.playing = true
}

func (s *pauseSprite) Stop() {
	s.playing = false
}

func (s *pauseSprite) Reset() {
	s.index = 0
}

type loopingSprite struct {
	Sprite
}

var _ Sprite = (*loopingSprite)(nil)

func (s *loopingSprite) Tick() {
	if !s.IsPlaying() {
		return
	}
	s.Sprite.Tick()
	if !s.Sprite.IsPlaying() {
		s.Sprite.Reset()
		s.Play()
	}
}

type rateSprite struct {
	Sprite
	rate  int
	index int
}

var _ Sprite = (*rateSprite)(nil)

func (s *rateSprite) Tick() {
	s.index++
	if s.index == s.rate {
		s.index = 0
		s.Sprite.Tick()
	}
}

type concatSprite struct {
	sprites []Sprite
	index   int
}

var _ Sprite = (*concatSprite)(nil)

func (s *concatSprite) Frame() *ebiten.Image {
	return s.sprites[s.index].Frame()
}

func (s *concatSprite) IsPlaying() bool {
	return s.sprites[s.index].IsPlaying()
}

func (s *concatSprite) Tick() {
	if !s.IsPlaying() {
		return
	}
	s.sprites[s.index].Tick()
	if !s.sprites[s.index].IsPlaying() {
		if s.index < len(s.sprites)-1 {
			s.index++
		}
	}
}

func (s *concatSprite) Play() {
	s.sprites[s.index].Play()
}

func (s *concatSprite) Stop() {
	s.sprites[s.index].Stop()
}

func (s *concatSprite) Reset() {
	for _, ss := range s.sprites {
		ss.Reset()
	}
	s.index = 0
}
