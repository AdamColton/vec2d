package vec2d

import (
	"math"
)

// P represents a polar coordinate defined by a magnitude and angle (in radians)
type P struct {
	M, A float64
}

// F converts a Polar coordinate to a Cartesean coordinate
func (p P) F() F {
	s, c := math.Sincos(p.A)
	return F{c * p.M, s * p.M}
}

// Add returns p + p2
func (p P) Add(p2 P) P {
	return p.F().Add(p2.F()).P()
}

// Subtract returns p + p2
func (p P) Subtract(p2 P) P {
	return p.F().Subtract(p2.F()).P()
}

const deg2rad = Pi / 180

// Deg converts degrees to radians
func Deg(d float64) float64 { return d * deg2rad }
