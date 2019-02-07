package grid

import (
	"github.com/adamcolton/vec2d"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
)

func TestDenseGrid(t *testing.T) {
	size := vec2d.I{10, 10}
	generator := func(pt vec2d.I) interface{} {
		return rand.Intn(9) + 1
	}
	g := NewDenseGrid(size, generator)
	assert.NotNil(t, g)
	for iter, pt, ok := g.Size.FromOrigin().Start(); ok; pt, ok = iter.Next() {
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
	g := NewDenseGrid(size, generator)
	proc := PluralProcessor(func(i interface{}) int { return i.(int) })
	for i := 0; i < 10; i++ {
		out := NewDenseGrid(size, nil)
		Process(g, out, proc)
		g = out
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
	g := NewDenseGrid(vec2d.I{3, 3}, gen)
	var f float64

	for iter, ok := IterAll(g, &f).Start(); ok; ok = iter.Next() {
		assert.Equal(t, gen(iter.I()).(float64), f)
	}

	// interface type matches interface
	gen2 := func(pt vec2d.I) interface{} {
		v := gen(pt).(float64)
		return &v
	}
	g = NewDenseGrid(vec2d.I{3, 3}, gen2)
	for iter, ok := IterAll(g, &f).Start(); ok; ok = iter.Next() {
		assert.Equal(t, gen(iter.I()).(float64), f)
	}
}

func TestFlood(t *testing.T) {
	gen := func(pt vec2d.I) interface{} {
		return pt.X < 2 && pt.Y < 2
	}
	g := NewDenseGrid(vec2d.I{4, 4}, gen)

	ds := []vec2d.I{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	pts := Flood(g, origin, ds, func(pt vec2d.I, g Grid) bool {
		return g.Get(pt).(bool)
	})
	assert.Len(t, pts, 4)

	pts = Flood(g, vec2d.I{3, 3}, ds, func(pt vec2d.I, g Grid) bool {
		return g.Get(pt).(bool)
	})
	assert.Len(t, pts, 0)

	pts = Flood(g, vec2d.I{3, 3}, ds, func(pt vec2d.I, g Grid) bool {
		return !g.Get(pt).(bool)
	})
	assert.Len(t, pts, 12)
}

func TestInterface(t *testing.T) {
	var g Grid
	generator := func(pt vec2d.I) interface{} {
		return rand.Intn(9) + 1
	}
	g = NewDenseGrid(vec2d.I{10, 10}, generator)
	assert.NotNil(t, g)
}

func TestSparseGrid(t *testing.T) {
	size := vec2d.I{10, 10}
	generator := func(pt vec2d.I) interface{} {
		return pt.Area()
	}
	g := NewSparseGrid(size, generator)
	var _ Grid = g
	assert.NotNil(t, g)
	assert.Len(t, g.Data, 0)
	assert.Equal(t, 2, g.Get(vec2d.I{1, 2}))
	assert.Len(t, g.Data, 1)
}

func TestFormatter(t *testing.T) {
	f := Formatter{
		Separator:  " | ",
		AlignRight: true,
	}
	generator := func(pt vec2d.I) interface{} {
		return pt.Area()
	}
	str := f.Format(NewDenseGrid(vec2d.I{10, 10}, generator))
	assert.Equal(t, 90, strings.Count(str, "|"))
}
