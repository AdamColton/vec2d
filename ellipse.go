package vec2d

import (
	"math"
)

// Ellipse fulfills Path and describes an elliptic arc. Start defines where the
// arc starts and Arc defines the arc length, in radians. Start defaults to 0
// which is the point on the arc that intersects the ray from foci-1 to foci-2
type Ellipse struct {
	Start, Arc float64
	c          F       // center
	sMa, sma   float64 // semi-major axis, semi-minor axis
	a, as, ac  float64 // angle and it's sin, cos
}

// F returns the float64 vector at t. Fulfills Path.
func (e Ellipse) F(t float64) F {
	return e.ByAngle((e.Arc)*t + e.Start).Add(e.c)
}

// Tangent returns a tangent line at t. Fulfills Path.
func (e Ellipse) Tangent(t float64) Line {
	t = (e.Arc)*t + e.Start

	p0 := e.ByAngle(t).Add(e.c)
	p1 := e.ByAngle(t + math.Pi/2).Add(p0)

	return p0.LineTo(p1)
}

// ByAngle returns the vector at the given angle relative to the center
func (e Ellipse) ByAngle(a float64) F {
	// https://en.wikipedia.org/wiki/Parametric_equation#Ellipse
	st, ct := math.Sincos(a)
	return F{
		X: e.sMa*ct*e.ac - e.sma*st*e.as,
		Y: e.sMa*ct*e.as + e.sma*st*e.ac,
	}
}

// Foci of the ellipse
func (e Ellipse) Foci() (F, F) {
	fociLen := math.Sqrt(e.sMa*e.sMa - e.sma*e.sma)
	d := F{fociLen * e.ac, fociLen * e.as}
	return e.c.Subtract(d), e.c.Add(d)
}

// Center of the ellipse
func (e Ellipse) Center() F {
	return e.c
}

// Axis returns the lengths of the major and minor axis of the ellipse
func (e Ellipse) Axis() (major, minor float64) {
	return e.sMa, e.sma
}

// NewEllipse returns an ellipse with foci f1 and f2 and a minor radius of r.
// The perimeter point that corresponds to an angle of 0 will be 1/4 rotation
// going from f1 to f2, which will lie along the minor axis. So an ellipse with
// foci (0,0) and (0,2) with a minor radius of 1 will have angle 0 at point
// (1,1).
func NewEllipse(f1, f2 F, r float64) Ellipse {
	e := Ellipse{
		c:   f1.Add(f2).ScalarMultiply(0.5),
		Arc: math.Pi * 2,
	}
	d := f2.Subtract(f1)
	e.a = d.Angle()
	e.as, e.ac = math.Sincos(e.a)

	e.sma = r
	e.sMa = F{d.Mag(), 2 * r}.Mag() / 2

	return e
}
