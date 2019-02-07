package grid

import (
	"github.com/adamcolton/vec2d"
)

// DenseGrid is backed by a slice and should have an entry for every node in
// the grid.
type DenseGrid struct {
	Size vec2d.I
	Data []interface{}
}

// NewDenseGrid creates a DenseGrid with Data allocated at the correct size. If
// a Generator is given, the data will be populated.
func NewDenseGrid(size vec2d.I, generator Generator) *DenseGrid {
	g := &DenseGrid{
		Size: size,
		Data: make([]interface{}, size.Area()),
	}
	if generator != nil {
		for iter, pt, ok := g.Size.FromOrigin().Start(); ok; pt, ok = iter.Next() {
			g.Data[iter.Idx()] = generator(pt)
		}
	}
	return g
}

// Get a value, fulfills Grid
func (g *DenseGrid) Get(pt vec2d.I) interface{} {
	return g.Data[g.Size.Idx(pt.Mod(g.Size))]
}

// Set a value, fulfills Grid
func (g *DenseGrid) Set(pt vec2d.I, val interface{}) {
	g.Data[g.Size.Idx(pt.Mod(g.Size))] = val
}

// GetSize of the DenseGrid, fulfills Grid
func (g *DenseGrid) GetSize() vec2d.I {
	return g.Size
}
