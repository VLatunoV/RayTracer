package camera

import (
	"github.com/VLatunoV/RayTracer/geometry"
	"github.com/VLatunoV/RayTracer/util"
	"math"
)

type PerspectiveCamera struct {
	Transform   util.Transform
	AspectRatio util.Float
	FOV         util.Float
	frontDir    util.Vec3
	rightDir    util.Vec3
	downDir     util.Vec3
}

// MakePerspectiveCamera returns a PerspectiveCamera with specified aspect ratio and FOV in degrees
// Its position is at the origin with the viewing direction along the Z axis
func MakePerspectiveCamera(aspectRatio, fov float64) PerspectiveCamera {
	width := util.Float(2.0 * math.Tan(util.DegreeToRad(util.Float(fov))))
	height := width / util.Float(aspectRatio)
	return PerspectiveCamera{
		Transform:   util.GetIdentityTransform(),
		AspectRatio: util.Float(aspectRatio),
		FOV:         util.Float(fov),

		frontDir: util.Vec3{
			Z: 1.0,
		},
		rightDir: util.Vec3{
			X: width,
		},
		downDir: util.Vec3{
			Y: -height,
		},
	}
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
		Pos: c.Transform.Translate,
		Dir: (&newDir).Normalized(),
	}
}
