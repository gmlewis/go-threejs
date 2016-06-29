package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusBufferGeometry represents a torusbuffergeometry.
type TorusBufferGeometry struct{ p *js.Object }

// TorusBufferGeometry returns a TorusBufferGeometry object.
func (t *Three) TorusBufferGeometry() *TorusBufferGeometry {
	p := t.ctx.Get("TorusBufferGeometry")
	return &TorusBufferGeometry{p: p}
}

// New returns a new TorusBufferGeometry object.
func (t *TorusBufferGeometry) New(radius, tube, radialSegments, tubularSegments, arc float64) *TorusBufferGeometry {
	p := t.p.New(radius, tube, radialSegments, tubularSegments, arc)
	return &TorusBufferGeometry{p: p}
}

