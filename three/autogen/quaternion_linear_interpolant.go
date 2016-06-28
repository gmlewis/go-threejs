package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// QuaternionLinearInterpolant represents a quaternionlinearinterpolant.
type QuaternionLinearInterpolant struct{ p *js.Object }

// QuaternionLinearInterpolant returns a quaternionlinearinterpolant object.
func (t *Three) QuaternionLinearInterpolant() *QuaternionLinearInterpolant {
	p := t.ctx.Get("QuaternionLinearInterpolant")
	return &QuaternionLinearInterpolant{p: p}
}

// NewQuaternionLinearInterpolant returns a new quaternionlinearinterpolant object.
func (t *QuaternionLinearInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *QuaternionLinearInterpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &QuaternionLinearInterpolant{p: p}
}

// Interpolate_ TODO description.
func (q *QuaternionLinearInterpolant) Interpolate_(i1, t0, t, t1 float64) *QuaternionLinearInterpolant {
	q.p.Call("interpolate_", i1, t0, t, t1)
	return q
}

