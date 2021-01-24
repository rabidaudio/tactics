package sprite

import (
	"github.com/hajimehoshi/ebiten"
)

type Template struct {
	frames []*ebiten.Image
	rate   int
	loop   bool
}

type Sprite struct {
	template *Template
	index    int
	playing  bool
	complete func()
}

func NewTemplate(frames ...*ebiten.Image) Template {
	return Template{frames: frames, rate: 1}
}

func (s Template) Rate(rate int) Template {
	s.rate = rate
	return s
}

func (s Template) Reversed() Template {
	reversed := make([]*ebiten.Image, len(s.frames))
	for i, f := range s.frames {
		reversed[len(s.frames)-i-1] = f
	}
	s.frames = reversed
	return s
}

func (s Template) Loop(loop bool) Template {
	s.loop = loop
	return s
}

func (s Template) Append(other Template) Template {
	return NewTemplate(append(s.frames, other.frames...)...)
}

func (s Template) Sprite() *Sprite {
	return &Sprite{template: &s, playing: true, index: 0}
}

func (s *Sprite) OnComplete(callback func()) *Sprite {
	s.complete = callback
	return s
}

func (s *Sprite) Frame() *ebiten.Image {
	return s.template.frames[s.index/s.template.rate]
}

func (s *Sprite) Tick() {
	if !s.playing {
		return
	}
	if s.index == (s.template.rate*len(s.template.frames))-1 {
		if s.template.loop {
			s.Reset()
		} else {
			if s.complete != nil {
				s.complete()
				s.complete = nil
			}
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
