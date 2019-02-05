package grid

import (
	"github.com/adamcolton/vec2d"
	"reflect"
)

type Iter struct {
	g              *DenseGrid
	iter           vec2d.IntIterator
	start, end, pt vec2d.I
	i              interface{}
	t              reflect.Type
}

func (g *DenseGrid) Iter(start, end vec2d.I, i interface{}) *Iter {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		panic("Iter requires a pointer")
	}

	return &Iter{
		start: start,
		end:   end,
		i:     i,
		g:     g,
		t:     t,
	}
}

func (g *DenseGrid) IterAll(i interface{}) *Iter {
	return g.Iter(origin, g.Size, i)
}

func (i *Iter) Next() bool {
	var ok bool
	if i.iter == nil {
		i.iter, i.pt, ok = i.start.To(i.end).Start()
	} else {
		i.pt, ok = i.iter.Next()
	}
	if ok {
		v := reflect.ValueOf(i.g.Get(i.pt))
		if v.Type() == i.t {
			reflect.ValueOf(i.i).Elem().Set(v.Elem())
		} else if v.Type() == i.t.Elem() {
			reflect.ValueOf(i.i).Elem().Set(v)
		} else {
			panic("types don't match")
		}
	}
	return ok
}

func (i *Iter) Pt() vec2d.I {
	return i.pt
}

func (i *Iter) Idx() int {
	if i.iter == nil {
		return -1
	}
	return i.iter.Idx()
}
