package main

import (
	vec "github.com/madshov/data-structures/algebraic"
)

type Shape interface {
	Intersect(*Ray) (float64, float64)
	GetNormalVector(*vec.Vector) *vec.Vector
	GetSurfaceColor() *vec.Vector
	IsReflective() bool
	Transparence() float64
}

type shape struct {
	surfaceColor vec.Vector
	isReflective bool
	transparence float64
}

func (s *shape) GetSurfaceColor() *vec.Vector {
	return &s.surfaceColor
}

func (s *shape) IsReflective() bool {
	return s.isReflective
}

func (s *shape) Transparence() float64 {
	return s.transparence
}
