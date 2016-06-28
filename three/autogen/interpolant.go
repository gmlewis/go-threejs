package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Interpolant represents an interpolant.
type Interpolant struct{ p *js.Object }

// Interpolant returns an interpolant object.
func (t *Three) Interpolant() *Interpolant {
	p := t.ctx.Get("Interpolant")
	return &Interpolant{p: p}
}

// NewInterpolant returns a new interpolant object.
func (t *Interpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *Interpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &Interpolant{p: p}
}

// Evaluate TODO description.
func (i *Interpolant) Evaluate(t float64) *Interpolant {
	i.p.Call("evaluate", t)
	return i
}

// GetSettings_ TODO description.
func (i *Interpolant) GetSettings_() *Interpolant {
	i.p.Call("getSettings_")
	return i
}

// CopySampleValue_ TODO description.
func (i *Interpolant) CopySampleValue_(index float64) *Interpolant {
	i.p.Call("copySampleValue_", index)
	return i
}

// Interpolate_ TODO description.
func (i *Interpolant) Interpolate_(i1, t0, t, t1 float64) *Interpolant {
	i.p.Call("interpolate_", i1, t0, t, t1)
	return i
}

// IntervalChanged_ TODO description.
func (i *Interpolant) IntervalChanged_(i1, t0, t1 float64) *Interpolant {
	i.p.Call("intervalChanged_", i1, t0, t1)
	return i
}

