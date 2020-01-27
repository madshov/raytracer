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
	hitPnt := ray.Origo.Add(ray.Direction.Multiplied(tnear))
	// Find normal at the intersection point
	normalHitPnt := sh.GetNormalVector(hitPnt)
	// Normalize normal
	normalHitPnt.Normalized()
	//fmt.Printf("Normal: %#v\n", normalHitPnt)
	inside := false

	if ray.Direction.Dot(normalHitPnt) > 0 {
		normalHitPnt = normalHitPnt.Multiplied(-1)
		inside = true
	}

	surfaceClr := v.NewZeroVector3d()

	if (sh.IsReflective() || sh.Transparence() > 0.0) && depth < 5 {
		facingRatio := -1 * ray.Direction.Dot(normalHitPnt)

		// change the mix value to tweak the effect
		fresnelEft := 1*0.1 + math.Pow(1-facingRatio, 3)*(1-0.1)

		// -2aN + I where N is the normal vector to the sphere and intersection point. a is N dot product I, and I is the inverse of the direction of the incoming ray
		//reflectDir := normalHitPnt.Multiplied(ray.Direction.Dot(normalHitPnt) * -2).Add(&ray.Direction)
		reflecDir := ray.Direction.Subtract(normalHitPnt.Multiplied(2).Multiplied(ray.Direction.Dot(normalHitPnt)))
		reflecDir.Normalized()

		reflec := Ray{
			Origo:     *hitPnt.Add(normalHitPnt.Multiplied(1e-4)),
			Direction: *reflecDir,
		}

		reflectClr := reflec.Trace(shapes, lights, depth+1)
		refractClr := v.NewZeroVector3d()

		if sh.Transparence() > 0.0 {
			eta := 1.1

			if !inside {
				eta = 1.0 / eta
			}

			cosi := normalHitPnt.Multiplied(-1).Dot(&ray.Direction)
			k := 1.0 - eta*eta*(1.0-cosi*cosi)

			refractDir := ray.Direction.Multiplied(eta).Add(normalHitPnt.Multiplied(eta*cosi - math.Sqrt(k)))
			refractDir.Normalized()

			refract := Ray{
				Origo:     *hitPnt.Subtract(normalHitPnt.Multiplied(1e-4)),
				Direction: *refractDir,
			}

			refractClr = refract.Trace(shapes, lights, depth+1)
		}

		v := reflectClr.Multiplied(fresnelEft)
		w := refractClr.Multiplied((1.0 - fresnelEft) * sh.Transparence())
		//fmt.Printf("%#v\n", reflecClr)
		surfaceClr = v.Add(w).Multiply(sh.GetSurfaceColor())
	} else {
		for _, light := range lights {
			// Find light direction from intersection point
			lightDir := light.Center.Subtract(hitPnt)
			// Normalize direction
			lightDir.Normalized()

			//transmission := v.NewVector3d(1.0, 1.0, 1.0)

			surfaceClr = surfaceClr.Add(sh.GetSurfaceColor().Multiplied(math.Max(0.0, normalHitPnt.Dot(lightDir))).Multiply(&light.EmissionColor))
		}
	}

	return surfaceClr
}
