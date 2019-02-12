package vec2d

import (
	"math"
)

// Triangle represented by 3 Float64 vectors
type Triangle [3]F

// Contains returns true if point f is inside the triangle
func (t Triangle) Contains(f F) bool {
	// If a point is inside the triangle, the sign of the cross product from the
	// point to each vertex will be the same. But a cross product of exactly 0
	// doesn't have a sign.
	s1 := t[0].Subtract(t[1])
	r1 := f.Subtract(t[1])
	s2 := t[1].Subtract(t[2])
	r2 := f.Subtract(t[2])

	c1 := s1.Cross(r1)
	c2 := s2.Cross(r2)
	if !(c1 >= 0 && c2 >= 0) && !(c1 <= 0 && c2 <= 0) {
		return false
	}

	s3 := t[2].Subtract(t[0])
	r3 := f.Subtract(t[0])
	c3 := s3.Cross(r3)
	return (c2 >= 0 && c3 >= 0) || (c2 <= 0 && c3 <= 0)
}

// SignedArea of the triangle
func (t Triangle) SignedArea() float64 {
	v1 := t[0].Subtract(t[1])
	v2 := t[0].Subtract(t[2])
	return 0.5 * v1.Cross(v2)
}

// Area of the triangle
func (t Triangle) Area() float64 {
	return math.Abs(t.SignedArea())
}

// Perimeter of the triangle
func (t Triangle) Perimeter() float64 {
	return t[0].Distance(t[1]) + t[1].Distance(t[2]) + t[2].Distance(t[0])
}

// Centroid returns the center of mass of the triangle
func (t Triangle) Centroid() F {
	// The center of a triangle is found by drawing a line from one corner to the
	// center of the opposite side and going 2/3 of the way.
	midpoint := t[0].Add(t[1]).ScalarMultiply(0.5)
	return t[2].LineTo(midpoint)(2.0 / 3.0)
}

// F returns a parametric point inside the triangle.
func (t Triangle) F(t0, t1 float64) F {
	m := t[0].LineTo(t[1])(0.5)
	p0 := t[0].LineTo(m)(t0)
	p1 := t[2].LineTo(t[1])(t0)
	return p0.LineTo(p1)(t1)
}

// CircumscribedCircle returns a circle whose perimeter touches all the vertexes
// of the triangle
func (t Triangle) CircumscribedCircle() Circle {
	c := Triangulate(t[0], t[1], t[2])
	r := t[0].Distance(c)
	return NewCircle(c, r)
}

// InscribeCircle returns a circlue whose perimeter touches each side of the
// triangle in exactly one place
func (t Triangle) InscribeCircle() Circle {
	ab := t[1].Subtract(t[0]).P()
	ac := t[2].Subtract(t[0]).P()
	p1 := P{100, (ab.A + ac.A) / 2}
	l1 := t[0].LineTo(t[0].Add(p1.F()))

	bc := t[2].Subtract(t[1]).P()
	ba := t[0].Subtract(t[1]).P()
	p2 := P{100, (bc.A + ba.A) / 2}
	l2 := t[1].LineTo(t[1].Add(p2.F()))

	l3 := t[1].LineTo(t[2])

	tc, _ := l1.Intersection(l2)
	c := l1(tc)
	tr, _ := l1.Intersection(l3)
	r := c.Distance(l1(tr))
	return NewCircle(c, r)
}
