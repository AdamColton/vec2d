package vec2d

import (
	"math"
	"strconv"
	"strings"
)

// I is a 2D int vector. It can be used to represent a point or the difference
// between two points.
type I struct {
	X, Y int
}

// Add returns i + i2
func (i I) Add(i2 I) I {
	return I{i.X + i2.X, i.Y + i2.Y}
}

// Subtract returns i - i2
func (i I) Subtract(i2 I) I {
	return I{i.X - i2.X, i.Y - i2.Y}
}

// Area returns i.X*i.Y
func (i I) Area() int {
	return i.X * i.Y
}

// Abs takes the Abs of X and Y
func (i I) Abs() I {
	if i.X < 0 {
		i.X = -i.X
	}
	if i.Y < 0 {
		i.Y = -i.Y
	}
	return i
}

// Idx returns the index of a grid represented as a slice
func (i I) Idx(i2 I) int {
	return i2.Y*i.X + i2.X
}

// InvIdx returns the point that corresponds to an index
func (i I) InvIdx(idx int) I {
	x := idx % i.X
	idx -= x
	return I{x, idx / i.X}
}

// In takes I.In(A, B) and returns true if I is between A (inclusive) and B
// (exclusive)
func (i I) In(a, b I) bool {
	if a.X > b.X {
		if i.X <= b.X || i.X > a.X {
			return false
		}
	} else {
		if i.X >= b.X || i.X < a.X {
			return false
		}
	}
	if a.Y > b.Y {
		if i.Y <= b.Y || i.Y > a.Y {
			return false
		}
	} else {
		if i.Y >= b.Y || i.Y < a.Y {
			return false
		}
	}
	return true
}

// Mod performs a modulus operation so that the returned point is between the
// origin and i.
func (i I) Mod(i2 I) I {
	i = I{i.X % i2.X, i.Y % i2.Y}
	if i.X < 0 {
		i.X += i2.X
	}
	if i.Y < 0 {
		i.Y += i2.Y
	}
	return i
}

// F converts an int vector to a float64 vector
func (i I) F() F {
	return F{float64(i.X), float64(i.Y)}
}

// Angle returns the angle in radians
func (i I) Angle() float64 {
	return math.Atan2(float64(i.Y), float64(i.X))
}

// Mag returns the magnitude (distance to origin) of the vector
func (i I) Mag() float64 {
	return i.F().Mag()
}

// ScalarMultiply returns I{i.X*sclr, i.Y*sclr}
func (i I) ScalarMultiply(sclr int) I {
	i.X *= sclr
	i.Y *= sclr
	return i
}

// Distance returns the distance between to points
func (i I) Distance(i2 I) float64 {
	return i.Subtract(i2).Mag()
}

// String fulfills Stringer, returns the vector as "(X, Y)"
func (i I) String() string {
	return strings.Join([]string{
		"(",
		strconv.FormatInt(int64(i.X), 10),
		", ",
		strconv.FormatInt(int64(i.Y), 10),
		")",
	}, "")
}

// SliceTo returns []I over every point from a (inclusive) to b (exclusive).
func (i I) SliceTo(i2 I) []I {
	d := i.Subtract(i2)
	l := make([]I, 0, d.X*d.Y)
	dx, dy := 1, 1
	cx, cy := lt, lt
	if i.X > i2.X {
		dx = -1
		cx = gt
	}
	if i.Y > i2.Y {
		dy = -1
		cy = gt
	}
	for y := i.Y; cy(y, i2.Y); y += dy {
		for x := i.X; cx(x, i2.X); x += dx {
			l = append(l, I{x, y})
		}
	}
	return l
}

// IntIterator provides an interface for iterating over point ranges
type IntIterator interface {
	Next() (I, bool)
	Idx() int
}

// XIterator is return from .To or .FromOrigin, it iterators by incrementing X
// first.
type XIterator struct {
	to, cur        I
	x, dx, dy      int
	checkX, checkY func(int, int) bool
	idx            int
}

// To is used to iterate over points. Given A.To(B) it will hit all the points
// in the rectangle between inclusive of A and exclusive of B. The iterator will
// move across a full row before moving to the next column.
func (i I) To(i2 I) (IntIterator, I, bool) {
	r := &XIterator{
		to:     i2,
		cur:    i,
		x:      i.X,
		dx:     1,
		dy:     1,
		checkX: gte,
		checkY: gte,
	}
	if i.X > i2.X {
		r.dx = -1
		r.checkX = lte
	}
	if i.Y > i2.Y {
		r.dy = -1
		r.checkY = lte
	}
	return r, r.cur, true
}

func lte(a, b int) bool {
	return a <= b
}

func gte(a, b int) bool {
	return a >= b
}

func lt(a, b int) bool {
	return a < b
}

func gt(a, b int) bool {
	return a > b
}

// Next fulfills the IntIterator interface. It returns the next point and if
// iteration is done.
func (xi *XIterator) Next() (I, bool) {
	xi.idx++
	xi.cur.X += xi.dx
	if xi.checkX(xi.cur.X, xi.to.X) {
		xi.cur.X = xi.x
		xi.cur.Y += xi.dy
		if xi.checkY(xi.cur.Y, xi.to.Y) {
			return xi.cur, false
		}
	}
	return xi.cur, true
}

// Idx fulfills the IntIterator interface. It returns the index of the current
// point.
func (xi *XIterator) Idx() int { return xi.idx }

// FromOrigin returns an iterator from the origin (incluse) to i (exclusive).
func (i I) FromOrigin() (IntIterator, I, bool) {
	return I{0, 0}.To(i)
}
