package vec2d

import (
	"math"
	"strconv"
	"strings"
)

type Ier interface {
	I() I
}

// I is a 2D int vector. It can be used to represent a point or the difference
// between two points.
type I struct {
	X, Y int
}

// I allow I to fulfill Ier
func (i I) I() I { return i }

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

func (i I) Multiply(i2 I) I {
	i.X *= i2.X
	i.Y *= i2.Y
	return i
}

func (i I) Divide(i2 I) I {
	i.X /= i2.X
	i.Y /= i2.Y
	return i
}

func (i I) Cross(i2 I) int {
	return i.X*i2.Y - i2.X*i.Y
}

func (i I) Dot(i2 I) int {
	return i.X*i2.X + i.Y*i2.Y
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

// Idx returns the index of a point on a grid. The grid is defined by i and the
// origin and the index is the point i2.
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
