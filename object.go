package main

import (
	vec "github.com/madshov/data-structures/algebraic"
)

type Object interface {
	Intersect(*Ray) (float64, float64)
	GetNormalVector(*vec.Vector) *vec.Vector
	GetSurfaceColor() *vec.Vector
	IsReflective() bool
	Transparence() float64
}

type object struct {
	surfaceColor vec.Vector
	isReflective bool
	transparence float64
}

func (o *object) GetSurfaceColor() *vec.Vector {
	return &o.surfaceColor
}

func (o *object) IsReflective() bool {
	return o.isReflective
}

func (o *object) Transparence() float64 {
	return o.transparence
}
