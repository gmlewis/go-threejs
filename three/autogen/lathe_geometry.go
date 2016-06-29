package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LatheGeometry represents a mesh with axial symmetry.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/LatheGeometry
type LatheGeometry struct{ p *js.Object }

// LatheGeometry returns a LatheGeometry object.
func (t *Three) LatheGeometry() *LatheGeometry {
	p := t.ctx.Get("LatheGeometry")
	return &LatheGeometry{p: p}
}

// New returns a new LatheGeometry object.
//
//     points — Slice of Vector2s.
//     segments — the number of circumference segments to generate. Default is 12.
//     phiStart — the starting angle in radians. Default is 0.
//     phiLength — the radian (0 to 2PI) range of the lathed section 2PI is a closed lathe, less than 2PI is a portion. Default is 2*PI
func (t *LatheGeometry) New(points []*js.Object, segments int, phiStart, phiLength float64) *LatheGeometry {
	p := t.p.New(points, segments, phiStart, phiLength)
	return &LatheGeometry{p: p}
}
