package vec2d

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTo(t *testing.T) {
	a := I{1, 2}
	b := I{3, 4}
	expected := []I{{1, 2}, {2, 2}, {1, 3}, {2, 3}}
	for iter, p, ok := a.To(b).Start(); ok; p, ok = iter.Next() {
		assert.Equal(t, expected[iter.Idx()], p)
	}
}

func TestToSamePoint(t *testing.T) {
	a := I{1, 2}
	for iter, _, ok := a.To(a).Start(); ok; _, ok = iter.Next() {
		t.Error("Should hit no points")
	}
}

func TestIterSlice(t *testing.T) {
	a := I{1, 2}
	b := I{3, 4}
	expected := []I{{1, 2}, {2, 2}, {1, 3}, {2, 3}}
	assert.Equal(t, expected, a.To(b).Slice())
}

func TestMod(t *testing.T) {
	size := I{3, 4}
	in := []I{
		{1, 2},
		{-1, -1},
	}
	expect := []I{
		{1, 2},
		{2, 3},
	}
	for i, p := range in {
		assert.Equal(t, expect[i], p.Mod(size))
	}
}

func TestIdx(t *testing.T) {
	size := I{5, 7}
	idx := size.Idx(I{3, 4})
	assert.Equal(t, 23, idx)
	p := size.InvIdx(idx)
	assert.Equal(t, I{3, 4}, p)
}

func TestIn(t *testing.T) {
	tests := []struct {
		i, a, b  I
		expected bool
	}{
		{I{1, 1}, I{0, 0}, I{2, 2}, true},
		{I{0, 0}, I{0, 0}, I{2, 2}, true},
		{I{2, 0}, I{0, 0}, I{2, 2}, false},
		{I{0, 2}, I{0, 0}, I{2, 2}, false},
		{I{2, 2}, I{2, 2}, I{0, 0}, true},
		{I{0, 0}, I{2, 2}, I{0, 0}, false},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, test.i.In(test.a, test.b))
	}
}

func TestStandardInt(t *testing.T) {
	// An angle of 0 should be +X
	assert.Equal(t, 0.0, I{1, 0}.Angle())
	assert.Equal(t, Pi/2, I{0, 1}.Angle())
}

func TestEach(t *testing.T) {
	p := I{3, 3}.FromOrigin()
	assert.Equal(t, 9, p.Area())
	fs := make([]float64, p.Area())
	p.Each(func(idx int, i I) {
		fs[idx] = i.Mag()
	})
	assert.Equal(t, math.Sqrt2, fs[4])
}

func TestIterChan(t *testing.T) {
	a := I{1, 2}
	b := I{3, 4}
	expected := []I{{1, 2}, {2, 2}, {1, 3}, {2, 3}}
	for i := range a.To(b).Chan() {
		assert.Equal(t, expected[0], i)
		expected = expected[1:]
	}
}
