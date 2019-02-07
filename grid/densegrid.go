package grid

import (
	"github.com/adamcolton/vec2d"
)

var origin vec2d.I
var dirs = vec2d.I{-1, -1}.To(vec2d.I{2, 2}).Slice()

type Generator func(pt vec2d.I) interface{}

type DenseGrid struct {
	Size vec2d.I
	Data []interface{}
}

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

func (g *DenseGrid) Get(pt vec2d.I) interface{} {
	return g.Data[g.Size.Idx(pt.Mod(g.Size))]
}

func (g *DenseGrid) Set(pt vec2d.I, val interface{}) {
	g.Data[g.Size.Idx(pt.Mod(g.Size))] = val
}

func (g *DenseGrid) GetSize() vec2d.I {
	return g.Size
}
