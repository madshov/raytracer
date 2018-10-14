package main

import (
	"fmt"
	"math"
	"os"
	"bufio"
	"github.com/madshov/data-structures/vector"
)

func render(objects []Object, lights []Light) {
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

 	// Create a buffered writer from the file.
	bufferedWriter := bufio.NewWriter(file)

	// Write header string to buffer.
	bufferedWriter.WriteString(fmt.Sprintf("P3\n %d %d\n255\n", width, height));

	// For each pixel in canvas.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xx = (2 * ((float64(x) + 0.5) * invWidth) - 1) * angle * aspectRatio
      		yy = (1 - 2 * ((float64(y) + 0.5) * invHeight)) * angle

			rayDirection := vector.NewVector3d(xx, yy, -1)
			rayDirection.Normalized();

			// Create ray vector.
			ray := Ray{
				Origo: vector.Vector3d{
					X: 0.0, 
					Y: 0.0, 
					Z: 0.0,
				},
				Direction: *rayDirection,
			}
			// Trace ray vector.
			color := ray.Trace(objects, lights, 0)
			// Create string of colors (RGB), and write string to bufer.
			bufferedWriter.WriteString(fmt.Sprintf("%d %d %d ", 
				int(math.Min(1, color.X) * 255), 
				int(math.Min(1, color.Y) * 255), 
				int(math.Min(1, color.Z) * 255)
			))
		}
	}

	// Write memory buffer to file.
	bufferedWriter.Flush()
}

func main() {
	var objects = make([]Object, 3)

	//for _ := range objects {
    objects[0] = &Sphere{
		Radius: 3,
		Center: vector.Vector3d{X:-0, Y:-0, Z:-20.0},
		SurfaceColor: vector.Vector3d{X:1.00, Y:0.32, Z:0.36},
		Reflection: false,
		Transparency: 0,
	}
	objects[1] = &Sphere{
		Radius: 5,
		Center: vector.Vector3d{X:2.5, Y:2.5, Z:-30.0},
		SurfaceColor: vector.Vector3d{X:0.0, Y:0.4, Z:1.0},
		Reflection: false,
		Transparency: 0.0,
	}
	objects[2] = &Sphere{
		Radius: 0.3,
		Center: vector.Vector3d{X:1, Y:0.3, Z:-5.0},
		SurfaceColor: vector.Vector3d{X:0.40, Y:0.32, Z:0.36},
		Reflection: false,
		Transparency: 0,
	}
	//}

	var lights = make([]Light, 1, 1)

	lights[0] = Light{
		Center: vector.Vector3d{X:20.0, Y:30.0, Z:10.0},
		EmissionColor: vector.Vector3d{X:1.0, Y:1.0, Z:1.0},
	}

	render(objects, lights)
}
