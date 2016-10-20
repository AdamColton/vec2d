## Vec2D

This is a simple 2D vector library. Vectors of the same type can be directly
compared and will be true if they are the same point, even if they are different
instances. This also means they can be used as keys in maps. This is not true
for pointers to vectors

```Go
I{1,2} == I{1,2} // true
F{3,4} == F{3,4} // true
&I{1,2} == &I{1,2} // false
```

The Vectors are also useful as embedded fields.

Documentation available at [https://godoc.org/github.com/AdamColton/vec2d]

### Lines
Lines are represented as parametric equations rather than slope intercept form.
This makes is easier to deal with vertical lines. It also allows points to be
referenced by their index points.

```go
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
