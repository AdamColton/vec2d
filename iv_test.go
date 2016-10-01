package vec2d

import (
	"testing"
)

func TestTo(t *testing.T) {
	a := I2d{1, 2}
	b := I2d{3, 4}
	expect := []I2d{{1, 2}, {1, 3}, {2, 2}, {2, 3}}
	for p := range a.To(b) {
		if p == expect[0] {
			expect = expect[1:]
		} else {
			t.Error("Expected: " + expect[0].String() + " Got: " + p.String())
		}
	}
}

func TestSliceTo(t *testing.T) {
	a := I2d{1, 2}
	b := I2d{3, 4}
	expect := []I2d{{1, 2}, {1, 3}, {2, 2}, {2, 3}}
	for i, p := range a.SliceTo(b) {
		if p != expect[i] {
			t.Error("Expected: " + expect[i].String() + " Got: " + p.String())
		}
	}
}
