package grid

import (
	"github.com/adamcolton/vec2d"
	"reflect"
)

// Iter is a helper for iterating over a grid. Iter uses reflection for type
// checking. When Iter is constructed, it will be passed a reference value that
// must be a pointer. Each time Next is called, that pointer will be updated.
type Iter struct {
	g       Grid
	intIter vec2d.IntIterator
	ref     interface{}
	t       reflect.Type
}

// NewIter will iterate over the grid using the provided IntIterator, updating
// the ref at each node.
func NewIter(g Grid, intIter vec2d.IntIterator, ref interface{}) *Iter {
	t := reflect.TypeOf(ref)
	if t.Kind() != reflect.Ptr {
		panic("Iter requires a pointer")
	}

	return &Iter{
		ref:     ref,
		g:       g,
		t:       t,
		intIter: intIter,
	}
}

// IterRange will iterate over the sub-grid defined by the points start and end.
func IterRange(g Grid, start, end vec2d.I, ref interface{}) *Iter {
	return NewIter(g, start.To(end), ref)
}

// IterAll will iterate over the entire grid
func IterAll(g Grid, ref interface{}) *Iter {
	return NewIter(g, g.GetSize().FromOrigin(), ref)
}

// Start the iterator
func (i *Iter) Start() (*Iter, bool) {
	_, pt, ok := i.intIter.Start()
	if ok {
		i.setRef(pt)
	}
	return i, ok
}

func (i *Iter) setRef(pt vec2d.I) {
	v := reflect.ValueOf(i.g.Get(pt))
	if v.Type() == i.t {
		reflect.ValueOf(i.ref).Elem().Set(v.Elem())
	} else if v.Type() == i.t.Elem() {
		reflect.ValueOf(i.ref).Elem().Set(v)
	} else {
		panic("types don't match")
	}
}

// Next increments the iterator and update the reference pointer, using
// reflection for type checking. It returns a bool indicating if the iterator
// is complete.
func (i *Iter) Next() bool {
	pt, ok := i.intIter.Next()
	if ok {
		i.setRef(pt)
	}
	return ok
}

// Pt returns the current point that iterator is on.
func (i *Iter) I() vec2d.I {
	return i.intIter.I()
}

// Idx returns the index value of the current point.
func (i *Iter) Idx() int {
	return i.intIter.Idx()
}

// Done returns true if the iterator is done
func (i *Iter) Done() bool {
	return i.intIter.Done()
}
