package vec2d

import (
	"math"
)

// NewBezierCurve returns a Bezier Curve defined by a list of control points
func NewBezierCurve(points ...F) Curve {
	l := len(points)
	l64 := float64(l - 1)
	return func(t float64) F {
		if t == 1 {
			return points[l-1]
		}
		if t == 0 {
			return points[0]
		}

		// B(t) = ∑ binomialCo(l,i) * (1-t)^(l-i) * t^(i) * points[i]
		// let s = (1-t)^(l-i) * t^(i)
		// then s[i] = s[i-1] * t/(1-t)
		// and s[0] = (1-t) ^ l

		ti := 1 - t
		s := math.Pow(ti, l64)
		sd := t / ti
		var pt F
		for i, p := range points {
			b := binomialCo(l64, float64(i))
			pt = pt.Add(p.ScalarMultiply(s * b))
			s *= sd
		}
		return pt
	}
}

// NewBezierTangent returns a function that returns the tangents of the Bezier
// Curve that would be described by the same points. The tangent curve is itself
// a Bezier Curve.
func NewBezierTangent(points ...F) Curve {
	qs := DiffPoints(points...)
	return NewBezierCurve(qs...)
}

// DiffPoints takes a list of points and returns the difference between each. So
// the first value returned will be points[0].Subtrace(points[1]). This also
// means the returned slice will be one less than the values given.
func DiffPoints(points ...F) []F {
	l := len(points) - 1
	dps := make([]F, l)
	prev := points[0]
	for i, p := range points[1:] {
		dps[i] = p.Subtract(prev)
		prev = p
	}
	return dps
}

// abuse F to compute binomialCo - this is a terrible idea, which is why it's
// not exposed
var binomialCoMemo = make(map[F]float64)

func binomialCo(n, k float64) float64 {
	b, ok := binomialCoMemo[F{n, k}]
	if ok {
		return b
	}

	// https://math.stackexchange.com/questions/202554/how-do-i-compute-binomial-coefficients-efficiently
	if k > n {
		b = math.NaN()
	} else if k == 0 {
		b = 1
	} else if k > n/2 {
		b = binomialCo(n, n-k)
	} else {
		b = n * binomialCo(n-1, k-1) / k
	}

	binomialCoMemo[F{n, k}] = b
	return b
}

// BezierPath fulfils Path for Bezier curves.
type BezierPath struct {
	ps      []F
	curve   Curve
	tangent Curve
}

// NewBezierPath creates a BezierPath from a list of points
func NewBezierPath(ps ...F) BezierPath {
	return BezierPath{
		ps:      ps,
		curve:   NewBezierCurve(ps...),
		tangent: NewBezierTangent(ps...),
	}
}

// F returns the float64 vector at t. Fulfills Path.
func (bp BezierPath) F(t float64) F {
	return bp.curve(t)
}

// Tangent does things and Fulfills Path.
func (bp BezierPath) Tangent(t float64) F {
	return bp.tangent(t)
}

// Points returns a copy of the points that define the curve.
func (bp BezierPath) Points() []F {
	cp := make([]F, len(bp.ps))
	copy(cp, bp.ps)
	return cp
}
