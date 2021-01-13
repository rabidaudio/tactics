package units

import (
	"image"
)

type Anim struct {
	m float64
	b int
	t int
}

// pos = 5 tile
// step = 2 tiles
// pixels = 32
// walk speed = 12 ticks / tile

// t=0 ... pos = 5
// t=1 ... pos = 5 + (t/12) = 5.08333 tiles = 81.333 pixels = 81 pixels
// t=2 ... pos = 5 + (t/12) = 82.666
// t=3 ... 84
// t=11 ... 94.666
// t=12 ... 96

// start=5*16=80
// end=7*16=112
// rate=16/12=1.33 pixels/tick
// last_t=(112-80)/(16/12)

// t=24 .. pos=112-(24*(16/12))

// start=112
// end=80
// rate=1.33
// last_t=(80-112)/(16/12)
// t=24  80 - (24*(-16/12))
// t=12  80 - (12*(-16/12))
// t=0   80 - (0*(-16/12))

func (a *Anim) Tick() int {
	if a.t == 0 {
		return a.b
	}
	a.t--
	return int(float64(a.b) - (float64(a.t) * a.m))
}

func (a *Anim) Value() int {
	if a.t == 0 {
		return a.b
	}
	return int(float64(a.b) - (float64(a.t) * a.m))
}

func Animate(start, end int, rate float64) Anim {
	if rate <= 0 {
		panic("rate must be positive")
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

func (a *Anim2D) Value() image.Point {
	return image.Pt(a.x.Value(), a.y.Value())
}

func Animate2D(start, end image.Point, rate float64) Anim2D {
	return Anim2D{
		x: Animate(start.X, end.X, rate),
		y: Animate(start.Y, end.Y, rate),
	}
}

// // APoint is an animatable point which
// // can update its value over time
// type APoint struct {
// 	image.Point
// 	start  image.Point
// 	target image.Point
// 	t      int
// 	rate   int
// }

// func AP(x, y int) APoint {
// 	return APoint{
// 		Point:  image.Point{X: x, Y: y},
// 		target: image.Point{X: x, Y: y},
// 	}
// }

// func (ap *APoint) Tick() {
// 	if !ap.IsAnimating() {
// 		return
// 	}
// 	ap.t++
// 	ap.Point.X = lerp(ap.start.X, ap.target.X, ap.t, ap.rate)
// 	ap.Point.Y = lerp(ap.start.Y, ap.target.Y, ap.t, ap.rate)
// }

// func (ap *APoint) IsAnimating() bool {
// 	return !ap.Point.Eq(ap.target)
// }

// func (ap *APoint) AnimatePlus(offset image.Point, in Tick) {
// 	ap.AnimateTo(ap.Point.Add(offset), in)
// }

// func (ap *APoint) AnimateTo(to image.Point, in Tick) {
// 	ap.start = ap.Point
// 	ap.target = to
// 	ap.t = 0
// 	ap.rate = int(in)
// }

// func lerp(start, end, t, rate int) int {
// 	s := float64(start)
// 	e := float64(end)
// 	tt := float64(t) / float64(rate)
// 	return int(s*(1-tt) + e*tt)
// }
