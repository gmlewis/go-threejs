package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Points is a class for displaying particles in the form of variable size points.
//
// http://threejs.org/docs/index.html#Reference/Objects/Points
type Points struct{ p *js.Object }

// Points returns a Points object.
func (t *Three) Points() *Points {
	p := t.ctx.Get("Points")
	return &Points{p: p}
}

// New returns a new Points object.
//
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
func (t *Points) New(geometry, material *js.Object) *Points {
	p := t.p.New(geometry, material)
	return &Points{p: p}
}
