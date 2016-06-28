package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OctahedronGeometry represents an octahedrongeometry.
type OctahedronGeometry struct{ p *js.Object }

// OctahedronGeometry returns an OctahedronGeometry object.
func (t *Three) OctahedronGeometry() *OctahedronGeometry {
	p := t.ctx.Get("OctahedronGeometry")
	return &OctahedronGeometry{p: p}
}

// New returns a new OctahedronGeometry object.
func (t *OctahedronGeometry) New(radius, detail float64) *OctahedronGeometry {
	p := t.p.New(radius, detail)
	return &OctahedronGeometry{p: p}
}

