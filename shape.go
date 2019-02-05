package vec2d

// Surface functions describe 2D area parametrically. They should have the
// following properties
// * All points on the perimeter should have either t0==0 or t1==0
// * The surface should have no creases
// * F(ta0,ta1)==F(tb0,tb1) --> ta0==tb0 && ta1==tb1
type Surface func(t0, t1 float64) F

// Surfacer is object with a 2D parametric method F
type Surfacer interface {
	F(t0, t1 float64) F
}

// Shape is a 2D parametric surface that can describe additional information
// about the surface.
type Shape interface {
	Surfacer
	Area() float64
	SignedArea() float64
	Perimeter() float64
	Contains(F) bool
	Centroid() F
}
