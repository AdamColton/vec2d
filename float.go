package vec2d

import (
	"fmt"
	"math"
)

// F2d is a 2D float64 vector. It can be used to represent a point or the
// difference between two points.
type F2d struct {
	X, Y float64
}

// Returns a + b
func (a F2d) Add(b F2d) F2d {
	a.X += b.X
	a.Y += b.Y
	return a
}

// Returns a - b
func (a F2d) Subtract(b F2d) F2d {
	a.X -= b.X
	a.Y -= b.Y
	return a
}

// Returns FV{a.X * b.X, a.Y * b.Y}
func (a F2d) Multiply(b F2d) F2d {
	a.X *= b.X
	a.Y *= b.Y
	return a
}

func (a F2d) ScalarMultiply(sclr float64) F2d {
	a.X *= sclr
	a.Y *= sclr
	return a
}

// Returns the angle in radians
func (v F2d) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Returns the magnitude (distance to origin) of the vector
func (v F2d) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Returns a new Vector rotated by r radians
func (v F2d) Rotate(r float64) F2d {
	m := v.Mag()
	r += v.Angle()
	return F2d{math.Cos(r) * m, math.Sin(r) * m}
}

// Returns the distance between to points
func (a F2d) Distance(b F2d) float64 {
	return a.Subtract(b).Mag()
}

// Converts a float64 vector to an int vector. Will always round down.
func (v F2d) I2d() I2d {
	return I2d{int(v.X), int(v.Y)}
}

// Converts a float64 vector to a Polar vector
func (v F2d) P() P {
	return P{v.Mag(), v.Angle()}
}

// Fulfills Stringer, returns the vector as "(X, Y)"
func (v F2d) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

// Takes the motion path described by mStart t(0) and mEnd t(1) and finds the
// time that the motion intersects the line segment described by sStart to sEnd.
// If there is an intersection and it happens between t(0) and t(1), the value
// will be returned, otherwise NaN will be returned
func MotionSurfaceIntersection(mStart, mEnd, sStart, sEnd F2d) float64 {
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
				} else {
					return math.NaN()
				}
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

// Returns a point that is equadistant from all 3 points
func Triangulate(a, b, c F2d) F2d {
	abd := b.Y - a.Y
	acd := c.Y - a.Y

	if abd == 0 {
		x := (b.X + a.X) / 2
		if acd == 0 {
			if (c.X+a.X)/2 == x {
				// at least 2 points are identical
				return F2d{x, 0}
			} else {
				// 3 unique points on a line, no equadistant point exists
				return F2d{math.NaN(), math.NaN()}
			}
		}
		acm := (a.X - c.X) / acd
		acb := (c.X*c.X - a.X*a.X + c.Y*c.Y - a.Y*a.Y) / (2 * acd)
		y := acm*x + acb
		return F2d{x, y}
	}

	if acd == 0 {
		x := (c.X + a.X) / 2
		abm := (a.X - b.X) / abd
		abb := (b.X*b.X - a.X*a.X + b.Y*b.Y - a.Y*a.Y) / (2 * abd)
		y := abm*x + abb
		return F2d{x, y}
	}

	abm := (a.X - b.X) / abd
	abb := (b.X*b.X - a.X*a.X + b.Y*b.Y - a.Y*a.Y) / (2 * abd)

	acm := (a.X - c.X) / acd
	acb := (c.X*c.X - a.X*a.X + c.Y*c.Y - a.Y*a.Y) / (2 * acd)

	d := (abm - acm)
	if d == 0 {
		// 3 unique points on a line, no equadistant point exists
		return F2d{math.NaN(), math.NaN()}
	}
	x := (acb - abb) / d
	y := abm*x + abb
	return F2d{x, y}
}
