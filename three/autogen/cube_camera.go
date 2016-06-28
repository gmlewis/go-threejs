package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubeCamera represents a cube camera.
type CubeCamera struct{ p *js.Object }

// CubeCamera returns a CubeCamera object.
func (t *Three) CubeCamera() *CubeCamera {
	p := t.ctx.Get("CubeCamera")
	return &CubeCamera{p: p}
}

// New returns a new CubeCamera object.
func (t *CubeCamera) New(near, far, cubeResolution float64) *CubeCamera {
	p := t.p.New(near, far, cubeResolution)
	return &CubeCamera{p: p}
}
