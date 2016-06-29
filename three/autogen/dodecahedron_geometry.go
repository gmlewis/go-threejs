package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DodecahedronGeometry represents a dodecagedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/DodecahedronGeometry
type DodecahedronGeometry struct{ p *js.Object }

// DodecahedronGeometry returns a DodecahedronGeometry object.
func (t *Three) DodecahedronGeometry() *DodecahedronGeometry {
	p := t.ctx.Get("DodecahedronGeometry")
	return &DodecahedronGeometry{p: p}
}

// New returns a new DodecahedronGeometry object.
//
//     radius — Radius of the dodecahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds vertices making it no longer a dodecahedron.
func (t *DodecahedronGeometry) New(radius, detail float64) *DodecahedronGeometry {
	p := t.p.New(radius, detail)
	return &DodecahedronGeometry{p: p}
}
