package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DodecahedronGeometry represents a dodecahedrongeometry.
type DodecahedronGeometry struct{ p *js.Object }

// DodecahedronGeometry returns a dodecahedrongeometry object.
func (t *Three) DodecahedronGeometry() *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry")
	return &DodecahedronGeometry{p: p}
}

// NewDodecahedronGeometry returns a new dodecahedrongeometry object.
func (t *DodecahedronGeometry) New(radius, detail float64) *DodecahedronGeometry {
	p := t.p.New(radius, detail)
	return &DodecahedronGeometry{p: p}
}

