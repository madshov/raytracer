package main

import (
  "github.com/madshov/data-structures/vector"
)

type Object interface {
  Intersect(*Ray) (float64, float64)
  GetNormalVector(*vector.Vector3d) (*vector.Vector3d)
  GetSurfaceColor() (*vector.Vector3d)
  IsReflective() (bool)
  hasTransparency() (float64)
  String() (string)
}
