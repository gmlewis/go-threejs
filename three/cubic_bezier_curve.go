// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubicBezierCurve represents a cubicbeziercurve.
type CubicBezierCurve struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CubicBezierCurve) JSObject() *js.Object { return c.p }

// CubicBezierCurve returns a CubicBezierCurve JavaScript class.
func (t *Three) CubicBezierCurve() *CubicBezierCurve {
	p := t.ctx.Get("CubicBezierCurve")
	return CubicBezierCurveFromJSObject(p)
}

// CubicBezierCurveFromJSObject returns a wrapped CubicBezierCurve JavaScript class.
func CubicBezierCurveFromJSObject(p *js.Object) *CubicBezierCurve {
	return &CubicBezierCurve{p: p}
}

// NewCubicBezierCurve returns a new CubicBezierCurve object.
func (t *Three) NewCubicBezierCurve(v0, v1, v2, v3 float64) *CubicBezierCurve {
	p := t.ctx.Get("CubicBezierCurve").New(v0, v1, v2, v3)
	return CubicBezierCurveFromJSObject(p)
}

// GetPoint TODO description.
func (c *CubicBezierCurve) GetPoint(t float64) *CubicBezierCurve {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CubicBezierCurve) GetTangent(t float64) *CubicBezierCurve {
	c.p.Call("getTangent", t)
	return c
}
