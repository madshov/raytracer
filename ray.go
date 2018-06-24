package main

import (
  //"fmt"
  //"strconv"
  "math"
  "github.com/madshov/raytracer/vector"
)

type Ray struct {
  Origo vector.Vector3d
	Direction vector.Vector3d
}

func (ray *Ray) Trace(shapes []Shape, lights []Light, depth int) (*vector.Vector3d) {
  tnear := math.Inf(0)
  
  var sphere Shape
  t0, t1 :=  math.Inf(0), math.Inf(0)

	for _, shape := range shapes {
    // Calculate intersection parameters.
		t0, t1 = shape.Intersect(ray)

		if t0 < 0 {
			t0 = t1
		}

    // If object is closer than previous.
		if t0 < tnear {
			tnear = t0;
			sphere = shape;
		}
	}

  // If no object found, return background color.
	if math.IsInf(tnear, 0) {
		return vector.NewVector3d(1.0, 1.0, 1.0)
	}

 	// Find point of intersection
	hitPoint := ray.Origo.Add(ray.Direction.Multiplied(tnear))
	// Find normal at the intersection point
	normalHitPoint := sphere.GetNormalVector(hitPoint)
	// Normalize normal
	normalHitPoint.Normalized()
//fmt.Printf("Normal: %#v\n", normalHitPoint)
	inside := false

	if ray.Direction.Dot(normalHitPoint) > 0 {
	   normalHitPoint = normalHitPoint.Multiplied(-1)
		 inside = true
	}

	surfaceColor := vector.NewZeroVector3d()

  if (sphere.IsReflective() || sphere.hasTransparency() > 0.0) && depth < 5 {
      facingRatio := -1 * ray.Direction.Dot(normalHitPoint)

      // change the mix value to tweak the effect
      fresnelEffect := 1 * 0.1 + math.Pow(1 - facingRatio, 3) * (1 - 0.1)

      // -2aN + I where N is the normal vector to the sphere and intersection point. a is N dot product I, and I is the inverse of the direction of the incoming ray
      //reflectionDirection := normalHitPoint.Multiplied(ray.Direction.Dot(normalHitPoint) * -2).Add(&ray.Direction)
      reflectionDirection := ray.Direction.Subtract(normalHitPoint.Multiplied(2).Multiplied(ray.Direction.Dot(normalHitPoint)))
      reflectionDirection.Normalized()

      reflection := Ray{
				Origo: *hitPoint.Add(normalHitPoint.Multiplied(1e-4)),
				Direction: *reflectionDirection,
			}

      reflectionColor := reflection.Trace(shapes, lights, depth + 1)
      refractionColor := vector.NewZeroVector3d()

      if sphere.hasTransparency() > 0.0 {
        eta := 1.1

        if !inside {
          eta = 1.0 / eta
        }

        cosi := normalHitPoint.Multiplied(-1).Dot(&ray.Direction)
        k := 1.0 - eta * eta * (1.0 - cosi * cosi)

        refractionDirection := ray.Direction.Multiplied(eta).Add(normalHitPoint.Multiplied(eta * cosi - math.Sqrt(k)))
        refractionDirection.Normalized()

        refraction := Ray{
  				Origo: *hitPoint.Subtract(normalHitPoint.Multiplied(1e-4)),
  				Direction: *refractionDirection,
  			}

        refractionColor = refraction.Trace(shapes, lights, depth + 1)
      }

      v := reflectionColor.Multiplied(fresnelEffect)
      w := refractionColor.Multiplied((1.0 - fresnelEffect) * sphere.hasTransparency())
      //fmt.Printf("%#v\n", reflectionColor)
      surfaceColor = v.Add(w).Multiply(sphere.GetSurfaceColor())
  } else {
    for _, light := range lights {
      // Find light direction from intersection point
      lightDirection := light.Center.Subtract(hitPoint)
      // Normalize direction
      lightDirection.Normalized()

      //transmission := vector.NewVector3d(1.0, 1.0, 1.0)

      surfaceColor = surfaceColor.Add(sphere.GetSurfaceColor().Multiplied(math.Max(0.0, normalHitPoint.Dot(lightDirection))).Multiply(&light.EmissionColor))
    }
  }

	return surfaceColor
}
