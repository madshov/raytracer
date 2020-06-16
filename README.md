# raytracer
This is a simple app for doing backwards ray tracing of 3D objects on a given 2D canvas. Currently only spheres are rendered. Each sphere is defined by a radius and center coordinates, and with a given color, and a few other properties like transparency and reflection. Each ray is traced 5 times using a recursive strategy, before tracing it to the light source.

## Specifications
- Canvas size: 640 x 480 pixels.
- Fresnel equation used to calculate the mixing of refraction and reflexion.
- Output: The app produces a 24-bit color image in a ppm file called `scene.ppm`.
