package grid

import (
	"github.com/adamcolton/vec2d"
)

var origin vec2d.I
var dirs = vec2d.I{-1, -1}.SliceTo(vec2d.I{2, 2})

type Generator func(pt vec2d.I) interface{}

type Grid struct {
	Size vec2d.I
	Data []interface{}
}

func New(size vec2d.I, generator Generator) *Grid {
	g := &Grid{
		Size: size,
		Data: make([]interface{}, size.Area()),
	}
	if generator != nil {
		for iter, pt, ok := g.Size.FromOrigin(); ok; pt, ok = iter.Next() {
			g.Data[iter.Idx()] = generator(pt)
		}
	}
	return g
}

func (g *Grid) Get(pt vec2d.I) interface{} {
	return g.Data[g.Size.Idx(pt.Mod(g.Size))]
}

func (g *Grid) Set(pt vec2d.I, val interface{}) {
	g.Data[g.Size.Idx(pt.Mod(g.Size))] = val
}

type Processor func(pt vec2d.I, g *Grid) interface{}

func (g *Grid) Process(processor Processor) *Grid {
	out := &Grid{
		Size: g.Size,
		Data: make([]interface{}, g.Size.Area()),
	}

	for iter, pt, ok := g.Size.FromOrigin(); ok; pt, ok = iter.Next() {
		out.Data[iter.Idx()] = processor(pt, g)
	}

	return out
}
