package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LinearInterpolant represents a linearinterpolant.
type LinearInterpolant struct{ p *js.Object }

// LinearInterpolant returns a linearinterpolant object.
func (t *Three) LinearInterpolant() *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant")
	return &LinearInterpolant{p: p}
}

// NewLinearInterpolant returns a new linearinterpolant object.
func (t *LinearInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *LinearInterpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &LinearInterpolant{p: p}
}

// Interpolate_ TODO description.
func (l *LinearInterpolant) Interpolate_(i1, t0, t, t1 float64) *LinearInterpolant {
	l.p.Call("interpolate_", i1, t0, t, t1)
	return l
}

