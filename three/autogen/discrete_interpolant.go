package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DiscreteInterpolant represents a discreteinterpolant.
type DiscreteInterpolant struct{ p *js.Object }

// DiscreteInterpolant returns a discreteinterpolant object.
func (t *Three) DiscreteInterpolant() *DiscreteInterpolant {
	p := t.ctx.Get("DiscreteInterpolant")
	return &DiscreteInterpolant{p: p}
}

// NewDiscreteInterpolant returns a new discreteinterpolant object.
func (t *DiscreteInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *DiscreteInterpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &DiscreteInterpolant{p: p}
}

// Interpolate_ TODO description.
func (d *DiscreteInterpolant) Interpolate_(i1, t0, t, t1 float64) *DiscreteInterpolant {
	d.p.Call("interpolate_", i1, t0, t, t1)
	return d
}

