package vec2d

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestLineTo(t *testing.T) {
	line := F{1, 3}.LineTo(F{2, 4})
	assert.Equal(t, F{1, 3}, line(0))
	assert.Equal(t, F{2, 4}, line(1))
	assert.Equal(t, -1.0, line.AtX(0))
	assert.Equal(t, -3.0, line.AtY(0))
	assert.Equal(t, 2.0, line.B())
	assert.Equal(t, 1.0, line.M())
}

func TestBisect(t *testing.T) {
	line := F{1, 2}.Bisect(F{0, 3})
	assert.EqualValues(t, F{0.5, 2.5}, line(0))
	assert.Equal(t, 2.0, line.B())
	assert.Equal(t, 1.0, line.M())

	p1, p2 := F{0.000000, 89.000000}, F{44.452365, 56.126156}
	line = p1.Bisect(p2)
	sx := []float64{0, 22.226183 / 2, 22.226183}
	for _, x := range sx {
		p := F{x, line(line.AtX(x)).Y}
		assert.InDelta(t, p1.Distance(p)-p2.Distance(p), 0, 1e-10)
	}
}

func TestLineIntersect(t *testing.T) {
	testCases := []struct {
		name      string
		points    []F
		expectNaN bool
	}{
		{
			name:   "normal line",
			points: []F{F{0, 1}, F{1, 2}, F{1, 0}, F{2, 3}},
		},
		{
			name:   "first line is vertical",
			points: []F{F{0, 1}, F{0, 2}, F{1, 0}, F{2, 3}},
		},
		{
			name:   "second line is vertical",
			points: []F{F{0, 1}, F{1, 2}, F{1, 0}, F{1, 3}},
		},
		{
			name:      "lines are parallel",
			points:    []F{F{0, 1}, F{1, 2}, F{1, 2}, F{2, 3}},
			expectNaN: true,
		},
		{
			name:      "first line is a point",
			points:    []F{F{0, 1}, F{0, 1}, F{1, 2}, F{2, 3}},
			expectNaN: true,
		},
	}

	for _, tc := range testCases {
		l0 := tc.points[0].LineTo(tc.points[1])
		l1 := tc.points[2].LineTo(tc.points[3])
		t0, t1 := l0.Intersection(l1)
		if tc.expectNaN {
			assert.True(t, math.IsNaN(t0), "intersection NaN t0", tc.name)
			assert.True(t, math.IsNaN(t1), "intersection NaN t1", tc.name)
		} else {
			assert.Equal(t, l0(t0), l1(t1), "intersection", tc.name)
		}
	}
}

func TestClosest(t *testing.T) {
	l := F{1, 2}.LineTo(F{1, 1})
	p := l.Closest(F{0, 0})
	assert.Equal(t, F{1, 0}, p)

	l = F{0, 0}.LineTo(F{1, 1})
	p = l.Closest(F{0, 1})
	assert.Equal(t, F{0.5, 0.5}, p)

	l = F{0, 0}.LineTo(F{1, 3})
	p = l.Closest(F{-3, 1})
	assert.Equal(t, F{0, 0}, p)
}
