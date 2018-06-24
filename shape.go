package main

import (
  "github.com/madshov/raytracer/vector"
)

type Shape interface {
  Intersect(*Ray) (float64, float64)
  GetNormalVector(*vector.Vector3d) (*vector.Vector3d)
  GetSurfaceColor() (*vector.Vector3d)
  IsReflective() (bool)
  hasTransparency() (float64)
  String() (string)
}
