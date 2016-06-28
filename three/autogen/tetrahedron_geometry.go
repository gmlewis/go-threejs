package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TetrahedronGeometry represents a tetrahedrongeometry.
type TetrahedronGeometry struct{ p *js.Object }

// TetrahedronGeometry returns a tetrahedrongeometry object.
func (t *Three) TetrahedronGeometry() *TetrahedronGeometry {
	p := t.ctx.Get("TetrahedronGeometry")
	return &TetrahedronGeometry{p: p}
}

// NewTetrahedronGeometry returns a new tetrahedrongeometry object.
func (t *TetrahedronGeometry) New(radius, detail float64) *TetrahedronGeometry {
	p := t.p.New(radius, detail)
	return &TetrahedronGeometry{p: p}
}

