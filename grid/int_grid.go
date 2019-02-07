package grid

import (
	"github.com/adamcolton/vec2d"
)

// Special tools for int grids

// PluralProcessor returns a Processor that sets each cell to the value with
// the highest plurality around it
func PluralProcessor(getInt func(interface{}) int) Processor {
	return func(pt vec2d.I, g Grid) interface{} {
		counts := make(map[int]int)
		var bestVal, bestCount int
		for _, d := range dirs {
			v := getInt(g.Get(pt.Add(d)))
			c := counts[v] + 1
			counts[v] = c
			// preference is to not change
			if c > bestCount || (c == bestCount && d == origin) {
				bestVal = v
				bestCount = c
			}
		}
		return bestVal
	}
}
