package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// PolyhedronGeometry represents a solid with flat faces.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PolyhedronGeometry
type PolyhedronGeometry struct{ p *js.Object }

// PolyhedronGeometry returns a PolyhedronGeometry object.
func (t *Three) PolyhedronGeometry() *PolyhedronGeometry {
	p := t.ctx.Get("PolyhedronGeometry")
	return &PolyhedronGeometry{p: p}
}

// New returns a new PolyhedronGeometry object.
//
//     vertices — Array of points of the form [1,1,1, -1,-1,-1, ... ]
//     indices — Array of indices that make up the faces of the form [0,1,2, 2,3,0, ... ]
//     radius — Float - The radius of the final shape
//     detail — Integer - How many levels to subdivide the geometry. The more detail, the smoother the shape.
func (t *PolyhedronGeometry) New(vertices []float64, indices []int, radius float64, detail int) *PolyhedronGeometry {
	p := t.p.New(vertices, indices, radius, detail)
	return &PolyhedronGeometry{p: p}
}
