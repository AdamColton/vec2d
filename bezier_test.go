package vec2d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBezierCurve(t *testing.T) {
	c := NewBezierCurve(F{0, 0}, F{0.5, 1}, F{1, 0})
	assert.Equal(t, F{0, 0}, c(0))
	assert.Equal(t, F{1, 0}, c(1))
	assert.Equal(t, F{0.5, 0.5}, c(0.5))

	c = NewBezierCurve(F{0, 0}, F{0, 1}, F{1, 1}, F{1, 0})
	assert.Equal(t, F{0, 0}, c(0))
	assert.Equal(t, F{1, 0}, c(1))
	assert.Equal(t, F{0.5, 0.75}, c(0.5))

	c = NewBezierCurve(F{0, 0}, F{0, 1}, F{1, -1}, F{1, 0})
	assert.Equal(t, F{0, 0}, c(0))
	assert.Equal(t, F{1, 0}, c(1))
	assert.Equal(t, F{0.5, 0}, c(0.5))
	assert.Equal(t, c(0.75).X, 1-c(0.25).X)
	assert.Equal(t, c(0.75).Y, -c(0.25).Y)
}

func TestBezierTangent(t *testing.T) {
	c := NewBezierTangent(F{0, 0}, F{0, 5}, F{5, 5}, F{5, 0})
	assert.Equal(t, 0.0, c(0).X)
	assert.Equal(t, 0.0, c(1).X)
	assert.Equal(t, 0.0, c(0.5).Y)
}

func TestBinomialCo(t *testing.T) {
	assert.Equal(t, 1.0, binomialCo(1, 1))
	assert.Equal(t, 2.0, binomialCo(2, 1))
	assert.Equal(t, 3.0, binomialCo(3, 1))
	assert.Equal(t, 6.0, binomialCo(4, 2))
}
