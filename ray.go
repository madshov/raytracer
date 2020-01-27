package main

import (
	"math"

	v "github.com/madshov/data-structures/vector"
)

type Ray struct {
	Origo     v.Vector3d
	Direction v.Vector3d
}

func (ray *Ray) Trace(shapes []Shape, lights []Light, depth int) *v.Vector3d {
	tnear := math.Inf(0)

	var sh Shape
	t0, t1 := math.Inf(0), math.Inf(0)

	for _, shape := range shapes {
		// Calculate intersection parameters.
		t0, t1 = shape.Intersect(ray)

		if t0 < 0 {
			t0 = t1
		}

		// If object is closer than previous.
		if t0 < tnear {
			tnear = t0
			sh = shape
		}
	}

	// If no object found, return background color.
	if math.IsInf(tnear, 0) {
		return v.NewVector3d(1.0, 1.0, 1.0)
	}

	// Find point of intersection
	hitPoint := ray.Origo.Add(ray.Direction.Multiplied(tnear))
	// Find normal at the intersection point
	normalHitPoint := sh.GetNormalVector(hitPoint)
	// Normalize normal
	normalHitPoint.Normalized()
	//fmt.Printf("Normal: %#v\n", normalHitPoint)
	inside := false

	if ray.Direction.Dot(normalHitPoint) > 0 {
		normalHitPoint = normalHitPoint.Multiplied(-1)
		inside = true
	}

	surfaceClr := v.NewZeroVector3d()

	if (sh.IsReflective() || sh.Transparence() > 0.0) && depth < 5 {
		facingRatio := -1 * ray.Direction.Dot(normalHitPoint)

		// change the mix value to tweak the effect
		fresnelEffect := 1*0.1 + math.Pow(1-facingRatio, 3)*(1-0.1)

		// -2aN + I where N is the normal vector to the sphere and intersection point. a is N dot product I, and I is the inverse of the direction of the incoming ray
		//reflectionDirection := normalHitPoint.Multiplied(ray.Direction.Dot(normalHitPoint) * -2).Add(&ray.Direction)
		reflecDir := ray.Direction.Subtract(normalHitPoint.Multiplied(2).Multiplied(ray.Direction.Dot(normalHitPoint)))
		reflecDir.Normalized()

		reflec := Ray{
			Origo:     *hitPoint.Add(normalHitPoint.Multiplied(1e-4)),
			Direction: *reflecDir,
		}

		reflecClr := reflec.Trace(shapes, lights, depth+1)
		refracClr := v.NewZeroVector3d()

		if sh.Transparence() > 0.0 {
			eta := 1.1

			if !inside {
				eta = 1.0 / eta
			}

			cosi := normalHitPoint.Multiplied(-1).Dot(&ray.Direction)
			k := 1.0 - eta*eta*(1.0-cosi*cosi)

			refractionDirection := ray.Direction.Multiplied(eta).Add(normalHitPoint.Multiplied(eta*cosi - math.Sqrt(k)))
			refractionDirection.Normalized()

			refraction := Ray{
				Origo:     *hitPoint.Subtract(normalHitPoint.Multiplied(1e-4)),
				Direction: *refractionDirection,
			}

			refracClr = refraction.Trace(shapes, lights, depth+1)
		}

		v := reflecClr.Multiplied(fresnelEffect)
		w := refracClr.Multiplied((1.0 - fresnelEffect) * sh.Transparence())
		//fmt.Printf("%#v\n", reflecClr)
		surfaceClr = v.Add(w).Multiply(sh.GetSurfaceColor())
	} else {
		for _, light := range lights {
			// Find light direction from intersection point
			lightDirection := light.Center.Subtract(hitPoint)
			// Normalize direction
			lightDirection.Normalized()

			//transmission := v.NewVector3d(1.0, 1.0, 1.0)

			surfaceClr = surfaceClr.Add(sh.GetSurfaceColor().Multiplied(math.Max(0.0, normalHitPoint.Dot(lightDirection))).Multiply(&light.EmissionColor))
		}
	}

	return surfaceClr
}
