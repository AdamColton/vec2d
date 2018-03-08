package vec2d

// Curve is a parametric function that takes a single float value and returns
// a 2D-float64 point
type Curve func(t float64) F

// Tangent takes a parametric point and returns the tangent line. For now, all
// that matters is that the point is on the line and the slope is tangent, but
// later there may be more requirements that t=0 and t=1 have meaning related to
// the second derivative.
type Tangent func(t float64) Line

// A Path is a curve that can also return tangent lines
type Path interface {
	F(t float64) F
	Tangent(t float64) Line
}
