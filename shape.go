package vec2d

// Surface functions describe 2D area parametrically. They should be continuous
// and not contain any internal creases (such as mapping a circle radially
// would have).
type Surface func(t0, t1 float64) F

type Surfacer interface {
	F(t0, t1 float64) F
}

type Shape interface {
	F(t0, t1 float64) F
	Area() float64
	SignedArea() float64
	Perimeter() float64
	Contains(F) bool
	Centroid() F
}
