package units

import (
	"image"
)

type Anim struct {
	m        float64
	b        int
	t        int
	callback func()
}

func (a *Anim) Tick() int {
	defer func() {
		if a.t == 0 && a.callback != nil {
			a.callback()
			a.callback = nil
		}
	}()
	if a.t == 0 {
		return a.b
	}
	a.t--
	return int(float64(a.b) - (float64(a.t) * a.m))
}

func (a *Anim) IsMoving() bool {
	return a.t > 0
}

func Animate(start, end int, rate float64, callback func()) Anim {
	if rate <= 0 {
		panic("rate must be positive")
	}
	if end < start {
		rate = rate * -1
	}
	t := int(float64(end-start) / rate)
	if t == 0 {
		return Anim{b: end}
	}
	return Anim{
		b:        end,
		m:        rate,
		t:        t,
		callback: callback,
	}
}

type Anim2D struct {
	x        Anim
	y        Anim
	callback func()
}

func (a *Anim2D) Tick() image.Point {
	return image.Pt(a.x.Tick(), a.y.Tick())
}

func (a *Anim2D) IsMoving() bool {
	return a.x.IsMoving() || a.y.IsMoving()
}

func (a Anim2D) childCallback() {
	if a.callback == nil || a.x.IsMoving() || a.y.IsMoving() {
		return
	}
	a.callback()
	a.callback = nil
}

func Animate2D(start, end image.Point, rate float64, callback func()) Anim2D {
	a := Anim2D{callback: callback}
	a.x = Animate(start.X, end.X, rate, a.childCallback)
	a.y = Animate(start.Y, end.Y, rate, a.childCallback)
	return a
}
