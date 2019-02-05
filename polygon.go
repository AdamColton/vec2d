package vec2d

import (
	"math"
	"sort"
	"strings"
)

// Polygon represents a Convex Polygon and fulfills Shape.
type Polygon []F

// FillCurve returns a Curve segment of the Polygon Surface.
func (p Polygon) FillCurve(t0 float64) Curve {
	n := (len(p) - 1)
	h := n - (n / 2) // half rounded up
	as := LineSegments(p[0:h])
	bs := LineSegments(p[h:])
	a := as.F(t0)
	b := bs.F(1 - t0)
	return Curve(a.LineTo(b))
}

// F fulfils Surface
func (p Polygon) F(t0, t1 float64) F {
	return p.FillCurve(t0)(t1)
}

// String lists the points as a string.
func (p Polygon) String() string {
	strs := make([]string, len(p))
	for i, f := range p {
		strs[i] = f.String()
	}
	return strings.Join(strs, ":")
}

// SignedArea returns the Area and may be negative depending on the polarity.
func (p Polygon) SignedArea() float64 {
	var s float64
	prev := p[len(p)-1]
	for _, cur := range p {
		s += prev.Cross(cur)
		prev = cur
	}
	return s / 2
}

// Area of the polygon
func (p Polygon) Area() float64 {
	return math.Abs(p.SignedArea())
}

// Centroid returns the center of mass of the polygon
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

// Contains returns true of the point f is inside of the polygon
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

// Perimeter returns the total length of the perimeter
func (p Polygon) Perimeter() float64 {
	var sum float64
	prev := p[0]
	for _, f := range p[1:] {
		sum += prev.Distance(f)
		prev = f
	}
	sum += prev.Distance(p[0])
	return sum
}

// GetAngles returns the counter-clockwise and clockwise angles of the polygon
// by index.
func (p Polygon) GetAngles() ([]int, []int) {
	var ccw []int
	var cw []int
	prevIdx := len(p) - 1
	prevF := p[prevIdx]
	prevPlr := p[prevIdx-1].Subtract(prevF).P()
	for i, f := range p {
		curPlr := prevF.Subtract(f).P()
		a := curPlr.A - prevPlr.A
		if a < 0 {
			a += 2 * math.Pi
		}
		if a > math.Pi {
			cw = append(cw, prevIdx)
		} else if a < math.Pi {
			ccw = append(ccw, prevIdx)
		}
		prevPlr, prevF, prevIdx = curPlr, f, i
	}
	return ccw, cw
}

// CountAngles returns the number of counter clockwise and clockwise angles
func (p Polygon) CountAngles() (int, int) {
	var ccw int
	var cw int
	prevF := p[len(p)-1]
	prevPlr := p[len(p)-2].Subtract(prevF).P()
	for _, f := range p {
		curPlr := prevF.Subtract(f).P()
		a := curPlr.A - prevPlr.A
		if a < 0 {
			a += 2 * math.Pi
		}
		if a > math.Pi {
			cw++
		} else if a < math.Pi {
			ccw++
		}
		prevPlr, prevF = curPlr, f
	}
	return ccw, cw
}

// Convex returns True if the polygon contains a convex angle.
func (p Polygon) Convex() bool {
	ccw, cw := p.CountAngles()
	return ccw == 0 || cw == 0
}

// CounterClockwise returns true if the points defining the polygon proceed
// counterclockwise.
func (p Polygon) CounterClockwise() bool {
	var sum float64

	prevIdx := len(p) - 1
	prevF := p[prevIdx]
	prevPlr := p[prevIdx-1].Subtract(prevF).P()
	for i, f := range p {
		curPlr := prevF.Subtract(f).P()
		a := curPlr.A - prevPlr.A
		if a < 0 {
			a += 2 * math.Pi
		}
		sum += a

		prevPlr, prevF, prevIdx = curPlr, f, i
	}
	sum -= math.Pi * 2
	return sum > -1E-10 && sum < 1E-10
}

// FindTriangles returns the index sets of the polygon broken up into triangles.
// Given a unit square it would return [[0,1,2], [0,2,3]] which means that
// the square can be broken up in to 2 triangles formed by the points at those
// indexes.
func (p Polygon) FindTriangles() [][3]int {
	var ts [][3]int

	idxMp := make([]int, len(p))
	for i := range p {
		idxMp[i] = i
	}

	for {
		if len(idxMp) == 3 {
			ts = append(ts, [3]int{idxMp[0], idxMp[1], idxMp[2]})
			break
		}

		cur := make(Polygon, len(idxMp))
		for i, idx := range idxMp {
			cur[i] = p[idx]
		}

		for i0 := 0; true; i0++ {
			i1 := (i0 + 1) % len(idxMp)
			i2 := (i0 + 2) % len(idxMp)
			ln := cur[i0].LineTo(cur[i2])
			if !cur.Contains(ln(0.5)) {
				continue
			}
			if _, idx, _ := p.Intersects(ln); idx != -1 {
				continue
			}
			ts = append(ts, [3]int{idxMp[i0], idxMp[i1], idxMp[i2]})
			idxMp = append(idxMp[0:i1], idxMp[i1+1:]...)
			break
		}
	}

	return ts
}

// Intersects returns the first side that is intersected by the given
// lineSegment, returning the parametic t for the lineSegment, the index of the
// side and the parametric t of the side
func (p Polygon) Intersects(lineSegment Line) (lineT float64, idx int, sideT float64) {
	lineT = math.NaN()
	idx = -1
	sideT = math.NaN()
	ln := len(p) - 1
	for i, f := range p {
		side := f.LineTo(p[(i+1)%ln])
		t0, t1 := lineSegment.Intersection(side)
		if t0 > 0 && t0 < 1 && t1 > 0 && t1 < 1 {
			if math.IsNaN(lineT) || lineT > t0 {
				lineT = t0
				idx = i
				sideT = t1
			}
		}
	}
	return
}

// NonIntersecting returns false if any two sides intersect. This requires
// O(N^2) time to check.
func (p Polygon) NonIntersecting() bool {
	side := make([]Line, len(p))
	prev := p[len(p)-1]
	for i, f := range p {
		side[i] = prev.LineTo(f)
		prev = f
	}
	// Each side needs to be check against each non-adjacent side with a greater
	// index.
	for i, si := range side[:len(side)-2] {
		for j, sj := range side[i+2:] {
			if i == 0 && j == len(side)-3 {
				// Do not check if first line segment intersects last line segment
				continue
			}
			t, _ := si.Intersection(sj)
			if !math.IsNaN(t) && t > 0 && t < 1.0 {
				return false
			}
		}
	}
	return true
}

// Reverse the order of the points defining the polygon
func (p Polygon) Reverse() Polygon {
	out := make([]F, len(p))
	l := len(p) - 1
	m := (len(p) + 1) / 2 //+1 causes round up
	for i := 0; i < m; i++ {
		out[i], out[l-i] = p[l-i], p[i]
	}
	return out
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
// vertexes is not known. It does not fulfill shape.
type PolarPolygon []P

// Len returns the number of points, fulfills sort.Interface
func (p PolarPolygon) Len() int { return len(p) }

// Less fulfills sort.Interface
func (p PolarPolygon) Less(i, j int) bool { return p[i].A < p[j].A }

// Swap fulfills sort.Interface
func (p PolarPolygon) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Polygon converts the PolarPolygon to a Polygon
func (p PolarPolygon) Polygon(center F) Polygon {
	sort.Sort(p)
	ply := make(Polygon, len(p))
	for i, plr := range p {
		ply[i] = center.Add(plr.F())
	}
	return ply
}

// RectangleToPoints takes two points and returns a Polygon representing a
// rectangle.
func RectangleToPoints(p1, p2 F) Polygon {
	return Polygon{
		p1,
		F{p2.X, p1.Y},
		p2,
		F{p1.X, p2.Y},
	}
}

// RectanglePointWidthLength takes one point, a width and a length and returns
// a Polygon rectangle.
func RectanglePointWidthLength(point F, width, length float64) Polygon {
	return Polygon{
		point,
		F{point.X + width, point.Y},
		F{point.X + width, point.Y + length},
		F{point.X, point.Y + length},
	}
}

// RegularPolygonRadius constructs a regular polygon. The radius is measured
// from the center of each side.
func RegularPolygonRadius(center F, radius, angle float64, sides int) Polygon {
	ps := make(Polygon, sides)
	p := P{radius, angle}
	da := Tau / float64(sides)
	for i := range ps {
		ps[i] = p.F().Add(center)
		p.A += da
	}
	return ps
}

const (
	// Tau constant
	Tau = math.Pi * 2
	// Pi constant
	Pi = math.Pi
)

var rpclC = math.Sin(Pi/2) / (2)

// RegularPolygonSideLength constructs a regular polygon defined by the length
// of the sides.
func RegularPolygonSideLength(center F, sideLength, angle float64, sides int) Polygon {
	// A right triangle is formed with the hypotenuse being length r (which we
	// want to find), one angle being 360°/(2n) and the opposite side being length
	// (s/2). So the sine law gives us:
	// r / sin(90°) = (s/2) / sin(180°/n)
	// which is gives us
	// r = (s*sin(90°)) / (2*sin(180°/n))
	// r = (sin(90°)/2) * (r/sin(180°/n))
	// so (sin(90°)/2) comes out as a constant, rpclC
	ao := Pi / float64(sides)
	r := (rpclC * sideLength) / math.Sin(ao)
	// rotate backwards so the first line is tangent to the X axis
	angle -= ao
	return RegularPolygonRadius(center, r, angle, sides)
}

// ConcavePolygon represents a Polygon with at least one concave angle.
type ConcavePolygon struct {
	concave   Polygon
	regular   Polygon
	triangles [][2]Triangle
}

// GetTriangles takes triangle indexes from FindTriangles and returns a slice
// of triangles. This can be used to map one polygon to another with the same
// number of sides.
func GetTriangles(triangles [][3]int, p Polygon) []Triangle {
	ts := make([]Triangle, len(triangles))
	for i, t := range triangles {
		ts[i][0] = p[t[0]]
		ts[i][1] = p[t[1]]
		ts[i][2] = p[t[2]]
	}
	return ts
}

// NewConcavePolygon converts a Polygon to a ConcavePolygon
func NewConcavePolygon(concave Polygon) ConcavePolygon {
	regular := RegularPolygonRadius(F{}, 1, 0, len(concave))
	tIdxs := concave.FindTriangles()
	cts := GetTriangles(tIdxs, concave)
	rts := GetTriangles(tIdxs, regular)
	ts := make([][2]Triangle, len(tIdxs))
	for i := range cts {
		ts[i][0] = rts[i]
		ts[i][1] = cts[i]
	}

	return ConcavePolygon{
		concave:   concave,
		regular:   regular,
		triangles: ts,
	}
}

// F fulfils Surface
func (c ConcavePolygon) F(t0, t1 float64) F {
	f := c.regular.F(t0, t1)

	for _, ts := range c.triangles {
		if !ts[0].Contains(f) {
			continue
		}
		tfrm, _ := TriangleTransform(ts[0], ts[1])
		return tfrm.Apply(f)
	}

	// point is on perimeter
	for _, ts := range c.triangles {
		if ts[0][0].LineTo(ts[0][1]).Closest(f).Distance(f) < 1E-5 ||
			ts[0][1].LineTo(ts[0][2]).Closest(f).Distance(f) < 1E-5 ||
			ts[0][2].LineTo(ts[0][0]).Closest(f).Distance(f) < 1E-5 {
			tfrm, _ := TriangleTransform(ts[0], ts[1])
			return tfrm.Apply(f)
		}
	}

	return F{}
}

// Area of the polygon
func (c ConcavePolygon) Area() float64 { return c.concave.Area() }

// SignedArea returns the Area and may be negative depending on the polarity.
func (c ConcavePolygon) SignedArea() float64 { return c.concave.SignedArea() }

// Perimeter returns the total length of the perimeter
func (c ConcavePolygon) Perimeter() float64 { return c.concave.Perimeter() }

// Contains returns true of the point f is inside of the polygon
func (c ConcavePolygon) Contains(f F) bool { return c.concave.Contains(f) }

// Centroid returns the center of mass of the polygon
func (c ConcavePolygon) Centroid() F { return c.concave.Centroid() }
