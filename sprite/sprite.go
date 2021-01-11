package sprite

import (
	"github.com/hajimehoshi/ebiten"
)

// A Sprite plays an animation
type Sprite struct {
	frames  []*ebiten.Image
	index   int
	Playing bool
	Loop    bool
	Rate    int
}

func (s *Sprite) Frame() *ebiten.Image {
	return s.frames[s.index/s.Rate]
}

func (s *Sprite) Step() {
	if !s.Playing {
		return
	}
	s.index++
	if s.index >= (len(s.frames) * s.Rate) {
		if s.Loop {
			s.index = 0
		} else {
			s.index--
			s.Playing = false
		}
	}
}

func (s *Sprite) Reset() {
	s.index = 0
}

func (s *Sprite) WithRate(r int) *Sprite {
	s.Rate = r
	return s
}

func (s *Sprite) PlayState(playing, loop bool) *Sprite {
	s.Playing = playing
	s.Loop = loop
	return s
}
