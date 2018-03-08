package grid

import (
	"github.com/adamcolton/vec2d"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGrid(t *testing.T) {
	size := vec2d.I{10, 10}
	generator := func(pt vec2d.I) interface{} {
		return rand.Intn(9) + 1
	}
	g := New(size, generator)
	assert.NotNil(t, g)
	for iter, pt, ok := g.Size.FromOrigin(); ok; pt, ok = iter.Next() {
		i, ok := g.Data[iter.Idx()].(int)
		assert.True(t, ok)
		assert.NotEqual(t, 0, i)
		assert.Equal(t, i, g.Get(pt).(int))
	}
}

func TestPluralProcessor(t *testing.T) {
	size := vec2d.I{10, 10}
	generator := func(pt vec2d.I) interface{} {
		return rand.Intn(100)
	}
	g := New(size, generator)
	proc := PluralProcessor(func(i interface{}) int { return i.(int) })
	for i := 0; i < 10; i++ {
		g = g.Process(proc)
	}
}

func TestIter(t *testing.T) {
	// interface type matches element
	gen := func(pt vec2d.I) interface{} {
		if (pt.X == 0 && pt.Y == 0) || (pt.X == 2 && pt.Y == 2) {
			return 1.0
		}
		if pt.X == 1 && pt.Y == 1 {
			return 0.0
		}
		return 0.5
	}
	g := New(vec2d.I{3, 3}, gen)
	var f float64
	iter := g.IterAll(&f)
	for iter.Next() {
		assert.Equal(t, gen(iter.Pt()).(float64), f)
	}

	// interface type matches interface
	gen2 := func(pt vec2d.I) interface{} {
		v := gen(pt).(float64)
		return &v
	}
	g = New(vec2d.I{3, 3}, gen2)
	iter = g.IterAll(&f)
	for iter.Next() {
		assert.Equal(t, gen(iter.Pt()).(float64), f)
	}
}

func TestFlood(t *testing.T) {
	gen := func(pt vec2d.I) interface{} {
		return pt.X < 2 && pt.Y < 2
	}
	g := New(vec2d.I{4, 4}, gen)

	ds := []vec2d.I{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	pts := g.Flood(origin, ds, func(pt vec2d.I, g *Grid) bool {
		return g.Get(pt).(bool)
	})
	assert.Len(t, pts, 4)

	pts = g.Flood(vec2d.I{3, 3}, ds, func(pt vec2d.I, g *Grid) bool {
		return g.Get(pt).(bool)
	})
	assert.Len(t, pts, 0)

	pts = g.Flood(vec2d.I{3, 3}, ds, func(pt vec2d.I, g *Grid) bool {
		return !g.Get(pt).(bool)
	})
	assert.Len(t, pts, 12)
}