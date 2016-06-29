package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotBufferGeometry represents a torusknotbuffergeometry.
type TorusKnotBufferGeometry struct{ p *js.Object }

// TorusKnotBufferGeometry returns a TorusKnotBufferGeometry object.
func (t *Three) TorusKnotBufferGeometry() *TorusKnotBufferGeometry {
	p := t.ctx.Get("TorusKnotBufferGeometry")
	return &TorusKnotBufferGeometry{p: p}
}

// New returns a new TorusKnotBufferGeometry object.
func (t *TorusKnotBufferGeometry) New(radius, tube, tubularSegments, radialSegments, p, q float64) *TorusKnotBufferGeometry {
	t = t.p.New(radius, tube, tubularSegments, radialSegments, p, q)
	return &TorusKnotBufferGeometry{p: t}
}
