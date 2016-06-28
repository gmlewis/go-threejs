package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotGeometry represents a torusknotgeometry.
type TorusKnotGeometry struct{ p *js.Object }

// TorusKnotGeometry returns a TorusKnotGeometry object.
func (t *Three) TorusKnotGeometry() *TorusKnotGeometry {
	p := t.ctx.Get("TorusKnotGeometry")
	return &TorusKnotGeometry{p: p}
}

// New returns a new TorusKnotGeometry object.
func (t *TorusKnotGeometry) New(radius, tube, tubularSegments, radialSegments, p, q, heightScale float64) *TorusKnotGeometry {
	t = t.p.New(radius, tube, tubularSegments, radialSegments, p, q, heightScale)
	return &TorusKnotGeometry{p: t}
}
