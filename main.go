package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"

	vec "github.com/madshov/data-structures/algebraic"
)

func render(objects []Object, lights []Light, bw *bufio.Writer) {
	width := 640
	height := 480

	var invWidth, invHeight, fov, aspectRatio, angle float64
	invWidth = 1 / float64(width)
	invHeight = 1 / float64(height)
	fov = 30
	aspectRatio = float64(width) / float64(height)
	angle = math.Tan(math.Pi * 0.5 * fov / 180.0)

	var xx, yy float64

	// Write header string to buffer.
	bw.WriteString(fmt.Sprintf("P3\n %d %d\n255\n", width, height))

	// For each pixel in canvas.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			xx = (2*((float64(x)+0.5)*invWidth) - 1) * angle * aspectRatio
			yy = (1 - 2*((float64(y)+0.5)*invHeight)) * angle

			rayDir, err := vec.NewVector(3, xx, yy, -1)
			if err != nil {
				log.Fatal("cannot create ray vector")
			}
			rayDir.Normalize()

			// Create ray vector.
			origo, err := vec.NewZeroVector(3)
			if err != nil {
				log.Fatal("cannot create origo vector")
			}

			ray := Ray{
				Origo: *origo,
				Dir:   *rayDir,
			}
			// Trace ray vector.
			clr := ray.Trace(objects, lights, 0)

			// Create string of colors (RGB), and write string to bufer.
			bw.WriteString(fmt.Sprintf("%d %d %d ",
				int(math.Min(1, clr.GetCoord(0))*255),
				int(math.Min(1, clr.GetCoord(1))*255),
				int(math.Min(1, clr.GetCoord(2))*255),
			))
		}
	}

	// Write memory buffer to file.
	bw.Flush()
}

func main() {
	var (
		objects []Object
		lights  []Light
	)

	v0, _ := vec.NewVector(3, 0.0, 0.0, -20.0)
	c0, _ := vec.NewVector(3, 1.00, 0.32, 0.36)

	objects = append(objects, &sphere{
		4,
		*v0,
		object{
			surfaceColor: *c0,
			isReflective: true,
			transparence: 0.5,
		},
	})

	v1, _ := vec.NewVector(3, 5.0, 0.0, -25.0)
	c1, _ := vec.NewVector(3, 0.65, 0.77, 0.97)

	objects = append(objects, &sphere{
		3,
		*v1,
		object{
			surfaceColor: *c1,
			isReflective: true,
			transparence: 0.0,
		},
	})

	v2, _ := vec.NewVector(3, 5.0, -1.0, -15.0)
	c2, _ := vec.NewVector(3, 0.90, 0.76, 0.46)
	objects = append(objects, &sphere{
		2,
		*v2,
		object{
			surfaceColor: *c2,
			isReflective: true,
			transparence: 0,
		},
	})

	l0, _ := vec.NewVector(3, 0.0, 20.0, -30.0)
	ec0, _ := vec.NewVector(3, 3.0, 3.0, 3.0)

	lights = append(lights, Light{
		Center:        *l0,
		EmissionColor: *ec0,
	})

	f, err := os.Create("./scene.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a buffered writer from the file.
	buffWriter := bufio.NewWriter(f)

	render(objects, lights, buffWriter)
}
