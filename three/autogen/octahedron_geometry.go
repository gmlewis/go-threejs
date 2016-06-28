package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OctahedronGeometry represents an octahedrongeometry.
type OctahedronGeometry struct{ p *js.Object }

// OctahedronGeometry returns an octahedrongeometry object.
func (t *Three) OctahedronGeometry() *OctahedronGeometry {
	p := t.ctx.Get("OctahedronGeometry")
	return &OctahedronGeometry{p: p}
}

// NewOctahedronGeometry returns a new octahedrongeometry object.
func (t *OctahedronGeometry) New(radius, detail float64) *OctahedronGeometry {
	p := t.p.New(radius, detail)
	return &OctahedronGeometry{p: p}
}

