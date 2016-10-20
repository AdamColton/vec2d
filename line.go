package vec2d

import (
	"math"
)

// Line is a parametric equation. It takes a value for t and returns a point
// along the line. Using a parametric equation makes it easier to handle
// vertical lines
type Line func(t float64) F

// LineTo returns a line where t=0 returns the from point and t=1 returns the to
// point
func (f F) LineTo(to F) Line {
	d := to.Subtract(f)
	return func(t float64) F { return F{f.X + d.X*t, f.Y + d.Y*t} }
}

// Bisect returns a line that bisects points a and b. All points on the line are
// equadistant from both a and b. At t=0, the mid-point is returned.
func (f F) Bisect(f2 F) Line {
	c := f.Add(f2).ScalarMultiply(0.5)
	d := c.Add(F{f.Y - f2.Y, f2.X - f.X})
	return c.LineTo(d)
}

// AtX Returns the value of t at x. May return Inf.
func (l Line) AtX(x float64) float64 {
	x0 := l(0).X
	dx := l(1).X - x0
	return (x - x0) / dx
}

// AtY returns the value of t at y. May return Inf.
func (l Line) AtY(y float64) float64 {
	y0 := l(0).Y
	dy := l(1).Y - y0
	return (y - y0) / dy
}

// B returns the X intercept, from the form y = m*x + b
func (l Line) B() float64 { return l(l.AtX(0)).Y }

// M returns the slope, from the form y = m*x + b
func (l Line) M() float64 {
	p0, p1 := l(0), l(1)
	return (p0.Y - p1.Y) / (p0.X - p1.X)
}

// Intersection returns the points at which the lines intersect. Two values are
// returned that indicate the index points at the line. If the lines do not
// intersect, NaN will be returned for both values.
func (l Line) Intersection(l2 Line) (float64, float64) {
	a0, b0 := l(0), l2(0)
	a1, b1 := l(1), l2(1)
	da, db := a1.Subtract(a0), b1.Subtract(b0)

	d := db.Y*da.X - da.Y*db.X
	if d == 0 {
		// lines do not intersect
		return math.NaN(), math.NaN()
	}

	tb := (da.Y*(b0.X-a0.X) + da.X*(a0.Y-b0.Y)) / d
	if da.X != 0 {
		ta := (b0.X + db.X*tb - a0.X) / da.X
		return ta, tb
	} else if da.Y == 0 {
		// la is not a line but a point, doing something like
		// p.LineTo(p) would create this
		return math.NaN(), math.NaN()
	}
	ta := (b0.Y + db.Y*tb - a0.Y) / da.Y
	return ta, tb
}
