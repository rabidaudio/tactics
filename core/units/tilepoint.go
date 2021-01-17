package units

import (
	"image"
	"strconv"
)

// A TPoint is in units of tiles, which is a standard unit
// of the game.
type TPoint image.Point

// TP is sugar for constructing new TPoints
func TP(x, y int) TPoint {
	return TPoint{X: x, Y: y}
}

func TPFromPoint(pt image.Point) TPoint {
	return TPoint(pt.Div(TileSize))
}

// IP converts tile units to image units
func (tp TPoint) IP() image.Point {
	return image.Point(tp).Mul(TileSize)
}

func (p TPoint) XY() (int, int) {
	return p.X, p.Y
}

func (p TPoint) Bound(x0, y0, x1, y1 int) TPoint {
	return TPoint{X: bound(p.X, x0, x1), Y: bound(p.Y, y0, y1)}
}

func bound(val, min, max int) int {
	if val < min {
		return min
	} else if val > max {
		return max
	} else {
		return val
	}
}

func Bound(pt image.Point, to image.Rectangle) image.Point {
	return image.Point{
		X: bound(pt.X, to.Min.X, to.Max.X),
		Y: bound(pt.Y, to.Min.Y, to.Max.Y),
	}
}

// The following methods are copied directly from image.Point

// String returns a string representation of p like "(3,4)".
func (p TPoint) String() string {
	return "(" + strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y) + ")"
}

// Add returns the vector p+q.
func (p TPoint) Add(q TPoint) TPoint {
	return TPoint{p.X + q.X, p.Y + q.Y}
}

// Sub returns the vector p-q.
func (p TPoint) Sub(q TPoint) TPoint {
	return TPoint{p.X - q.X, p.Y - q.Y}
}

// Mul returns the vector p*k.
func (p TPoint) Mul(k int) TPoint {
	return TPoint{p.X * k, p.Y * k}
}

// Div returns the vector p/k.
func (p TPoint) Div(k int) TPoint {
	return TPoint{p.X / k, p.Y / k}
}

// In reports whether p is in r.
func (p TPoint) In(r image.Rectangle) bool {
	return r.Min.X <= p.X && p.X < r.Max.X &&
		r.Min.Y <= p.Y && p.Y < r.Max.Y
}

// Mod returns the point q in r such that p.X-q.X is a multiple of r's width
// and p.Y-q.Y is a multiple of r's height.
func (p TPoint) Mod(r image.Rectangle) TPoint {
	w, h := r.Dx(), r.Dy()
	p = p.Sub(TPoint(r.Min))
	p.X = p.X % w
	if p.X < 0 {
		p.X += w
	}
	p.Y = p.Y % h
	if p.Y < 0 {
		p.Y += h
	}
	return p.Add(TPoint(r.Min))
}
