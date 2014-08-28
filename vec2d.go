package vec2d

import (
  "math"
)

type Vector struct {
  X,Y float64
}

func Vec(x,y float64) *Vector {
  return &Vector{x, y}
}

func (a *Vector) Add(b *Vector) (*Vector) {
  return &Vector{a.X + b.X, a.Y + b.Y}
}

func (a *Vector) Subtract(b *Vector) (*Vector) {
  return &Vector{a.X-b.X, a.Y-b.Y}
}

func (v *Vector) Angle() (float64){
  return math.Atan2(v.Y, v.X)
}

func Polar(m, a float64) (*Vector) {
  return &Vector{math.Cos(a)*m, math.Sin(a)*m}
}

func (v *Vector) Mag() (float64) {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Rotate(a float64){
  m := v.Mag()
  a += v.Angle()
  v.X = math.Cos(a)*m
  v.Y = math.Sin(a)*m
}

/*
The following is a solution to these parametric equations
  sStart.X + S*dS.X = mStart.X + M*dM.X
  sStart.Y + S*dS.Y = mStart.Y + M*dM.Y

  Most of the complexity arises from checking for zeros in denominators
  and using a different equation if they are found.
*/
func MotionSurfaceIntersection(mStart, mEnd, sStart, sEnd *Vector) float64 {
  dM := mEnd.Subtract(mStart)
  dS := sEnd.Subtract(sStart)
  var S, M float64
  if (dM.Y == 0) {
    if (dS.Y == 0) {
      if (mStart.Y != sStart.Y || dS.X == dM.X){
        return math.NaN()
      }
      S = (mStart.X - sStart.X) / (dS.X - dM.X)
      M = S
    } else {
      if (dM.X == 0) {
        if (mStart.X == sStart.X && mStart.Y == sStart.Y){
          return 0
        } else {
          return math.NaN()
        }
      }
      if (dS.Y == 0) {
        return math.NaN()
      }
      S = (mStart.Y - sStart.Y) / dS.Y
      M = (sStart.X + S*dS.X - mStart.X) / dM.X
    }
  } else {
    if (dS.Y == 0) {
      M = (sStart.Y - mStart.Y) / dM.Y
      S = (mStart.X + M*dM.X - sStart.X) / dS.X
    } else {
      if ( dM.X/dM.Y == dS.X/dS.Y ) {
        //TODO slopes are parallel check which end it hits first
        return math.NaN() //this isn't right, but it prevent an error and it's an edgecase
      }
      S = ( (dM.X/dM.Y) * (sStart.Y - mStart.Y) - sStart.X + mStart.X) / (dS.X - (dM.X*dS.Y/dM.Y))
      M = (sStart.Y + S*dS.Y - mStart.Y) / dM.Y
    }
  }
  if (S >= 0 && S <= 1 && M >= 0 && M <= 1){
    return M
  }
  return math.NaN()
}