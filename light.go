package main

import (
  "raytracer/vector"
)

type Light struct {
  Center vector.Vector3d
  EmissionColor vector.Vector3d
}
