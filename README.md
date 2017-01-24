# RayTracer
An educational ray tracer. It is meant to be a complete application, but can be used as a library-like package.
## As a minimum it will have these features
### Geometry
* intersect with basic objects (plane, sphere, rectangle)
* intersect with triangle mesh
* translate / rotate / scale

### Textures
* diffuse
* specular

### Camera
* perspective

### Lighting
* hard shadows

### Scene
* load from file

## As a follow-up it can be better improved with
### Geometry
* bounding box
* k-d trees

### Textures
* normal
* bitmap
* reflection
* refraction
* layered

### Camera
* parallel
* fish eye

### Lighting
* soft shadows
* global illumination

### Scene
* save scenes
* save renderings
* environment map

## Building
The project uses [GL](https://github.com/go-gl/gl)/[GLFW](https://github.com/go-gl/glfw) for visualization. Make sure you have the correct dependencies installed before building.
A simple **go build** should do the job.
If you plan on building more than once, it might be better to use **go build -i .** instead.