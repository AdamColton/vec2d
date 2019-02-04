package vec2d

import (
	"math"
)

// EllipseArc fulfills Path and describes an elliptic arc. Start defines where the
// arc starts and Length defines the arc length, in radians. Start defaults to 0
// which is the point on the arc that intersects the ray from foci-1 to foci-2
type EllipseArc struct {
	Start, Length float64
	c             F       // center
	sMa, sma      float64 // semi-major axis, semi-minor axis
	a, as, ac     float64 // angle and it's sin, cos
}

// F returns the float64 vector at t. Fulfills Path.
func (e EllipseArc) F(t float64) F {
	return e.ByAngle((e.Length)*t + e.Start).Add(e.c)
}

// Tangent returns a tangent line at t. Fulfills Path.
func (e EllipseArc) Tangent(t float64) F {
	t = (e.Length)*t + e.Start
	return e.ByAngle(t + math.Pi/2)
}

// ByAngle returns the vector at the given angle relative to the center
func (e EllipseArc) ByAngle(a float64) F {
	// https://en.wikipedia.org/wiki/Parametric_equation#Ellipse
	st, ct := math.Sincos(a)
	return F{
		X: e.sMa*ct*e.ac - e.sma*st*e.as,
		Y: e.sMa*ct*e.as + e.sma*st*e.ac,
	}
}

// Foci of the ellipse
func (e EllipseArc) Foci() (F, F) {
	fociLen := math.Sqrt(e.sMa*e.sMa - e.sma*e.sma)
	d := F{fociLen * e.ac, fociLen * e.as}
	return e.c.Subtract(d), e.c.Add(d)
}

// Center of the ellipse
func (e EllipseArc) Center() F {
	return e.c
}

// Axis returns the lengths of the major and minor axis of the ellipse
func (e EllipseArc) Axis() (major, minor float64) {
	return e.sMa, e.sma
}

// NewEllipseArc returns an EllipseArc with foci f1 and f2 and a minor radius of
// r. The perimeter point that corresponds to an angle of 0 will be 1/4 rotation
// going from f1 to f2, which will lie along the minor axis. So an ellipse with
// foci (0,0) and (0,2) with a minor radius of 1 will have angle 0 at point
// (1,1).
func NewEllipseArc(f1, f2 F, r float64) EllipseArc {
	e := EllipseArc{
		c:      f1.Add(f2).ScalarMultiply(0.5),
		Length: math.Pi * 2,
	}
	d := f2.Subtract(f1)
	e.a = d.Angle()
	e.as, e.ac = math.Sincos(e.a)

	e.sma = r
	e.sMa = F{d.Mag(), 2 * r}.Mag() / 2

	return e
}

// Ellipse fulfills Shape
type Ellipse struct {
	perimeter EllipseArc
}

// NewEllipse returns an Ellipse with foci f1 and f2 and a minor radius of r.
// The perimeter point that corresponds to an angle of 0 will be 1/4 rotation
// going from f1 to f2, which will lie along the minor axis. So an ellipse with
// foci (0,0) and (0,2) with a minor radius of 1 will have angle 0 at point
// (1,1).
func NewEllipse(f1, f2 F, r float64) Ellipse {
	return Ellipse{
		perimeter: NewEllipseArc(f1, f2, r),
	}
}

// FillCurve returns a curve that lies inside of the ellipse. All curves in the
// range 0.0 to 1.0 will fill the entire ellipse.
func (e Ellipse) FillCurve(t float64) Curve {
	f0, f1 := e.perimeter.Foci()
	tFrom := Triangle{
		e.perimeter.F(1.0 / 8.0),
		e.perimeter.F(3.0 / 8.0),
		f0.Bisect(f1)(0.5),
	}

	//0 ==> 1/8  0.5 ==> 0  1 ==> -1/8
	t = t*-0.25 + 0.125

	tTo := Triangle{
		e.perimeter.F(t),
		e.perimeter.F(0.5 - t),
		f0.Bisect(f1)(t * 4),
	}
	tfrm, _ := TriangleTransform(tFrom, tTo)

	return func(t float64) F {
		t = (t / 4.0) + (1.0 / 8.0)
		return tfrm.Apply(e.perimeter.F(t))
	}
}

// F fulfils Surface taking two parametric points and returning a point in the
// ellipse.
func (e Ellipse) F(t0, t1 float64) F {
	return e.FillCurve(t0)(t1)
}

// Area returns the area of the Ellipse
func (e Ellipse) Area() float64 {
	a := e.SignedArea()
	if a < 0 {
		a = -a
	}
	return a
}

// SignedArea returns the area of the ellipse, though the value may be negative
// depending on polarity.
func (e Ellipse) SignedArea() float64 {
	return e.perimeter.sMa * e.perimeter.sma * math.Pi
}

// Perimeter returns the length of the perimeter of the ellipse. This value may
// have a slight error that will grow as the ellipse is elongated.
func (e Ellipse) Perimeter() float64 {
	d, s := (e.perimeter.sMa - e.perimeter.sma), (e.perimeter.sMa + e.perimeter.sma)
	h := (d * d) / (s * s)
	p := 1 + ((3 * h) / (10 + math.Sqrt(4-3*h)))
	p *= math.Pi * s
	return p
}

// Centroid returns the center of the ellipse
func (e Ellipse) Centroid() F {
	return e.perimeter.c
}

// Contains returns true of the point f is contained in the ellipse
func (e Ellipse) Contains(f F) bool {
	a := f.Subtract(e.perimeter.c).Angle() - e.perimeter.a
	p := e.perimeter.ByAngle(a).Add(e.perimeter.c)

	return e.perimeter.c.Distance(f) < e.perimeter.c.Distance(p)
}

// Arc returns the EllipseArc defined by the perimeter
func (e Ellipse) Arc() EllipseArc {
	return e.perimeter
}

// Circle fulfills Shape.
type Circle struct {
	e Ellipse
}

// NewCircle returns a circle defined by a center and radius
func NewCircle(center F, radius float64) Circle {
	return Circle{e: NewEllipse(center, center, radius)}
}

// Radius returns the radius of the circle
func (c Circle) Radius() float64 {
	return c.e.perimeter.sma
}

// FillCurve returns a curve that lies inside of the circle. All curves in the
// range 0.0 to 1.0 will fill the entire circle.
func (c Circle) FillCurve(t float64) Curve { return c.e.FillCurve(t) }

// F returns the float64 vector at t. Fulfills Path.
func (c Circle) F(t0, t1 float64) F { return c.e.F(t0, t1) }

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return c.e.perimeter.sma * c.e.perimeter.sma * math.Pi
}

// SignedArea returns the area of the ellipse, though the value may be negative
// depending on polarity.
func (c Circle) SignedArea() float64 { return c.e.SignedArea() }

// Perimeter length of the circle
func (c Circle) Perimeter() float64 {
	return 2 * c.e.perimeter.sma * math.Pi
}

// Contains returns true of the point f is contained in the ellipse
func (c Circle) Contains(f F) bool {
	return c.e.perimeter.c.Distance(f) <= c.e.perimeter.sma
}

// Centroid returns the center of the circle
func (c Circle) Centroid() F { return c.e.perimeter.c }

// Arc returns the EllipseArc that represent the perimeter of the circle
func (c Circle) Arc() EllipseArc { return c.e.perimeter }
