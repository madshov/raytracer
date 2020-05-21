package main

import (
	"math"

	vec "github.com/madshov/data-structures/algebraic"
)

type Ray struct {
	Origo vec.Vector
	Dir   vec.Vector
}

func (ray *Ray) Trace(objects []Object, lights []Light, depth int) *vec.Vector {
	var (
		obj    Object
		tnear  = math.Inf(0)
		t0, t1 = math.Inf(0), math.Inf(0)
		inside bool
	)

	for _, object := range objects {
		// calculate intersection parameters
		t0, t1 = object.Intersect(ray)

		if t0 < 0 {
			t0 = t1
		}

		// if object is closer than previous
		if t0 < tnear {
			tnear = t0
			obj = object
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
	nIntersectPnt := obj.GetNormalVector(intersectPnt)
	// Normalize normal
	nIntersectPnt.Normalize()

	if ray.Dir.Dot(nIntersectPnt) > 0 {
		nIntersectPnt = nIntersectPnt.Scale(-1)
		inside = true
	}

	surfaceClr, _ := vec.NewZeroVector(3)

	if (obj.IsReflective() || obj.Transparence() > 0.0) && depth < 5 {
		facingRatio := -1 * ray.Dir.Dot(nIntersectPnt)

		// change the mix value to tweak the fresnel effect
		a, b := 0.1, 1.0
		fresnelEft := b*a + math.Pow(1-facingRatio, 3)*(1.0-a)

		// a is N dot product I, and I is the inverse of the direction of the
		// incoming ray
		// calculate reflection direction
		reflectDir := ray.Dir.Sub(nIntersectPnt.Scale(2).Scale(ray.Dir.Dot(nIntersectPnt)))
		reflectDir.Normalize()

		reflect := Ray{
			Origo: *(intersectPnt.Add(nIntersectPnt.Scale(1e-4))),
			Dir:   *reflectDir,
		}

		// trace reflection rays
		reflectClr := reflect.Trace(objects, lights, depth+1)

		refractClr, _ := vec.NewZeroVector(3)

		if obj.Transparence() > 0.0 {
			eta := 1.1

			if !inside {
				eta = 1.0 / eta
			}

			cosi := -1 * (nIntersectPnt.Dot(&ray.Dir))
			k := 1.0 - (eta * eta * (1.0 - cosi*cosi))

			// calculate refraction direction
			refractDir := ray.Dir.Scale(eta).Add(nIntersectPnt.Scale(eta*cosi - math.Sqrt(k)))
			refractDir.Normalize()

			refract := Ray{
				Origo: *(intersectPnt.Sub(nIntersectPnt.Scale(1e-4))),
				Dir:   *refractDir,
			}

			// trace refraction ray
			refractClr = refract.Trace(objects, lights, depth+1)
		}

		c1 := reflectClr.Scale(fresnelEft)
		c2 := refractClr.Scale(1.0 - fresnelEft).Scale(obj.Transparence())
		surfaceClr = (c1.Add(c2)).Mul(obj.GetSurfaceColor())
	} else {
		for _, l := range lights {
			// Find light direction from intersection point
			lightDir := l.Center.Sub(intersectPnt)
			// Normalize direction
			lightDir.Normalize()

			transmission, _ := vec.NewVector(3, 1.0, 1.0, 1.0)
			for _, object := range objects {
				light := Ray{
					Origo: *(intersectPnt.Add(nIntersectPnt.Scale(1e-4))),
					Dir:   *lightDir,
				}

				t0, _ := object.Intersect(&light)
				if !math.IsInf(t0, 0) {
					transmission, _ = vec.NewZeroVector(3)
					break
				}
			}

			max := math.Max(0.0, nIntersectPnt.Dot(lightDir))
			f := (obj.GetSurfaceColor().Mul(transmission)).Scale(max)
			surfaceClr = surfaceClr.Add(f.Mul(&l.EmissionColor))
		}
	}

	return surfaceClr
}
