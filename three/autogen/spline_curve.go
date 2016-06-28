package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SplineCurve represents a splinecurve.
type SplineCurve struct{ p *js.Object }

// SplineCurve returns a splinecurve object.
func (t *Three) SplineCurve() *SplineCurve {
	p := t.ctx.Get("SplineCurve")
	return &SplineCurve{p: p}
}

// NewSplineCurve returns a new splinecurve object.
func (t *SplineCurve) New(points /* array of Vector2 */ float64) *SplineCurve {
	p := t.p.New(points /* array of Vector2 */)
	return &SplineCurve{p: p}
}

// GetPoint TODO description.
func (s *SplineCurve) GetPoint(t float64) *SplineCurve {
	s.p.Call("getPoint", t)
	return s
}

