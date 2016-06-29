package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// OctahedronGeometry represents an octahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/OctahedronGeometry
type OctahedronGeometry struct{ p *js.Object }

// OctahedronGeometry returns an OctahedronGeometry object.
func (t *Three) OctahedronGeometry() *OctahedronGeometry {
	p := t.ctx.Get("OctahedronGeometry")
	return &OctahedronGeometry{p: p}
}

// New returns a new OctahedronGeometry object.
//
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more vertices making it no longer an icosahedron. When detail is greater than 1, it's effectively a sphere.
func (t *OctahedronGeometry) New(radius, detail float64) *OctahedronGeometry {
	p := t.p.New(radius, detail)
	return &OctahedronGeometry{p: p}
}
