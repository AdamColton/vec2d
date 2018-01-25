package vec2d

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestF(t *testing.T) {
	v := F{5, 6}
	assert.Equal(t, v.X, 5.0)
	assert.Equal(t, v.Y, 6.0)
}

func TestAdd(t *testing.T) {
	v1 := F{5, 6}
	v2 := F{1, 2}
	assert.Equal(t, F{6, 8}, v1.Add(v2))
}

func TestRef(t *testing.T) {
	v1 := &F{5, 6}
	v2 := &F{1, 2}
	assert.Equal(t, F{6, 8}, v1.Add(*v2))
	// Show the values are not changed
	assert.Equal(t, F{5, 6}, *v1)
}

func TestArea(t *testing.T) {
	assert.Equal(t, 12.0, F{3, 4}.Area())
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, F{10, 20}, F{5, 2}.Multiply(F{2, 10}))
}

func TestScalarMultiply(t *testing.T) {
	assert.Equal(t, F{6, 10}, F{3, 5}.ScalarMultiply(2))
}

func TestAngle(t *testing.T) {
	v := P{1, 1}.F()
	assert.Equal(t, 1.0, v.Angle())
}

func TestSubtract(t *testing.T) {
	v1 := F{6, 4}
	v2 := F{2, 1}
	assert.Equal(t, F{4, 3}, v1.Subtract(v2))
}

func TestRotate(t *testing.T) {
	v := P{1, 1}.F()
	v = v.Rotate(1)
	assert.Equal(t, 2.0, v.Angle())
}

func TestMag(t *testing.T) {
	v := P{2, 1}.F()
	assert.Equal(t, 2.0, v.Mag())
}

func TestDistance(t *testing.T) {
	assert.Equal(t, 5.0, F{1, 2}.Distance(F{4, 6}))
}

func TestCross(t *testing.T) {
	assert.Equal(t, -2.0, F{1, 2}.Cross(F{4, 6}))
}

func TestIntersect(t *testing.T) {
	ms := F{0, 1}
	me := F{2, 1}
	ss := F{1, 0}
	se := F{1, 2}
	assert.Equal(t, 0.5, MotionSurfaceIntersection(ms, me, ss, se))

	ms = F{2, 2}
	me = F{10, 6}
	ss = F{8, 1}
	se = F{6, 8}
	assert.Equal(t, 0.625, MotionSurfaceIntersection(ms, me, ss, se))
}

func TestEmbed(t *testing.T) {
	type Foo struct {
		F
		name string
	}

	foo1 := &Foo{
		F:    F{3, 1},
		name: "Adam",
	}
	foo2 := &Foo{
		F:    F{3, 1},
		name: "Adam Colton",
	}
	assert.True(t, foo1.F == foo2.F)
}

func TestTriangulage(t *testing.T) {
	testCases := []struct {
		center F
		m      float64
		angles []float64
	}{
		{
			center: F{2, 3},
			m:      2.0,
			angles: []float64{1, 2, 3},
		},
		{
			center: F{1, 4},
			m:      1.23,
			angles: []float64{0, math.Pi, 3},
		},
	}

	for _, tc := range testCases {
		p1 := tc.center.Add(P{tc.m, tc.angles[0]}.F())
		p2 := tc.center.Add(P{tc.m, tc.angles[1]}.F())
		p3 := tc.center.Add(P{tc.m, tc.angles[2]}.F())

		assert.InDelta(t, 0, tc.center.Distance(Triangulate(p1, p2, p3)), 1e-10)
	}

	p1 := F{0, 0}
	p2 := F{0, 0}
	p3 := F{0, 0}

	assert.InDelta(t, 0, F{0, 0}.Distance(Triangulate(p1, p2, p3)), 1e-10)

	p3 = F{2, 2}
	assert.InDelta(t, 0, F{1, 1}.Distance(Triangulate(p1, p2, p3)), 1e-10)
}
