package grid

import (
	"github.com/adamcolton/vec2d"
)

type SparseGrid struct {
	Size vec2d.I
	Data map[vec2d.I]interface{}
	gen  Generator
}

func NewSparseGrid(size vec2d.I, generator Generator) *SparseGrid {
	g := &SparseGrid{
		Size: size,
		Data: make(map[vec2d.I]interface{}),
		gen:  generator,
	}
	return g
}

func (g *SparseGrid) Get(pt vec2d.I) interface{} {
	pt = pt.Mod(g.Size)
	v, ok := g.Data[pt]
	if !ok && g.gen != nil {
		v = g.gen(pt)
		g.Data[pt] = v
	}
	return v
}

func (g *SparseGrid) Set(pt vec2d.I, val interface{}) {
	g.Data[pt.Mod(g.Size)] = val
}

func (g *SparseGrid) GetSize() vec2d.I {
	return g.Size
}
