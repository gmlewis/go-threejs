// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// QuaternionLinearInterpolant represents a quaternionlinearinterpolant.
type QuaternionLinearInterpolant struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *QuaternionLinearInterpolant) JSObject() *js.Object { return q.p }

// QuaternionLinearInterpolant returns a QuaternionLinearInterpolant JavaScript class.
func (t *Three) QuaternionLinearInterpolant() *QuaternionLinearInterpolant {
	p := t.ctx.Get("QuaternionLinearInterpolant")
	return &QuaternionLinearInterpolant{p: p}
}

// NewQuaternionLinearInterpolant returns a new QuaternionLinearInterpolant object.
func (t *Three) NewQuaternionLinearInterpolant(parameterPositions, sampleValues, sampleSize, resultBuffer float64) *QuaternionLinearInterpolant {
	p := t.ctx.Get("QuaternionLinearInterpolant").New(parameterPositions, sampleValues, sampleSize, resultBuffer)
	return &QuaternionLinearInterpolant{p: p}
}

// Interpolate TODO description.
func (q *QuaternionLinearInterpolant) Interpolate(i1, t0, t, t1 float64) *QuaternionLinearInterpolant {
	q.p.Call("interpolate_", i1, t0, t, t1)
	return q
}
