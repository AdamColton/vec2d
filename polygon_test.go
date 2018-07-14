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
