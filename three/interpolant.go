// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Interpolant represents an interpolant.
type Interpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (i *Interpolant) JSObject() *js.Object { return i.p }

// Interpolant returns an Interpolant JavaScript class.
func (t *Three) Interpolant() *Interpolant {
	p := t.ctx.Get("Interpolant")
	return &Interpolant{p: p}
}

// NewInterpolant returns a new Interpolant object.
func (t *Three) NewInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *Interpolant {
	p := t.ctx.Get("Interpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &Interpolant{p: p}
}

// Evaluate TODO description.
func (i *Interpolant) Evaluate(t float64) *Interpolant {
	i.p.Call("evaluate", t)
	return i
}

// GetSettings TODO description.
func (i *Interpolant) GetSettings() *Interpolant {
	i.p.Call("getSettings_")
	return i
}

// CopySampleValue TODO description.
func (i *Interpolant) CopySampleValue(index float64) *Interpolant {
	i.p.Call("copySampleValue_", index)
	return i
}

// Interpolate TODO description.
func (i *Interpolant) Interpolate(i1, t0, t, t1 float64) *Interpolant {
	i.p.Call("interpolate_", i1, t0, t, t1)
	return i
}

// IntervalChanged TODO description.
func (i *Interpolant) IntervalChanged(i1, t0, t1 float64) *Interpolant {
	i.p.Call("intervalChanged_", i1, t0, t1)
	return i
}
