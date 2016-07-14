// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// QuadraticBezierCurve represents a quadraticbeziercurve.
type QuadraticBezierCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (q *QuadraticBezierCurve) JSObject() *js.Object { return q.p }

// QuadraticBezierCurve returns a QuadraticBezierCurve JavaScript class.
func (t *Three) QuadraticBezierCurve() *QuadraticBezierCurve {
	p := t.ctx.Get("QuadraticBezierCurve")
	return QuadraticBezierCurveFromJSObject(p)
}

// QuadraticBezierCurveFromJSObject returns a wrapped QuadraticBezierCurve JavaScript class.
func QuadraticBezierCurveFromJSObject(p *js.Object) *QuadraticBezierCurve {
	return &QuadraticBezierCurve{p: p}
}

// NewQuadraticBezierCurve returns a new QuadraticBezierCurve object.
func (t *Three) NewQuadraticBezierCurve(v0, v1, v2 float64) *QuadraticBezierCurve {
	p := t.ctx.Get("QuadraticBezierCurve").New(v0, v1, v2)
	return QuadraticBezierCurveFromJSObject(p)
}

// GetPoint TODO description.
func (q *QuadraticBezierCurve) GetPoint(t float64) *QuadraticBezierCurve {
	q.p.Call("getPoint", t)
	return q
}

// GetTangent TODO description.
func (q *QuadraticBezierCurve) GetTangent(t float64) *QuadraticBezierCurve {
	q.p.Call("getTangent", t)
	return q
}
