package vec2d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPolygonSignedArea(t *testing.T) {
	tri := Triangle{
		F{0, 0},
		F{1, 0},
		F{0, 1},
	}
	p := Polygon(tri[:])
	assert.InDelta(t, tri.SignedArea(), p.SignedArea(), 1e-10)

	square := Polygon{
		F{0, 0},
		F{1, 0},
		F{1, 1},
		F{0, 1},
	}
	assert.InDelta(t, 1.0, square.SignedArea(), 1e-10)
}

func TestPolygonCentroid(t *testing.T) {
	square := Polygon{
		F{0, 0},
		F{1, 0},
		F{1, 1},
		F{0, 1},
	}
	assert.Equal(t, F{0.5, 0.5}, square.Centroid())
}

func TestPolygonContains(t *testing.T) {
	p := Polygon{
		F{0, 0},
		F{2, 2},
		F{1, 0},
		F{2, -2},
	}
	assert.True(t, p.Contains(F{0.5, 0}))
	assert.False(t, p.Contains(F{0.5, 1}))
	assert.True(t, p.Contains(F{1, 1}))
	assert.False(t, p.Contains(F{2, 0}))
}

func TestPolygonSurface(t *testing.T) {
	ti := I{20, 20}
	outside := NewCircle(F{}, 2).Arc()
	for sides := 3; sides < 7; sides++ {
		p := RegularPolygonSideLength(F{}, 1, 0, sides)
		// Note that we're iterating over interior points only, not perimeter points
		// where t0 or t1 = 0.0 or 1.0.
		for iter, i, ok := ti.FromOrigin().Start(); ok; i, ok = iter.Next() {
			tf := i.F().ScalarMultiply(.05).Add(F{.025, .025})
			f := p.F(tf.X, tf.Y)
			assert.True(t, p.Contains(f))
		}
		for t0 := 0.0; t0 < 1.0; t0 += 0.05 {
			f := outside.F(t0)
			assert.False(t, p.Contains(f))
		}
	}
}

func TestUnitSquareSurface(t *testing.T) {
	// The surface function of the unit square should map to itself
	unitSquare := Polygon{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	i := I{101, 101}
	i.FromOrigin().Each(func(idx int, i I) {
		f := i.F().ScalarMultiply(0.01)
		assert.InDelta(t, 0, f.Distance(unitSquare.F(f.X, f.Y)), 0.00001)
	})
}

func TestRegularPolygonSideLength(t *testing.T) {
	actual := RegularPolygonSideLength(F{}, 1, 0, 4)
	expected := Polygon{F{0.5000, -0.5000}, {0.5000, 0.5000}, {-0.5000, 0.5000}, {-0.5000, -0.5000}}
	for i, p := range expected {
		assert.InDelta(t, 0, p.Distance(actual[i]), 1E-10)
	}
}

func TestRegularPolygonRadius(t *testing.T) {
	actual := RegularPolygonRadius(F{}, 1, 0, 4)
	expected := Polygon{F{1.0000, 0.0000}, {0.0000, 1.0000}, {-1.0000, 0.0000}, {-0.0000, -1.0000}}
	for i, p := range expected {
		assert.InDelta(t, 0, p.Distance(actual[i]), 1E-10)
	}
}

func TestCountAngles(t *testing.T) {
	p := Polygon{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	ccw, cw := p.CountAngles()
	assert.Equal(t, 4, ccw)
	assert.Equal(t, 0, cw)
}

func TestConvex(t *testing.T) {
	tt := []struct {
		Polygon Polygon
		Convex  bool
	}{
		{
			Polygon: Polygon{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			Convex:  true,
		},
		{
			Polygon: Polygon{{0, 0}, {0, 1}, {1, 1}, {1, 0}},
			Convex:  true,
		},
		{
			Polygon: Polygon{{0, 1}, {2, 2}, {1, 1}, {2, 0}},
			Convex:  false,
		},
		{
			Polygon: Polygon{{0, 0}, {1, 1}, {0, 2}, {2, 1}},
			Convex:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Polygon.String(), func(t *testing.T) {
			assert.Equal(t, tc.Convex, tc.Polygon.Convex())
		})
	}
}

func TestNonIntersecting(t *testing.T) {
	tt := []struct {
		Polygon         Polygon
		NonIntersecting bool
	}{
		{
			Polygon:         Polygon{{0, 0}, {1, 0}, {1, 1}},
			NonIntersecting: true,
		},
		{
			Polygon:         Polygon{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			NonIntersecting: true,
		},
		{
			Polygon:         Polygon{{0, 0}, {1, 1}, {1, 0}, {0, 1}},
			NonIntersecting: false,
		},
		{
			Polygon:         Polygon{{0, 1}, {0, 2}, {1, 0}, {2, 2}, {2, 1}},
			NonIntersecting: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Polygon.String(), func(t *testing.T) {
			assert.Equal(t, tc.NonIntersecting, tc.Polygon.NonIntersecting())
		})
	}
}

func TestReverse(t *testing.T) {
	tt := []struct {
		p        Polygon
		expected Polygon
	}{
		{
			p:        Polygon{{0, 0}, {1, 0}, {1, 1}},
			expected: Polygon{{1, 1}, {1, 0}, {0, 0}},
		},
		{
			p:        Polygon{{0, 0}, {1, 0}, {1, 1}, {0, 1}},
			expected: Polygon{{0, 1}, {1, 1}, {1, 0}, {0, 0}},
		},
	}

	for _, tc := range tt {
		t.Run(tc.p.String(), func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.p.Reverse())
		})
	}
}

func TestFindTriangles(t *testing.T) {
	p := Polygon{{0, 0}, {2, 1}, {0, 2}, {1, 1}}
	expected := [][3]int{
		{1, 2, 3},
		{0, 1, 3},
	}
	assert.Equal(t, expected, p.FindTriangles())
}
