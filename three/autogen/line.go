package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Line represents a continuous line.
//
// http://threejs.org/docs/index.html#Reference/Objects/Line
type Line struct{ p *js.Object }

// Line returns a Line object.
func (t *Three) Line() *Line {
	p := t.ctx.Get("Line")
	return &Line{p: p}
}

// New returns a new Line object.
//
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
func (t *Line) New(geometry []*js.Object, material *js.Object, mode float64) *Line {
	p := t.p.New(geometry, material, mode)
	return &Line{p: p}
}
