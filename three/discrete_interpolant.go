// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DiscreteInterpolant represents a discreteinterpolant.
type DiscreteInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (d *DiscreteInterpolant) JSObject() *js.Object { return d.p }

// DiscreteInterpolant returns a DiscreteInterpolant object.
func (t *Three) DiscreteInterpolant() *DiscreteInterpolant {
	p := t.ctx.Get("DiscreteInterpolant")
	return &DiscreteInterpolant{p: p}
}

// New returns a new DiscreteInterpolant object.
func (d *DiscreteInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *DiscreteInterpolant {
	p := d.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &DiscreteInterpolant{p: p}
}

// Interpolate_ TODO description.
func (d *DiscreteInterpolant) Interpolate_(i1, t0, t, t1 float64) *DiscreteInterpolant {
	d.p.Call("interpolate_", i1, t0, t, t1)
	return d
}
