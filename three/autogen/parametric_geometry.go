package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ParametricGeometry represents a parametricgeometry.
type ParametricGeometry struct{ p *js.Object }

// ParametricGeometry returns a parametricgeometry object.
func (t *Three) ParametricGeometry() *ParametricGeometry {
	p := t.ctx.Get("ParametricGeometry")
	return &ParametricGeometry{p: p}
}

// NewParametricGeometry returns a new parametricgeometry object.
func (t *ParametricGeometry) New(fn, slices, stacks float64) *ParametricGeometry {
	p := t.p.New(fn, slices, stacks)
	return &ParametricGeometry{p: p}
}
