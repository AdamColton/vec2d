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

type Processor func(pt vec2d.I, g *DenseGrid) interface{}

func (g *DenseGrid) Process(processor Processor) *DenseGrid {
	out := &DenseGrid{
		Size: g.Size,
		Data: make([]interface{}, g.Size.Area()),
	}

	for iter, pt, ok := g.Size.FromOrigin().Start(); ok; pt, ok = iter.Next() {
		out.Data[iter.Idx()] = processor(pt, g)
	}

	return out
}

func (g *DenseGrid) Flood(start vec2d.I, dirs []vec2d.I, include func(pt vec2d.I, g *DenseGrid) bool) []vec2d.I {
	var ret []vec2d.I
	seen := map[vec2d.I]bool{
		start: true,
	}
	q := []vec2d.I{start}
	for ln := len(q); ln > 0; ln = len(q) {
		pt := q[ln-1]
		q = q[:ln-1]
		if !include(pt, g) {
			continue
		}

		ret = append(ret, pt)
		for _, d := range dirs {
			dpt := pt.Add(d)
			if !seen[dpt] {
				seen[dpt] = true
				if dpt.In(origin, g.Size) {
					q = append(q, dpt)
				}
			}
		}

	}

	return ret
}
