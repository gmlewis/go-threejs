package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TetrahedronGeometry represents a tetrahedrongeometry.
type TetrahedronGeometry struct{ p *js.Object }

// TetrahedronGeometry returns a TetrahedronGeometry object.
func (t *Three) TetrahedronGeometry() *TetrahedronGeometry {
	p := t.ctx.Get("TetrahedronGeometry")
	return &TetrahedronGeometry{p: p}
}

// New returns a new TetrahedronGeometry object.
func (t *TetrahedronGeometry) New(radius, detail float64) *TetrahedronGeometry {
	p := t.p.New(radius, detail)
	return &TetrahedronGeometry{p: p}
}

