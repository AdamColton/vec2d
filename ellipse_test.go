package vec2d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEllipse(t *testing.T) {
	f1 := F{0, 0}
	f2 := F{2, 0}
	r := 1.0

	e := NewEllipse(f1, f2, r)
	assert.Equal(t, r, e.F(0.25).Distance(e.c))
	assert.InDelta(t, 0.0, e.F(0).Y, 0.000001)

	// Test by definition of ellipse
	p0 := e.F(0)
	d0 := f1.Distance(p0) + f2.Distance(p0)
	for i := 0.0; i <= 1.0; i += 0.2 {
		p := e.F(i)
		d := f1.Distance(p) + f2.Distance(p)
		assert.InDelta(t, d0, d, 0.000001)
	}

	// Get correct foci
	tf1, tf2 := e.Foci()
	assert.InDelta(t, 0, f1.Distance(tf1), 0.00001)
	assert.InDelta(t, 0, f2.Distance(tf2), 0.00001)

	var p Path
	p = e
	assert.NotNil(t, p)
}
