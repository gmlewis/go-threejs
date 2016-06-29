package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoxGeometry represents a quadrilateral geometry primitive.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/BoxGeometry
type BoxGeometry struct{ p *js.Object }

// BoxGeometry returns a BoxGeometry object.
func (t *Three) BoxGeometry() *BoxGeometry {
	p := t.ctx.Get("BoxGeometry")
	return &BoxGeometry{p: p}
}

// BoxGeometryOpts represents options passed to the BoxGeometry constructor.
type BoxGeometryOpts struct {
	widthSegments  int // Number of segmented faces along the width of the sides. Default is 1.
	heightSegments int // Number of segmented faces along the height of the sides. Default is 1.
	depthSegments  int // Number of segmented faces along the depth of the sides. Default is 1.
}

// New returns a new BoxGeometry object.
//
//     width — Width of the sides on the X axis.
//     height — Height of the sides on the Y axis.
//     depth — Depth of the sides on the Z axis.
func (t *BoxGeometry) New(width, height, depth float64, opts *BoxGeometryOpts) *BoxGeometry {
	var p *js.Object
	if opts != nil {
		p = t.p.New(width, height, depth, opts.widthSegments, opts.heightSegments, opts.depthSegments)
	} else {
		p = t.p.New(width, height, depth)
	}
	return &BoxGeometry{p: p}
}
