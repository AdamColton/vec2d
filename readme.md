## Vec2D

This is a simple 2D vector library. Vectors of the same type can be directly compared and will be true if they are the same point, even if they are different instances. This is not true for pointers to vectors

```Go
I2d{1,2} == I2d{1,2} // true
F2d{3,4} == F2d{3,4} // true
&I2d{1,2} == &I2d{1,2} // false
```

The Vectors are also very useful as embedded fields.

Documentation available at [https://godoc.org/github.com/AdamColton/vec2d]