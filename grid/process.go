package grid

import (
	"github.com/adamcolton/vec2d"
)

type Processor func(pt vec2d.I, g Grid) interface{}

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
