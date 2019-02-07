package grid

import (
	"github.com/adamcolton/vec2d"
	"reflect"
)

type Iter struct {
	g              Grid
	iter           vec2d.IntIterator
	start, end, pt vec2d.I
	ref            interface{}
	t              reflect.Type
}

func NewIter(g Grid, start, end vec2d.I, ref interface{}) *Iter {
	t := reflect.TypeOf(ref)
	if t.Kind() != reflect.Ptr {
		panic("Iter requires a pointer")
	}

	return &Iter{
		start: start,
		end:   end,
		ref:   ref,
		g:     g,
		t:     t,
	}
}

func IterAll(g Grid, i interface{}) *Iter {
	return NewIter(g, origin, g.GetSize(), i)
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
			reflect.ValueOf(i.ref).Elem().Set(v.Elem())
		} else if v.Type() == i.t.Elem() {
			reflect.ValueOf(i.ref).Elem().Set(v)
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
