package main

import (
	"fmt"
	"math"

	"github.com/madshov/data-structures/vector"
)

type Sphere struct {
	Radius       float64
	Center       vector.Vector3d
	SurfaceColor vector.Vector3d
	Reflection   bool
	Transparency float64
}

// Intersect determines whether a given ray intersect the object.
func (sphere *Sphere) Intersect(ray *Ray) (float64, float64) {
	L := ray.Origo.Subtract(&sphere.Center)
	a := ray.Direction.Dot(&ray.Direction)
	b := 2 * ray.Direction.Dot(L)
	c := sphere.Center.Dot(&sphere.Center) + ray.Origo.Dot(&ray.Origo) + -2*sphere.Center.Dot(&ray.Origo) - math.Pow(sphere.Radius, 2)

	var t0, t1 float64
	discr := b*b - 4*a*c

	if discr < 0 {
		t0 = math.Inf(0)
		t1 = t0
	} else if discr == 0 {
		t0 = -b / (2 * a)
		t1 = t0
	} else {
		t0 = (-b - math.Sqrt(discr)) / (2 * a)
		t1 = (-b + math.Sqrt(discr)) / (2 * a)
	}

	if t0 > t1 {
		tmp := t0
		t0 = t1
		t1 = tmp
	}

	return t0, t1
}

func (sphere *Sphere) GetNormalVector(point *vector.Vector3d) *vector.Vector3d {
	return point.Subtract(&sphere.Center)
}

func (sphere *Sphere) GetSurfaceColor() *vector.Vector3d {
	return &sphere.SurfaceColor
}

func (sphere *Sphere) IsReflective() bool {
	return sphere.Reflection
}

func (sphere *Sphere) hasTransparency() float64 {
	return sphere.Transparency
}

func (sphere Sphere) String() string {
	return fmt.Sprintf("Radius: %f, Center: (%f,%f,%f)", sphere.Radius, sphere.Center.X, sphere.Center.Y, sphere.Center.Z)
}
