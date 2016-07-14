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
func (l *LinearInterpolant) JSObject() *js.Object { return l.p }

// LinearInterpolant returns a LinearInterpolant JavaScript class.
func (t *Three) LinearInterpolant() *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant")
	return LinearInterpolantFromJSObject(p)
}

// LinearInterpolantFromJSObject returns a wrapped LinearInterpolant JavaScript class.
func LinearInterpolantFromJSObject(p *js.Object) *LinearInterpolant {
	return &LinearInterpolant{p: p}
}

// NewLinearInterpolant returns a new LinearInterpolant object.
func (t *Three) NewLinearInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *LinearInterpolant {
	p := t.ctx.Get("LinearInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return LinearInterpolantFromJSObject(p)
}

// Interpolate TODO description.
func (l *LinearInterpolant) Interpolate(i1, t0, t, t1 float64) *LinearInterpolant {
	l.p.Call("interpolate_", i1, t0, t, t1)
	return l
}
