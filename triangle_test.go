package vec2d

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTriangleArea(t *testing.T) {
	p1 := F{0, 0}
	p2 := F{1, 0}
	p3 := F{0, 1}
	tri := Triangle{p1, p2, p3}
	assert.InDelta(t, 0.5, tri.SignedArea(), 1E-10)
	assert.InDelta(t, 0.5, tri.Area(), 1E-10)
	tri = Triangle{p1, p3, p2}
	assert.InDelta(t, -0.5, tri.SignedArea(), 1E-10)
	assert.InDelta(t, 0.5, tri.Area(), 1E-10)
}

func TestTriangleContains(t *testing.T) {
	p1 := F{0, 0}
	p2 := F{1, 0}
	p3 := F{0, 1}
	tri := Triangle{p1, p2, p3}
	assert.True(t, tri.Contains(F{0.1, 0.1}))
	assert.True(t, tri.Contains(F{0, 0}))
	assert.False(t, tri.Contains(F{1, 1}))
}

func TestTrianglePerimeter(t *testing.T) {
	p1 := F{0, 0}
	p2 := F{1, 0}
	p3 := F{0, 1}
	tri := Triangle{p1, p2, p3}
	assert.Equal(t, 2+math.Sqrt2, tri.Perimeter())
}

func TestTriangleSurface(t *testing.T) {
	p1 := F{0, 0}
	p2 := F{1, 0}
	p3 := F{0, 1}
	tri := Triangle{p1, p2, p3}
	assert.Equal(t, p1, tri.F(0, 0))

	var s Shape
	s = tri
	assert.NotNil(t, s)
}

func TestCircumscribedCircle(t *testing.T) {
	p1 := F{0, 0}
	p2 := F{1, 0}
	p3 := F{0, 1}
	tri := Triangle{p1, p2, p3}
	c := tri.CircumscribedCircle()
	assert.InDelta(t, c.Radius(), c.Centroid().Distance(p1), 1E-10)
	assert.InDelta(t, c.Radius(), c.Centroid().Distance(p2), 1E-10)
	assert.InDelta(t, c.Radius(), c.Centroid().Distance(p3), 1E-10)
}
