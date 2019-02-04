package vec2d

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformApply(t *testing.T) {
	tt := []struct {
		T        Transformation
		F        F
		expected F
	}{
		{
			T: Transformation{
				X: F{1, 0},
				Y: F{0, 1},
			},
			F:        F{3, 4},
			expected: F{3, 4},
		},
		{
			T: Transformation{
				X: F{2, 0},
				Y: F{0, 3},
			},
			F:        F{1, 1},
			expected: F{2, 3},
		},
		{
			T: Transformation{
				X: F{2, 0},
				Y: F{0, 3},
			},
			F:        F{1, 0},
			expected: F{2, 0},
		},
		{
			T: Transformation{
				X: F{2, 0},
				Y: F{0, 3},
			},
			F:        F{0, 1},
			expected: F{0, 3},
		},
		{
			T: Transformation{
				X: F{0, 2},
				Y: F{3, 0},
			},
			F:        F{1, 0},
			expected: F{0, 2},
		},
		{
			T: Transformation{
				X: F{0, 2},
				Y: F{3, 0},
			},
			F:        F{0, 1},
			expected: F{3, 0},
		},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.expected, tc.T.Apply(tc.F))
	}
}

func TestTriangleTransform(t *testing.T) {
	tt := []struct {
		a, b Triangle
	}{
		{
			a: Triangle{
				F{0, 0},
				F{1, 0},
				F{0, 1},
			},
			b: Triangle{
				F{1, 1},
				F{2, 0},
				F{0, 1},
			},
		},
		{
			a: Triangle{
				F{0, 0},
				F{1, 0},
				F{0, 1},
			},
			b: Triangle{
				F{0, 1},
				F{0, 0},
				F{1, 0},
			},
		},
		// confirm 3 points in a line don't cause an error
		{
			a: Triangle{
				F{0, 0},
				F{1, 0},
				F{0, 1},
			},
			b: Triangle{
				F{1, 0},
				F{2, 0},
				F{3, 0},
			},
		},
	}

	for _, tc := range tt {
		tfrm, err := TriangleTransform(tc.a, tc.b)
		assert.NoError(t, err)

		for i := 0; i < 3; i++ {
			assert.Equal(t, tc.b[i], tfrm.Apply(tc.a[i]))
		}
	}

}
