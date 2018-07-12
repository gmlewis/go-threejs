// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubicBezierCurve3 represents a smooth 3D cubic bezier curve.
//
// http://threejs.org/docs/index.html#Reference/Extras.Curves/CubicBezierCurve3
type CubicBezierCurve3 struct{ *Curve }

// JSObject returns the underlying *js.Object.
func (c *CubicBezierCurve3) JSObject() *js.Object { return c.p }

// CubicBezierCurve3 returns a CubicBezierCurve3 JavaScript class.
func (t *Three) CubicBezierCurve3() *CubicBezierCurve3 {
	p := t.ctx.Get("CubicBezierCurve3")
	return CubicBezierCurve3FromJSObject(p)
}

// CubicBezierCurve3FromJSObject returns a wrapped CubicBezierCurve3 JavaScript class.
func CubicBezierCurve3FromJSObject(p *js.Object) *CubicBezierCurve3 {
	return &CubicBezierCurve3{CurveFromJSObject(p)}
}

// NewCubicBezierCurve3 returns a new CubicBezierCurve3 object.
func (t *Three) NewCubicBezierCurve3(v0, v1, v2, v3 *Vector3) *CubicBezierCurve3 {
	p := t.ctx.Get("CubicBezierCurve3").New(v0.p, v1.p, v2.p, v3.p)
	return CubicBezierCurve3FromJSObject(p)
}

// GetPoint TODO description.
func (c *CubicBezierCurve3) GetPoint(t float64) *CubicBezierCurve3 {
	c.p.Call("getPoint", t)
	return c
}

// GetTangent TODO description.
func (c *CubicBezierCurve3) GetTangent(t float64) *CubicBezierCurve3 {
	c.p.Call("getTangent", t)
	return c
}
