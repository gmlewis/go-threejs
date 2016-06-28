package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Spline represents a spline.
type Spline struct{ p *js.Object }

// Spline returns a spline object.
func (t *Three) Spline() *Spline {
	p := t.ctx.Get("Spline")
	return &Spline{p: p}
}

// NewSpline returns a new spline object.
func (t *Spline) New(points float64) *Spline {
	p := t.p.New(points)
	return &Spline{p: p}
}

