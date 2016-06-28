package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CurveUtils represents a curveutils.
type CurveUtils struct{ p *js.Object }

// CurveUtils returns a CurveUtils object.
func (t *Three) CurveUtils() *CurveUtils {
	p := t.ctx.Get("CurveUtils")
	return &CurveUtils{p: p}
}

// New returns a new CurveUtils object.
func (t *CurveUtils) New() *CurveUtils {
	p := t.p.New()
	return &CurveUtils{p: p}
}

// TangentQuadraticBezier TODO description.
func (c *CurveUtils) TangentQuadraticBezier(t, p0, p1, p2 float64) *CurveUtils {
	c.p.Call("tangentQuadraticBezier", t, p0, p1, p2)
	return c
}

// TangentCubicBezier TODO description.
func (c *CurveUtils) TangentCubicBezier(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentCubicBezier", t, p0, p1, p2, p3)
	return c
}

// TangentSpline TODO description.
func (c *CurveUtils) TangentSpline(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentSpline", t, p0, p1, p2, p3)
	return c
}

// Interpolate TODO description.
func (c *CurveUtils) Interpolate(p0, p1, p2, p3, t float64) *CurveUtils {
	c.p.Call("interpolate", p0, p1, p2, p3, t)
	return c
}

