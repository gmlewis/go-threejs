package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TorusKnotGeometry represents a torusknotgeometry.
type TorusKnotGeometry struct{ p *js.Object }

// TorusKnotGeometry returns a torusknotgeometry object.
func (t *Three) TorusKnotGeometry() *TorusKnotGeometry {
	p := t.ctx.Get("TorusKnotGeometry")
	return &TorusKnotGeometry{p: p}
}

// NewTorusKnotGeometry returns a new torusknotgeometry object.
func (t *TorusKnotGeometry) New(radius, tube, tubularSegments, radialSegments, p, q, heightScale float64) *TorusKnotGeometry {
	t = t.p.New(radius, tube, tubularSegments, radialSegments, p, q, heightScale)
	return &TorusKnotGeometry{p: t}
}
