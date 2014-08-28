package vec2d

import (
  "testing"
)

func TestVec(t *testing.T) {
  v := Vec(5, 6)
  if (v.X != 5 || v.Y != 6){
    t.Error("expected 5,6")
  }
}

func TestAdd(t *testing.T) {
  v1 := Vec(5, 6)
  v2 := Vec(1, 2)
  v := v1.Add(v2)
  if (v.X != 6 || v.Y != 8){
    t.Error("expected 6,8")
  }
}

func TestAngle(t *testing.T) {
  v := Polar(1,1)
  a := v.Angle()
  if (a != 1) {
    t.Error("expected 1, got ", a)
  }
}

func TestSubtract(t *testing.T) {
  v1 := Vec(6,4)
  v2 := Vec(2,1)
  v := v1.Subtract(v2)
  if (v.X != 4 || v.Y != 3) {
    t.Error("expected 4,3")
  }
}

func TestRotate(t *testing.T) {
  v := Polar(1,1)
  v.Rotate(1)
  a := v.Angle()
  if (a != 2) {
    t.Error("Expected 2, got ", a)
  }
}

func TestMag(t *testing.T) {
  v := Polar(2,1)
  m := v.Mag()
  if (m != 2) {
    t.Error("expected 2, got ", m)
  }
}

func TestIntersect1(t *testing.T) {
  ms := Vec(0,1)
  me := Vec(2,1)
  ss := Vec(1,0)
  se := Vec(1,2)
  i := MotionSurfaceIntersection(ms, me, ss, se)
  if (i != 0.5) {
    t.Error("expected 0.5, got ", i)
  }
}

func TestIntersect2(t *testing.T) {
  ms := Vec(2,2)
  me := Vec(10,6)
  ss := Vec(8,1)
  se := Vec(6,8)
  i := MotionSurfaceIntersection(ms, me, ss, se)
  if (i != 0.625) {
    t.Error("expected 0.625, got ", i)
  }
}


/*
func Test(t *testing.T) {
  if () {
    t.Error()
  }
}
*/