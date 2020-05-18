package main

import (
	"math"

	vec "github.com/madshov/data-structures/algebraic"
)

type Ray struct {
	Origo vec.Vector
	Dir   vec.Vector
}

func (ray *Ray) Trace(shapes []Shape, lights []Light, depth int) *vec.Vector {
	tnear := math.Inf(0)

	var sh Shape
	t0, t1 := math.Inf(0), math.Inf(0)

	for _, shape := range shapes {
		// calculate intersection parameters
		t0, t1 = shape.Intersect(ray)

		if t0 < 0 {
			t0 = t1
		}

		// if object is closer than previous
		if t0 < tnear {
			tnear = t0
			sh = shape
		}
	}

	// If no object found, return background color.
	if math.IsInf(tnear, 0) {
		bgClr, _ := vec.NewVector(3, 2.0, 2.0, 2.0)
		return bgClr
	}

	// Find point of intersection
	intersectPnt := ray.Origo.Add(ray.Dir.Scale(tnear))
	// Find normal at the intersection point
	nIntersectPnt := sh.GetNormalVector(intersectPnt)
	// Normalize normal
	nIntersectPnt.Normalize()
	//fmt.Printf("Normal: %#v\n", normalHitPnt)
	inside := false

	if ray.Dir.Dot(nIntersectPnt) > 0 {
		nIntersectPnt = nIntersectPnt.Scale(-1)
		inside = true
	}

	surfaceClr, _ := vec.NewZeroVector(3)

	if (sh.IsReflective() || sh.Transparence() > 0.0) && depth < 5 {
		facingRatio := -1 * ray.Dir.Dot(nIntersectPnt)

		// change the mix value to tweak the effect
		a, b := 0.1, 1.0
		fresnelEft := b*a + math.Pow(1-facingRatio, 3)*(1.0-a)

		// -2aN + I where N is the normal vector to the sphere and intersection
		// point
		// a is N dot product I, and I is the inverse of the direction of the
		// incoming ray
		reflectDir := ray.Dir.Sub(nIntersectPnt.Scale(2).Scale(ray.Dir.Dot(nIntersectPnt)))
		reflectDir.Normalize()

		reflect := Ray{
			Origo: *(intersectPnt.Add(nIntersectPnt).Scale(1e-4)),
			Dir:   *reflectDir,
		}

		reflectClr := reflect.Trace(shapes, lights, depth+1)
		refractClr, _ := vec.NewZeroVector(3)

		if sh.Transparence() > 0.0 {
			eta := 1.1

			if !inside {
				eta = 1.0 / eta
			}

			cosi := -1 * (nIntersectPnt.Dot(&ray.Dir))
			k := 1.0 - (eta * eta * (1.0 - cosi*cosi))

			refractDir := ray.Dir.Scale(eta).Add(nIntersectPnt.Scale(eta*cosi - math.Sqrt(k)))
			refractDir.Normalize()

			refract := Ray{
				Origo: *(intersectPnt.Sub(nIntersectPnt).Scale(1e-4)),
				Dir:   *refractDir,
			}

			refractClr = refract.Trace(shapes, lights, depth+1)
		}

		surfaceClr = reflectClr.Scale(fresnelEft).Add(refractClr.Scale((1.0 - fresnelEft) * sh.Transparence())).Mul(sh.GetSurfaceColor())
	} else {
		for _, l := range lights {
			// Find light direction from intersection point
			lightDir := l.Center.Sub(intersectPnt)
			// Normalize direction
			lightDir.Normalize()

			transmission, _ := vec.NewVector(3, 1.0, 1.0, 1.0)
			for _, shape := range shapes {
				light := Ray{
					Origo: *(intersectPnt.Add(nIntersectPnt).Scale(1e-4)),
					Dir:   *lightDir,
				}

				t0, _ := shape.Intersect(&light)
				if math.IsInf(t0, 0) {
					transmission, _ = vec.NewZeroVector(3)
					break
				}
			}

			max := math.Max(0.0, nIntersectPnt.Dot(lightDir))
			surfaceClr = surfaceClr.Add(sh.GetSurfaceColor().Mul(transmission).Scale(max).Mul(&l.EmissionClr))
		}
	}

	return surfaceClr
}
