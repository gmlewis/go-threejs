// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubicInterpolant represents a cubicinterpolant.
type CubicInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CubicInterpolant) JSObject() *js.Object { return c.p }

// CubicInterpolant returns a CubicInterpolant JavaScript class.
func (t *Three) CubicInterpolant() *CubicInterpolant {
	p := t.ctx.Get("CubicInterpolant")
	return &CubicInterpolant{p: p}
}

// NewCubicInterpolant returns a new CubicInterpolant object.
func (t *Three) NewCubicInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *CubicInterpolant {
	p := t.ctx.Get("CubicInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &CubicInterpolant{p: p}
}

// IntervalChanged TODO description.
func (c *CubicInterpolant) IntervalChanged(i1, t0, t1 float64) *CubicInterpolant {
	c.p.Call("intervalChanged_", i1, t0, t1)
	return c
}

// Interpolate TODO description.
func (c *CubicInterpolant) Interpolate(i1, t0, t, t1 float64) *CubicInterpolant {
	c.p.Call("interpolate_", i1, t0, t, t1)
	return c
}
