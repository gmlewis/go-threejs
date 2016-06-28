package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubicInterpolant represents a cubicinterpolant.
type CubicInterpolant struct{ p *js.Object }

// CubicInterpolant returns a cubicinterpolant object.
func (t *Three) CubicInterpolant() *CubicInterpolant {
	p := t.ctx.Get("CubicInterpolant")
	return &CubicInterpolant{p: p}
}

// NewCubicInterpolant returns a new cubicinterpolant object.
func (t *CubicInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *CubicInterpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &CubicInterpolant{p: p}
}

// IntervalChanged_ TODO description.
func (c *CubicInterpolant) IntervalChanged_(i1, t0, t1 float64) *CubicInterpolant {
	c.p.Call("intervalChanged_", i1, t0, t1)
	return c
}

// Interpolate_ TODO description.
func (c *CubicInterpolant) Interpolate_(i1, t0, t, t1 float64) *CubicInterpolant {
	c.p.Call("interpolate_", i1, t0, t, t1)
	return c
}

