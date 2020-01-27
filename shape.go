package main

import (
	v "github.com/madshov/data-structures/vector"
)

type Shape interface {
	Intersect(*Ray) (float64, float64)
	GetNormalVector(*v.Vector3d) *v.Vector3d
	GetSurfaceColor() *v.Vector3d
	IsReflective() bool
	Transparence() float64
}

type shape struct {
	surfaceColor v.Vector3d
	isReflective bool
	transparence float64
}

func (s *shape) GetSurfaceColor() *v.Vector3d {
	return &s.surfaceColor
}

func (s *shape) IsReflective() bool {
	return s.isReflective
}

func (s *shape) Transparence() float64 {
	return s.transparence
}
