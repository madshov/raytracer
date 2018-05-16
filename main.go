package main

import (
	"fmt"
	"math"
	"os"
	"bufio"
    "vector"
)

func render(shapes []raytracer.Shape, lights []raytracer.Light) {
	width := 640
	height := 480

	var invWidth, invHeight, fov, aspectRatio, angle float64
	invWidth = 1 / float64(width)
	invHeight = 1 / float64(height)
	fov = 30
	aspectRatio = float64(width) / float64(height)
	angle = math.Tan(math.Pi * 0.5 * fov / 180.0)

	var xx, yy float64

	//var buffer bytes.Buffer
	file, error := os.Create("test.ppm")
	defer file.Close()

	if error != nil {
		panic(error)
	}

 	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(file)

	// Write header string to buffer.
	bufferedWriter.WriteString(fmt.Sprintf("P3\n %d %d\n255\n", width, height));

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xx = (2 * ((float64(x) + 0.5) * invWidth) - 1) * angle * aspectRatio
      yy = (1 - 2 * ((float64(y) + 0.5) * invHeight)) * angle

			rayDirection := vector.NewVector3d(xx, yy, -1)
			rayDirection.Normalized();

			ray := raytracer.Ray{
				Origo: vector.Vector3d{0.0, 0.0, 0.0},
				Direction: *rayDirection,
			}

			// Trace ray.
			color := ray.Trace(shapes, lights, 0, false)
			// Create string of colors (RGB), and write string to bufer.
			bufferedWriter.WriteString(fmt.Sprintf("%d %d %d ", int(math.Min(1, color.X) * 255), int(math.Min(1, color.Y) * 255), int(math.Min(1, color.Z) * 255)))
		}
	}

	// Write memory buffer to file.
	bufferedWriter.Flush()
}

func main() {
	var shapes = make([]raytracer.Shape, 1, 1)

	//for _ := range shapes {
     shapes[0] = &raytracer.Sphere{
			 Radius: 3,
			 Center: vector.Vector3d{-0, -0, -20.0},
			 SurfaceColor: vector.Vector3d{1.00, 0.32, 0.36},
			 Reflection: true,
			 Transparency: 0.0,
		 }
//		 shapes[1] = &raytracer.Sphere{
//			 Radius: 1,
//			 Center: vector.Vector3d{ 0.5, 1.5, -30.0},
//			 SurfaceColor: vector.Vector3d{0.0, 0.0, 1.0},
//			 Reflection: false,
//			 Transparency: 0.0,
//			 }
	//}

	var lights = make([]raytracer.Light, 1, 1)

	lights[0] = raytracer.Light{
		Center: vector.Vector3d{10.0, 10.0, 10.0},
		EmissionColor: vector.Vector3d{1.0, 1.0, 1.0},
	}

	render(shapes, lights)
}
