package grid

import (
	"github.com/adamcolton/vec2d"
	"math/rand"
)

// diamond dirs
var dd = [][]vec2d.I{
	{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}},
	{{-1, 0}, {0, -1}, {1, 0}, {0, 1}},
}

type diamondGenerator struct {
	scale      int
	size       vec2d.I
	grid       []float64
	noise      float64
	noiseDecay float64
}

func newDiamondGenerator(startSize vec2d.I, iterations int, noise, noiseDecay float64) *diamondGenerator {
	gen := &diamondGenerator{
		scale:      int(1 << uint(iterations)),
		noise:      noise,
		noiseDecay: noiseDecay,
	}
	o := -gen.scale + 1
	ov := vec2d.I{o, o}
	gen.size = startSize.ScalarMultiply(gen.scale).Add(ov)
	gen.grid = make([]float64, gen.size.Area())
	return gen
}

// Diamond uses the diamond-square algorithm to generate a height map
func Diamond(startSize vec2d.I, iterations int, noiseDecay float64) *Grid {
	gen := newDiamondGenerator(startSize, iterations, 1.0, noiseDecay)

	// populate the initial points
	for iter, pt, ok := startSize.FromOrigin().Start(); ok; pt, ok = iter.Next() {
		p2 := pt.ScalarMultiply(gen.scale)
		idx := gen.size.Idx(p2)
		gen.grid[idx] = rand.Float64()
	}

	gen.fillAll()
	g := gen.getGrid()
	Normalize(g)
	return g
}

func (gen *diamondGenerator) getGrid() *Grid {
	g := &Grid{
		Size: gen.size,
		Data: make([]interface{}, len(gen.grid)),
	}
	for i, f := range gen.grid {
		g.Data[i] = f
	}
	return g
}

func (gen *diamondGenerator) fillAll() {
	for gen.scale >>= 1; gen.scale > 0; gen.scale >>= 1 {
		gen.noise *= gen.noiseDecay
		gen.fill(gen.scale, gen.scale, 0)
		gen.fill(gen.scale, 0, 1)
		gen.fill(0, gen.scale, 1)
	}
}

func (gen *diamondGenerator) fill(x0, y0, dirIdx int) {
	for y := y0; y < gen.size.Y; y += 2 * gen.scale {
		for x := x0; x < gen.size.X; x += 2 * gen.scale {
			gen.point(vec2d.I{x, y}, dirIdx)
		}
	}
}

func (gen *diamondGenerator) point(pt vec2d.I, dirIdx int) {
	var v, c float64
	for _, d := range dd[dirIdx] {
		p2 := pt.Add(d.ScalarMultiply(gen.scale))
		if !p2.In(origin, gen.size) {
			continue
		}
		c++
		v += gen.grid[gen.size.Idx(p2)]
	}
	gen.grid[gen.size.Idx(pt)] = (v / c) + (0.5-rand.Float64())*gen.noise
}

// Normalize takes a grid of float64 values and normalizes it so the lowest
// value is 0.0 and the highest is 1.0.
func Normalize(g *Grid) {
	min := 1.0
	max := 0.0
	for _, vi := range g.Data {
		v := vi.(float64)
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	d := max - min
	for i, vi := range g.Data {
		v := vi.(float64)
		g.Data[i] = (v - min) / d
	}
}

func Gradient(start *Grid, iterations int) *Grid {
	gen := newDiamondGenerator(start.Size, iterations, 0, 0)

	// populate the initial points
	for iter, pt, ok := start.Size.FromOrigin().Start(); ok; pt, ok = iter.Next() {
		p2 := pt.ScalarMultiply(gen.scale)
		idx := gen.size.Idx(p2)
		gen.grid[idx] = start.Get(pt).(float64)
	}

	gen.fillAll()
	return gen.getGrid()
}
