package vec2d

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestEllipseArc(t *testing.T) {
	f1 := F{0, 0}
	f2 := F{2, 0}
	r := 1.0

	e := NewEllipseArc(f1, f2, r)
	assert.Equal(t, r, e.F(0.25).Distance(e.c))
	assert.InDelta(t, 0.0, e.F(0).Y, 1E-10)

	// Test by definition of ellipse
	p0 := e.F(0)
	d0 := f1.Distance(p0) + f2.Distance(p0)
	for i := 0.0; i <= 1.0; i += 0.2 {
		p := e.F(i)
		d := f1.Distance(p) + f2.Distance(p)
		assert.InDelta(t, d0, d, 1E-10)
	}

	// Get correct foci
	tf1, tf2 := e.Foci()
	assert.InDelta(t, 0, f1.Distance(tf1), 1E-10)
	assert.InDelta(t, 0, f2.Distance(tf2), 1E-10)

	var p Path
	p = e
	assert.NotNil(t, p)
}

func TestEllipse(t *testing.T) {
	f1 := F{0, 0}
	f2 := F{2, 0}
	r := 1.0

	e := NewEllipse(f1, f2, r)
	expected := F{1, 0}
	actual := e.F(0.5, 0.5)
	assert.InDelta(t, 0, expected.Distance(actual), 1E-10)

	assert.Equal(t, math.Sqrt2*math.Pi, e.Area())

	e = NewEllipse(f1, f1, r)
	assert.Equal(t, 2*math.Pi*r, e.Perimeter())

	var s Shape
	s = e
	assert.NotNil(t, s)
}

func TestCircleArea(t *testing.T) {
	c := NewCircle(F{0, 0}, -1/math.Sqrt(math.Pi))
	assert.InDelta(t, 1, c.Area(), 1E-10)
	assert.InDelta(t, -1, c.SignedArea(), 1E-10)
}

func TestEllipseStandard(t *testing.T) {
	// Make sure the ellipse follows standards of Polar plane
	f1 := F{-1, 0}
	f2 := F{1, 0}
	r := 1.0
	e := NewEllipseArc(f1, f2, r)

	// An angle of 0 should be in the +X direction
	assert.Equal(t, F{math.Sqrt2, 0}, e.F(0))

	// 1/4 rotation should be +Y
	assert.InDelta(t, 0.0, e.F(0.25).Distance(F{0, 1}), 1E-10)
}
