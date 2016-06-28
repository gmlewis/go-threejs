package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PolyhedronGeometry represents a polyhedrongeometry.
type PolyhedronGeometry struct{ p *js.Object }

// PolyhedronGeometry returns a PolyhedronGeometry object.
func (t *Three) PolyhedronGeometry() *PolyhedronGeometry {
	p := t.ctx.Get("PolyhedronGeometry")
	return &PolyhedronGeometry{p: p}
}

// New returns a new PolyhedronGeometry object.
func (t *PolyhedronGeometry) New(vertices, indices, radius, detail float64) *PolyhedronGeometry {
	p := t.p.New(vertices, indices, radius, detail)
	return &PolyhedronGeometry{p: p}
}

