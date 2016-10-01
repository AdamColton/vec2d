package vec2d

import (
	"fmt"
)

// F2d is a 2D int vector. It can be used to represent a point or the difference
// between two points.
type I2d struct {
	X, Y int
}

// Returns a + b
func (a I2d) Add(b I2d) I2d {
	return I2d{a.X + b.X, a.Y + b.Y}
}

// Returns a - b
func (a I2d) Subtract(b I2d) I2d {
	return I2d{a.X - b.X, a.Y - b.Y}
}

// Takes the Abs of X and Y
func (v I2d) Abs() I2d {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	return v
}

// Converts an int vector to a float64 vector
func (v I2d) F2d() F2d {
	return F2d{float64(v.X), float64(v.Y)}
}

// Returns the angle in radians
func (v I2d) Angle() float64 {
	return v.F2d().Angle()
}

// Returns the magnitude (distance to origin) of the vector
func (v I2d) Mag() float64 {
	return v.F2d().Mag()
}

// Returns the distance between to points
func (a I2d) Distance(b I2d) float64 {
	return a.Subtract(b).Mag()
}

// Fulfills Stringer, returns the vector as "(X, Y)"
func (v I2d) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

// Returns a channel that will iterate over every point from a (incluse) to b
// (exclusive). This opens a Go routine, so be sure to read from the channel
// until it is closed.
func (a I2d) To(b I2d) <-chan I2d {
	ch := make(chan I2d)
	if b.X < a.X {
		a.X, b.X = b.X, a.X
	}
	if b.Y < a.Y {
		a.Y, b.Y = b.Y, a.Y
	}
	go ranger(a, b, ch)
	return ch
}

func ranger(a, b I2d, ch chan<- I2d) {
	for x := a.X; x < b.X; x++ {
		for y := a.Y; y < b.Y; y++ {
			ch <- I2d{x, y}
		}
	}
	close(ch)
}

// Returns []I2d over every point from a (incluse) to b (exclusive).
func (a I2d) SliceTo(b I2d) []I2d {
	d := a.Subtract(b)
	l := make([]I2d, 0, d.X*d.Y)
	for p := range a.To(b) {
		l = append(l, p)
	}
	return l
}
