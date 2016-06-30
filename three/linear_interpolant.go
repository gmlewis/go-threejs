// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LinearInterpolant represents a linearinterpolant.
type LinearInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *LinearInterpolant) JSObject() *js.Object { return t.p }

// LinearInterpolant returns a LinearInterpolant object.
func (t *Three) LinearInterpolant() *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant")
	return &LinearInterpolant{p: p}
}

// New returns a new LinearInterpolant object.
func (t *LinearInterpolant) New(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *LinearInterpolant {
	p := t.p.New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &LinearInterpolant{p: p}
}

// Interpolate_ TODO description.
func (l *LinearInterpolant) Interpolate_(i1, t0, t, t1 float64) *LinearInterpolant {
	l.p.Call("interpolate_", i1, t0, t, t1)
	return l
}
