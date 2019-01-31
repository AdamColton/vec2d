## Vec2D

The package provides tools for 2D vector math. There are 3 vector types, I
(integer), F (float64) and P(float64 polar).

Vectors of the same type can be directly compared and will be true if they are
the same point, even if they are different instances. This also means they can
be used as keys in maps. This is not true for pointers to vectors

```Go
I{1,2} == I{1,2} // true
F{3,4} == F{3,4} // true
&I{1,2} == &I{1,2} // false
```

The Vectors are also useful as embedded fields.

Documentation available at [https://godoc.org/github.com/AdamColton/vec2d]

### Curves and Paths
Curves are represented as parametric equations where t = 0.0 is the start of the
curve and t = 1.0 is the end of the curve.

A Path pairs together a curve and it's Tangent. 

#### Lines
Lines are represented as parametric equations rather than slope intercept form.
This makes is easier to deal with vertical lines. It also allows points to be
referenced by their index points.

```Go
l := p1.LineTo(p2)
l(0) == p1
l(1) == p2
idxX0 := l.AtX(0)
l(idxX0).X == 0
idxY0 := l.AtY(0)
l(idxY0).y == 0

l2 := p3.LineTo(p2)
i1, i2 := l.Intersection(l2)
l(i1) == p2
l2(i2) == p2
```

### Iterators
Right now there are only IntIterators, but there may be more in the future.
Here's how to use an iterator

```go
for iter,p,ok := a.To(b); ok; p,ok = iter.Next{
  fmt.Println(iter.Idx(), p)
}
```

### Surfaces and Shapes

Like Curves and Paths, these two describe the same sort of object but in
different degrees of detail.

A Surface takes two parameters and returns a point (type F) on the surface. On
a Shape, the F method does the same thing, but there are several other methods
that are useful.

### Extension
Many of these structures and interfaces are meant to be extended. Make sure that
if you cast something to a Line that it really is a line. It's possible to cast
a EllipticArc to a Line, but unless you're doing something with non-Euclidean
geometry, the results are probably not going to work out (and even then...).

### Notes on norms

#### CCW
Where possible and appropriate, data should proceed counter-clockwise for normal
data. For instance, when calling Rectangle2Points(p1, p2), if p1.X > p2.X and
p1.Y > p2.Y then the points returned will proceed counter clockwise, starting
from p1.

Also where possible, the 0 point of a rotation should be along the positive X
axis.