package vec2d

// BaseIntIterator provides the base methods that are needed to provide a full
// IntIterator.
type BaseIntIterator interface {
	Next() (I, bool)
	I() I
	Done() bool
	Idx() int
	Reset() (I, bool)
	Area() int
}

// IntIterator provides an interface for iterating over integer point ranges
type IntIterator interface {
	BaseIntIterator
	Start() (IntIterator, I, bool)
	Slice() []I
	Chan() <-chan I
	Each(fn func(int, I))
	Until(fn func(int, I) bool) bool
}

// xIterator is return from .To or .FromOrigin, it iterators by incrementing X
// first.
type xIterator struct {
	from, to, cur  I
	x, dx, dy      int
	checkX, checkY func(int, int) bool
	idx            int
	done           bool
}

// To is used to iterate over points. Given A.To(B) it will hit all the points
// in the rectangle between inclusive of A and exclusive of B. The iterator will
// move across a full row before moving to the next column.
func (i I) To(i2 I) IntIterator {
	r := &xIterator{
		from:   i,
		to:     i2,
		cur:    i,
		x:      i.X,
		dx:     1,
		dy:     1,
		checkX: gte,
		checkY: gte,
		done:   i == i2,
	}
	if i.X > i2.X {
		r.dx = -1
		r.checkX = lte
	}
	if i.Y > i2.Y {
		r.dy = -1
		r.checkY = lte
	}
	return IterBaseWrapper{r}
}

// FromOrigin returns an iterator from the origin (incluse) to i (exclusive).
func (i I) FromOrigin() IntIterator {
	return I{0, 0}.To(i)
}

func lte(a, b int) bool {
	return a <= b
}

func gte(a, b int) bool {
	return a >= b
}

func lt(a, b int) bool {
	return a < b
}

func gt(a, b int) bool {
	return a > b
}

// Done returns if the
func (xi *xIterator) Done() bool {
	return xi.done
}

// I returns the current point of the iterator
func (xi *xIterator) I() I {
	return xi.cur
}

// Next fulfills the IntIterator interface. It returns the next point and if
// iteration is done.
func (xi *xIterator) Next() (I, bool) {
	if xi.done {
		return xi.cur, false
	}
	xi.idx++
	xi.cur.X += xi.dx
	if xi.checkX(xi.cur.X, xi.to.X) {
		xi.cur.X = xi.x
		xi.cur.Y += xi.dy
		if xi.checkY(xi.cur.Y, xi.to.Y) {
			xi.done = true
		}
	}
	return xi.cur, !xi.done
}

// Idx fulfills the IntIterator interface. It returns the index of the current
// point.
func (xi *xIterator) Idx() int { return xi.idx }

// Area returns the numer of points the iterator will visit
func (xi *xIterator) Area() int {
	return xi.from.Subtract(xi.to).Abs().Area()
}

// Reset the iterator.
func (xi *xIterator) Reset() (I, bool) {
	xi.cur = xi.from
	xi.x = xi.from.X
	xi.idx = 0
	xi.done = xi.cur == xi.to
	return xi.cur, !xi.done
}

// IterBaseWrapper takes a type that fulfills BaseIntIterator and wraps it to
// fulfill IntIterator.
type IterBaseWrapper struct {
	BaseIntIterator
}

// Start is helper tha can be as the first portion of a classic for loop
func (base IterBaseWrapper) Start() (IntIterator, I, bool) {
	i, b := base.Reset()
	return base, i, b
}

// Slice returns all the points in the iterator as a slice.
func (base IterBaseWrapper) Slice() []I {
	s := make([]I, base.Area())
	for pt, ok := base.Reset(); ok; pt, ok = base.Next() {
		s[base.Idx()] = pt
	}
	return s
}

// Each calls the fn for each point in the iterator.
func (base IterBaseWrapper) Each(fn func(int, I)) {
	for pt, ok := base.Reset(); ok; pt, ok = base.Next() {
		fn(base.Idx(), pt)
	}
}

// Until calls fn against each point in the iterator until a point returns true.
// The bool indicates if a value returned true. The iterator will be left at the
// point that returned true.
func (base IterBaseWrapper) Until(fn func(int, I) bool) bool {
	for pt, ok := base.Reset(); ok; pt, ok = base.Next() {
		if fn(base.Idx(), pt) {
			return true
		}
	}
	return false
}

// Chan runs a go routine that will return the points of the iterator. When all
// the points are consumed the channel is closed. Failing to consume all the
// points will cause a Go routine leak.
func (base IterBaseWrapper) Chan() <-chan I {
	c := make(chan I)
	go func() {
		for pt, ok := base.Reset(); ok; pt, ok = base.Next() {
			c <- pt
		}
		close(c)
	}()
	return c
}
