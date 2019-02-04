package vec2d

// Curve is a parametric function that takes a single float value and returns
// a 2D-float64 point
type Curve func(t float64) F

// Curver is an object that has a method named F fulling the Curve interface
type Curver interface {
	F(t float64) F
}

// A Path is a curve that can also return tangent lines
type Path interface {
	F(t float64) F

	// Tangent takes a single parameter and returns a point that represents the
	// tangent to the Path at the same parameter.
	Tangent(t float64) F
}

// TangentLineFactory takes a Path and returns a tangent Line at point T.
func TangentLineFactory(p Path) func(t float64) Line {
	return func(t float64) Line {
		p0 := p.F(t)
		p1 := p.Tangent(t).Add(p0)
		return p0.LineTo(p1)
	}
}
