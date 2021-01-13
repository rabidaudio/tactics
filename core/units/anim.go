package units

import (
	"image"
)

type Anim struct {
	m float64
	b int
	t int
}

func (a *Anim) Tick() int {
	if !a.IsMoving() {
		return a.b
	}
	a.t--
	return int(float64(a.b) - (float64(a.t) * a.m))
}

func (a *Anim) IsMoving() bool {
	return a.t > 0
}

func Animate(start, end int, rate float64) Anim {
	if rate <= 0 {
		panic("rate must be positive")
	}
	if start == end {
		return Anim{t: 0, b: end}
	}
	if end < start {
		rate = rate * -1
	}
	return Anim{
		b: end,
		m: rate,
		t: int(float64(end-start) / rate),
	}
}

type Anim2D struct {
	x Anim
	y Anim
}

func (a *Anim2D) Tick() image.Point {
	return image.Pt(a.x.Tick(), a.y.Tick())
}

func (a *Anim2D) IsMoving() bool {
	return a.x.IsMoving() || a.y.IsMoving()
}

func Animate2D(start, end image.Point, rate float64) Anim2D {
	return Anim2D{
		x: Animate(start.X, end.X, rate),
		y: Animate(start.Y, end.Y, rate),
	}
}
