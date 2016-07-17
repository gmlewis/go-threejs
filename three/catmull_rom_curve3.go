// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CatmullRomCurve3 represents a smooth 3D spline curve from
// a series of points using the Catmull-Rom algorithm.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/CatmullRomCurve3
type CatmullRomCurve3 struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (c *CatmullRomCurve3) JSObject() *js.Object { return c.p }

// CatmullRomCurve3 returns a CatmullRomCurve3 JavaScript class.
func (t *Three) CatmullRomCurve3() *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3")
	return CatmullRomCurve3FromJSObject(p)
}

// CatmullRomCurve3FromJSObject returns a wrapped CatmullRomCurve3 JavaScript class.
func CatmullRomCurve3FromJSObject(p *js.Object) *CatmullRomCurve3 {
	return &CatmullRomCurve3{CurveFromJSObject(p)}
}

// NewCatmullRomCurve3 returns a new CatmullRomCurve3 object.
func (t *Three) NewCatmullRomCurve3(pts []*js.Object) *CatmullRomCurve3 {
	p := t.ctx.Get("CatmullRomCurve3").New(pts)
	return CatmullRomCurve3FromJSObject(p)
}
