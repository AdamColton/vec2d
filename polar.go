package vec2d

import (
	"math"
)

// P represents a polar coordinate defined by a magnitude and angle (in radians)
type P struct {
	M, A float64
}

// Converts a Polar coordinate to a Cartesean coordinate
func (p P) F2d() F2d {
	return F2d{math.Cos(p.A) * p.M, math.Sin(p.A) * p.M}
}

// Returns a + b
func (a P) Add(b P) P {
	return a.F2d().Add(b.F2d()).P()
}

// Returns a + b
func (a P) Subtract(b P) P {
	return a.F2d().Subtract(b.F2d()).P()
}

const deg2rad = math.Pi / 180

// Deg converts degrees to radians
func Deg(d float64) float64 { return d * deg2rad }
