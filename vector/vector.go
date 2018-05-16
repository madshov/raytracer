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
/*
func (s *Vector3d) Rotated(radians float64) *Vector3d {
	cr := math.Cos(radians)
	sr := math.Sin(radians)
	return &Vector3d{
		X: (s.X * cr) - (s.Y * sr),
		Y: (s.X * sr) + (s.Y * cr),
	}
}
*/
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
/*
func (s *Vector3d) Wrap(XLimit, YLimit float64) *Vector3d {
	var newX float64 = math.Mod(s.X, XLimit)
	var newY float64 = math.Mod(s.Y, YLimit)

	if newX < 0 {
		newX = newX + XLimit
	}
	if newY < 0 {
		newY = newY + YLimit
	}

	return &Vector3d{
		X: newX,
		Y: newY,
	}
}
*//*
// get the minimum distance between two points, taking into account wrapping
func (s *Vector3d) WrappedDistanceVector(to *Vector3d, XLimit, YLimit float64) *Vector3d {

	dX := math.Mod(XLimit+to.X-s.X, XLimit)
	if dX > XLimit/2 {
		dX = dX - XLimit
	}

	dY := math.Mod(YLimit+to.Y-s.Y, YLimit)
	if dY > YLimit/2 {
		dY = dY - YLimit
	}

	return &Vector3d{
		X: dX,
		Y: dY,
	}
}
*/
/*
func (s *Vector3d) Contain(XLimit, YLimit float64) *Vector3d {

	newX := math.Min(s.X, XLimit)
	if newX < 0 {
		newX = 0
	}

	newY := math.Min(s.Y, YLimit)
	if newY < 0 {
		newY = 0
	}

	return &Vector3d{
		X: newX,
		Y: newY,
	}
}
*/
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
/*
func NewRandomUnitVector() *Vector3d {
	unit := &Vector3d{
		X: 1,
		Y: 0,
	}
	return unit.Rotated(rand.Float64() * 2 * math.Pi)
}*/
