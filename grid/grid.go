package grid

import (
	"github.com/adamcolton/vec2d"
)

var origin vec2d.I
var dirs = vec2d.I{-1, -1}.To(vec2d.I{2, 2}).Slice()

// Generator is used to populate a grid
type Generator func(pt vec2d.I) interface{}

// Grid represents a grid of any type.
type Grid interface {
	Get(vec2d.I) interface{}
	Set(vec2d.I, interface{})
	GetSize() vec2d.I
}

// Flood takes a grid and beginning at start floods out in every direction in
// dirs, checking each point against include. If a point is included, it will
// flood out from there. All points that are included are returned as a slice.
func Flood(g Grid, start vec2d.I, dirs []vec2d.I, include func(pt vec2d.I, g Grid) bool) []vec2d.I {
	var ret []vec2d.I
	seen := map[vec2d.I]bool{
		start: true,
	}
	sz := g.GetSize()
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
				if dpt.In(origin, sz) {
					q = append(q, dpt)
				}
			}
		}
	}

	return ret
}
