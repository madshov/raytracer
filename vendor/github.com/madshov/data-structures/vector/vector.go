package vector

import (
	"math"
)

type Vector3d struct {
	X, Y, Z float64
}

func (s *Vector3d) Normalized() {
	mag := s.Magnitude()
	if mag != 0 {
		s.X = s.X / mag
		s.Y = s.Y / mag
		s.Z = s.Z / mag
	}
}

func (s *Vector3d) Magnitude() float64 {
	l2 := (s.X * s.X) + (s.Y * s.Y) + (s.Z * s.Z)
	if l2 == 0 {
		return 0
	}
	return math.Sqrt(l2)
}

func (s *Vector3d) MagnitudeSquared() float64 {
	return (s.X * s.X) + (s.Y * s.Y) + (s.Z * s.Z)
}

func (s *Vector3d) Subtract(v *Vector3d) *Vector3d {
	return &Vector3d{
		X: s.X - v.X,
		Y: s.Y - v.Y,
		Z: s.Z - v.Z,
	}
}

func (s *Vector3d) Add(v *Vector3d) *Vector3d {
	return &Vector3d{
		X: s.X + v.X,
		Y: s.Y + v.Y,
		Z: s.Z + v.Z,
	}
}

func (s *Vector3d) Multiply(v *Vector3d) *Vector3d {
	return &Vector3d{
		X: s.X * v.X,
		Y: s.Y * v.Y,
		Z: s.Z * v.Z,
	}
}

func (s *Vector3d) Divide(v *Vector3d) *Vector3d {
	return &Vector3d{
		X: s.X / v.X,
		Y: s.Y / v.Y,
		Z: s.Z / v.Z,
	}
}

func (s *Vector3d) Dot(v *Vector3d) float64 {
	return (s.X * v.X) + (s.Y * v.Y) + (s.Z * v.Z)
}

func (s *Vector3d) Multiplied(by float64) *Vector3d {
	return &Vector3d{
		X: s.X * by,
		Y: s.Y * by,
		Z: s.Z * by,
	}
}

func (s *Vector3d) Divided(by float64) *Vector3d {
	return &Vector3d{
		X: s.X / by,
		Y: s.Y / by,
		Z: s.Z / by,
	}
}

func NewVector3d(x, y, z float64) *Vector3d {
	return &Vector3d{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewZeroVector3d() *Vector3d {
	return &Vector3d{
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}
}
