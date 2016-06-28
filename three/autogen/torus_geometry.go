package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusGeometry represents a torusgeometry.
type TorusGeometry struct{ p *js.Object }

// TorusGeometry returns a TorusGeometry object.
func (t *Three) TorusGeometry() *TorusGeometry {
	p := t.ctx.Get("TorusGeometry")
	return &TorusGeometry{p: p}
}

// New returns a new TorusGeometry object.
func (t *TorusGeometry) New(radius, tube, radialSegments, tubularSegments, arc float64) *TorusGeometry {
	p := t.p.New(radius, tube, radialSegments, tubularSegments, arc)
	return &TorusGeometry{p: p}
}

