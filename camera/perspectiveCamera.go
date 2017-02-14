package camera

import (
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/util"
)

type PerspectiveCamera struct {
	transform   util.Transform
	aspectRatio util.Float
	fov         util.Float
	frontDir    util.Vec3
	rightDir    util.Vec3
	downDir     util.Vec3
}

// Return the default camera directions (front, right, down). Not normalized as their length matters.
func getDefaultOrientation(aspectRatio, fov util.Float) (util.Vec3, util.Vec3, util.Vec3) {
	width := 2.0 * util.Tan(util.DegreeToRad(fov/2.0))
	height := width / aspectRatio
	return util.Vec3{Z: 1.0}, util.Vec3{X: width}, util.Vec3{Y: -height}
}

// MakePerspectiveCamera returns a PerspectiveCamera with specified aspect ratio and FOV in degrees.
// Its position is at the origin with the viewing direction along the Z axis.
func MakePerspectiveCamera(aspectRatio, fov float64) PerspectiveCamera {
	front, right, down := getDefaultOrientation(util.Float(aspectRatio), util.Float(fov))
	return PerspectiveCamera{
		transform:   util.GetIdentityTransform(),
		aspectRatio: util.Float(aspectRatio),
		fov:         util.Float(fov),

		frontDir: front,
		rightDir: right,
		downDir:  down,
	}
}

func (c *PerspectiveCamera) SetTransform(t util.Transform) {
	c.transform = t
	c.transform.CalculateRotationMatrix()

	c.frontDir, c.rightDir, c.downDir = getDefaultOrientation(c.aspectRatio, c.fov)

	c.frontDir.ApplyTransform(c.transform)
	c.rightDir.ApplyTransform(c.transform)
	c.downDir.ApplyTransform(c.transform)
}

// GetRay returns the normalized ray passing through the camera grid at location (x, y), where x and y are the
// weights of the right and down directions starting from the top left corner.
// Example: GetRay(0, 0)     --> Ray through the top left corner
//          GetRay(0.5, 0.5) --> Ray through the center
//          GetRay(0, 1)     --> Ray through the top right corner
func (c *PerspectiveCamera) GetRay(x, y util.Float) geometry.Ray {
	r := util.Mult(c.rightDir, (x - 0.5))
	d := util.Mult(c.downDir, (y - 0.5))
	newDir := util.Add(c.frontDir, util.Add(r, d))
	return geometry.Ray{
		Pos: c.transform.Translate,
		Dir: newDir.Normalized(),
	}
}
