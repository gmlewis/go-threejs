package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineSegments represents a series of lines.
//
// http://threejs.org/docs/index.html#Reference/Objects/LineSegments
type LineSegments struct{ p *js.Object }

// LineSegments returns a LineSegments object.
func (t *Three) LineSegments() *LineSegments {
	p := t.ctx.Get("LineSegments")
	return &LineSegments{p: p}
}

// New returns a new LineSegments object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *LineSegments) New(geometry []*js.Object, material *js.Object) *LineSegments {
	p := t.p.New(geometry, material)
	return &LineSegments{p: p}
}
