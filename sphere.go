package main

import (
	"math"

	v "github.com/madshov/data-structures/vector"
)

type Sphere interface {
	Shape
}

type sphere struct {
	Radius float64
	Center v.Vector3d
	shape
}

// Intersect determines whether a given ray intersect the object.
func (s *sphere) Intersect(ray *Ray) (float64, float64) {
	l := ray.Origo.Subtract(&s.Center)
	a := ray.Direction.Dot(&ray.Direction)
	b := 2 * ray.Direction.Dot(l)
	c := s.Center.Dot(&s.Center) + ray.Origo.Dot(&ray.Origo) + -2*s.Center.Dot(&ray.Origo) - math.Pow(s.Radius, 2)

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
		t0, t1 = t1, t0
	}

	return t0, t1
}

func (s *sphere) GetNormalVector(point *v.Vector3d) *v.Vector3d {
	return point.Subtract(&s.Center)
}
