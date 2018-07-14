package vec2d

import (
	"math"
	"sort"
)

type Polygon []F

func (p Polygon) F(t0, t1 float64) F {
	n := (len(p) - 1)
	h := n - (n / 2) // half rounded up
	as := LineSegments(p[0:h])
	bs := LineSegments(p[h:])
	a := as.F(t0)
	b := bs.F(1 - t0)
	return a.LineTo(b)(t1)
}

func (p Polygon) SignedArea() float64 {
	var s float64
	prev := p[len(p)-1]
	for _, cur := range p {
		s += prev.X*cur.Y - cur.X*prev.Y
		prev = cur
	}
	return s / 2
}

func (p Polygon) Area() float64 {
	return math.Abs(p.SignedArea())
}

func (p Polygon) Centroid() F {
	var x, y, a float64
	prev := p[len(p)-1]
	for _, cur := range p {
		t := (prev.X*cur.Y - cur.X*prev.Y)
		x += (prev.X + cur.X) * t
		y += (prev.Y + cur.Y) * t
		a += t
		prev = cur
	}
	a = 1 / (3 * a)
	return F{x * a, y * a}
}

func (p Polygon) Contains(f F) bool {
	// https://en.wikipedia.org/wiki/Point_in_polygon#Ray_casting_algorithm
	ray := f.LineTo(f.Add(F{1, 0}))
	prev := p[len(p)-1]
	var itersects int
	for _, cur := range p {
		side := prev.LineTo(cur)
		ri, si := ray.Intersection(side)
		if ri > 0 && si >= 0 && si < 1 {
			itersects++
		}
		prev = cur
	}
	return itersects&1 == 1
}

// NewPolygon creates a Polygon. It may not work correctly if the average of the
// points does not lie inside the polygon (such as a cercent shape).
func NewPolygon(vertexes []F) Polygon {
	var c F
	for _, v := range vertexes {
		c = c.Add(v)
	}
	c = c.ScalarMultiply(1 / float64(len(vertexes)))
	pp := make(PolarPolygon, len(vertexes))
	for i, v := range vertexes {
		pp[i] = v.Subtract(c).P()
	}
	return pp.Polygon(c)
}

// PolarPolygon is useful in constructing a Polygon when the order of the
// vertexes is not known.
type PolarPolygon []P

func (p PolarPolygon) Len() int           { return len(p) }
func (p PolarPolygon) Less(i, j int) bool { return p[i].A < p[j].A }
func (p PolarPolygon) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p PolarPolygon) Polygon(center F) Polygon {
	sort.Sort(p)
	ply := make(Polygon, len(p))
	for i, plr := range p {
		ply[i] = center.Add(plr.F())
	}
	return ply
}

func Rectangle2Points(p1, p2 F) Polygon {
	return Polygon{
		p1,
		F{p2.X, p1.Y},
		p2,
		F{p1.X, p2.Y},
	}
}

func Rectangle2PointsWidthLength(p F, w, l float64) Polygon {
	return Polygon{
		p,
		F{p.X + w, p.Y},
		F{p.X + w, p.Y + l},
		F{p.X, p.Y + l},
	}
}

func RegularPolygonRadius(c F, r, a float64, n int) Polygon {
	ps := make(Polygon, n)
	p := P{r, a}
	da := Tau / float64(n)
	for i := range ps {
		ps[i] = p.F().Add(c)
		p.A += da
	}
	return ps
}

const (
	Tau = math.Pi * 2
	Pi  = math.Pi
)

var rpclC = math.Sin(Pi/2) / (2)

func RegularPolygonSideLength(c F, s, a float64, n int) Polygon {
	// A right triangle is formed with the hypotenuse being length r (which we
	// want to find), one angle being 360°/(2n) and the opposite side being length
	// (s/2). So the sine law gives us:
	// r / sin(90°) = (s/2) / sin(180°/n)
	// which is gives us
	// r = (s*sin(90°)) / (2*sin(180°/n))
	// r = (sin(90°)/2) * (r/sin(180°/n))
	// so (sin(90°)/2) comes out as a constant, rpclC
	ao := Pi / float64(n)
	r := (rpclC * s) / math.Sin(ao)
	// rotate backwards so the first line is tangent to the X axis
	a -= ao
	return RegularPolygonRadius(c, r, a, n)
}
