package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ClosedSplineCurve3 represents a closedsplinecurve3.
type ClosedSplineCurve3 struct{ p *js.Object }

// ClosedSplineCurve3 returns a ClosedSplineCurve3 object.
func (t *Three) ClosedSplineCurve3() *ClosedSplineCurve3 {
	p := t.ctx.Get("ClosedSplineCurve3")
	return &ClosedSplineCurve3{p: p}
}

// New returns a new ClosedSplineCurve3 object.
func (t *ClosedSplineCurve3) New(points float64) *ClosedSplineCurve3 {
	p := t.p.New(points)
	return &ClosedSplineCurve3{p: p}
}
