package vec2d

import (
	"math"
	"testing"
)

func TestF2d(t *testing.T) {
	v := F2d{5, 6}
	if v.X != 5 || v.Y != 6 {
		t.Error("expected 5,6")
	}
}

func TestAdd(t *testing.T) {
	v1 := F2d{5, 6}
	v2 := F2d{1, 2}
	v := v1.Add(v2)
	if v.X != 6 || v.Y != 8 {
		t.Error("expected 6,8")
	}
}

func TestRef(t *testing.T) {
	v1 := &F2d{5, 6}
	v2 := &F2d{1, 2}
	v := v1.Add(*v2)
	if v.X != 6 || v.Y != 8 {
		t.Error("expected 6,8")
	}
	if v1.X != 5 || v1.Y != 6 {
		t.Error("expected 6,8")
	}
}

func TestAngle(t *testing.T) {
	v := P{1, 1}.F2d()
	a := v.Angle()
	if a != 1 {
		t.Error("expected 1, got ", a)
	}
}

func TestSubtract(t *testing.T) {
	v1 := F2d{6, 4}
	v2 := F2d{2, 1}
	v := v1.Subtract(v2)
	if v.X != 4 || v.Y != 3 {
		t.Error("expected 4,3")
	}
}

func TestRotate(t *testing.T) {
	v := P{1, 1}.F2d()
	v = v.Rotate(1)
	a := v.Angle()
	if a != 2 {
		t.Error("Expected 2, got ", a)
	}
}

func TestMag(t *testing.T) {
	v := P{2, 1}.F2d()
	m := v.Mag()
	if m != 2 {
		t.Error("expected 2, got ", m)
	}
}

func TestIntersect1(t *testing.T) {
	ms := F2d{0, 1}
	me := F2d{2, 1}
	ss := F2d{1, 0}
	se := F2d{1, 2}
	i := MotionSurfaceIntersection(ms, me, ss, se)
	if i != 0.5 {
		t.Error("expected 0.5, got ", i)
	}
}

func TestIntersect2(t *testing.T) {
	ms := F2d{2, 2}
	me := F2d{10, 6}
	ss := F2d{8, 1}
	se := F2d{6, 8}
	i := MotionSurfaceIntersection(ms, me, ss, se)
	if i != 0.625 {
		t.Error("expected 0.625, got ", i)
	}
}

type Foo struct {
	F2d
	name string
}

func TestEmbed(t *testing.T) {
	foo1 := &Foo{
		F2d:  F2d{3, 1},
		name: "Adam",
	}
	foo2 := &Foo{
		F2d:  F2d{3, 1},
		name: "Adam Colton",
	}
	if foo1.F2d != foo2.F2d {
		t.Error("Expected equality")
	}
	if foo1.Distance(foo2.F2d) != 0 {
		t.Error("Expected distance 0")
	}
}

func TestTriangulage(t *testing.T) {
	testTrig(F2d{2, 3}, 2, 1, 2, 3, t)
	testTrig(F2d{1, 4}, 1.23, 0, math.Pi, 3, t)

}

func testTrig(p F2d, m, a1, a2, a3 float64, t *testing.T) {
	p1 := p.Add(P{m, a1}.F2d())
	p2 := p.Add(P{m, a2}.F2d())
	p3 := p.Add(P{m, a3}.F2d())

	got := Triangulate(p1, p2, p3)

	if p.Distance(got) > 0.001 {
		t.Error("Expected: ", p, "\nGot:", got)
	}
}

func TestLineTo(t *testing.T) {
	line := F2d{1, 3}.LineTo(F2d{2, 4})
	expected := F2d{1, 3}
	got := line(0)

	if expected != got {
		t.Error("Expected:", expected, "\nGot:", got)
	}

	expected = F2d{2, 4}
	got = line(1)

	if expected != got {
		t.Error("Expected:", expected, "\nGot:", got)
	}

	if line(line.AtX(0)).Y != 2.0 {
		t.Error("Wrong X at 0", line.AtX(0))
	}

	if line(line.AtY(0)).X != -2.0 {
		t.Error("Wrong Y at 0", line.AtY(0))
	}

	if line.B() != 2.0 {
		t.Error("Wrong B")
	}

	if line.M() != 1.0 {
		t.Error("Wrong M")
	}
}

func TestLineBetween(t *testing.T) {
	line := F2d{1, 2}.Bisect(F2d{0, 3})
	expected := F2d{0.5, 2.5}
	got := line(0)
	if expected != got {
		t.Error("Expected:", expected, "\nGot:", got)
	}

	if line.B() != 2.0 {
		t.Error("Wrong B")
	}

	if line.M() != 1.0 {
		t.Error("Wrong M")
	}

	p1, p2 := F2d{0.000000, 89.000000}, F2d{44.452365, 56.126156}
	line = p1.Bisect(p2)
	sx := []float64{0, 22.226183 / 2, 22.226183}
	for _, x := range sx {
		p := F2d{x, line(line.AtX(x)).Y}
		d1, d2 := p1.Distance(p), p2.Distance(p)
		dd := d1 - d2
		if dd*dd > 0.001 {
			t.Error("Not equadistant", x, d1, d2, p)
		}
	}
}

func TestLineIntersect(t *testing.T) {
	// normal line
	l1 := F2d{0, 1}.LineTo(F2d{1, 2})
	l2 := F2d{1, 0}.LineTo(F2d{2, 3})
	testLine(l1, l2, t)

	// l1 is vertical
	l1 = F2d{0, 1}.LineTo(F2d{0, 2})
	l2 = F2d{1, 0}.LineTo(F2d{2, 3})
	testLine(l1, l2, t)

	// l2 is vertical
	l1 = F2d{0, 1}.LineTo(F2d{1, 2})
	l2 = F2d{1, 0}.LineTo(F2d{1, 3})
	testLine(l1, l2, t)

	// lines are parallel
	l1 = F2d{0, 1}.LineTo(F2d{1, 2})
	l2 = F2d{1, 2}.LineTo(F2d{2, 3})
	t1, t2 := l1.Intersection(l2)
	if !math.IsNaN(t1) || !math.IsNaN(t2) {
		t.Error("Lines do not intersect, expected NaN")
	}

	// l1 is a point
	l1 = F2d{0, 1}.LineTo(F2d{0, 1})
	l2 = F2d{1, 2}.LineTo(F2d{2, 3})
	t1, t2 = l1.Intersection(l2)
	if !math.IsNaN(t1) || !math.IsNaN(t2) {
		t.Error("Lines do not intersect, expected NaN")
	}
}

func testLine(l1, l2 Line, t *testing.T) {
	t1, t2 := l1.Intersection(l2)

	if l1(t1) != l2(t2) {
		t.Error("Intersection incorrect:\n", t1, l1(t1), "\n", t2, l2(t2))
	}
}
