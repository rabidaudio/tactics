package units

import (
	"image"
)

// APoint is an animatable point which
// can update its value over time
type APoint struct {
	image.Point
	start  image.Point
	target image.Point
	t      int
	rate   int
}

func AP(x, y int) APoint {
	return APoint{
		Point:  image.Point{X: x, Y: y},
		target: image.Point{X: x, Y: y},
	}
}

func (ap *APoint) Tick() {
	if !ap.IsAnimating() {
		return
	}
	ap.t++
	ap.Point.X = lerp(ap.start.X, ap.target.X, ap.t, ap.rate)
	ap.Point.Y = lerp(ap.start.Y, ap.target.Y, ap.t, ap.rate)
}

func (ap *APoint) IsAnimating() bool {
	return !ap.Point.Eq(ap.target)
}

func (ap *APoint) AnimatePlus(offset image.Point, in Tick) {
	ap.AnimateTo(ap.Point.Add(offset), in)
}

func (ap *APoint) AnimateTo(to image.Point, in Tick) {
	ap.start = ap.Point
	ap.target = to
	ap.t = 0
	ap.rate = int(in)
}

func lerp(start, end, t, rate int) int {
	s := float64(start)
	e := float64(end)
	tt := float64(t) / float64(rate)
	return int(s*(1-tt) + e*tt)
}
