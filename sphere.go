package main

import (
	"math"

	vec "github.com/madshov/data-structures/algebraic"
)

type sphere struct {
	Radius float64
	Center vec.Vector
	object
}

// Intersect determines whether a given ray intersect the object, and returns
// two intersections. Both can be infinity which means no intersection.
func (s *sphere) Intersect(ray *Ray) (float64, float64) {
	// find vector from sphere center and ray origo,
	// and project it onto ray to get length
	l := s.Center.Sub(&ray.Origo)
	a := l.Dot(&ray.Dir)

	var (
		t0 = math.Inf(0)
		t1 = math.Inf(0)
	)

	// if length from origo to projection point is
	// negative, ray does not intersect
	if a < 0 {
		return t0, t1
	}

	// find length between sphere center and ray
	d := math.Sqrt(l.Dot(l) - math.Pow(a, 2))

	// if length is greater than squared radius,
	// ray does not intersect
	if d > s.Radius {
		return t0, t1
	}

	// find intersection(s)
	b := math.Sqrt(math.Pow(s.Radius, 2) - math.Pow(d, 2))

	t0, t1 = a-b, a+b

	return t0, t1
}

// GetNormalVector returns the normal vector of the sphere for a given point in
// space.
func (s *sphere) GetNormalVector(point *vec.Vector) *vec.Vector {
	return point.Sub(&s.Center)
}
