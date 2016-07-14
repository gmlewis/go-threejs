// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CurveUtils represents a curveutils.
type CurveUtils struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *CurveUtils) JSObject() *js.Object { return c.p }

// CurveUtils returns a CurveUtils JavaScript class.
func (t *Three) CurveUtils() *CurveUtils {
	p := t.ctx.Get("CurveUtils")
	return CurveUtilsFromJSObject(p)
}

// CurveUtilsFromJSObject returns a wrapped CurveUtils JavaScript class.
func CurveUtilsFromJSObject(p *js.Object) *CurveUtils {
	return &CurveUtils{p: p}
}

// NewCurveUtils returns a new CurveUtils object.
func (t *Three) NewCurveUtils() *CurveUtils {
	p := t.ctx.Get("CurveUtils").New()
	return CurveUtilsFromJSObject(p)
}

// TangentQuadraticBezier TODO description.
func (c *CurveUtils) TangentQuadraticBezier(t, p0, p1, p2 float64) *CurveUtils {
	c.p.Call("tangentQuadraticBezier", t, p0, p1, p2)
	return c
}

// TangentCubicBezier TODO description.
func (c *CurveUtils) TangentCubicBezier(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentCubicBezier", t, p0, p1, p2, p3)
	return c
}

// TangentSpline TODO description.
func (c *CurveUtils) TangentSpline(t, p0, p1, p2, p3 float64) *CurveUtils {
	c.p.Call("tangentSpline", t, p0, p1, p2, p3)
	return c
}

// Interpolate TODO description.
func (c *CurveUtils) Interpolate(p0, p1, p2, p3, t float64) *CurveUtils {
	c.p.Call("interpolate", p0, p1, p2, p3, t)
	return c
}
