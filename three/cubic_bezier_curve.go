package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubicBezierCurve represents a cubicbeziercurve.
type CubicBezierCurve struct{ p *js.Object }

// CubicBezierCurve returns a CubicBezierCurve object.
func (t *Three) CubicBezierCurve() *CubicBezierCurve {
	p := t.ctx.Get("CubicBezierCurve")
	return &CubicBezierCurve{p: p}
}

// New returns a new CubicBezierCurve object.
func (t *CubicBezierCurve) New(v0, v1, v2, v3 float64) *CubicBezierCurve {
	p := t.p.New(v0, v1, v2, v3)
	return &CubicBezierCurve{p: p}
}

// GetPoint TODO description.
func (c *CubicBezierCurve) GetPoint(t float64) *CubicBezierCurve {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CubicBezierCurve) GetTangent(t float64) *CubicBezierCurve {
	c.p.Call("getTangent", t)
	return c
}

