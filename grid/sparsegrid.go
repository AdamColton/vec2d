package grid

import (
	"github.com/adamcolton/vec2d"
)

// SparseGrid fulfills Grid storing the data in a map. If the Generator is set
// data will be generated lazily.
type SparseGrid struct {
	Size vec2d.I
	Data map[vec2d.I]interface{}
	Gen  Generator
}

// NewSparseGrid creates a sparse grid of the defined size.
func NewSparseGrid(size vec2d.I, generator Generator) *SparseGrid {
	g := &SparseGrid{
		Size: size,
		Data: make(map[vec2d.I]interface{}),
		Gen:  generator,
	}
	return g
}

// Get the data at the defined pt. If the Generator is not nil, it will generate
// the node and populate the data. If pt is outside of the grid, the modulus
// will be used.
func (g *SparseGrid) Get(pt vec2d.I) interface{} {
	pt = pt.Mod(g.Size)
	v, ok := g.Data[pt]
	if !ok && g.Gen != nil {
		v = g.Gen(pt)
		if v != nil {
			g.Data[pt] = v
		}
	}
	return v
}

// Set the value at pt. If pt is outside the grid, the modulus will be taken. If
// val is nil, the pt will be deleted from the underlying map.
func (g *SparseGrid) Set(pt vec2d.I, val interface{}) {
	if val == nil {
		delete(g.Data, pt.Mod(g.Size))
	} else {
		g.Data[pt.Mod(g.Size)] = val
	}
}

// GetSize returns the size of the SparseGrid
func (g *SparseGrid) GetSize() vec2d.I {
	return g.Size
}
