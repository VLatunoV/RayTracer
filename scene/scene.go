package scene

import (
	"github.com/VLatunoV/RayTracer/camera"
	"github.com/VLatunoV/RayTracer/light"
)

type Scene struct {
	Nodes  []Node
	Camera camera.PerspectiveCamera
	Lights []light.Light
}
