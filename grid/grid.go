package grid

import (
	"github.com/adamcolton/vec2d"
)

type Grid interface {
	Get(vec2d.I) interface{}
	Set(vec2d.I, interface{})
	GetSize() vec2d.I
}

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
