package grid

import (
	"github.com/adamcolton/vec2d"
)

// Processor can perform an operation on a grid by returning a new value for
// each point
type Processor func(pt vec2d.I, g Grid) interface{}

// Process takes a grid in and applies the processor to each node and populating
// the out grid. If no out grid is defined, a DenseGrid is created and returned.
func Process(in, out Grid, processor Processor) Grid {
	sz := in.GetSize()
	if out == nil {
		out = &DenseGrid{
			Size: sz,
			Data: make([]interface{}, sz.Area()),
		}
	}

	for iter, pt, ok := sz.FromOrigin().Start(); ok; pt, ok = iter.Next() {
		out.Set(pt, processor(pt, in))
	}

	return out
}
