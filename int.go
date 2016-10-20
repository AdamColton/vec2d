package vec2d

import (
	"fmt"
	"math"
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

// F converts an int vector to a float64 vector
func (i I) F() F {
	return F{float64(i.X), float64(i.Y)}
}

// Angle returns the angle in radians
func (i I) Angle() float64 {
	return math.Atan2(float64(i.Y), float64(i.X))
}

// Returns the magnitude (distance to origin) of the vector
func (i I) Mag() float64 {
	return i.F().Mag()
}

// Distance returns the distance between to points
func (i I) Distance(i2 I) float64 {
	return i.Subtract(i2).Mag()
}

// String fulfills Stringer, returns the vector as "(X, Y)"
func (i I) String() string {
	return fmt.Sprintf("(%d, %d)", i.X, i.Y)
}

// To returns a channel that will iterate over every point from a (incluse) to b
// (exclusive). This opens a Go routine, so be sure to read from the channel
// until it is closed.
func (i I) To(i2 I) <-chan I {
	ch := make(chan I)
	dx, dy := 1, 1
	if i2.X < i.X {
		dx = -1
	}
	if i2.Y < i.Y {
		dy = -1
	}
	go ranger(i, i2, dx, dy, ch)
	return ch
}

func ranger(a, b I, dx, dy int, ch chan<- I) {
	for x := a.X; x < b.X; x += dx {
		for y := a.Y; y < b.Y; y += dy {
			ch <- I{x, y}
		}
	}
	close(ch)
}

func (i I) FromOrigin() <-chan I {
	return I{0, 0}.To(i)
}

// SliceTo returns []I over every point from a (incluse) to b (exclusive).
func (i I) SliceTo(i2 I) []I {
	d := i.Subtract(i2)
	l := make([]I, 0, d.X*d.Y)
	for p := range i.To(i2) {
		l = append(l, p)
	}
	return l
}
