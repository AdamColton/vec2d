package vec2d

import (
	"fmt"
	"math"
)

// F is a 2D float64 vector. It can be used to represent a point or the
// difference between two points.
type F struct {
	X, Y float64
}

// Add returns f + f2
func (f F) Add(f2 F) F {
	f.X += f2.X
	f.Y += f2.Y
	return f
}

// Subtract returns f - f2
func (f F) Subtract(f2 F) F {
	f.X -= f2.X
	f.Y -= f2.Y
	return f
}

// Area returns f.X*f.Y
func (f F) Area() float64 {
	return f.X * f.Y
}

// Multiply returns F{f.X * f2.X, f.Y * f2.Y}
func (f F) Multiply(f2 F) F {
	f.X *= f2.X
	f.Y *= f2.Y
	return f
}

// ScalarMultiply returns F{f.X*sclr, f.Y*sclr}
func (f F) ScalarMultiply(sclr float64) F {
	f.X *= sclr
	f.Y *= sclr
	return f
}

// Angle returns the angle in radians
func (f F) Angle() float64 {
	return math.Atan2(f.Y, f.X)
}

// Mag returns the magnitude (distance to origin) of the vector
func (f F) Mag() float64 {
	return math.Sqrt(f.X*f.X + f.Y*f.Y)
}

// Rotate returns a new Vector rotated by r radians
func (f F) Rotate(r float64) F {
	m := f.Mag()
	r += f.Angle()
	return F{math.Cos(r) * m, math.Sin(r) * m}
}

// Distance returns the distance between to points
func (f F) Distance(f2 F) float64 {
	return f.Subtract(f2).Mag()
}

// I converts a float64 vector to an int vector. Will always round down.
func (f F) I() I {
	return I{int(f.X), int(f.Y)}
}

// P converts a float64 vector to a Polar vector
func (f F) P() P {
	return P{f.Mag(), f.Angle()}
}

// String fulfills Stringer, returns the vector as "(X, Y)"
func (f F) String() string {
	return fmt.Sprintf("(%f, %f)", f.X, f.Y)
}

// MotionSurfaceIntersection Takes the motion path described by mStart t(0) and
// mEnd t(1) and finds the time that the motion intersects the line segment
// described by sStart to sEnd. If there is an intersection and it happens
// between t(0) and t(1), the value will be returned, otherwise NaN will be
// returned
func MotionSurfaceIntersection(mStart, mEnd, sStart, sEnd F) float64 {
	/*
		The following is a solution to these parametric equations
		  sStart.X + S*dS.X = mStart.X + M*dM.X
		  sStart.Y + S*dS.Y = mStart.Y + M*dM.Y

		  Most of the complexity arises from checking for zeros in denominators
		  and using a different equation if they are found.
	*/
	dM := mEnd.Subtract(mStart)
	dS := sEnd.Subtract(sStart)
	var S, M float64
	if dM.Y == 0 {
		if dS.Y == 0 {
			if mStart.Y != sStart.Y || dS.X == dM.X {
				return math.NaN()
			}
			S = (mStart.X - sStart.X) / (dS.X - dM.X)
			M = S
		} else {
			if dM.X == 0 {
				if mStart.X == sStart.X && mStart.Y == sStart.Y {
					return 0
				}
				return math.NaN()
			}
			if dS.Y == 0 {
				return math.NaN()
			}
			S = (mStart.Y - sStart.Y) / dS.Y
			M = (sStart.X + S*dS.X - mStart.X) / dM.X
		}
	} else {
		if dS.Y == 0 {
			M = (sStart.Y - mStart.Y) / dM.Y
			S = (mStart.X + M*dM.X - sStart.X) / dS.X
		} else {
			if dM.X/dM.Y == dS.X/dS.Y {
				//TODO slopes are parallel check which end it hits first
				return math.NaN() //this isn't right, but it prevent an error and it's an edgecase
			}
			S = ((dM.X/dM.Y)*(sStart.Y-mStart.Y) - sStart.X + mStart.X) / (dS.X - (dM.X * dS.Y / dM.Y))
			M = (sStart.Y + S*dS.Y - mStart.Y) / dM.Y
		}
	}
	if S >= 0 && S <= 1 && M >= 0 && M <= 1 {
		return M
	}
	return math.NaN()
}

// Triangulate returns a point that is equadistant from all 3 points
func Triangulate(a, b, c F) F {
	abd := b.Y - a.Y
	acd := c.Y - a.Y

	if abd == 0 {
		x := (b.X + a.X) / 2
		if acd == 0 {
			if (c.X+a.X)/2 == x {
				// at least 2 points are identical
				return F{x, 0}
			}
			// 3 unique points on a line, no equadistant point exists
			return F{math.NaN(), math.NaN()}
		}
		acm := (a.X - c.X) / acd
		acb := (c.X*c.X - a.X*a.X + c.Y*c.Y - a.Y*a.Y) / (2 * acd)
		y := acm*x + acb
		return F{x, y}
	}

	if acd == 0 {
		x := (c.X + a.X) / 2
		abm := (a.X - b.X) / abd
		abb := (b.X*b.X - a.X*a.X + b.Y*b.Y - a.Y*a.Y) / (2 * abd)
		y := abm*x + abb
		return F{x, y}
	}

	abm := (a.X - b.X) / abd
	abb := (b.X*b.X - a.X*a.X + b.Y*b.Y - a.Y*a.Y) / (2 * abd)

	acm := (a.X - c.X) / acd
	acb := (c.X*c.X - a.X*a.X + c.Y*c.Y - a.Y*a.Y) / (2 * acd)

	d := (abm - acm)
	if d == 0 {
		// 3 unique points on a line, no equadistant point exists
		return F{math.NaN(), math.NaN()}
	}
	x := (acb - abb) / d
	y := abm*x + abb
	return F{x, y}
}
